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

### Single Node Behind Firewall

``` mermaid
graph LR
    
    intHost[Internal Hosts]
    intNet[[Internal\n Network]]
    intHost <-.Optional.-> intNet <-.Optional.-> snLAN 
    subgraph sn [Single Node]
        direction RL
        snWAN[WAN\nInterface]
        snLAN[LAN\nInterface]
    end
    firewall["Firewall/Router"]
    dmzNet[[DMZ\n Network]]
    firewall == Primary==> ISP1 
    firewall -. Failover/Backup .-> ISP2
    snWAN --> dmzNet --> firewall
```

### Clustered Nodes Behind Firewall

``` mermaid
graph LR
    subgraph internal[Internal Network]
       intHost[Internal Hosts]
    end
    intHost <-.Optional.-> cl1LAN & cl2LAN
    subgraph dmz [" "]
        subgraph Clustered Nodes
           subgraph Cluster-Node1
               cl1WAN[WAN\nInterface]
               cl1LAN[LAN\nInterface]
           end
           subgraph Cluster-Node2
               cl2WAN[WAN\nInterface]
               cl2LAN[LAN\nInterface]
           end
           Cluster-Node1 ~~~~ Cluster-Node2
        end
        firewall["Firewall/Router"]
    end
    firewall == Primary==> ISP1 
    firewall -. Failover/Backup .-> ISP2
    cl1WAN & cl2WAN ---> firewall
```


## Cluster WAN Interface to Different Networks 
Another method of providing redundancy takes advantage of Trustgrid [clustering]({{<ref "/docs/clusters">}}) by connecting each member of the cluster to a different ISP on their WAN interface. This could be done by:
* Directly attaching each member WAN interface to a different ISP handoff
* Connecting each member WAN interface to different DMZ private networks configured to use different ISPs for internet access
* A combination of public and private WAN networks

{{<alert color="info">}}Because the WAN interfaces are on different networks, this configuration **requires at least one additional LAN interface** to be configured for accessing internal resources and providing the [cluster heartbeat]({{<ref "/docs/clusters/#cluster-heartbeat-communication">}}) communication{{</alert>}}

### Cluster WAN Direct Connections to Multiple ISPs


### Cluster WAN to separate DMZ networks


### Cluster WAN using mix of public and DMZ networks
