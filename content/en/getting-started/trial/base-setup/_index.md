---
title: Initial Agent Setup
linkTitle: Initial Setup
weight: 20
description: "Deploy two agents and see traffic pass between them - 10 minutes"
---
In this section, we will cover the initial setup of two Trustgrid agents on separate Ubuntu hosts to demonstrate communication across a Trustgrid virtual network. Examples where this could be useful in the real world include:
- Enabling an application server secure access to a data source on a remote server
- Enabling a server the ability to connect via SSH to one or more remote servers without exposing any ports to the internet

## High-Level Steps
- Deploy two Trustgrid agents
- Initiate ping between devices
- Utilize tools to monitor traffic passing between agents

## Prerequisites
- [Trustgrid Trial Account]({{<relref "../request-trial">}})
- [Two instances of Ubuntu 22.04]({{<relref "./ubuntu2204">}}) with `sudo` privileges to install additional repositories and packages.
  - Both instances need to be able to make **outbound connections to the internet on ports 443 and 8443** and cannot be subject to TLS inspection that alters the certificate chain.
  - Ideally, these devices should not be able to communicate with each other directly. This is not a hard requirement.


## Understanding the Default Network
To facilitate a smooth trial a default [virtual network]({{<relref "/getting-started/basic-architecture#virtual-networks">}}) is created. This network uses the carrier-grade NAT address space 100.64.0.0/10 as an [IP Pool](). Agents are automatically assigned an IP address from this pool when they are attached to the virtual network and routes are automatically created to allow communication between agents on the same virtual network.


## Step 1 - Setup Agents
### Install First Agent
{{< readfile file="/tutorials/agent-deploy/ubuntu2204-install.md" >}}
#### Determine agent1 IP address
As part of registration, each agent is automatically assigned an IP address on the Trustgrid virtual network. We will need this IP address to confirm communication between the agents in later steps.
1. From the console of agent1, run the below command:{{<codeblock>}}ip address show dev trustgrid0{{</codeblock>}}
2. Look for an IP address starting with 100.64. 

{{<tgimg src="agent1-ip.png" width="90%" caption="Console showing the Trustgrid IP address of 100.64.0.1">}}

### Install Second Agent

Repeat the above steps on the second Ubuntu instance to [install the Trustgrid agent](#install-first-agent) with a name like "**agent2**" on the same Virtual Network and [determine agent2's IP address](#determine-agent1-ip-address)
This should return a different IP address in the same network.
{{<tgimg src="agent2-ip.png" width="90%" caption="Console showing the Trustgrid IP address of 100.64.0.2">}}

## Step 2 - Confirm communication

From `agent2` run the below command to ping `agent1` using its Trustgrid IP address:
{{<codeblock>}}ping -c 4 100.64.0.1{{</codeblock>}}
{{<alert color="info">}} The examples above and below assume the agents received the first two addresses in the pool. If the agents were assigned different IP addresses adjust the below commands accordingly. {{</alert>}}
This should generate 4 ping requests with successful responses showing traffic is traversing the Trustgrid network between the two agents.
{{<tgimg src="ping-agent1.png" width="60%" caption="Successful ping from agent2 to agent1">}}

Similarly, from the console of `agent1` run the below command to ping `agent2`:
{{<codeblock>}}ping -c 4 100.64.0.2{{</codeblock>}}
{{<tgimg src="ping-agent2.png" width="60%" caption="Successful ping from agent1 to agent2">}}

## Step 3 - View Flow Logs
[Flow logs]({{<relref "/help-center/flow-logs">}}) provide visibility Trustgrid provides into the traffic passing through agents. They show details of every connection including source, destination, protocols, and more.
### View on Overview
The flow logs for agents are shown on the Overview page for each node beneath the stats. 
{{<tgimg src="agent-icmp-flow-logs.png" width="90%" caption="ICMP flow logs for agents 1 and 2">}}
### View Flow Logs Table
Additionally, the flow logs table is available under History > Flow Logs. {{<tgimg src="flow-log-nav.png" width="30%">}}

This table includes the ability to perform [advanced searches]({{<relref "/help-center/flow-logs#advanced-search">}}) and [export]({{<relref "/help-center/flow-logs#exporting-flow-logs-to-csv">}}) the flow logs to csv. 

In the below screenshot, the TCP flow log from the [generated web traffic](#step-4---optional-generate-web-traffic) steps is visible including TCP flags.  TCP Flags can be very useful in troubleshooting failed connections in the past. If only the SYN flag is shown this indicates the destination ip:port did not respond to complete the TCP handshake.
{{<tgimg src="agent2-flow-logs.png" width="90%" caption="Flow logs table with search for agent2">}}

## Next Steps
Now that you have the basic agent setup complete and have seen traffic passing between the agents you can proceed to test [using Access Policies]({{<relref "/getting-started/trial/network-acl">}}) to control what traffic is allowed on the network.