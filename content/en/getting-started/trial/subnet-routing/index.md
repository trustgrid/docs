---
title: Configure Subnet Routing
linkTitle: Subnet Routing
description: Configure an agent to route traffic to resources on the local network
weight: 40
---
The tutorial up to this point only demonstrated the ability to send traffic between the IP addresses assigned to the trustgrid0 interfaces on each agent.  However, it's quite likely that security and operational constraints will prevent the Trustgrid agent from being installed directly on the target server.  Instead, the agent will need to route traffic destined for the target server's local network.

In this section, we will configure an agent to route traffic destined for another network through the Trustgrid virtual network. To avoid assumptions about the networks your agents are deployed in, we will use a loopback interface to simulate an adjacent network. But the steps shown here would apply equally if routing between real physical interfaces.


## High-Level Goals
- Create a loopback interface to simulate an adjacent network 
- Configure a route to send traffic destined for the adjacent network subnet through a Trustgrid agent
- Configure the agent to forward the virtual network traffic to the real subnet
- Configure a NAT to obscure the real target address

## Prerequisites
- Two agents installed and confirmed communicating (pingable) on the default virtual network as set up in previous tutorials
    - [Initial Setup]({{<relref "/getting-started/trial/base-setup">}})
    - [Access Policies]({{<relref "/getting-started/trial/access-policies">}})

## Step 1 - Configure Loopback Interface and Start Server
On **agent1** we will create a loopback interface and assign it an IP address from another subnet. This will simulate an adjacent network that the agent needs to route traffic to.
{{<alert color="info">}}The assigned IP address `172.31.255.1/32` can be replaced by another IP if it conflicts with an IP or route on either of the agents.{{</alert>}}

1. Gain console access to **agent1**. 
1. Run the below commands to add a loopback interface called `lo2` and assign an IP address `172.16.255.1`. {{<codeblock>}}sudo ip link add name lo2 type dummy
sudo ip addr add 172.31.255.1/32 dev lo2
sudo ip link set lo2 up
{{</codeblock>}}
1. Start a netcat listener on the lo2 interface using the below commands: {{<codeblock>}}nc -k -l -p 5000 -s 172.31.255.1{{</codeblock>}}

## Step 2 - Configure access to an adjacent network

### Configure IP Forwarding on agent1
The first step is to configure **agent1** to forward traffic to networks other than its own trustgrid0 interface IP. By default, the Linux operating system will only accept packets on an interface if the destination IP is assigned to that interface. Enabling IP forwarding allows the OS to manage packets destined for other IPs and networks. 

1. In the Trustgrid portal Nodes page select **agent1**
1. Navigate to the VPN panel {{<tgimg src="agent1-vpn-nav.png" width="40%">}}
1. Click the box to "Enable Forwarding" {{<tgimg src="agent1-enable-forwarding.png" width="40%">}}
1. Click "Save"

### Configure Virtual Network Route to agent1
Now that agent1 has an "adjacent network" of interest, the virtual network needs to be told that packets destined for this network should be routed to agent1. We will do this with a [virtual network route.]({{<relref "/docs/domain/virtual-networks/routes">}})

1. In the Trustgrid portal, navigate to the "Virtual Network" page and select the **default** network. 
1. On the Routes panel click the "Add Route" button. {{<tgimg src="add-routes-button.png" width="60%">}}
1. Configure the route with the following settings
    1. Destination: Node: agent1
    1. Destination CIDR: `172.31.255.1/32`
    1. Metric: `1`
    1. Description: agent1 loopback
    {{<tgimg src="agent1-route.png" width="95%">}}
1. Click Save at the top of the page.
1. Navigate to the Review Changes page and click Apply Changes. Confirm with Yes when prompted.  {{<tgimg src="agent1-route-confirm.png" width="90%" >}}

## Configure OS Route on agent2
Next, we need to tell **agent2** to use the trustgrid0 interface to route traffic destined for the adjacent network. This is done by adding a local route in the operating systems that will forward traffic for the target IP, 172.31.255.1, via the trustgrid0 interface and flow through the virtual network tunnel.
1. Gain console access to **agent2**
1. Run the command: {{<codeblock>}}sudo ip route add 172.31.255.1/32 dev trustgrid0{{</codeblock>}}
1. You can confirm this is effective with the command: {{<codeblock>}}ip route{{</codeblock>}}
1. You should see the route listed as shown below: {{<tgimg src="os-route.png" width="80%" alt="ip route output with the last line showing '172.31.255.1 dev trustgrid0 scope link metric 1'" caption="Route for 172.31.255.1 on the trustgrid0 interface">}}

## Confirm Connectivity
### Test ICMP Connectivity
We can now test ICMP connectivity using `ping`.
1. On **agent2**, run the command: {{<codeblock>}}ping -c 3 172.31.255.1{{</codeblock>}}
You should see replies from the IP on agent1's lo2 interface.
{{<tgimg src="routed-icmp-success.png" width="80%" alt="successful pings to 172.31.255.1" caption="Successfully routed pings">}}

This confirms that traffic destined for the adjacent network IP is:
1. Routed by the OS on agent2 out the trustgrid0 interface
1. Flows over the virtual network to agent1
1. Is received and forwarded by agent1 to the lo2 interface
1. Replies return over the same path

### Test TCP Connectivity
Similarly, we can verify TCP traffic is being routed and forwarded as expected.  Similar to the [access policy tutorial]({{<relref "../access-policies">}}), we will test connectivity to the netcat listener but in this case, we will bind the listener to the loopback interface IP.

#### Update Access Policies
In the [access policy tutorial]({{<relref "../access-policies">}}) we specified a rule that only allowed TCP access to port 5000 when the destination IP was agent1's trustgrid0 IP, `100.64.0.1`. Since we are now routing traffic to agent1's loopback IP `172.31.255.1` instead, we need to update the policy to allow access to that IP/port instead.
1. In the Trustgrid portal navigate to the Virtual Networks page and select the default network. 
1. Click the edit icon to the right of the TCP 5000 rule previously created. {{<tgimg src="edit-acl.png" width="95%">}}
1. Change the Destination to `172.31.255.1/32` and click Save.{{<tgimg src="tcp-acl-change.png" width="70%">}}
1. Confirm the rule change was applied by navigating to the Review Changes page and clicking Apply Changes. Confirm with Yes when prompted. {{<tgimg src="tcp-acl-change-review.png" width="90%">}}

#### Confirm TCP Connectivity
Now we can start our listener and test connectivity.
1. On **agent1** run the command: {{<codeblock>}}nc -k -l -p 5000 -s 172.31.255.1{{</codeblock>}}
1. On **agent2** and confirm we can connect to port 5000 on agent1's loopback ip with the command: {{<codeblock>}}nc -vz -w 1 172.31.255.1 5000{{</codeblock>}}
1. You should see a successful connection message like below. {{<tgimg src="tcp-connect-success.png" width="70%">}}

