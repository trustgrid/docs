---
title: Overview
linkTitle: Overview
description: Agent node statistics and recent flow logs
---

The overview page for agent-based nodes shows several graphs and provides a table of recent flow logs.

{{<tgimg src="agent-overview.png" width="95%" caption="Example Overview panel for an agent">}}

## Graphs

- **Node Performance** - This section shows graphs of CPU, memory, disk, and network usage over time for the agent node.
- **Agent Process Performance** - This section shows the CPU, memory, and disk IO usage of the actual Trustgrid agent process. 
- **Traffic Throughput** - This section shows the total sent and received traffic in megabits per second (Mbps) for all interfaces on the agent.
- **VPN Traffic Throughput** - This section shows the total sent and received traffic specifically over the virtual network tunnel interface in megabits per second (Mbps). 

{{<alert color="warning" title="Note about agent statistics">}} **Agent-based** nodes report their statistics to the Trustgrid control plane periodically which are then processed and available for display. As such there can be a delay of up to 5 minutes before the graphs show a data point.  **Appliance-based** nodes can report their stats in realtime when online and connected to the portal. {{</alert>}}

## Recent Flow Logs
This section gives a quick view of the most recent reported flow logs and provides a link to the [Flow Logs]({{<relref "/help-center/flow-logs">}}) page for more advanced filtering and analysis.