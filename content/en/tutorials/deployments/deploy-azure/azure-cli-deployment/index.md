---
title: "Azure CLI (az) Deployment"
linkTitle: "AZ CLI Deployment"
description: "Details how to deploy a Trustgrid appliance and related resources using the Azure CLI (az) commands."
---


Deploying the Trustgrid appliance requires certain resources to be created in Azure before the appliance VM can be deployed. This guide will walk through deploying those prerequisite resources using the Azure CLI (az) commands and then deploying the Trustgrid appliance VM itself. 

## Prerequisites
In addition to the [Azure Appliance Requirements]({{<relref "/tutorials/deployments/deploy-azure#azure-requirements">}}) you will need:

- [Azure CLI installed](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli) or access to [Azure Cloud Shell](https://learn.microsoft.com/en-us/azure/cloud-shell/get-started?tabs=azurecli) and logged in with `az login`
- SSH public key



## Deploy Azure VM as Trustgrid Node

First, Ensure you are authenticated to Azure and have selected the desired subscript using `az account set --subscription "MySubscription"` and ensure all the [Azure Appliance Requirements]({{<relref "/tutorials/deployments/deploy-azure#azure-requirements">}}) exist before proceeding. 

### Declare Required Variables

{{<alert color="info">}}The below commands are presented in Bash syntax for compatibility with other operating systems. The Azure Cloud Shell can be changed to Bash mode, but if you prefer Powershell you'll need to remove the `export` part of the command and add a `$` to each variable name. e.g. `export location="centralus"` needs to be input as `$location="centralus"` {{</alert>}}

| Variable | Example Bash Command | Description |
|-|-|-|
| location | `export location="centralus"` | <ul><li>Change **centralus** to the desired Azure region for deploying resources</li><li>See the table of [supported regions]({{<relref "/tutorials/deployments/deploy-azure#supported-regions">}}) for locations with the required Trustgrid image published</li> |
| resourceGroup | `export resourceGroup="myResourceGroup"` | Change **myResourceGroup** to the name of your target resource group |
| vNetName | `export vNetName="myVnet"` | Change **myVnet** to the name of your target virtual network where the Trustgrid appliance will be connected |
| outsideSubnet | `export outsideSubnet="outside"` | Change **outside** to the name of your target WAN subnet. This must exist in the virtual network defined by myVnetName |
| insideSubnet | `export insideSubnet="inside"` | Change **inside** to the name of your target LAN subnet. This must exist in the virtual network defined by myVnetName |
| name | `export name="myNode"` | Change **myNode** to the desired name of your Trustgrid appliance virtual machine being created |
| size | `export size="Standard_B2s"` | Change **Standard_B2s** to the desired VM size. See [Instance sizes]({{<relref "/tutorials/deployments/deploy-azure#instance-size">}}) for supported options and sizing recommendation. |
| osDiskSize | `export osDiskSize=30` | Change **30** to the desired size in GB for the OS disk. Minimum recommended is 30GB. |


Here are all the commands together without values so that you can copy (a button will appear if you hover in the top right), edit and paste into your terminal.

```bash
export location=""
export resourceGroup=""
export vNetName=""
export outsideSubnet=""
export insideSubnet=""
export name=""
export size="Standard_B2s"
```

### Capture Latest Trustgrid Image ID

```bash
export imageID=$(az sig image-version list-community --public-gallery-name trustgrid-45680719-9aa7-43b9-a376-dc03bcfdb0ac --gallery-image-definition trustgrid-node-2204-prod --location $location --output json 2>/dev/null | jq -r 'sort_by(.name)| reverse | .[0].uniqueId')
```
The above command sets `imageID` to the latest Trustgrid image ID from the Trustgrid image gallery in the Azure region specified by `location`.

You can use the command `echo $imageID` to see the result. it should look something like the below:
```bash
sales@Azure:~$ echo $imageID
/CommunityGalleries/trustgrid-45680719-9aa7-43b9-a376-dc03bcfdb0ac/Images/trustgrid-node-2204-prod/Versions/2.17.1
```

### Prepare SSH Key in Azure
 After registration SSH is only accessible via the Trustgrid portal, but Azure requires an SSH key be associated with the VM to allow SSH access on creation. 

 There are a few ways to handle this in Azure, but it is important you end up with a variable `sshKeyName` containing the name of an existing SSH key resource in Azure.


 {{< alert color="info" >}}The SSH public keys in examples below are not valid and are for demonstration purposes only.{{< /alert >}}

 #### Reference an Existing Azure SSH key resource
 If you have an existing SSH key resource in Azure you can export its name as a variable:
 ```bash
 export sshKeyName="MyKeyName"
 ```

 #### Create new key and Azure SSH key resource
 The below example creates a new SSH key pair and uses that to create a new Azure SSH key resource associated with the resource group:

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
{{<alert color="warning">}}Note the location where the Private key was saved and ensure you save that key somewhere secure just in case SSH access is required later. Without the Private key, the public key is unusable.{{</alert>}}

Here is the command without output or values
```bash
az sshkey create --name $sshKeyName --resource-group $resourceGroup --location $location
```

#### Create new Azure SSH Key resource with existing public key
If you have an existing SSH key pair already, you can import the public key to create a new Azure SSH key resource:
```bash
sales@Azure:~$ az sshkey create --name $sshKeyName --resource-group $resourceGroup --location $location --public-key "ssh-rsa ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC6fak3eFAKEKEY8uRMxw+GfkT4xFGkUlFakeyDTNeM59C3z7g1cTtABCDEF12345678910ABCDksVaNfFakeKeyInHere+UoXUQtDZghF38I6t7S58xvZX7Gdr+W4qvlhLP7YYQXdMwFakEyXoS+ZrmdFakeyZZFakeKeyi2V8U1tzlvF5M= existing-key"
{
  "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DOCSEXAMPLE/providers/Microsoft.Compute/sshPublicKeys/myNewSSHKey",
  "location": "eastus",
  "name": "myNewSSHKey",
  "publicKey": "ssh-rsa ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC6fak3eFAKEKEY8uRMxw+GfkT4xFGkUlFakeyDTNeM59C3z7g1cTtABCDEF12345678910ABCDksVaNfFakeKeyInHere+UoXUQtDZghF38I6t7S58xvZX7Gdr+W4qvlhLP7YYQXdMwFakEyXoS+ZrmdFakeyZZFakeKeyi2V8U1tzlvF5M= existing-key",
  "resourceGroup": "DOCSEXAMPLE",
  "tags": null,
  "type": null
}
```
Here is the command without output or values
```bash
az sshkey create --name $sshKeyName --resource-group $resourceGroup --location $location --public-key ""
```

### Create Network Interfaces
This step will create the network interfaces for the Trustgrid appliance VM to connect to the subnets. As part of this process we will create a public IP to be attached to the outside interface.

```bash
sales@Azure:~$ az network public-ip create --name $name-pubIP --resource-group $resourceGroup --location $location --sku "Standard" --allocation-method "Static"
[Coming breaking change] In the coming release, the default behavior will be changed as follows when sku is Standard and zone is not provided: For zonal regions, you will get a zone-redundant IP indicated by zones:["1","2","3"]; For non-zonal regions, you will get a non zone-redundant IP indicated by zones:null.
{
  "publicIp": {
    "ddosSettings": {
      "protectionMode": "VirtualNetworkInherited"
    },
    "etag": "W/\"f5e5d94b-c37d-4b48-a7c6-ed28566769f3\"",
    "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/publicIPAddresses/docs-node-pubIP",
    "idleTimeoutInMinutes": 4,
    "ipAddress": "20.25.9.228",
    "ipTags": [],
    "location": "eastus",
    "name": "docs-node-pubIP",
    "provisioningState": "Succeeded",
    "publicIPAddressVersion": "IPv4",
    "publicIPAllocationMethod": "Static",
    "resourceGroup": "DocsExample",
    "resourceGuid": "40e63960-f8d6-4749-92fc-318e4eb4bcde",
    "sku": {
      "name": "Standard",
      "tier": "Regional"
    },
    "type": "Microsoft.Network/publicIPAddresses"
  }
sales@Azure:~$ az network nic create --resource-group $resourceGroup --location $location --name $name-outside --vnet-name $vNetName --subnet $outsideSubnet --accelerated-networking false --public-ip-address $name-pubIP
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
    "etag": "W/\"ad9c7ff6-a706-4387-92df-95ca9a1961e9\"",
    "hostedWorkloads": [],
    "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/networkInterfaces/docs-node-outside",
    "ipConfigurations": [
      {
        "etag": "W/\"ad9c7ff6-a706-4387-92df-95ca9a1961e9\"",
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
    "nicType": "Standard",
    "provisioningState": "Succeeded",
    "resourceGroup": "DocsExample",
    "resourceGuid": "a47697e7-0dfc-4272-8087-fa8f42b4a3d1",
    "tapConfigurations": [],
    "type": "Microsoft.Network/networkInterfaces",
    "vnetEncryptionSupported": false
  }
}
sales@Azure:~$ az network nic create --resource-group $resourceGroup --location $location --name $name-inside --vnet-name $vNetName --subnet $insideSubnet  --accelerated-networking false --ip-forwarding true
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
    "etag": "W/\"0a2b4cad-77a6-41be-aac7-e7396d2af6a3\"",
    "hostedWorkloads": [],
    "id": "/subscriptions/#######-####-####-####-#########/resourceGroups/DocsExample/providers/Microsoft.Network/networkInterfaces/docs-node-inside",
    "ipConfigurations": [
      {
        "etag": "W/\"0a2b4cad-77a6-41be-aac7-e7396d2af6a3\"",
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
    "nicType": "Standard",
    "provisioningState": "Succeeded",
    "resourceGroup": "DocsExample",
    "resourceGuid": "3fe216df-f52c-40fd-90b4-a8eaa9de1a7f",
    "tapConfigurations": [],
    "type": "Microsoft.Network/networkInterfaces",
    "vnetEncryptionSupported": false
  }
}
```

Here are the commands without output
```bash
az network public-ip create --name $name-pubIP --resource-group $resourceGroup --location $location --sku "Standard" --allocation-method "Static"
az network nic create --resource-group $resourceGroup --location $location --name $name-outside --vnet-name $vNetName --subnet $outsideSubnet --accelerated-networking false --public-ip-address $name-pubIP
az network nic create --resource-group $resourceGroup --location $location --name $name-inside --vnet-name $vNetName --subnet $insideSubnet  --accelerated-networking false --ip-forwarding true
```

### Create Trustgrid Appliance VM

Finally we can deploy the actual VM


 ## Example Environment Setup Walk Through
 
 Below walks through creating an brand new resource group, virtual network, subnets and associated resources needed to deploy a Trustgrid appliance using Azure CLI commands.  This is rarely appropriate for production environments but demonstrates the process.

{{<alert color="info">}}The output below may be modified to obscure unique ids for security purposes{{</alert>}}

 ### Create Resource Group

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

 ### Create Virtual Network

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

Here are the actual commands again without values or output

```bash
export vNetName=""
export outsideSubnet=""
export insideSubnet=""
az network vnet create --name $vNetName --resource-group $resourceGroup --location $location --address-prefix 192.168.0.0/16
az network vnet subnet create --name $outsideSubnet --resource-group $resourceGroup --vnet-name $vNetName --address-prefix 192.168.1.0/24
az network vnet subnet create --name $insideSubnet --resource-group $resourceGroup --vnet-name $vNetName --address-prefix 192.168.2.0/24
```