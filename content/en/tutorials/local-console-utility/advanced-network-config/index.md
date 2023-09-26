---
title: Advanced Network Configuration
linkTitle: Advanced Network Config
description: For changing IP, Speed/Duplex and MTU when not possible via control plane
---

{{<alert color="warning">}} The setting below should be changed via the Trustgrid portal [interface configuration]({{<ref "/docs/nodes/interfaces#configuration">}}) if at all possible. These tools should only be used if the correct configuration is preventing successful communication to the control plane.. {{</alert>}}


## Modify Network Settings
If the Trustgrid service is unreachable the normal [Network Configuration]({{<ref "/tutorials/local-console-utility/#changing-a-trustgrid-node-wan-ip-via-trustgrid-console">}}) will be unavailable and it may be necessary to manually edit the network settings on the device to change the IP and DNS configuration. This can be done via the Advanced Network Configuration > Modify Network Settings option. 

The file you are editing is a [netplan configuration file](https://netplan.io/reference). Only the WAN/outside interface should be configured via Netplan as all other interfaces are managed by Trustgrid. The file format is [YAML](https://yaml.org/) so it is important that all sections are indented correctly. Each level is two spaces.

1. From the Advanced Network Configuration select the **Modify Network Settings** option {{<tgimg src="adv-net-conf-menu.png" caption="Modify Network Settings option" width="80%">}}
1. The file will be opened in a basic text editor. Use the arrow keys to navigate to the correct section and edit as needed. Once complete use **TAB** to move to **OK** (or **Cancel** to exit without saving). {{<tgimg src="adv-net-conf-edit.png" caption="Edit the file" width="80%">}}
1. You will be prompted to apply the changes. Select **Yes** to apply the changes. {{<tgimg src="adv-net-conf-apply.png" caption="Apply the changes" width="80%">}}
1. You will then be asked if you want to save the changes permanently. Select **Yes** to save the changes. {{<tgimg src="adv-net-conf-save.png" caption="Save the changes" width="80%">}}
1. The system will apply the changes and wait about 20 seconds for the network to stabilize.  {{<tgimg src="adv-net-conf-stable.png" caption="Wait for the network to stabilize" width="80%">}}
1. After editing the file the change must be applied and then the [Trustgrid service restarted]({{<ref "/tutorials/local-console-utility/troubleshooting#restart-node-service">}}).

### Deleting the runtime file
A temporary file is created when the network settings are changed but not permanently saved. If this file exists you will see an option to delete it from the Advanced Network Configuration menu. Selecting this option will delete the file and the changes will be lost.
{{<tgimg src="adv-net-conf-del-run.png" caption="Delete the Netplan Runtime File" width="80%">}}

## Change Interface Speed/Duplex
By default all interfaces are set to auto-negotiate speed and duplex settings. However, if the other interface is also set for auto negotiate it is common to end up with mismatched settings. (e.g. the port is configured for 100 full and the interface negotiates to 100 half) This can lead to performance problems and interface errors until the correct settings are hard set on the interface.  

1. From the Advanced Network Configuration select the **Change Interface Speed/Duplex** option
1. Select the desired interface from the list. {{<tgimg src="speed-select-int.png" alt="Speed/Duplex Select Interface" width="80%">}}
1. Select the desired speed from the list. {{<tgimg src="speed-select-speed.png" alt="Select Speed of 1000, 100 or 10" width="80%">}}
1. Select the desired duplex from the list. {{<tgimg src="duplex-select.png" alt="Duplex options Full or Half" width="80%">}}
1. Select save to apply the changes.

## Change Interface MTU
By default all interfaces set themselves to the typical MTU size for their environments (1500 for physical, VMware and Hyper-V, 9000 for AWS and Azure). However, if any device between the Trustgrid node and the resources it connects to (both control and data plane) is configured to use a smaller MTU size then this can cause problems. 

{{<alert color="warning">}} To make the below change permanent it will be necessary to update the interface settings in the Trustgrid portal/api. {{</alert>}}

To change the MTU size of the interface:
1. From the Advanced Network Configuration select the **Change Interface MTU** option
1. Select the desired interface from the list. {{<tgimg src="mtu-interface-select.png" alt="MTU Select Interface" width="80%">}}
1. Enter the desired MTU size and select OK to apply the changes. {{<tgimg src="mtu-enter-mtu.png" alt="MTU Enter MTU" width="80%">}}
