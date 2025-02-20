---
title: "Azure Route Failover"
linkTitle: "Route Failover"
description: "Route failover for Azure appliances"
---
## How it works
Trustgrid provides the ability for a cluster of Azure based appliances to published routes to one or more Azure route table. When failover occurs, the Trustgrid appliance will automatically update the next hop IP address in the route table to the active appliance in the cluster.

### Graceful Failover
This describes the process of a graceful failover when the active member of the node is changed but both members are online and working normally.
1. The Trustgrid appliance that is relinquishing the active role will remove the Azure Route Table Entry(ies) from the route table(s).
1. The Trustgrid appliance that is gaining the active role will create a new Azure Route Table Entry(ies) and associate it with the route table(s).

### Ungraceful Failover
This describes the process of an ungraceful failover when [cluster health conditions]({{<relref "/docs/clusters#cluster-member-health-conditions">}}) prevent the active member of the node from functioning.
1. After the specified [Cluster Timeout]({{<relref "/docs/clusters#cluster-timeout">}}) period has elapsed the Trustgrid appliance that is taking the active role will remove the Azure Route Table Entry(ies) from the route table(s).
1. The now active Trustgrid appliance will create a new Azure Route Table Entry and associate its own interface IP.

## Requirements for HA Route Failover
- [One or more route tables](#azure-routing-table).  By default, the node attempts to use the route table associated with the subnet of the interface where the Azure Route Table Entry is defined.
- [Trustgrid Nodes need Azure permissions to modify route tables](#permissions-required-for-cluster-route-failover)

### Azure Routing Table

An Azure routing table resource needs to be associated with the LAN interface's subnet.

#### View LAN Subnet Routing Table

1. In the Azure Portal search for Virtual Networks and select the service
1. From the list of Virtual Networks select your target Virtual Network
1. From the navigation panel select Subnets
1. Select your inside/private subnet that is attached to the LAN interface of your Trustgrid VMs
1. There should be a route table
{{<tgimg src="azure-route-table.png" alt="Example subnet showing assigned Route Table" width="60%">}}

#### Create Route Table for LAN Subnet

1. If there is no Route Table associated with your LAN/inside/private subnet you will need to add it.
1. In the Azure portal search for Route Tables and select the service
1. Click the +Create button
1. Select the Resource Group that contains your Virtual Network and VMs
1. Select the Region that your VMs are deployed in
1. Give the Route Table a name consistent with your naming conventions
1. (Optional) change the Propagate Gateway routes option. {{<tgimg src="propagate-routes.png" alt="Propagate Routes Option" width="70%">}}
1. Click Review + Create, review then click Review + Create again
1. Repeat the above steps to “View LAN Subnet Routing Table” and change the route table from None to the newly created Route Table.
1. Save the change

### Permissions Required for Cluster Route Failover

Copy this sample json file for use in creating a custom role with the required permissions. See process below.

> The assignableScopes section will need to be modified to represent the subscription or resource group where the Trustgrid appliances and any route tables they will modify reside.

{{<highlight json>}}
{
    "properties": {
        "roleName": "tg-route-table",
        "description": "manage azure route table",
        "assignableScopes": [
        ],
        "permissions": [
            {
                "actions": [
                    "Microsoft.Network/networkWatchers/nextHop/action",
                    "Microsoft.Network/networkInterfaces/effectiveRouteTable/action",
                    "Microsoft.Network/routeTables/routes/delete",
                    "Microsoft.Network/routeTables/routes/write",
                    "Microsoft.Network/routeTables/routes/read",
                    "Microsoft.Network/routeTables/join/action",
                    "Microsoft.Network/routeTables/delete",
                    "Microsoft.Network/routeTables/write",
                    "Microsoft.Network/routeTables/read",
                    "Microsoft.Network/networkInterfaces/read",
                    "Microsoft.Network/virtualNetworks/read",
                    "Microsoft.Compute/virtualMachines/read"
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

A custom role needs to be created in the Azure subscription that allows the Trustgrid nodes to update the route table when failover occurs

##### Create the custom role
1. In the Azure portal search for “Subscriptions” and select the Subscriptions service
1. Select the subscription that contains the Trustgrid VMs 
1. Select “Access control (IAM),” then click “+Add”, then “Add custom role” {{<tgimg src="../add-custom-role.png" alt="Add Custom Role" width="80%">}}
1. Save the JSON above to a file named `azure-custom-role-sample.json`.
1. Select “Start from JSON” and from the file selector, select the downloaded json file. {{<tgimg src="../create-custom-role.png" alt="Create Custom Role" width="80%">}}
1. Optionally, update the role name to meet your internal naming conventions.
1. Click `Next`. 
1. On the Permissions page you will see the permissions that will be granted. Click `Next` again.
1. On the Assignable Scopes page click +Add Assignable Scope
    1. From the Type select Resource Group
    1. From the Subscription, select the subscription containing your VMs and virtual networks.
    1. From the Select pane on the right search for and select the Resource Group containing you VM’s {{<tgimg src="../select-group.png" alt="Select Resource Group" width="90%">}}
    1. Click `Select` and then `Next`.
    1. Repeat for any other resource groups that contain route tables that will be modified by the Trustgrid nodes.  Alternatively, you can set the assignable scopes to the entire subscription(s) that container the VMs and route tables. 
1. On the JSON page, click the `Next` button.
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
1. Repeat the above steps for any other resource groups that contain route tables that will be modified by the Trustgrid nodes.

{{<alert color="info">}} Azure permission changes can take a few minutes to go into effect and a reboot of the VMs is required to pickup.{{</alert>}}

