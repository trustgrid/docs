---
title: "Azure CLI (az) Deployment"
linkTitle: "AZ CLI Deployment"
description: "Details how to deploy a Trustgrid appliance and related resources using the Azure CLI (az) commands."
---


Deploying the Trustgrid appliance requires certain resources to be created in Azure before the appliance VM can be deployed. This guide will walk through deploying those prerequisite resources using the Azure CLI (az) commands and then deploying the Trustgrid appliance VM itself. 

See [AZ CLI Example Walkthrough]({{<relref "./az-cli-example">}}) to see the output of running the below commands, and see how you could create the [Azure Appliance Requirements]({{<relref "/tutorials/deployments/deploy-azure#azure-requirements">}}) with the `az` cli in advance if they do not exist.

## Prerequisites
In addition to the [Azure Appliance Requirements]({{<relref "/tutorials/deployments/deploy-azure#azure-requirements">}}) you will need:

- [Azure CLI installed](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli) or access to [Azure Cloud Shell](https://learn.microsoft.com/en-us/azure/cloud-shell/get-started?tabs=azurecli) and logged in with `az login`
- SSH public key



## Deploy Azure VM as Trustgrid Appliance

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
The above command sets `imageID` to the latest Trustgrid image ID from the Trustgrid image gallery in the Azure region specified by `location`.

You can use the command `echo $imageID` to see the result. it should look something like the below:
```bash
sales@Azure:~$ echo $imageID
/CommunityGalleries/trustgrid-45680719-9aa7-43b9-a376-dc03bcfdb0ac/Images/trustgrid-node-2204-prod/Versions/2.17.1
```

### Prepare SSH Key in Azure
 After registration, SSH is only accessible via the Trustgrid portal, but Azure requires an SSH key to be associated with the VM to allow SSH access on creation. 

 There are a few ways to handle this in Azure, but it is important you end up with a variable `sshKeyName` containing the name of an existing SSH key resource in Azure.


 {{< alert color="info" >}}The SSH public keys in the examples below are not valid and are for demonstration purposes only.{{< /alert >}}

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
This step will create the network interfaces for the Trustgrid appliance VM to connect to the subnets. As part of this process, we will create these additional resources:
- Public IP for the outside interface.
- Network Security Group for the outside interface and an explicit Outbound rule for the [required connectivity to the Trustgrid Control Plane]({{<relref "/help-center/kb/site-requirements#trustgrid-control-plane">}}). Additional rules will need to be added to allow the node to connect to:
  - Outbound rules for data plane gateway IPs and ports if the appliance is acting as an edge/client device
  - Inbound rules if the appliance will be acting as a data plane or ZTNA gateway
- Network Security Group for the inside interface with no additional rules. After deployment, this security group could be extended to allow for required communication to internal resources.

{{<alert color="info">}} Both security groups will be created with the [default security group rules](https://learn.microsoft.com/en-us/azure/virtual-network/network-security-groups-overview#default-security-rules) {{</alert>}}

Here are the commands without output
```bash
az network public-ip create --name $name-pubIP \
  --resource-group $resourceGroup --location $location \
  --sku "Standard" --allocation-method "Static"

az network nsg create --name $name-outside \
  --resource-group $resourceGroup --location $location

az network nsg rule create --name "TGControlPlaneTCP" \
  --nsg-name $name-outside \
  --resource-group $resourceGroup \
  --priority 1000 \
  --access 'Allow' \
  --description 'TCP ports required Trustgrid Control Plane networks' \
  --destination-address-prefixes '35.171.100.16/28' '34.223.12.192/28' \
  --destination-port-ranges 443 8443 \
  --direction 'Outbound' \
  --protocol 'Tcp' 

az network nsg create --name $name-inside \
  --resource-group $resourceGroup --location $location

az network nic create --resource-group $resourceGroup \
 --location $location --name $name-outside \
  --vnet-name $vNetName --subnet $outsideSubnet \
  --accelerated-networking false --public-ip-address $name-pubIP \
  --network-security-group $name-outside
  
az network nic create --resource-group $resourceGroup \
  --location $location --name $name-inside \
  --vnet-name $vNetName --subnet $insideSubnet \
  --accelerated-networking false --ip-forwarding true \
  --network-security-group $name-inside
```
See the example output for all of the above commands in the [AZ CLI Example Walkthrough](./az-cli-example/#create-network-interfaces)

### Create Trustgrid Appliance VM

Finally, we can deploy the actual VM with the two below commands

This command creates the VM:
```bash
az vm create \
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
```

This command enables boot diagnostics so that the serial console can be accessed.
```bash
az vm boot-diagnostics enable --name $name --resource-group $resourceGroup 
```
See the example output in the [AZ CLI Example Walkthrough](./az-cli-example/#create-trustgrid-appliance-vm)

### Get MAC Address

To make [Console login]({{<relref "/tutorials/local-console-utility#logging-in">}}) easier run the below command to get the outside interface's MAC address:

```bash
az network nic show --name $name-outside \
  --resource-group $resourceGroup | jq '.macAddress' | \
   sed 's/-/:/g' | tr 'A-F' 'a-f'
```

Example output:
```bash
sales@Azure:~$ az network nic show --name $name-outside \
  --resource-group $resourceGroup | jq '.macAddress' | \
   sed 's/-/:/g' | tr 'A-F' 'a-f'
"00:0d:3a:9d:a3:2d"
```
### Register the VM

You can now start the [console registration process]({{<relref "/tutorials/local-console-utility/remote-registration">}}) and if you have portal access complete the activation process. 

## Additional Steps for HA Clusters
If deploying an HA cluster there are additional steps required:

### Create Additional VM Appliance
Repeat the steps above to deploy a second VM appliance using different names but in the same resource group and vNet. This will be the secondary node in the HA cluster. 

Assuming you still have the same variables declared from the above deployment you will just need to update the `$name` variable to a new value like shown in the example below.
```
export name="newName"
```

Then you can proceed with creating the [network interface resources](#create-network-interfaces) and then [create the vm appliance](#create-trustgrid-appliance-vm).

### Create Role for Route Management
The steps below create a custom role that has permission to manipulate route tables:

First, download the template JSON file to your environment.
```bash
curl -o tg-route-role.json https://raw.githubusercontent.com/trustgrid/trustgrid-infra-as-code/main/azure/resources/cluster-role-template/tg-route-role.json
```

Use the command below to get your subscription ID:
```bash
az account show --query id -o tsv
```
Use an editor such as `vi` or `nano` to modify the template and replace the placeholders `REPLACE` in the template with your subscription ID. 

Or, you can perform both the above steps in a single command:
```bash
sed -i "s/REPLACE/$(az account show --query id -o tsv)/g" tg-route-role.json
```

Optionally, you can edit the Name and Description fields as well.

Then create the role using:
```bash
az role definition create --role-definition @tg-route-role.json
```

See [AZ CLI Example Walkthrough](./az-cli-example/#create-role-for-route-management) for output.

{{<alert color="info">}} The role created above can only be assigned to resources in the current subscription. If your Azure account topology requires the Trustgrid appliance to be able to manage route tables in other subscriptions you'll need to create the role in those accounts as well and assign (see below) the role accordingly.{{</alert>}}

###  Assign Role for Route Management
Finally, we need to assign this role.  Below are the commands to perform this process without values for the names of the two VMs (vm1 and vm2) for the resource group. 

```bash
export vm1=""
export vm2=""
export resourceGroup=""
export roleName="Trustgrid HA Route Role"
export subscription=$(az account show --query id -o tsv)
export vm1ID=$(az vm show --name $vm1 --resource-group $resourceGroup --query 'identity.principalId' -o tsv)
export vm2ID=$(az vm show --name $vm2 --resource-group $resourceGroup --query 'identity.principalId' -o tsv)
export roleID=$(az role definition list --query "[?roleName=='$roleName'].name | [0]" -o tsv)
az role assignment create --role "Trustgrid HA Route Role" --assignee $vm1ID --scope "/subscriptions/$subscription/resourceGroups/$resourceGroup"
az role assignment create --role "Trustgrid HA Route Role" --assignee $vm2ID --scope "/subscriptions/$subscription/resourceGroups/$resourceGroup"

```
{{<alert color="warning">}} 
The above assumes you didn't change the `Name` settings in the tg-route-role.json file when [creating the role definition]({{<relref "#create-role-for-route-management">}}). If you did you need to update the `roleName` variable before exporting.
{{</alert>}}
See [AZ CLI Example Walkthrough](./az-cli-example/#assign-role-for-route-management) for output.
{{<alert color="info">}} The above commands scope the role to the resource group that contains the node appliance VMs. If your Azure account and resource group topology requires the nodes to be able to update route tables in other resources groups you'll need to repeat the role assignment commands either at a higher level (like subscription) or with a different value for `resourceGroup`. {{</alert>}}



