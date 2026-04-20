---
tags: ["gcp"]
linkTitle: "Deploy to GCP"
title: "Deploy a Trustgrid Node in GCP"
---

Trustgrid nodes in GCP are deployed from a published Trustgrid machine image. Each node requires two network interfaces — a **Management** interface for control plane communication and tunnel traffic, and a **Data** interface for internal data traffic. Each interface must be attached to a separate VPC network. Nodes can be registered automatically on first boot using a license key, or manually using the GCP Serial Console.

## Prerequisites

### GCP Project Setup

The **Compute Engine API** must be enabled in your GCP project. This API is required to create and manage VM instances and associated resources (disks, network interfaces, firewall rules, and VPC routes). To enable it:

```bash
gcloud services enable compute.googleapis.com --project=PROJECT_ID
```

You will also need two VPC networks, each with a subnet in the target deployment zone:
- **Management network** — internet-facing, used for management traffic to the Trustgrid control plane and tunnel traffic to remote nodes.
- **Data network** — private, used for internal data traffic and, in clustered deployments, for inter-node communication.

### Machine Type

The validated and supported machine type family for Trustgrid nodes in GCP is **e2**. Any e2 instance size may be used depending on your deployment requirements. If your deployment requires a machine type outside the e2 family, contact Trustgrid support before proceeding.

### Networking

#### Management Network

The Management network must allow outbound access to the Trustgrid control plane. See [Network Requirements for All Nodes]({{<relref "/help-center/kb/site-requirements#network-requirements-for-all-nodes">}}) for the full list of required IPs and ports.

{{<alert color="warning">}}
Inbound TCP/UDP 8443 on the Management network is only required if the node will act as a **gateway** (a node that terminates tunnels from remote edge nodes). Edge nodes do not require any inbound Management firewall rules.
{{</alert>}}

#### Data Network

{{<alert>}}
The Data network firewall rule for TCP 9000 is only required for **clustered deployments**. Standalone nodes do not need this rule.
{{</alert>}}

##### MTU Configuration

{{<alert color="warning">}}
**Temporary workaround:** GCP's default VPC MTU of 1460 bytes must be manually configured on the Data (LAN) interface. Without this, traffic may silently drop or underperform due to packet fragmentation.

After the node is registered and visible in the portal, set the MTU on the Data interface to **1460** under the node's [interface settings]({{<relref "/docs/nodes/appliances/interfaces#mtu">}}).

This step will no longer be required once the next node release is available, which will detect and apply the correct MTU automatically.
{{</alert>}}

The following table summarizes the firewall rules required across both networks:

| Network | Direction | Protocol | Ports | Source/Destination | Purpose |
|---------|-----------|----------|-------|--------------------|---------|
| Management | Egress | TCP | 443, 8443 | Trustgrid control plane IPs — see [Network Requirements]({{<relref "/help-center/kb/site-requirements#network-requirements-for-all-nodes">}}) | Control plane communication |
| Management | Ingress | TCP/UDP | 8443 | 0.0.0.0/0 | Data plane tunnel traffic (gateway nodes that terminate tunnels from remote edge nodes only) |
| Data | Ingress | TCP | 9000 | Data subnet CIDR | Cluster communication (clustered nodes only) |
| Data | Egress | TCP | 9000 | Data subnet CIDR | Cluster communication (clustered nodes only) |

---

## Deployment Options

Choose one of the following methods to deploy and register the node:

### Option 1: Automated Deployment via gcloud CLI

In this path, you create the node in the Trustgrid portal first to obtain a license key, then pass the key as instance metadata when deploying the VM. The node registers automatically with the Trustgrid control plane on first boot.

#### Step 1: Obtain a License Key from the Trustgrid Portal

Generate a license key for the new node from the [Nodes page]({{<relref "/docs/nodes">}}) in the Trustgrid portal. Save this key — you will need it in the next step.

{{<alert>}}
The node will not appear in the portal until the VM instance successfully completes registration with the Trustgrid control plane.
{{</alert>}}

#### Step 2: Deploy the VM

Run the following command, replacing all placeholder values:

```bash
gcloud compute instances create NODE_NAME \
  --project=PROJECT_ID \
  --zone=ZONE \
  --machine-type=MACHINE_TYPE \
  --image-family=trustgrid-node \
  --image-project=trustgrid-images \
  --boot-disk-size=30GB \
  --can-ip-forward \
  --network-interface=network=MANAGEMENT_NETWORK,subnet=MANAGEMENT_SUBNET,nic-type=GVNIC \
  --network-interface=network=DATA_NETWORK,subnet=DATA_SUBNET,no-address,nic-type=GVNIC \
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
The GCP zone to deploy into (e.g., `us-central1-a`). Must match the zone where your Management and Data subnets are configured.
{{% /field %}}
{{% field "MACHINE_TYPE" %}}
The e2 instance size appropriate for your deployment (e.g., `e2-standard-2`). See [Machine Type](#machine-type) above.
{{% /field %}}
{{% field "MANAGEMENT_NETWORK / MANAGEMENT_SUBNET" %}}
The name of your internet-facing VPC network and its subnet.
{{% /field %}}
{{% field "DATA_NETWORK / DATA_SUBNET" %}}
The name of your private VPC network and its subnet. Note the `no-address` flag — the Data interface should not have an external IP.
{{% /field %}}
{{% field "LICENSE_KEY" %}}
The license key generated from the Trustgrid portal.
{{% /field %}}
{{</fields>}}

On first boot, the node will automatically:

1. Detect the Trustgrid environment from the license key
2. Generate a unique node identity and keys
3. Register with the Trustgrid control plane
4. Reboot and connect

Once registration is complete, the node will appear as online in the portal and is ready to use.

---

### Option 2: Manual Deployment with Serial Console Registration

In this path, you deploy the VM without a license key and then register the node manually using the GCP Serial Console and the Trustgrid remote registration utility. Use this approach if you need to register a node without pre-generating a license key from the portal.

#### Step 1: Deploy the VM

```bash
gcloud compute instances create NODE_NAME \
  --project=PROJECT_ID \
  --zone=ZONE \
  --machine-type=MACHINE_TYPE \
  --image-family=trustgrid-node \
  --image-project=trustgrid-images \
  --boot-disk-size=30GB \
  --can-ip-forward \
  --network-interface=network=MANAGEMENT_NETWORK,subnet=MANAGEMENT_SUBNET,nic-type=GVNIC \
  --network-interface=network=DATA_NETWORK,subnet=DATA_SUBNET,no-address,nic-type=GVNIC \
  --service-account=tg-node@PROJECT_ID.iam.gserviceaccount.com \
  --scopes=https://www.googleapis.com/auth/compute \
  --tags=tg-node \
  --metadata=serial-port-enable=true
```

The `--metadata=serial-port-enable=true` flag is required to allow access to the GCP Serial Console for the registration step below.

#### Step 2: Register via Serial Console

Once the VM is running, open the **Serial Console** for the instance from the GCP console (navigate to **Compute Engine → VM instances**, click the instance name, then select **Connect to serial console**).

{{<alert>}}
When prompted to log in at the serial console, use username `tgadmin` and the **GCP instance name** (the value of `NODE_NAME`) as the password.
{{</alert>}}

{{<alert color="warning">}}
Completing remote registration requires access to the Trustgrid portal. This step must be performed by someone with portal access — either the end customer or Trustgrid support.
{{</alert>}}

From the serial console, follow the Trustgrid remote registration process to register the node with the control plane. See [Remote Registration]({{<relref "/tutorials/local-console-utility/remote-registration">}}) for full instructions.

---

### Option 3: Terraform Deployment

The [trustgrid-infra-as-code](https://github.com/trustgrid/trustgrid-infra-as-code) repository provides purpose-built Terraform modules for deploying Trustgrid nodes in GCP. These modules encapsulate the prerequisite setup (IAM, firewall rules, compute instances) and are the recommended approach for repeatable or infrastructure-as-code deployments.

#### Available Modules

| Module | Path | Purpose |
|--------|------|---------|
| `trustgrid_single_node` | `gcp/terraform/modules/compute/trustgrid_single_node` | Deploys a Trustgrid node VM with two NICs, IP forwarding, and optional automatic registration |
| `trustgrid_node_service_account` | `gcp/terraform/modules/iam/trustgrid_node_service_account` | Creates a dedicated GCP service account for Trustgrid nodes |
| `trustgrid_cluster_route_role` | `gcp/terraform/modules/iam/trustgrid_cluster_route_role` | Creates and binds the custom IAM role required for HA cluster route failover |
| `trustgrid_mgmt_firewall` | `gcp/terraform/modules/network/trustgrid_mgmt_firewall` | Creates egress firewall rules for control plane, DNS, and metadata server access on the management network |
| `trustgrid_gateway_firewall` | `gcp/terraform/modules/network/trustgrid_gateway_firewall` | Creates the ingress firewall rule for gateway nodes to accept data-plane tunnel connections |

#### Examples

Complete, ready-to-use examples are maintained in the repository:

- [Single node (manual registration)](https://github.com/trustgrid/trustgrid-infra-as-code/tree/main/gcp/terraform/examples/single-node-manual)
- [Gateway cluster (HA)](https://github.com/trustgrid/trustgrid-infra-as-code/tree/main/gcp/terraform/examples/gateway-cluster-ha)
- [Gateway cluster (HA, full)](https://github.com/trustgrid/trustgrid-infra-as-code/tree/main/gcp/terraform/examples/gateway-cluster-ha-full)

{{<alert>}}
After deploying via Terraform, you will still need to configure the Data interface MTU to 1460. See [MTU Configuration]({{<relref "/tutorials/deployments/deploy-gcp/#mtu-configuration">}}).
{{</alert>}}

---

## Cluster Configuration

When deployed as a [cluster]({{<relref "/docs/clusters">}}), Trustgrid nodes manage GCP VPC routes to provide automatic cluster failover. The active cluster master creates a route in the Data VPC pointing the VPN network CIDR to its own Data IP as the next hop. On failover, the new master deletes the old route and creates a replacement pointing to itself.

### Requirements

- **IP forwarding** must be enabled on both nodes (`--can-ip-forward` in the deploy command).
- **Both nodes** must be on the same Data VPC network.

### IAM — Service Account for Cluster Route Failover

Trustgrid nodes require a GCP service account with permission to create and delete VPC routes to manage cluster failover. This must be configured before deploying the VMs.

#### Step 1: Create the custom IAM role

```bash
gcloud iam roles create tgNodeRouteManager \
  --project=PROJECT_ID \
  --title="TrustGrid Node Route Manager" \
  --description="Allows TrustGrid nodes to manage VPC routes for cluster failover" \
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

Replace `PROJECT_ID` with your GCP project ID in all three commands. The service account must be attached to each node VM via the `--service-account` flag in the deploy command.

### Data Network Route

GCP assigns `/32` addresses to VM interfaces and only programs a default route on the first NIC (the Management interface). The Data interface does not automatically have a route for its own subnet. Without an explicit route, cluster nodes cannot reach each other directly over the Data network.

A single cluster interface route must be configured for the Data network subnet on the Data interface — both nodes share this route. This is configured in the Trustgrid portal under the cluster's [interface settings]({{<relref "/docs/nodes/appliances/interfaces">}}), not via `gcloud`.
