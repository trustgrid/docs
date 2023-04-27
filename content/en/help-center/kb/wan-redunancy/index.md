---
title: "WAN/ISP Redundancy Configurations"
linkTitle: "WAN/ISP Redundancy"
date: 2023-04-21
weight: 200
description: Shows different way Trustgrid nodes can utilize multiple WAN or ISP connections for redundancy
---

{{<pageinfo>}} The sections below provide examples of how Trustgrid nodes can be deployed to provide WAN/ISP redundancy{{</pageinfo>}}

## Behind Firewall/Router with Multiple ISP Connections
In this configuration the Trustgrid WAN interfaces are behind a firewall or router that has two independent ISP connections to provide internet access. 

The firewall or router is responsible for either failing over outbound traffic in the event of an ISP failure, or to route different nodes to utilize specific ISP connections. 

The WAN interface of the Trustgrid node would utilize **private IPs** in this configuration that are NAT'd to public IPs by the firewall/router.
{{<alert color="info">}}This is the only configuration that:
* Supports **Single Node** deployments
* Supports **Single Interface** configurations
{{</alert>}}

{{<alert color="warning">}} This configuration is **not recommended for nodes acting as gateways**.  In this configuration devices are NAT'd to a different public IP depending on which ISP is active.  A gateway node can only announce a single IP or DNS address for its clients to connect to. It would be possible to announce a DNS record and some form of dynamic DNS health check to update the record if the ISP changed. But that change would have to propagate which could take some time.{{</alert>}}

### Single Node Behind Firewall
This configuration shows a single node connected to a firewall or router that has two ISP connected to it. 

``` mermaid
graph LR
    
    intHost[Internal Hosts]
    intNet[[Internal\n Network]]
    intHost <-.Optional.-> intNet <-.Optional.-> snLAN 
    subgraph sn [Single Node]
        snWAN[WAN\nInterface]
        snLAN[LAN\nInterface]
        snLAN ~~~ snWAN
    end
    firewall["Firewall/Router"]
    dmzNet[[DMZ\n Network]]
    firewall == Primary==> ISP1([ISP 1])
    firewall -. Failover/Backup .-> ISP2([ISP 2])
    snWAN --> dmzNet --> firewall
    classDef tgnode fill:#346ed9,color:white
    class sn tgnode
```

### Clustered Nodes Behind Firewall

This configuration shows a cluster of Trustgrid nodes connected to a firewall or router that has two ISPs connected to it.

``` mermaid
graph LR
    intHost[Internal Hosts]
    intNet[[Internal\n Network]]
    intHost <-.Optional.-> intNet <-.Optional.-> cl1LAN & cl2LAN
    subgraph cluster [" "]
       subgraph Cluster-Node1
           cl1WAN[WAN\nInterface]
           cl1LAN[LAN\nInterface]
           cl1LAN ~~~ cl1WAN
       end
       subgraph Cluster-Node2
           cl2WAN[WAN\nInterface]
           cl2LAN[LAN\nInterface]
           cl2LAN ~~~ cl2WAN
       end
    end
    dmzNet[[DMZ\n Network]]    
    firewall["Firewall/Router"]
    firewall == Primary==> ISP1([ISP 1])
    firewall -. Failover/Backup .-> ISP2([ISP 2])
    cl1WAN & cl2WAN --> dmzNet --> firewall
    classDef tgnode fill:#346ed9,color:white
    class Cluster-Node1,Cluster-Node2 tgnode
    style cluster fill:#c8c8c8
```


## Cluster WAN Interface to Different Networks 
Another method of providing redundancy takes advantage of Trustgrid [clustering]({{<ref "/docs/clusters">}}) by connecting each member of the cluster to a different ISP on their WAN interface. This could be done by:
* Directly attaching each member WAN interface to a different ISP handoff
* Connecting each member WAN interface to different DMZ private networks configured to use different ISPs for internet access
* A combination of public and private WAN networks

In this configuration, the master/active member of the node will determine which ISP is being utilized.  If one ISP is preferred you will need to have operational procedures in place to ensure the member node connected to it stays the master node. It is recommended that you keep the [cluster mode configured as Automatic Failback]({{<ref "/docs/clusters#cluster-mode">}}) and your organization establishes a [tag]({{<ref "/docs/nodes/tags">}}) to designate the preferred member.  

{{<alert color="info">}}Because the WAN interfaces are on different networks, this configuration **requires at least one additional LAN interface** to be configured for accessing internal resources and providing the [cluster heartbeat]({{<ref "/docs/clusters/#cluster-heartbeat-communication">}}) communication{{</alert>}}

### Cluster WAN Direct Connections to Multiple ISPs
This configuration shows Cluster-Node1's WAN interface directly connected to one ISP.  And Cluster-Node2's WAN interface is connected directly to another ISP. 

``` mermaid
graph LR
    intHost[Internal Hosts]
    intNet[[Internal\n Network]]
    intHost <---> intNet <---> cl1LAN & cl2LAN
    subgraph cluster [" "]
        subgraph Cluster-Node1
           cl1WAN[WAN\nInterface]
           cl1LAN[LAN\nInterface]
           cl1LAN ~~~ cl1WAN
        end
        subgraph Cluster-Node2
           cl2WAN[WAN\nInterface]
           cl2LAN[LAN\nInterface]
           cl2LAN ~~~ cl2WAN
        end
    end
    cl1WAN ==> ISP1([ISP 1])
    cl2WAN ==> ISP2([ISP 2])
    classDef tgnode fill:#346ed9,color:white
    class Cluster-Node1,Cluster-Node2 tgnode
    style cluster fill:#c8c8c8

```

### Cluster WAN to separate DMZ networks

This configuration shows Cluster-Node1's WAN interface is connected to one DMZ/private network that is behind a firewall connected to one ISP. And Cluster-Node2's WAN interface is connected to a second DMZ network behind a separate firewall connected to a different ISP.

``` mermaid
graph LR
    intHost[Internal Hosts]
    intNet[[Internal\n Network]]
    intHost <---> intNet <---> cl1LAN & cl2LAN
    subgraph cluster [" "]
       subgraph Cluster-Node1
           cl1WAN[WAN\nInterface]
           cl1LAN[LAN\nInterface]
           cl1LAN ~~~ cl1WAN
       end
       subgraph Cluster-Node2
           cl2WAN[WAN\nInterface]
           cl2LAN[LAN\nInterface]
           cl2LAN ~~~ cl2WAN
       end
    end
    dmzNet1[[DMZ\n Network 1]]
    dmzNet2[[DMZ\n Network 2]]    
    firewall1["Firewall/Router 1"]
    firewall2["Firewall/Router 2"]
    cl1WAN --> dmzNet1 --> firewall1 ==> ISP1([ISP 1])
    cl2WAN --> dmzNet2 --> firewall2 ==> ISP2([ISP 2])
    classDef tgnode fill:#346ed9,color:white
    class Cluster-Node1,Cluster-Node2 tgnode
    style cluster fill:#c8c8c8
```

{{<alert color="info">}} Note: While shown with two separate firewall/routers, this could also be accomplished with a single firewall/router configured to route each DMZ network to a different ISP. {{</alert>}}

### Cluster WAN using mix of public and DMZ networks

This configuration shows Cluster-Node1's WAN interface is connected to a private DMZ network behind a firewall that is using ISP 1.  Cluster-Node2's WAN interface is connected directly to ISP 2.

``` mermaid
graph LR
    intHost[Internal Hosts]
    intNet[[Internal\n Network]]
    intHost <---> intNet <---> cl1LAN & cl2LAN
    subgraph cluster [" "]
       subgraph Cluster-Node1
           cl1WAN[WAN\nInterface]
           cl1LAN[LAN\nInterface]
           cl1LAN ~~~ cl1WAN
       end
       subgraph Cluster-Node2
           cl2WAN[WAN\nInterface]
           cl2LAN[LAN\nInterface]
           cl2LAN ~~~ cl2WAN
       end
    end
    dmzNet1[[DMZ\n Network 1]]    
    firewall1["Firewall/Router 1"]
    cl1WAN --> dmzNet1 --> firewall1 ==> ISP1([ISP 1])
    cl2WAN ==> ISP2([ISP 2])
    classDef tgnode fill:#346ed9,color:white
    class Cluster-Node1,Cluster-Node2 tgnode
    style cluster fill:#c8c8c8
```