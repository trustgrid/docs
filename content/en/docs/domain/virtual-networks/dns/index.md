---
linkTitle: DNS
Title: "DNS"
weight: 50
---

The DNS feature allows nodes in the virtual network to act as DNS servers for ingress VPN traffic. The node can resolve queries custom defines DNS zones and records, or forward queries to upstream DNS servers.

## DNS Settings
{{<fields>}}
{{<field "Enable/Disable DNS" >}}This field enables or disables the DNS feature on the virtual network.{{</field>}}
{{<field "DNS Server IP">}}This is the IP address that should be used as the resolver. There is no need to create a VPN route for this IP. If the ingress node was the [DNS feature enabled]({{<ref "/docs/nodes/appliances/vpn/dns">}}) it will automatically handle responses for the query.{{</field>}}
{{</fields>}}

{{<tgimg src="vnet-dns-settings.png" caption="DNS Settings" width="60%">}}

After making any changes to the DNS settings you will need to [review and apply changes]({{<ref "/docs/domain/virtual-networks/review-changes">}}) for the updates to take effect.

## DNS Zones
Optionally you can create custom DNS zones and either:
- Configure specific DNS records to resolve to IPs/hosts on the virtual network
- Configure a node as a resolver for this zone. Requests will be forwarded to that node and resolved by its DNS servers or configured upstream servers.

{{<tgimg src="vnet-dns-zones.png" caption="DNS Zone Settings" width="40%">}}


### Adding a DNS Zone
To add a custom DNS zone:
1. Click the "Add Zone" button.
1. Provide a name for the zone (e.g. example.com).
1. Optionally, provide a description for the zone.
1. [Review and apply changes]({{<ref "/docs/domain/virtual-networks/review-changes">}}) to save the zone.

{{<tgimg src="vnet-dns-add-zone.png" caption="Adding a DNS Zone" width="50%">}}


Once a zone has been added, you can configure either a DNS resolver or specific DNS records. 

{{<alert color="warning">}} You cannot have both custom DNS records and a resolver configured for the same zone. {{</alert>}}

### Configuring a DNS Resolver

1. Click on the zone name to enter its configuration page.
1. Under "Resolver", select a node from the dropdown.
1. Click "Save" to save the resolver configuration.
1. [Review and apply changes]({{<ref "/docs/domain/virtual-networks/review-changes">}}) to update the DNS configuration.

### Configuring DNS Records
DNS records allow mapping names to IP addresses for hosts on the virtual network. DNS records have the following fields:

{{<fields>}}
{{<field "Name">}}The hostname being mapped (e.g. www). The zone name will be automatically appended.{{</field>}}
{{<field "Record Type">}}
- A - for records that should resolve to an IP address
- CNAME - for alias/canonical name records
{{</field>}}
{{<field "Value">}}The IP address or hostname being mapped to for A/CNAME records respectively.{{</field>}}
{{<field "TTL">}}Time To Live - how long records may be cached by other resolvers.{{</field>}}
{{</fields>}}

To add a DNS record:

1. Click on the zone name to enter its configuration page.
1. Under "DNS Records", click "Add Record"
1. Provide the required fields {{<tgimg src="vnet-dns-add-record.png" caption="Add DNS Record prompt" width="60%">}}
1. Click save.
1. (Optional) repeat to add more records to the current zone.
1. [Review and apply changes]({{<ref "/docs/domain/virtual-networks/review-changes">}}) to save the record(s).