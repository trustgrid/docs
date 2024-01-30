Nodes are available in two form factors: 
- [**Appliance**]({{<relref "/docs/nodes/appliances" >}}) - This is a combination of a fully managed operating system and Trustgrid software pre-installed. 
- **Agent** - This is a software package that can installed on supported operating systems. Users can install additional software and configure the operating system as needed. 

| Functionality | Agent | Appliance |
|-|-|-|
| Gateway Capabilities | None | Data Plane or ZTNA |
| Operating System | [Multiple Supported OS]({{<relref "/tutorials/agent-deploy#supported-operating-systems">}}), install additional software | Fully managed OS and Trustgrid service, no additional software permitted |
| Updates | Managed by user via OS native tools | OS and Trustgrid updates managed via Trustgrid |
| Interface IP Management | Managed by user via OS native tools | Configurable via Trustgrid portal or [local console]({{<relref "/tutorials/local-console-utility/">}})
| Layer 3 VPN | Single Virtual Network connectivity | Support for multiple Virtual Networks |
| Layer 4 Proxy | Full Support | Full Support |
| Compute | Commands only | Commands, Containers or VMs supported |