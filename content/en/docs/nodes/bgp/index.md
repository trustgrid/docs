---
title: BGP
linkTitle: BGP
weight: 1
description: BPG Server running on Trustgrid Nodes
---

## Summary
Trustgrid nodes can be configured to connect to an external BGP router to advertise and receive routes. This allows Trustgrid networks to be integrated into an existing routing fabric. 

The BGP panel is listed under the Network panel group.

## Configuration 
### Router Configuration
These are the settings used by the local BGP router.

{{<fields>}}
{{<field "Status" >}}  
Enabled/Disabled - Whether the BGP server is running on the node.{{</field >}}
{{<field "ID">}}The IP addressed used to identify the BGP router. This is usually the IP on the interface used for BGP communication.{{</field >}}
{{<field "ASN">}}The Autonomous System Number that identifies the router.{{</field >}}
{{<field "Client">}}Determines if the BGP router is in client mode {{</field>}}
{{</fields>}}

{{<tgimg src="bgp-config.png" caption="BGP router settings" alt="table with BGP router settings" width="40%">}}


### Peer Group Configuration
This page is used to configure BGP peers that this router will establish connections with. All peers associated with the group share the same Export and Import settings.

To add a new group:
1. Click **+Add Peer Group** and provide a name.
1. Click Add. 
{{<tgimg src="add-peer-group.png" caption="Add Peer Group button" width="40%">}}

After your group is created you will need to:
* Add peers - these are the other BGP routers your node will connect to
* Define import policies - these limit what routes are accepted from peers
* Define export policies - these limit what routes are advertised to peers

#### Add Peers

Use the **+Add Peer** button to add a BGP peer to the group. Provide:

{{<fields>}}
{{<field "Name">}}User friendly name for the peer{{</field>}}
{{<field "ASN">}}The Autonomous System Number that identifies the peer identifies itself with{{</field>}}
{{<field "Secret">}} (Optional) Secret passphrase used to authenticate with the peer{{</field>}}
{{<field "IP">}}IP address that the peer can be reached at{{</field>}}
{{</fields>}}

{{<tgimg src="add-peer.png" caption="Example Add Peer Dialog" width="40%">}}

#### Define Import Policies
The BGP server on the node will only accept routes advertised by peers if they match an import policy prefix.

On the Imports panel use the **+Add Import Policy** button to add a new policy.  

{{<fields>}}
{{<field "Name">}}The user friendly name of the policy{{</field>}}
{{<field "Action">}}
* Allow - advertised routes matching the defined prefixes will be added to the local routing table
* Deny - advertised routes matching the defined prefixes prefixes will be rejected
{{</field>}}
{{<field "Description">}}(optional) Additional information describing the policy{{</field>}}
{{</fields>}}

Click on the newly created policy to **+Add Prefix** 

{{<fields>}}
{{<field "Prefix">}}CIDR notation of a network used to match with advertised routes{{</field>}}
{{<field "Exact">}} 
* Yes - Requires the advertised route to match the subnet length of the prefix defined above
* No - Allows routes with the same network but have different prefix lengths to match
{{</field>}}
{{<field "Description">}}(optional) Additional information describing the prefix {{</field>}}
{{</fields>}}

#### Define Export Policies
A BGP export policy controls which routes are advertised and sent to external BGP peers.

One the Exports panel use the **+Add Export Policy** to create a new policy

{{<fields>}}
{{<field "Name">}}The user friendly name of the policy{{</field>}}
{{<field "Cluster">}}(Yes/No) If the node is a member of a cluster this setting will determine if it should only advertise the configured prefixes when it is the active member of the cluster.{{</field>}}
{{<field "Action">}} Determines if matching route prefixes will be advertised or not:
* Allow - Configured prefixes will be advertised to the peer group members
* Deny - Explicitly prevent prefixes from being advertised
{{</field>}}
{{</fields>}}

Click the newly created policy and then click **+Add Prefix**

{{<fields>}}
{{<field "Prefix">}}CIDR notation of a network to be advertised{{</field>}}
{{<field "Description">}}(optional) Additional information describing the prefix {{</field>}}
{{</fields>}}

## Management Tools

### Restart BGP Server
{{<tgimg src="restart-bgp.png" caption="Restart BGP button" width="40%">}}
It is sometimes necessary to completely restart the BGP server to clear any issues or force new settings to go into effect immediately. The **Restart BGP** button is available in the tools section of the BGP panel.

Once issued you should see the confirmation message: Restart request sent.
{{<tgimg src="restart-bgp-output.png" caption="BGP Restart Confirmation" width="30%">}}

### BGP Status
{{<tgimg src="bgp-status.png" caption="BGP Status button" width="40%">}}

The BGP Status tool allows you to see information about the state of BGP peering.

{{<tgimg src="bgp-status-output.png" caption="BGP Status Output" width="80%">}}
{{<fields>}}
{{<field "Status">}}Shows if the peer is actively connected or not{{</field>}}
{{<field "Connected">}}If connected, shows how long the peering has been established{{</field>}}
{{<field "ASN">}}The Autonomous System Number that identifies the router{{</field>}}
{{<field "Peer IP">}}The IP address of the peer router{{</field>}}
{{<field "Advertised Routes">}}Lists the routes being advertised to the peer router{{</field>}}
{{<field "Recieved Routes">}}Hovering over the view button displays the routes received from the peer and their associated metrics{{</field>}}
{{</fields>}}
