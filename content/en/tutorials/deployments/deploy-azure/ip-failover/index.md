---
title: "Azure IP Failover"
linkTitle: "IP Failover"
description: "IP Failover for Azure appliances"
---
## How it works
Trustgrid provides the ability for a floating IP address to be assigned to a Trustgrid cluster. When failover occurs, the floating IP address will automatically be assigned to the active appliance in the cluster. This feature works by creating a new Azure IP Configuration and associating it with the interface of the active cluster node.  

### Graceful Failover
This describes the process of a graceful failover when the active member of the node is changed but both members are online and working normally. 
1. The Trustgrid appliance that is relinquishing the active role will remove the Azure IP Configuration from the interface.
1. The Trustgrid appliance that is gaining the active role will create a new Azure IP Configuration and associate it with the interface.

### Ungraceful Failover
This describes the process of an ungraceful failover when [cluster health conditions]({{<relref "/docs/clusters#cluster-member-health-conditions">}}) prevent the active member of the node from functioning. 
1. After the specified [Cluster Timeout]({{<relref "/docs/clusters#cluster-timeout">}}) period has elapsed the Trustgrid appliance that is taking the active role will remove the Azure IP Configuration from the interface on the prior active node.
1. The now active Trustgrid appliance will create a new Azure IP Configuration and associate its own interface.

## Requirements for HA IP Failover
- Permissions to create and manage the Azure IP Configuration
- An unused private IP address in the subnet of the interface where the Azure IP Configuration is defined

### Permissions Required for Cluster IP Failover
Below details how to create a custom role definition with the minimum required permissions, that can be assigned to the Trustgrid appliance via the Azure Portal. Alternatively, you can use the Build-In Role "Network Contributor" role.

The Trustgrid appliances need permissions to:
- Create a new IP Configuration for themselves. This requires permissions to the resource group where the Trustgrid appliances and the attached subnet are deployed.
- Delete an IP Configuration for their cluster peer.
- Associate the IP Configuration with their interface.
- Associate the IP Configuration with any Security Group or Application Security Groups that are associated with the interface/virtual machine.

> The assignableScopes section will need to be modified to represent the subscription or resource group where the Trustgrid appliances and their attached subnets are deployed.
{{<highlight json>}}
{
    "properties": {
        "roleName": "tg-cluster-ip-failover",
        "description": "Manage Trustgrid Cluster IP Failover",
        "assignableScopes": [
        ],
        "permissions": [
            {
                "actions": [
                    "Microsoft.Network/networkInterfaces/read",
                    "Microsoft.Network/networkInterfaces/write",
                    "Microsoft.Network/networkInterfaces/ipconfigurations/read",
                    "Microsoft.Network/networkInterfaces/ipconfigurations/join/action",
                    "Microsoft.Network/networkSecurityGroups/join/action",
                    "Microsoft.Network/virtualNetworks/read",
                    "Microsoft.Network/virtualNetworks/subnets/read",
                    "Microsoft.Compute/virtualMachines/read",
                    "Microsoft.Network/virtualNetworks/subnets/join/action",
                    "Microsoft.Network/applicationSecurityGroups/joinIpConfiguration/action"
                ],
                "notActions": [],
                "dataActions": [],
                "notDataActions": []
            }
        ]
    }
}
{{</highlight>}}

#### Create and Assign Custom Role via Azure Portal

A custom role needs to be created in the Azure subscription that allows the Trustgrid nodes to update the IP Configuration when failover occurs

##### Create the custom role
1. In the Azure portal search for “Subscriptions” and select the Subscriptions service
1. Select the subscription that contains the Trustgrid VMs 
1. Select “Access control (IAM),” then click “+Add”, then “Add custom role” {{<tgimg src="../add-custom-role.png" alt="Add Custom Role" width="80%">}}
1. Save the JSON above to a file named `azure-custom-role-sample.json`.
1. Select “Start from JSON” and from the file selector, select the downloaded json file. {{<tgimg src="../create-custom-role.png" alt="Create Custom Role" width="80%">}}
1. Optionally, update the role name to meet your internal naming conventions.
1. Click **Next**. 
1. On the Permissions page you will see the permissions that will be granted. Click **Next** again.
1. On the Assignable Scopes page click +Add Assignable Scope
    1. From the Type select Resource Group
    1. From the Subscription, select the subscription containing your VMs.
    1. From the Select pane on the right search for and select the Resource Group containing you VM’s {{<tgimg src="../select-group.png" alt="Select Resource Group" width="90%">}}
    1. Click **Select** and then **Next**.
    1. Repeat for the resource group containing the VMs' virtual network (if different than the VMs).  Alternatively, you can set the assignable scopes to the entire subscription that contain the VMs and virtual network.
1. On the JSON page, click the **Next** button.
1. Click Review + Create, then click Create.
##### Assign the custom role to your Trustgrid VM’s system-assigned
1. In the Azure portal search for Resource Groups and select the service
1. Select your target Resource Group (start with the group containing your Trustgrid VMs)
1. Select the Access Control (IAM) panel, then click +Add, then “Add role assignment” {{<tgimg src="../add-role.png" alt="Add Role" width="80%">}}
1. Search for and select the desired role and click Next {{<tgimg src="../role-list.png" alt="Role List" width="80%">}}
1. Under “Assign access to” select “Managed Identity” then click +Select members {{<tgimg src="../members.png" alt="Select Members" width="80%">}}
1. From the Managed Identity dropdown select Virtual Machine
1. Select the identity for your first Trustgrid VM {{<tgimg src="../select-vms.png" alt="Select VMs" width="80%">}}
1. Click select. 
1. Click +Select members again and repeat with your second Trustgrid VM
1. Click “Review + Assign” then “Review + Assign” a second time
1. If the VMs' virtual network is in a different resource group, repeat the above steps for the virtual network resource group.

{{<alert color="info">}} Azure permission changes can take a few minutes to go into effect and a reboot of the VMs is required to pickup.{{</alert>}}
