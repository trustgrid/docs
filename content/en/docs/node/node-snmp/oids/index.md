---
Title: "OIDs"
Date: 2023-1-12
---
#### SNMP Standard OIDs
Trustgrid exposes data through standard OIDs to provide information about the below resources.  Most monitoring software will detect these automatically.

- CPU

- Memory

- Disk 

- Network Interface

#### Trustgrid Custom OIDs
The following custom data points are made available under Trustgrid’s Enterprise OID .1.3.6.1.4.1.53560 which will need to be used as the prefix for the below OIDs

- `.2.2.1` - Lists VPN Stats in a table each virtual network is a row in the table (signified with the suffix .# for that networks row number) with a value for each metric

    - `.1.#` - Lists the interface name associated with the below Virtual network name.  

    - `.2.#` - Lists the Virtual Network names associated with this node.  

    - `.3.#` - Lists the Number of New Flows

    - `.4.#` - Lists the Number of Active Flows  

    - `.5.#` - Lists the Inbound TCP Octets since last polling

    - `.6.#` - Lists the Outbound TCP Octets since last polling

    - `.7.#` - Lists the Inbound UDP Octets since last polling

    - `.8.#` - Lists the Outbound UDP Octet since last polling

    - `.9.#` - Lists the Inbound ICMP Octets since last polling

    - `.10.#` - Lists the Outbound ICMP Octets since last polling

- `.2.3.0` - Lists the Control Plane connectivity status. Returns UP if connected and DOWN if disconnected

- `.2.4.0` - Lists Data Plane connectivity status. Returns UP if connected to all gateways and DOWN if disconnected from at least one configured gateway

- `.2.5.1` - Table with name and value of all tags associated with the node

- `.2.6.0` - Integer showing the number of available data plane connections (known gateways if an edge node is queried, known edge node clients if a gateway node is queried)

- `.2.7.0` - Integer showing the number of active data plane connections

- `.2.8.0` - String showing the cluster status of the node. Possible values are PRIMARY, SECONDARY, or `NOT_CLUSTERED`

#### Example output
The output below shows the device has two virtual networks (mesh & LRE-NETWORK) both attached to the same physical interface (enp0s20f1).  The mesh network has no activity.  The LRE-NETWORK has significant TCP traffic and a large number of flows. 

{{< highlight json >}}
iso.3.6.1.4.1.53560.2.2.1.1.1 = STRING: "enp3s0"
iso.3.6.1.4.1.53560.2.2.1.1.2 = STRING: "enp3s0"
iso.3.6.1.4.1.53560.2.2.1.1.3 = STRING: "enp3s0"
iso.3.6.1.4.1.53560.2.2.1.2.1 = STRING: "Default-Route"
iso.3.6.1.4.1.53560.2.2.1.2.2 = STRING: "mesh"
iso.3.6.1.4.1.53560.2.2.1.2.3 = STRING: "LRE-NETWORK"
iso.3.6.1.4.1.53560.2.2.1.3.1 = INTEGER: 0
iso.3.6.1.4.1.53560.2.2.1.3.2 = INTEGER: 0
iso.3.6.1.4.1.53560.2.2.1.3.3 = INTEGER: 6562
iso.3.6.1.4.1.53560.2.2.1.4.1 = INTEGER: 0
iso.3.6.1.4.1.53560.2.2.1.4.2 = INTEGER: 0
iso.3.6.1.4.1.53560.2.2.1.4.3 = INTEGER: -8328
iso.3.6.1.4.1.53560.2.2.1.5.1 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.5.2 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.5.3 = Counter64: 5821341
iso.3.6.1.4.1.53560.2.2.1.6.1 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.6.2 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.6.3 = Counter64: 132753210
iso.3.6.1.4.1.53560.2.2.1.7.1 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.7.2 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.7.3 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.8.1 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.8.2 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.8.3 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.9.1 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.9.2 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.9.3 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.10.1 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.10.2 = Counter64: 0
iso.3.6.1.4.1.53560.2.2.1.10.3 = Counter64: 0
iso.3.6.1.4.1.53560.2.3.0 = STRING: "UP"
iso.3.6.1.4.1.53560.2.4.0 = STRING: "UP"
iso.3.6.1.4.1.53560.2.5.1.1.1 = STRING: "test"
iso.3.6.1.4.1.53560.2.5.1.1.2 = STRING: "enviro"
iso.3.6.1.4.1.53560.2.5.1.2.1 = STRING: "null"
iso.3.6.1.4.1.53560.2.5.1.2.2 = STRING: "LRE"
iso.3.6.1.4.1.53560.2.6.0 = INTEGER: 8
iso.3.6.1.4.1.53560.2.7.0 = INTEGER: 8
iso.3.6.1.4.1.53560.2.8.0 = STRING: "NOT_CLUSTERED"
{{< /highlight >}}