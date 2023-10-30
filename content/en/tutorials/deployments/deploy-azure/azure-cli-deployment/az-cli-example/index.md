---
title: "Example Walkthrough  of Azure CLI (az) Deployment"
linkTitle: "AZ CLI Example Walkthrough"
description: "Shows how to create all prerequisite resources and the output of running the deployment commands"
---
 
 Below walks through creating a brand new resource group, virtual network, subnets and associated resources needed to deploy a Trustgrid appliance using Azure CLI commands.  This is rarely appropriate for production environments but demonstrates the process.

 It then walks through all the steps listed on the [AZ CLI Deployment]({{<relref "../">}}) and shows the output of running the commands.

{{<alert color="info">}}The output below may be modified to obscure unique ids for security purposes and truncated due to the verbosity of some response.{{</alert>}}

 ## Create Resource Group

First you'll need to decide on the [location to deploy]({{<relref "/tutorials/deployments/deploy-azure#supported-regions">}})

 ```shell 
sales@Azure:~$ export location="eastus"
sales@Azure:~$ export resourceGroup="DocsExample"
sales@Azure:~$ az group create --resource-group $resourceGroup --location $location
{
  "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample",
  "location": "eastus",
  "managedBy": null,
  "name": "DocsExample",
  "properties": {
    "provisioningState": "Succeeded"
  },
  "tags": null,
  "type": "Microsoft.Resources/resourceGroups"
}
 ```

 ## Create Virtual Network

 Create a virtual network with two subnets - one for the outside interface and one for the inside interface of the appliance. In this example we are using the 192.168.0.0/16 address space but you can customize as needed. 


```bash
sales@Azure:~$ export vNetName="TrustgridvNet"
export outsideSubnet="tg-outside"
export insideSubnet="tg-inside"
sales@Azure:~$ az network vnet create --name $vNetName --resource-group $resourceGroup --location $location --address-prefix 192.168.0.0/16
{
  "newVNet": {
    "addressSpace": {
      "addressPrefixes": [
        "192.168.0.0/16"
      ]
    },
    "enableDdosProtection": false,
    "etag": "W/\"1#######-####-####-####-#########\"",
    "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/virtualNetworks/TrustgridvNet",
    "location": "eastus",
    "name": "TrustgridvNet",
    "provisioningState": "Succeeded",
    "resourceGroup": "DocsExample",
    "resourceGuid": "#######-####-####-####-#########",
    "subnets": [],
    "type": "Microsoft.Network/virtualNetworks",
    "virtualNetworkPeerings": []
  }

}
sales@Azure:~$ az network vnet subnet create --name $outsideSubnet --resource-group $resourceGroup --vnet-name $vNetName --address-prefix 192.168.1.0/24
{
  "addressPrefix": "192.168.1.0/24",
  "delegations": [],
  "etag": "W/\"#######-####-####-####-#########\"",
  "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/virtualNetworks/TrustgridvNet/subnets/tg-outside",
  "name": "tg-outside",
  "privateEndpointNetworkPolicies": "Disabled",
  "privateLinkServiceNetworkPolicies": "Enabled",
  "provisioningState": "Succeeded",
  "resourceGroup": "DocsExample",
  "type": "Microsoft.Network/virtualNetworks/subnets"
}
sales@Azure:~$ az network vnet subnet create --name $insideSubnet --resource-group $resourceGroup --vnet-name $vNetName --address-prefix 192.168.2.0/24
{
  "addressPrefix": "192.168.2.0/24",
  "delegations": [],
  "etag": "W/\"468e9585-daa2-4908-9484-9708e4cac8d4\"",
  "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/virtualNetworks/TrustgridvNet/subnets/tg-inside",
  "name": "tg-inside",
  "privateEndpointNetworkPolicies": "Disabled",
  "privateLinkServiceNetworkPolicies": "Enabled",
  "provisioningState": "Succeeded",
  "resourceGroup": "DocsExample",
  "type": "Microsoft.Network/virtualNetworks/subnets"
}
```

Here are the actual commands again without values or output to make it easier to copy and paste for your use:

```bash
export vNetName=""
export outsideSubnet=""
export insideSubnet=""
az network vnet create --name $vNetName --resource-group $resourceGroup --location $location --address-prefix 192.168.0.0/16
az network vnet subnet create --name $outsideSubnet --resource-group $resourceGroup --vnet-name $vNetName --address-prefix 192.168.1.0/24
az network vnet subnet create --name $insideSubnet --resource-group $resourceGroup --vnet-name $vNetName --address-prefix 192.168.2.0/24
```

### Optional - Create Route Table for Inside Subnet
When Trustgrid appliances are deployed in an HA cluster the inside interface will need a route

## Create Remaining Resources and VM
At this point we have all the prerequisites to proceed with the steps [outlined above to deploy the Trustgrid appliance]({{<relref "..#deploy-azure-vm-as-trustgrid-appliance">}}). Below shows the commands being run with example output. 

### Declare Variables
First, declare all the variables and capture the Image ID:
```bash
export location="eastus"
export resourceGroup="DocsExample"
export vNetName="TrustgridvNet"
export outsideSubnet="tg-outside"
export insideSubnet="tg-inside"
export name="docs-node"
export size="Standard_B2s"
export osDiskSize=30
```
### Capture Latest Trustgrid Image ID
```bash
export imageID=$(az sig image-version list-community \
  --public-gallery-name trustgrid-45680719-9aa7-43b9-a376-dc03bcfdb0ac \
  --gallery-image-definition trustgrid-node-2204-prod \
  --location $location \
  --output json 2>/dev/null | jq -r 'sort_by(.name)| reverse | .[0].uniqueId')
```

### Prepare SSH Key Resource
Using the steps defined in the [Prepare SSH Key](#prepare-ssh-key-in-azure) we will create a new key pair. If you have an existing public key you want to use the output and commands will be different. 

 ```bash
sales@Azure:~$ export sshKeyName="myNewSSHKey"
 az sshkey create --name $sshKeyName --resource-group $resourceGroup --location $location
No public key is provided. A key pair is being generated for you.
Private key is saved to "/home/sales/.ssh/1698355223_61326".
Public key is saved to "/home/sales/.ssh/1698355223_61326.pub".
{
  "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DOCSEXAMPLE/providers/Microsoft.Compute/sshPublicKeys/myNewSSHKey",
  "location": "eastus",
  "name": "myNewSSHKey",
  "publicKey": "ssh-rsa ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC6fak3eFAKEKEY8uRMxw+GfkT4xFGkUlFakeyDTNeM59C3z7g1cTtABCDEF12345678910ABCDksVaNfFakeKeyInHere+UoXUQtDZghF38I6t7S58xvZX7Gdr+W4qvlhLP7YYQXdMwFakEyXoS+ZrmdFakeyZZFakeKeyi2V8U1tzlvF5M= generated-by-azure",
  "resourceGroup": "DOCSEXAMPLE",
  "tags": null,
  "type": null
}
 ```


### Create Network Interfaces
Using the steps in [Create Network Interfaces](#create-network-interfaces) above we create a Public IP and two NICs and associated Network Security Groups.

```bash
sales@Azure:~$ az network public-ip create --name $name-pubIP \
  --resource-group $resourceGroup --location $location \
  --sku "Standard" --allocation-method "Static"
[Coming breaking change] In the coming release, the default behavior will be changed as follows when sku is Standard and zone is not provided: For zonal regions, you will get a zone-redundant IP indicated by zones:["1","2","3"]; For non-zonal regions, you will get a non zone-redundant IP indicated by zones:null.
{
  "publicIp": {
    "ddosSettings": {
      "protectionMode": "VirtualNetworkInherited"
    },
    "etag": "W/\"d976ed14-bf3d-4f1a-964d-9ff7dd0783b4\"",
    "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/publicIPAddresses/docs-node-pubIP",
    "idleTimeoutInMinutes": 4,
    "ipAddress": "20.168.241.169",
    "ipTags": [],
    "location": "eastus",
    "name": "docs-node-pubIP",
    "provisioningState": "Succeeded",
    "publicIPAddressVersion": "IPv4",
    "publicIPAllocationMethod": "Static",
    "resourceGroup": "DocsExample",
    "resourceGuid": "73fc5b1b-bc3e-4670-9af2-65545bb381b6",
    "sku": {
      "name": "Standard",
      "tier": "Regional"
    },
    "type": "Microsoft.Network/publicIPAddresses"
  }
}

sales@Azure:~$ 
sales@Azure:~$ az network nsg create --name $name-outside \
  --resource-group $resourceGroup --location $location
{
  "NewNSG": {
    "defaultSecurityRules": [
      {
        "access": "Allow",
        "description": "Allow inbound traffic from all VMs in VNET",
        "destinationAddressPrefix": "VirtualNetwork",
...Truncated....
        "sourceAddressPrefixes": [],
        "sourcePortRange": "*",
        "sourcePortRanges": [],
        "type": "Microsoft.Network/networkSecurityGroups/defaultSecurityRules"
      }
    ],
    "etag": "W/\"57d825f4-e5eb-407f-8eec-ad9f3be06486\"",
    "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/networkSecurityGroups/docs-node-outside",
    "location": "eastus",
    "name": "docs-node-outside",
    "provisioningState": "Succeeded",
    "resourceGroup": "DocsExample",
    "resourceGuid": "3796f81b-d50c-4d43-89ed-f41da1523fd8",
    "securityRules": [],
    "type": "Microsoft.Network/networkSecurityGroups"
  }
}
sales@Azure:~$ az network nsg rule create --name "TGControlPlaneTCP" \
  --nsg-name $name-outside \
  --resource-group $resourceGroup \
  --priority 1000 \
  --access 'Allow' \
  --description 'TCP ports required Trustgrid Control Plane networks' \
  --destination-address-prefixes '35.171.100.16/28' '34.223.12.192/28' \
  --destination-port-ranges 443 8443 \
  --direction 'Outbound' \
  --protocol 'Tcp' 
{
  "access": "Allow",
  "description": "TCP ports required Trustgrid Control Plane networks",
  "destinationAddressPrefixes": [
    "35.171.100.16/28",
    "34.223.12.192/28"
  ],
  "destinationPortRanges": [
    "443",
    "8443"
  ],
  "direction": "Outbound",
  "etag": "W/\"a26f3973-ed39-4990-a87e-657815c43721\"",
  "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/networkSecurityGroups/docs-node-outside/securityRules/TGControlPlaneTCP",
  "name": "TGControlPlaneTCP",
  "priority": 1000,
  "protocol": "Tcp",
  "provisioningState": "Succeeded",
  "resourceGroup": "DocsExample",
  "sourceAddressPrefix": "*",
  "sourceAddressPrefixes": [],
  "sourcePortRange": "*",
  "sourcePortRanges": [],
  "type": "Microsoft.Network/networkSecurityGroups/securityRules"
}
sales@Azure:~$
sales@Azure:~$ az network nsg create --name $name-inside \
  --resource-group $resourceGroup --location $location
{
  "NewNSG": {
    "defaultSecurityRules": [
      {
        "access": "Allow",
        "description": "Allow inbound traffic from all VMs in VNET",
        "destinationAddressPrefix": "VirtualNetwork",
...truncated...
        "type": "Microsoft.Network/networkSecurityGroups/defaultSecurityRules"
      }
    ],
    "etag": "W/\"4d057391-2af0-468d-87b3-482eaafd6ef9\"",
    "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/networkSecurityGroups/docs-node-inside",
    "location": "eastus",
    "name": "docs-node-inside",
    "provisioningState": "Succeeded",
    "resourceGroup": "DocsExample",
    "resourceGuid": "a8f23fcc-6460-4642-aa12-ac5f54a75108",
    "securityRules": [],
    "type": "Microsoft.Network/networkSecurityGroups"
  }
}
sales@Azure:~$ az network nic create --resource-group $resourceGroup \
 --location $location --name $name-outside \
  --vnet-name $vNetName --subnet $outsideSubnet \
  --accelerated-networking false --public-ip-address $name-pubIP \
  --network-security-group $name-outside
{
  "NewNIC": {
    "auxiliaryMode": "None",
    "auxiliarySku": "None",
    "disableTcpStateTracking": false,
    "dnsSettings": {
      "appliedDnsServers": [],
      "dnsServers": [],
      "internalDomainNameSuffix": "w0n1hkdkbiyu1mglxakd4jdtth.bx.internal.cloudapp.net"
    },
    "enableAcceleratedNetworking": false,
    "enableIPForwarding": false,
    "etag": "W/\"3ddba8ac-99b2-41a8-8ca3-47138ba69b58\"",
    "hostedWorkloads": [],
    "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/networkInterfaces/docs-node-outside",
    "ipConfigurations": [
      {
        "etag": "W/\"3ddba8ac-99b2-41a8-8ca3-47138ba69b58\"",
        "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/networkInterfaces/docs-node-outside/ipConfigurations/ipconfig1",
        "name": "ipconfig1",
        "primary": true,
        "privateIPAddress": "192.168.1.4",
        "privateIPAddressVersion": "IPv4",
        "privateIPAllocationMethod": "Dynamic",
        "provisioningState": "Succeeded",
        "publicIPAddress": {
          "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/publicIPAddresses/docs-node-pubIP",
          "resourceGroup": "DocsExample"
        },
        "resourceGroup": "DocsExample",
        "subnet": {
          "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/virtualNetworks/TrustgridvNet/subnets/tg-outside",
          "resourceGroup": "DocsExample"
        },
        "type": "Microsoft.Network/networkInterfaces/ipConfigurations"
      }
    ],
    "location": "eastus",
    "name": "docs-node-outside",
    "networkSecurityGroup": {
      "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/networkSecurityGroups/docs-node-outside",
      "resourceGroup": "DocsExample"
    },
    "nicType": "Standard",
    "provisioningState": "Succeeded",
    "resourceGroup": "DocsExample",
    "resourceGuid": "e863aac3-a69c-489d-9b1e-64bf61028f52",
    "tapConfigurations": [],
    "type": "Microsoft.Network/networkInterfaces",
    "vnetEncryptionSupported": false
  }
}
sales@Azure:~$ 
sales@Azure:~$ az network nic create --resource-group $resourceGroup \
  --location $location --name $name-inside \
  --vnet-name $vNetName --subnet $insideSubnet \
  --accelerated-networking false --ip-forwarding true \
  --network-security-group $name-inside
{
  "NewNIC": {
    "auxiliaryMode": "None",
    "auxiliarySku": "None",
    "disableTcpStateTracking": false,
    "dnsSettings": {
      "appliedDnsServers": [],
      "dnsServers": [],
      "internalDomainNameSuffix": "w0n1hkdkbiyu1mglxakd4jdtth.bx.internal.cloudapp.net"
    },
    "enableAcceleratedNetworking": false,
    "enableIPForwarding": true,
    "etag": "W/\"7918658f-76fa-4462-ab5e-bbbabf387016\"",
    "hostedWorkloads": [],
    "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/networkInterfaces/docs-node-inside",
    "ipConfigurations": [
      {
        "etag": "W/\"7918658f-76fa-4462-ab5e-bbbabf387016\"",
        "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/networkInterfaces/docs-node-inside/ipConfigurations/ipconfig1",
        "name": "ipconfig1",
        "primary": true,
        "privateIPAddress": "192.168.2.4",
        "privateIPAddressVersion": "IPv4",
        "privateIPAllocationMethod": "Dynamic",
        "provisioningState": "Succeeded",
        "resourceGroup": "DocsExample",
        "subnet": {
          "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/virtualNetworks/TrustgridvNet/subnets/tg-inside",
          "resourceGroup": "DocsExample"
        },
        "type": "Microsoft.Network/networkInterfaces/ipConfigurations"
      }
    ],
    "location": "eastus",
    "name": "docs-node-inside",
    "networkSecurityGroup": {
      "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/networkSecurityGroups/docs-node-inside",
      "resourceGroup": "DocsExample"
    },
    "nicType": "Standard",
    "provisioningState": "Succeeded",
    "resourceGroup": "DocsExample",
    "resourceGuid": "08488814-0337-4742-9751-fd125ea1acb3",
    "tapConfigurations": [],
    "type": "Microsoft.Network/networkInterfaces",
    "vnetEncryptionSupported": false
  }
}

```

### Create Trustgrid Appliance VM

Finally, we can run the command to [create the azure VM](#create-trustgrid-appliance-vm)


```bash
sales@Azure:~$ az vm create \
  --resource-group $resourceGroup \
  --name $name \
  --image $imageID \
  --accept-term \
  --admin-user 'ubuntu' \
  --assign-identity '[system]' \
  --authentication-type 'ssh' \
  --nics $name-outside $name-inside \
  --nic-delete-option 'Delete' \
  --os-disk-size-gb $osDiskSize \
  --os-disk-delete-option 'Delete' \
  --size $size \
  --ssh-key-name $sshKeyName 
No access was given yet to the 'docs-node', because '--scope' was not provided. You should setup by creating a role assignment, e.g. 'az role assignment create --assignee <principal-id> --role contributor -g DocsExample' would let it access the current resource group. To get the pricipal id, run 'az vm show -g DocsExample -n docs-node --query "identity.principalId" -otsv'
{
  "fqdns": "",
  "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Compute/virtualMachines/docs-node",
  "identity": {
    "systemAssignedIdentity": "27f3a3a1-a7f2-462f-a697-844eb9f5c249",
    "userAssignedIdentities": {}
  },
  "location": "eastus",
  "macAddress": "00-0D-3A-9D-A3-2D,00-0D-3A-9D-A2-32",
  "powerState": "VM running",
  "privateIpAddress": "192.168.1.4,192.168.2.4",
  "publicIpAddress": "20.168.241.169",
  "resourceGroup": "DocsExample",
  "zones": ""
}
sales@Azure:~$ az vm boot-diagnostics enable --name $name --resource-group $resourceGroup 
sales@Azure:~$
```
{{<alert color="info">}} The warning that "No access was given yet" can be ignored.{{</alert>}}


## Additional Steps for HA Clusters

For high availability clusters, additional steps would be needed.

First, you'd need to create a [second VM with a new name](../#create-additional-vm-appliance). The output should be similar to the above.

### Create Role for Route Management
The steps below will show creating the custom role required for the VMs to be able to modify the route tables in their resource group. 

```bash
sales@Azure:~$ curl -o tg-route-role.json https://raw.githubusercontent.com/trustgrid/trustgrid-infra-as-code/main/azure/resources/cluster-role-template/tg-route-role.json
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1007  100  1007    0     0   7407      0 --:--:-- --:--:-- --:--:--  7459
sales@Azure:~$ sed -i "s/REPLACE/$(az account show --query id | tr -d '\"')/g" tg-route-role.json
sales@Azure:~$az role definition create --role-definition @tg-route-role.json
Readonly attribute type will be ignored in class <class 'azure.mgmt.authorization.v2022_05_01_preview.models._models_py3.RoleDefinition'>
{
  "assignableScopes": [
    "/subscriptions/#######-####-####-####-#########"
  ],
  "createdBy": null,
  "createdOn": "2023-10-30T18:29:30.705079+00:00",
  "description": "Allows clustered Trustgrid nodes to perform HA routing actions",
  "id": "/subscriptions/#######-####-####-####-#########/providers/Microsoft.Authorization/roleDefinitions/a06d4811-1d33-42ef-b7dc-c92e038246f6",
  "name": "a06d4811-1d33-42ef-b7dc-c92e038246f6",
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
      "condition": null,
      "conditionVersion": null,
      "dataActions": [],
      "notActions": [],
      "notDataActions": []
    }
  ],
  "roleName": "Trustgrid HA Route Role",
  "roleType": "CustomRole",
  "type": "Microsoft.Authorization/roleDefinitions",
  "updatedBy": "bcc767f1-e232-4cdf-8144-5dcc3d2ae88b",
  "updatedOn": "2023-10-30T18:29:30.705079+00:00"
}
```

### Assign Role for Route Management
After you've created the above custom role it needs to be assigned to the virtual machines [managed system identities](https://learn.microsoft.com/en-us/entra/identity/managed-identities-azure-resources/overview#managed-identity-types). 

For this example we will assume the role was created with the name "Trustgrid HA Route Role" and the VMs are named `docs-node` and `docs-node2`. We will export a new variable for each and also make sure we have variables for the resource group and subscription values.
```bash
sales@Azure:~$ export vm1="docs-node"
sales@Azure:~$ export vm2="docs-node2"
sales@Azure:~$ export resourceGroup="DocsExample"
sales@Azure:~$ export subscription=$(az account show --query id -o tsv)
sales@Azure:~$ export vm1ID=$(az vm show --name $vm1 --resource-group $resourceGroup --query 'identity.principalId' -o tsv)
sales@Azure:~$ export vm2ID=$(az vm show --name $vm2 --resource-group $resourceGroup --query 'identity.principalId' -o tsv)
sales@Azure:~$ az role assignment create --role "Trustgrid HA Route Role" --assignee $vm1ID --scope "/subscriptions/$subscription/resourceGroups/$resourceGroup"
{
  "condition": null,
  "conditionVersion": null,
  "createdBy": null,
  "createdOn": "2023-10-30T18:56:44.664663+00:00",
  "delegatedManagedIdentityResourceId": null,
  "description": null,
  "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Authorization/roleAssignments/0799638c-57ce-4c3b-8aec-d20de77b449d",
  "name": "0799638c-57ce-4c3b-8aec-d20de77b449d",
  "principalId": "27f3a3a1-a7f2-462f-a697-844eb9f5c249",
  "principalType": "ServicePrincipal",
  "resourceGroup": "DocsExample",
  "roleDefinitionId": "/subscriptions/#######-####-####-####-#########/providers/Microsoft.Authorization/roleDefinitions/a06d4811-1d33-42ef-b7dc-c92e038246f6",
  "scope": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample",
  "type": "Microsoft.Authorization/roleAssignments",
  "updatedBy": "bcc767f1-e232-4cdf-8144-5dcc3d2ae88b",
  "updatedOn": "2023-10-30T18:56:45.306683+00:00"
}
sales@Azure:~$ az role assignment create --role "Trustgrid HA Route Role" --assignee $vm2ID --scope "/subscriptions/$subscription/resourceGroups/$resourceGroup"
{
  "condition": null,
  "conditionVersion": null,
  "createdBy": null,
  "createdOn": "2023-10-30T18:56:44.664663+00:00",
  "delegatedManagedIdentityResourceId": null,
  "description": null,
  "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Authorization/roleAssignments/0799638c-57ce-4c3b-8aec-d20de77b449d",
  "name": "0799638c-57ce-4c3b-8aec-d20de77b449d",
  "principalId": "27f3a3a1-a7f2-462f-a697-844eb9f5c249",
  "principalType": "ServicePrincipal",
  "resourceGroup": "DocsExample",
  "roleDefinitionId": "/subscriptions/#######-####-####-####-#########/providers/Microsoft.Authorization/roleDefinitions/a06d4811-1d33-42ef-b7dc-c92e038246f6",
  "scope": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample",
  "type": "Microsoft.Authorization/roleAssignments",
  "updatedBy": "bcc767f1-e232-4cdf-8144-5dcc3d2ae88b",
  "updatedOn": "2023-10-30T18:56:45.306683+00:00"
}
```

az role assignment create --role $roleName --assignee $vm1ID --scope "/subscriptions/$subscription/resourceGroups/$resourceGroup"
