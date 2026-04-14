---
title: "AWS Route Tables Interface Tool"
linkTitle: "AWS Route Tables"
weight: 55
description: "Query AWS route tables associated with an interface to troubleshoot route failover and destination routing"
tags: ["aws", "networking", "routing", "troubleshooting", "interface-tools"]
categories: ["interface-tools"]
---

{{% pageinfo %}}
The AWS Route Tables interface tool queries AWS for route tables associated with the selected interface and shows the current route definitions. Use it when validating cluster route failover behavior or investigating unexpected routing.
{{% /pageinfo %}}

## Accessing AWS Route Tables

1. Open the target node in the Trustgrid portal.
1. Go to **Network** -> **Interfaces**.
1. In **Interface Tools**, click **AWS Route Tables**.
{{<tgimg src="aws-route-tables-button.png" width="95%" caption="Interface Tools with the AWS Route Tables button highlighted">}}
1. In the service prompt, confirm the interface and click **Execute**.
1. Review the returned route table data in the **AWS Route Tables** dialog.
{{<tgimg src="aws-route-tables-modal.png" width="75%" caption="AWS Route Tables results for the selected interface">}}

{{<alert title="Note" color="info">}}
The **AWS Route Tables** button is shown for AWS-hosted nodes with compatible interface support. If the button is missing on an AWS node, upgrade the device to a version that includes AWS Route Tables support.
{{</alert>}}

## What The Tool Shows

The AWS Route Tables dialog includes a row per route table associated with the selected interface.

{{<fields>}}
{{<field "ID">}}AWS route table identifier (for example, `rtb-...`).{{</field>}}
{{<field "VPC ID">}}VPC containing the route table.{{</field>}}
{{<field "Routes">}}Count of routes currently present in that table.{{</field>}}
{{<field "View Routes">}}Open the route list for a specific route table.{{</field>}}
{{</fields>}}

## Common Troubleshooting Uses

- Confirm route tables are discoverable from the node's selected interface.
- Verify expected destination CIDRs exist before and after failover testing.
- Check that route entries moved to the active node as expected.
- Quickly compare what AWS currently has against intended design.

## Related Pages

- [Node Interfaces]({{<relref "/docs/nodes/appliances/interfaces">}})
