---
tags: ["gcp"]
linkTitle: "Deploy to GCP"
title: "Deploy a Trustgrid Node in GCP"
---

Standing up a Trustgrid node in GCP uses a published Trustgrid machine image and the `gcloud` CLI. Trustgrid nodes in GCP use two network interfaces — a WAN interface for management and tunnel traffic, and a LAN interface for internal data traffic. Each interface must be on a separate VPC network.

## Prerequisites

### GCP Project Setup

- Compute Engine API must be enabled in your GCP project.
- Two VPC networks, each with a subnet in the target deployment zone:
  - **WAN network** — internet-facing, used for management traffic to the Trustgrid control plane and tunnel traffic to remote nodes.
  - **LAN network** — private, used for internal data traffic and, in clustered deployments, for inter-node communication.

### Machine Type

The validated and supported machine type for Trustgrid nodes in GCP is **`e2-medium`**. If your deployment requires a different machine type, contact Trustgrid support before proceeding.

### Networking

#### WAN Network

The WAN network must allow outbound access to the Trustgrid control plane. See [Network Requirements for All Nodes]({{<relref "/help-center/kb/site-requirements#network-requirements-for-all-nodes">}}) for the full list of required IPs and ports.

{{<alert color="warning">}}
Inbound TCP 8443 on the WAN network is only required if the node will act as a **gateway**. Edge nodes do not require any inbound WAN firewall rules.
{{</alert>}}

#### LAN Network

{{<alert>}}
The LAN firewall rules for TCP 9000 and ICMP are only required for **clustered deployments**. Standalone nodes do not need these rules.
{{</alert>}}

The following table summarizes the firewall rules required across both networks:

| Network | Direction | Protocol | Ports | Source/Destination | Purpose |
|---------|-----------|----------|-------|--------------------|---------|
| WAN | Egress | TCP | 80, 443, 8443 | Trustgrid control plane IPs | Control plane communication |
| WAN | Ingress | TCP | 8443 | 0.0.0.0/0 | Gateway tunnel traffic (gateway nodes only) |
| LAN | Ingress | TCP | 9000 | LAN subnet CIDR | Cluster communication (clustered nodes only) |
| LAN | Ingress | ICMP | — | LAN subnet CIDR | Cluster health checks (clustered nodes only) |
| LAN | Egress | TCP | 9000 | LAN subnet CIDR | Cluster communication (clustered nodes only) |
| LAN | Egress | ICMP | — | 0.0.0.0/0 | Cluster health checks (clustered nodes only) |

### IAM — Service Account for HA Route Failover

In [clustered deployments]({{<relref "/docs/clusters">}}), Trustgrid nodes automatically manage GCP VPC routes to handle failover. This requires a service account bound to a custom IAM role with the minimum permissions needed to create and delete routes.

{{<alert>}}
If you are deploying a standalone (non-clustered) node, you can skip this section. The service account is only required for HA cluster route failover.
{{</alert>}}

#### Step 1: Create the custom IAM role

```bash
gcloud iam roles create tgNodeRouteManager \
  --project=PROJECT_ID \
  --title="TrustGrid Node Route Manager" \
  --description="Allows TrustGrid nodes to manage VPC routes for HA failover" \
  --permissions=compute.routes.list,compute.routes.get,compute.routes.create,compute.routes.delete,compute.networks.updatePolicy
```

#### Step 2: Create the service account

```bash
gcloud iam service-accounts create tg-node \
  --project=PROJECT_ID \
  --display-name="TrustGrid Node"
```

#### Step 3: Bind the role to the service account

```bash
gcloud projects add-iam-policy-binding PROJECT_ID \
  --member="serviceAccount:tg-node@PROJECT_ID.iam.gserviceaccount.com" \
  --role="projects/PROJECT_ID/roles/tgNodeRouteManager"
```

Replace `PROJECT_ID` with your GCP project ID in all three commands.

---

## Deployment Options

Choose one of the following methods to deploy and register the node:

### Option 1: Automated Deployment via gcloud CLI

In this path, you create the node in the Trustgrid portal first to obtain a license key, then pass the key as instance metadata when deploying the VM. The node registers automatically on first boot — no additional portal steps are required after the VM is running.

#### Step 1: Create a Node in the Trustgrid Portal

Add a new node in the Trustgrid portal. When the node is created, a license key will be generated and copied to your clipboard. Save this key — you will need it in the next step and it cannot be reissued without recreating the node.

{{<alert>}}
The node will not appear as active in the portal until the VM boots and completes registration.
{{</alert>}}

#### Step 2: Deploy the VM

Run the following command, replacing all placeholder values:

```bash
gcloud compute instances create NODE_NAME \
  --project=PROJECT_ID \
  --zone=ZONE \
  --machine-type=e2-medium \
  --image-family=trustgrid-node \
  --image-project=trustgrid-images \
  --boot-disk-size=30GB \
  --can-ip-forward \
  --network-interface=network=WAN_NETWORK,subnet=WAN_SUBNET,nic-type=GVNIC \
  --network-interface=network=LAN_NETWORK,subnet=LAN_SUBNET,no-address,nic-type=GVNIC \
  --service-account=tg-node@PROJECT_ID.iam.gserviceaccount.com \
  --scopes=https://www.googleapis.com/auth/compute \
  --tags=tg-node \
  --metadata=tg-license-key=LICENSE_KEY,serial-port-enable=true
```

{{<fields>}}
{{% field "NODE_NAME" %}}
The name for the GCP VM instance.
{{% /field %}}
{{% field "PROJECT_ID" %}}
Your GCP project ID.
{{% /field %}}
{{% field "ZONE" %}}
The GCP zone to deploy into (e.g., `us-central1-a`). Must match the zone where your WAN and LAN subnets are configured.
{{% /field %}}
{{% field "WAN_NETWORK / WAN_SUBNET" %}}
The name of your internet-facing VPC network and its subnet.
{{% /field %}}
{{% field "LAN_NETWORK / LAN_SUBNET" %}}
The name of your private VPC network and its subnet. Note the `no-address` flag — the LAN interface should not have an external IP.
{{% /field %}}
{{% field "LICENSE_KEY" %}}
The license key generated when you created the node in the Trustgrid portal.
{{% /field %}}
{{</fields>}}

On first boot, the node will automatically:

1. Detect the Trustgrid environment from the license key
2. Generate a unique node identity and keys
3. Register with the Trustgrid control plane
4. Reboot and connect

Once registration is complete, the node will appear as online in the portal and is ready to use.

---

### Option 2: Manual Deployment with Remote Console Registration

In this path, you deploy the VM without a license key and then register the node manually using the GCP Serial Console and the Trustgrid remote registration utility. Use this approach if you need to register a node without pre-generating a license key from the portal.

#### Step 1: Deploy the VM

```bash
gcloud compute instances create NODE_NAME \
  --project=PROJECT_ID \
  --zone=ZONE \
  --machine-type=e2-medium \
  --image-family=trustgrid-node \
  --image-project=trustgrid-images \
  --boot-disk-size=30GB \
  --can-ip-forward \
  --network-interface=network=WAN_NETWORK,subnet=WAN_SUBNET,nic-type=GVNIC \
  --network-interface=network=LAN_NETWORK,subnet=LAN_SUBNET,no-address,nic-type=GVNIC \
  --service-account=tg-node@PROJECT_ID.iam.gserviceaccount.com \
  --scopes=https://www.googleapis.com/auth/compute \
  --tags=tg-node \
  --metadata=serial-port-enable=true
```

The `--metadata=serial-port-enable=true` flag is required to allow access to the GCP Serial Console for the registration step below.

#### Step 2: Register via Serial Console

Once the VM is running, open the **Serial Console** for the instance from the GCP console (navigate to **Compute Engine → VM instances**, click the instance name, then select **Connect to serial console**).

From the serial console, follow the Trustgrid remote registration process to register the node with the control plane. See [Remote Registration]({{<relref "/tutorials/local-console-utility/remote-registration">}}) for full instructions.

---

## HA Cluster Configuration

When deployed as a [cluster]({{<relref "/docs/clusters">}}), Trustgrid nodes manage GCP VPC routes to provide automatic failover. The active cluster master creates a route in the LAN VPC pointing the VPN network CIDR to its own LAN IP as the next hop. On failover, the new master deletes the old route and creates a replacement pointing to itself.

### Requirements

- **IP forwarding** must be enabled on both nodes (`--can-ip-forward` in the deploy command).
- **Both nodes** must be on the same LAN VPC network.
- The **service account** (`tg-node@PROJECT_ID.iam.gserviceaccount.com`) with the `tgNodeRouteManager` role must be attached to each node. See [IAM — Service Account for HA Route Failover](#iam--service-account-for-ha-route-failover) above.

### LAN Subnet Route

GCP assigns `/32` addresses to VM interfaces and only programs a default route on the first NIC (the WAN interface). The LAN interface does not automatically have a route for its own subnet. Without an explicit route, cluster nodes cannot reach each other directly over the LAN network.

Each node in the cluster must have a route configured for the LAN subnet on its LAN interface. This is configured within the Trustgrid portal on the cluster's LAN interface settings — it is not a `gcloud` step.
