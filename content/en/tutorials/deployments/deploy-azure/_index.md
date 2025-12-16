---
title: "Deploy to Azure"
---

## Azure Requirements
- An Azure subscription. If you don't have one, create a [free account](https://azure.microsoft.com/free/)
- An Azure resource group to deploy the resources into
- An Azure Virtual Network (vNet) with at least two subnets:
  - An "outside" subnet for the appliance to connect to the Trustgrid control plane and data plane gateways, and accept incoming connections if the Azure Trustgrid appliance will be acting as a [data plane gateway]({{< relref "/docs/nodes/appliances/gateway" >}})
	- For clustered appliances, this subnet must allow communication between the two appliances on for the cluster heartbeat (typically TCP 9000).
	- Clustered appliances will also need to be able to reach the Azure Management API endpoints. See [Azure Management API endpoints](https://learn.microsoft.com/en-us/azure/api-management/api-management-howto-ip-addresses) for ranges or utilize the [ApiManagement service tag](https://learn.microsoft.com/en-us/azure/virtual-network/service-tags-overview#available-service-tags) in your security group rules. See [Public Cloud Appliance Requirements]({{<ref "/help-center/kb/site-requirements#public-cloud-appliance-requirements">}}) for more details.
  - An "inside" subnet for communicating with other virtual machines and services within the Azure vNet
  - **(For Clustered Appliances)** An [Azure routing table]({{<relref "#azure-route-table">}}) associated with the "inside" subnet. 

  {{<alert color="warning" title="Azure Extension Conflicts">}} 
  * Azure Backup - Like most network virtual appliances the Trustgrid appliance does not support being backed up using Azure Backup. The act of taking a snapshot with Azure Backup will disrupt appliance operations and may cause data plane disruption and cluster failovers. 
  * Defender for Servers - The Trustgrid appliance OS is a minimal hardened Linux distribution. Installing Defender for Servers or similar endpoint protection software on the appliance may cause system instability and is not supported.
  {{</alert>}}

## VM Requirements

| Requirement | Description      |
| ----------- | ---------------- |
| Disk Size   | At least 30 GB   |
| Interfaces  | <ul><li>1 Public with a Public IP address</li><li>1 Private</li><ul> |
| CPU & RAM   | See [Instance Type below]({{< relref ".#instance-size" >}}) for recommendations |

### Instance Size

Trustgrid has validated using the [B-series burstable - Azure Virtual Machines](https://learn.microsoft.com/en-us/azure/virtual-machines/sizes-b-series-burstable) instance type.

VPN throughput is tied to CPU the recommended size depends on roles, expected throughput.

- For Gateway nodes expecting up to ~200Mbps throughput, Trustgrid recommends the Standard_B4ms or larger
- For Edge nodes expecting less than 100Mbps throughput, Trustgrid recommends the Standard_B2s or Standard_B2ms or larger

### Interfaces

One WAN interface with a public IP and one LAN interface on a private subnet. The nodes will need to be able to route to all required hosts/applications that need to communicate across the Trustgrid virtual network.

The LAN interface needs to have **IP Forwarding Enabled** in order to forward the traffic across the tunnel.

![IP Forwarding](azure-ip-config.png)

See [Azure virtual network traffic routing](https://learn.microsoft.com/en-us/azure/virtual-network/virtual-networks-udr-overview).

### Supported Regions

The Trustgrid official community image, `trustgrid-node-2204-prod`, in the public gallery `trustgrid-45680719-9aa7-43b9-a376-dc03bcfdb0ac` is currently published in the following regions. If you need to deploy in another region, please contact Trustgrid support. If you are not a direct customer of Trustgrid, please check with your vendor that is utilizing Trustgrid to have them contact support.

| Region Display Name | Region Name |
|-|-|
|East US|eastus|
|East US 2|eastus2|
|Central US|centralus|
|North Central US|northcentralus|
|South Central US|southcentralus|
|West US|westus|
|West US 2|westus2|
|West US 3|westus3|

### Network Access

For gateways:

- Outbound internet access to the [Trustgrid control plane networks]({{<ref "/help-center/kb/site-requirements">}}) and ability to resolve public DNS names.
- Inbound access required is the TCP port defined for the Trustgrid gateway service to listen on. Edge nodes will connect to the gateways public IP and port defined. The default port used is 8443.

For edge nodes:

- Outbound internet access to the [Trustgrid control plane networks]({{<ref "/help-center/kb/site-requirements">}}), outbound access to the IP and ports of the Trustgrid gateways, and ability to resolve public DNS names.
- No inbound access is required on the public interface.

For all clustered nodes:

- The cluster heartbeat runs on the LAN/inside interface on TCP Port 9000. This port will need to be open between both Trustgrid Gateways for failover to work correctly.

### Other VM Requirements
* [System-assigned Managed Identity](https://learn.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/qs-configure-portal-windows-vm) needs to be enabled for both VMs in the cluster. {{<tgimg src="system-managed-identity.png" alt="System-assigned Managed Identity" width="80%">}}
* [Boot Diagnostics](https://learn.microsoft.com/en-us/azure/virtual-machines/boot-diagnostics) needs to be enabled to allow access via the Serial Console

## Deployment Process

One of more Virtual Machines will need to be deployed into the target Azure subscription to act as the Trustgrid nodes using the official community image.  Then the [remote registration process]({{<relref "/tutorials/local-console-utility/remote-registration">}}) can be used to activate the nodes in the Trustgrid portal.

### Participants

* Site Tech - User(s) with permissions and skills to deploy new instances in Azure, create the required Managed System Identity shown above, and make changes in Azure to allow the required network connectivity
* Trustgrid User - User with permissions to [Activate nodes]({{<relref "/tutorials/local-console-utility/remote-registration#portal-activation-process">}}) in the Trustgrid portal (or API)

> If the Site Tech is not part of the organization that is a Trustgrid's direct customer, Trustgrid's professional service team will need documented approval from that customer before proceeding with assisting in the deployment.

### High-Level Process

1. The Site Tech should be able to complete the following steps independently:
	1. Build out prerequisite resources including Resource Groups, vNets, subnets and routing tables in Azure
	- For single node deployments: 
		1. Create VM Instances based of the official Trustgrid community image
	- For clustered deployments:
		1. Create a [routing table for the in LAN interface subnet]({{<relref "#azure-routing-table">}}) if it does not already exist
		1. Create two VM Instances based of the official Trustgrid community image
		1. Create the Azure IAM role as defined above
	1. Use the Azure VM Serial Console to start the registration process, this code then needs to be communicated securely to the Trustgrid User. 
1. Trustgrid Tech - 
	1. Activate the device with the target organization
	1. Confirm healthy functionality and connectivity to the required gateways
	1. Configure the nodes as needed (e.g. clustering, VPN, L4proxy)


### Deployment Methods

- [Deploy via Azure command line tool]({{<relref "/tutorials/deployments/deploy-azure/">}})
- [Deploy with Bicep Modules](https://github.com/trustgrid/trustgrid-infra-as-code/tree/main/azure/bicep) - This is a collection of Bicep modules that can be used to deploy Trustgrid nodes in Azure. It includes modules for deploying the virtual machines as well as creating IAM roles and networking resources.
- [Deploy with Terraform](https://github.com/trustgrid/trustgrid-infra-as-code/tree/main/azure/terraform) - This is a collection of Terraform modules that can be used to deploy Trustgrid nodes in Azure. It includes modules for deploying the virtual machines as well as creating IAM roles and networking resources.

## High Availability
Trustgrid supports two methods for supporting high availability networking connectivity via clustered Trustgrid nodes in Azure. These methods can be used together or independently.

-------------------------
| Method | Description | Common Use Cases |
| --- | --- | --- |
| [Route failover]({{<relref "route-failover">}}) | Publishes routes to the Azure route table associated an interface or specified route tables. Automatically adjusts the `Next hop IP address` to point to the active node. | <ul><li>Environments with only a few route tables that need to be adjusted</li></ul>|
| [IP failover]({{<relref "ip-failover">}})| Assigns a floating IP address to the interface of the active Trustgrid node. | <ul><li>Environments with many route tables.<sup>1</sup></li><li>Environments using Azure Virtual WAN.</li><li>Using [connectors]({{<relref "/docs/nodes/shared/connectors/">}}) on the Azure cluster.</li></ul>|
-------------------------
<sup>1</sup> Route based failover requires the nodes to have correct Azure permission to modify each route table. Environments with many route tables would have to grant permission to the containers (resource groups, subscriptions) for each table which makes maintaining least-privilege access difficult.
