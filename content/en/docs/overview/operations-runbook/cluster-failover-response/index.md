---
Title: "Cluster Failover Response"
Date: 2023-1-9
Tags: ["cluster", "help", "troubleshoot"]
---

{{% pageinfo %}}
When the active (master) member of the [cluster]({{< ref "docs/cluster" >}}) goes unhealthy the standby member will take over the active role. This process should be automatic and not require manual intervention. However, in certain circumstances, such as the unexpected failover of a public gateway, it is worth investigating to confirm traffic is in a healthy state.
{{% /pageinfo %}}

### Possible Messages

- Master role assumed - failover

  - Indicates the master role has moved from the designated master (primary) to the backup/secondary [node]({{< ref "docs/node" >}})

- Master role reclaimed by expected master

  - Indicates the master role has returned to the designated master

### Failover Process

Below is a brief description events that occur during a failover process:

- The [node]({{< ref "docs/node" >}}) assuming the master role will ARP to the network that it now owns the [Cluster]({{< ref "docs/cluster" >}}) Virtual IP (VIP)

- The [Domain route]({{< ref "docs/domain/routes" >}}) table will update that the assuming [node]({{< ref "docs/node" >}}) should receive all traffic for the [cluster]({{< ref "docs/cluster" >}}) (_clustername-master_)

- The assuming [node]({{< ref "docs/node" >}}) will load all NAT entries associated with the [cluster]({{< ref "docs/cluster" >}})

### Response Process

After a failover or failback it is necessary to verify that traffic is flowing appropriately.

1. Login to the portal and navigate to the affected [cluster’s]({{< ref "docs/cluster" >}}) page

2. Verify that only a single [node]({{< ref "docs/node" >}}) shows as master

3. On the `Configuration` → `Network` tab note the [cluster]({{< ref "docs/cluster" >}}) VIP

![img](cluster-virtual-ip2.png)

1.  Click on the indicated current master

    a. Verify the VPN Route Table shows

        i. Navigate to the `Configuration` → `VPN` tab

        ii. Launch the "View Virtual Route Table" tool

      ![img](virtual-network-tools.png)

        iii. Verify that routes show as “available true”

    ![img](routing-tables.png)

           1. If the cluster is a gateway cluster there may be many routes and not all be active, just confirm many show as available.

           2. The route to the management VIP for the other node in the cluster will always be false.

    b. Verify traffic is flowing through the appropriate node.

        i. Navigate to the `Configuration` → `Network` tab

        ii. Confirm the interface associated with the Cluster VIP is selected

           3. For single interface nodes: ETH0 / Network Adapter 1 - WAN Adapter

           4. For dual interface nodes: ETH1 / Network Adapter 2 - LAN Adapter

        iii. Open the `Sniff Interface Traffic` tool.

            5. Set the filter to “host clusterVIP” without quotes and replacing clusterVIP with the appropriate Cluster virtual IP. Click `start session`.

            6. Confirm that you see traffic flowing through the interface. **Continue monitoring for several minutes to confirm the traffic is maintained.**

            7. Leave the Sniff Interface Tool running while completing the next step

        iv. Repeat steps i-iii on the node that **is not** currently indicated as the master.

            1.You want to verify there is **no traffic for the cluster VIP running through the non-master node**.

            8. You may see a periodic ARP from the cluster master, but that should be it.

    1. Compare traffic volume before and after the failover

        a. If the event was a failover and failback compare traffic from the current master

        b. If the event was a failover but has not failed back you will need to compare traffic volume on the current master to the volume on the previous master

            i. e.g if traffic failed from node1 to node2, compare node1’s traffic prior to the failover to the volume of traffic on node2 after the failover
