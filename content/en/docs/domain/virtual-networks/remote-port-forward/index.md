---
title: Virtual Network Port Forwards
linkTitle: Port Forwards
type: docs
---
{{% pageinfo %}} Virtual Network Port Forwards allow a virtual IP address to be associated with a [TCP Layer 4 service]({{<ref "docs/nodes/shared/services">}}) on a remote node or cluster without the need to specify a route for that virtual IP address. {{% /pageinfo %}}

{{<tgimg src="remote-port-forward.png" caption="Example Virtual Network Port Forward" width="80%">}}

## Description
As noted above, a Virtual Network Port Forward (or remote port forward) associates a [service]({{<ref "docs/nodes/shared/services">}}) defined on a node or cluster with a virtual IP address within the virtual network without needing to specifically [route]({{<ref "docs/domain/virtual-networks/routes">}}) that IP to the node/cluster.  

Any node attached to the virtual network that receives VPN traffic destined for the configure IP and port will proxy the connection and send the traffic to the target node which will then send the traffic to the IP and port configured for the target service. 

{{<fields>}}
{{<field "Virtual IP Address">}}Virtual IP address assigned to the port forward. {{</field>}}
{{<field "Virtual Port">}} TCP port on which the port forward will listen for connections. {{</fields>}}
{{<field "Destination Node">}}Node or Cluster where the target **service** is defined.{{</field>}}
{{<field "Destination Service">}}Service that traffic for the **virtual IP** and **virtual port** will be forwarded to. Alternately, you can specify an IP and port local to the **destination node**.{{</field>}}
{{</fields>}}
