---
title: Troubleshoot High Tunnel Latency
linkTitle: Latency Troubleshooting
date: 2024-10-21
---
## Active Incident Triage Steps
These steps should be performed as soon as possible **while** the incident is occurring.

{{<alert color="warning">}}Latency issues are frequently temporary and if you do not gather this information at the time of the incident it will prevent deeper analysis.{{</alert>}}

### 1. Enable Hop Monitoring
**Requires permissions to modify the edge node config. If your user is read-only skip to [Step 2](#2-run-mtr-to-reported-gateway).**

Hop monitoring uses TCP SYN packets to continuously monitor the path between an edge node and all it's attached gateways. Follow the steps the [Enable Network Hop Monitoring]({{<relref "/tutorials/gateway-tools/monitoring-network-hops-to-peers#enabling-network-hop-monitoring">}}) to enable hop monitoring.


### 2. Run MTR To Reported Gateway
This step should 

1. Navigate to the [data plane panel]({{<relref "/docs/nodes/appliances/data-plane">}}) edge node (not the gateway) that is reporting high latency.
1. From the list of peers click the peer gateway with which the edge node is experiencing high latency.
1. Click the "MTR" button on the right side of the table in that row. This will open a prompt with the gateway's IP address and port pre-populated and the protocol set to TCP.
1. Change the following settings:
    - Change the DNS field to either "Show IPs and Hostnames" or "Show IPs only". Note: if the node is having DNS health issues showing host names will slow down the response from MTR.
    - Change ping count to 30.  This will give you more data points with a potential to identify the hop in the path introducing problems.
    - Select "Use Report Mode". This will cause the MTR to run in a silent mode, but when it completes the results stay visible in the window.
    - Click Execute. {{<tgimg src="mtr-settings.png" alt="MTR Settings" width="60%" caption="Recommended MTR Settings">}}
1. When the MTR window first opens it will show only the word "Start" followed by the date and time (UTC timezone). This will stay for approximately 30-40 seconds while the data is gathered in the background. This is normal. {{<tgimg src="mtr-start.png" alt="MTR Start" width="60%" caption="MTR Start">}}.
1. After MTR completes testing the results will be displayed and the words "Session Closed" will appear. Capture this data either with a screenshot or by copying and pasting the text into a text file.{{<tgimg src="mtr-results.png" alt="MTR Results" width="80%" caption="MTR Results">}}

{{<alert color="primary" title="What if the MTR Results are blank?">}} MTR and Trustgrid's hop monitoring feature both rely on the node receiving ICMP TTL (time to live) exceeded messages from the various routers (or hops) along the path.
- If you are seeing no responses or only the first and last hop, then it is likely a firewall in front of the node is blocking the ICMP responses. You would need to work with the site technical contact to allow these response back to the node. Allowing this traffic differs depending on the type of firewall but search for "allowing traceroute responses" or "allowing ICMP responses" to find the correct setting is usually helpful. 
- If you are seeing responses up to a certain point and then several missing/blank hops this is because there is no requirement for all routers to respond.  Public clouds like AWS and Azure do not respond once the traffic has entered their networks. 
{{</alert>}}

### 3. Check Latency to Other Gateways
The Data Plane panel shows the tunnel latency to all connected peers.

1. If the reported gateway is a member of a cluster, click the other name of the other cluster member to see if it is impacted as well.  
    - It is common that the latency will be more severe on the active/master member of the cluster because it is under more load with more packets that could be impacted.
    - If high latency is seen to both members of the cluster, then it is likely a problem with the internet path than an issue with the gateway or edge node itself.
2. If the node is connected to other gateways in different regions, check the latency to those gateways as well.
    - If the latency to the other gateways is also high, this indicates a problem closer to the edge node like the site ISP. 

### 4. (Clustered Edge Nodes Only) Check Latency to Cluster Peer
If the reporting edge node is a member of a cluster, check the latency to the other cluster member.
1. Navigate to the Data Plane panel on the other member of the clusters.
1. Repeat [Step 3](#3-check-latency-to-other-gateways). to see if the other member of the cluster is impacted as well.
    - If the latency to the other cluster member is also high, this indicates a problem closer to the edge node like the site ISP.
    - If the latency to the other cluster member is normal, check to see if it is using a different ISP than the reporting edge node by viewing the ISP field in the [Location]({{<relref "/docs/nodes/shared/location">}}) panel. **If the cluster member is using a different ISP, making the non-impacted member the active member of the cluster will likely restore normal functionality while the impacted member's ISP issue is investigated further.**

## Post Incident Triage Steps
If the incident has resolve itself there will be a limit to the information that can be gathered.  For example, running MTR on the impacted node will not provide any information because the node is no longer experiencing the issue. The steps below are intended to determine what information is available and set up having more information if the issue reoccurs. 

Before starting, determine the start and end time of the incident.  Be aware that the portal will display most things using your browsers local timezone, but timestamps send via[alarm channels]({{<relref "/docs/alarms/channels">}}) will be in UTC time. As are any timestamps in [debug logs]({{<relref "/help-center/ops-logs/debug-logs">}}).

### 1. Enable Hop Monitoring
**Requires permissions to modify the edge node config. If your user is read-only skip to [Step 2](#2-run-mtr-to-reported-gateway).**

Hop monitoring uses TCP SYN packets to continuously monitor the path between an edge node and all it's attached gateways. Follow the steps the [Enable Network Hop Monitoring]({{<relref "/tutorials/gateway-tools/monitoring-network-hops-to-peers#enabling-network-hop-monitoring">}}) to enable hop monitoring. Enabling after the incident will provide more data if the issue reoccurs.  

If the impacted edge node is a member of a cluster, enable hop monitoring on the other member of the cluster as well.

### 2. Use Data Plane Panel to View Latency to Gateways
The Data Plane panel can be used to view the tunnel latency to all connected peers( and hop monitoring data if enabled ) up to the past 90 days.  To view data from the past:
1. Navigate to the Data Plane panel on the impacted edge node.
2. Click on one of the impacted gateway's name to view the latency to that gateway.
3. By default the data for the last hour is show. Under Monitoring click **Custom** in the time selector. {{<tgimg src="post-incident-investigate-1.png" alt="Data Plane panel Monitoring time selector">}}
4. Use the time selector to select a time range that includes the incident and click **Apply**. It is a good idea to include 15-30 minutes before the incident started and after the incident ended. {{<tgimg src="post-incident-investigate-2.png" alt="Data Plane panel custom time selector">}}
5. This will show the latency to the selected gateway over the selected time range. 
    - If hop monitoring was enabled, clicking any point on the graph will update to show the data collected at that time. {{<tgimg src="post-incident-investigate-3.png" alt="Data Plane panel latency graph">}}
    - If you select other peers/gateways, the latency to those gateways will be shown during the same time period.
        - If the gateway is a member of a cluster, view the other member of the cluster to see if it is also impacted.
        - Select other gateways in other data centers/regions to see if they are also impacted.


## Latency Issue Matrix
The below matrix will assume that 
- Edge1 is reporting high latency to East-gw1
- East-gw1 is the active member of a cluster that includes East-gw2, both in the same data center/availability zone
- Edge1 is also connected to West-gw1, the active member of a cluster that includes West-gw2, both in the same data center/availability zone
- Most VPN/Layer 4 traffic is between the East gateways and Edge1


In each column under the gateway name is what would be shown in the Data Plane panel of Edge1 during the incident (see [Step 2](#2-use-data-plane-panel-to-view-latency-to-gateways) for how to view other times).

| East-gw1 | East-gw2 | West-gw1 | West-gw2 | Likely Cause | 
| ---|---|---|---|---| 
| Consistent high latency | Periodic spikes of high latency | Periodic spikes of high latency  | Periodic spikes of high latency | This indicates an issue with Edge1 or its ISP. <ul><li> Check for high CPU and Memory utilization on Edge1.</li><li>Look at bandwidth usage, if it is plateauing at a certain level (e.g. 50Mbps) it could indicate the internet connection at the site is fully utilized or something is doing traffic shaping.</li><li>Work with the local network IT contact to have them open a ticket with their ISP.</li></ul> |
| Consistent high latency | Periodic spikes of high latency | No latency | No Latency | This indicates a problem with the internet path between the East gateways and Edge1. |
| Consistent high latency |  No latency | No Latency | No latency | This indicates either an issue with East-gw1 or a problem with the internet path between the East gateways and Edge1. <ul><li> Check for high CPU and Memory utilization on East-gw1.</li><li>Look at bandwidth usage, if it is plateauing at a certain level (e.g. 50Mbps) it could indicate the internet connection at the site is fully utilized or something is doing traffic shaping.</li><li> Pull [debug logs]({{<relref "/help-center/ops-logs/debug-logs">}}) from East-gw1 and Edge1 and look for errors. </li> <li> You could force a [peer disconnect]({{<relref "/docs/nodes/appliances/data-plane#peer-disconnect">}}) but this will disrupt traffic over the connection </li></ul> |

If Edge1 is the active member of a cluster **and both members of the cluster are using the same ISP** (typically using the public IPs will be the same or adjacent), the expected behavior would be that both members see the latency, but the latency will be higher and more consistent on Edge1 because it has more traffic.

If Edge1 is the active member of a cluster **and the member use different ISPs** (typically using the public IPs will be different) a few scenarios are possible:
- If the local ISP for Edge1 is the source of the issue, you would not see any latency on the other member of the cluster.  It would be recommended to promote the other member as the active member of the cluster.
- If the issue lies closer to the gateways internet connect, you wil likely see latency on both members of the cluster but it will be more pronounced on the active member because it has more traffic.


## What is High Tunnel Latency?
High tunnel latency is a condition where the time it takes for data to travel between gateway and edge nodes is significantly longer than normal. This can be caused by a variety of factors, including packet loss and latency on the internet between the gateway and edge nodes, or tunnel traffic exceeding the hardware capabilities of the gateway or edge node. 

Packet loss and latency can occur on the internet due to various factors, including network congestion, inadequate bandwidth, and routing inefficiencies. When data is transmitted between a source internet source provider (ISP) and a destination ISP, the two may have different peering relationships—either direct peering, where they exchange traffic directly, or transit relationships, where one ISP pays another to carry its traffic. These relationships can influence the speed and reliability of data transfer. For instance, if there is congestion at a peering point or if the transit path is suboptimal, packets may be delayed or lost altogether, leading to higher latency and poor VPN performance.

### Latency TCP vs UDP Tunnels
While both TCP and UDP traffic on the internet would be subject to the same latency issues, the way the data is transmitted and received can impact the perceived latency. Packet loss has a more significant impact on a TCP-based VPN tunnel compared to a UDP-based tunnel due to the fundamental differences in how these protocols handle data transmission. 

TCP is designed for reliability, ensuring that all packets are delivered in the correct order and without errors. When packet loss occurs in a TCP tunnel, the protocol triggers retransmission of lost packets, which can lead to increased latency and reduced overall throughput. If the application utilizing the tunnel also retransmits lost packets, the latency can be further exacerbated.

In contrast, UDP prioritizes speed over reliability, allowing packets to be sent without waiting for acknowledgments. This means that while UDP-based VPNs may experience some data loss, they can maintain better performance and lower latency, making them more resilient to packet loss in real-time applications provided the applications utilizing the tunnel properly manage retransmission of lost packets.






