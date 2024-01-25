---
title: Implement Access Policy Restrictions
linkTitle: Access Policies
description: Configure access policies to limit the traffic allowed on the virtual network
weight: 30
---

When designing a network from a zero-trust, least-privilege mindset all agents and appliances shouldn't have unrestricted access to each other and their adjacent network resources. Trustgrid allows for granular control over what traffic is permitted across the virtual network using [access policies]({{<relref "/docs/domain/virtual-networks/access-policy">}}).  The default virtual network was created with an "Allow All" policy to facilitate initial connectivity during the trial period. 

In this section, we will update the policy to restrict traffic between the two agents. 

## High-Level Steps
- Generate TCP traffic between the agents
- Create an access policy to deny all traffic and verify traffic is blocked
- Create additional policies to allow specific traffic

## Prerequisites
- Two agents installed and confirmed communicating (pingable) on the default virtual network as [set up in previous tutorials]({{<relref "/getting-started/trial/base-setup">}})
- Netcat ( `nc` ) installed on both agents to generate and capture traffic. This utility is commonly installed by default in most Linux OSes, but may need to be installed manually.

## Step 1 - Generate Traffic
{{<alert color="info">}} Note: The below assumes **agent1** was assigned `100.64.0.1` and **agent2** was assigned `100.64.0.2`. If your agents were assigned different IP addresses update the below commands accordingly.{{</alert>}}

### Start TCP Netcat listener on agent1
Netcat (or `nc`) is a versatile networking tool that can function as a client to send data to other hosts or as a listener to receive data on a specified port. We will first configure it to listen on agent1.
1. Gain console access to **agent1**. 
1. Run the below command to start netcat listening on TCP port 5000. {{<codeblock>}}nc -k -l 5000{{</codeblock>}}
1. Leave this terminal open as netcat will continue listening in the background and will display any data sent to port 5000.

### Send Traffic from agent2
1. Gain console access to **agent2**.
1. Run the below command to send a message from **agent2** to **agent1**.  {{<codeblock>}} echo "Hello from $(hostname)!" | nc -q 1 100.64.0.1 5000 {{</codeblock>}}
1. After about a second you will be returned to the command prompt for **agent2** with no output.
1. Return to the console for **agent1** and you should see a message like `Hello from agent2!` {{<tgimg src="tcp-hello.png" width="50%" caption="Netcat listener output on agent1">}}

### Stop TCP Netcat listener on agent1
1. Return to the console for **agent1** and press Ctrl-C to terminate the netcat listener that was started in Step 1.


### Confirm Ping (ICMP) Traffic
Before implementing access policy restrictions, confirm ping (ICMP) traffic is still allowed between the agents:
1. On **agent1** run: {{<codeblock>}}ping -c 3 100.64.0.2{{</codeblock>}}
2. On **agent2** run: {{<codeblock>}}ping -c 3 100.64.0.1{{</codeblock>}}
You should see ping responses from the other agent, confirming ICMP traffic is still permitted by the default "Allow All" policy.

## Step 2 - Deny All Traffic
### Update Existing Access Policy
Now we will update the default policy to deny all traffic and confirm TCP communication is blocked:
1. In the Trustgrid portal, return to the home page and select the Virtual Networks option from the left navigation menu. {{<tgimg src="home-vnet.png" width="40%" >}}
1. From the Virtual Networks table, select the **default** virtual network. {{<tgimg src="default-vnet.png" width="50%" >}}
1. Click the Access Policies option from the navigation menu on the left. {{<tgimg src="vnet-acl-menu.png" width="50%" >}}
1. The existing "Allow All" policy will be displayed.  Click the edit button on the far right.  {{<tgimg src="default-acl-allow-all.png" width="90%">}}
1. Change the policy action to "Reject", update the description to say 'deny all traffic' and click Save. {{<tgimg src="default-acl-reject-all.png" width="70%">}}
    1. In addition to the "Allow" and "Reject" actions is the option to "Drop".  This will also deny traffic but without sending a response back to the source. Reject is preferred **for the trial** because it provides feedback that traffic was blocked. This may not be the desired setting in a production environment. 
1. All changes at the virtual network level must be reviewed and applied before being broadcast to nodes. Click the Review Changes from the left navigation menu. You will see the previous and new settings.  Click the "Apply Changes" button and then confirm when prompted. {{<tgimg src="acl-deny-apply.png" width="90%">}}

### Confirm ICMP Traffic is Blocked
1. On **agent1** run: {{<codeblock>}}ping -c 3 100.64.0.2{{</codeblock>}}
2. On **agent2** run: {{<codeblock>}}ping -c 3 100.64.0.1{{</codeblock>}}
You should see ping responses from the other agents are now filtered. {{<tgimg src="icmp-filtered.png" width="75%" caption="Terminal showing ICMP filtered responses">}}

### Confirm TCP Traffic is Blocked
1. On **agent1**, run the following command to start netcat listening on TCP port 5000. {{<codeblock>}}nc -k -l 5000{{</codeblock>}}
1. On **agent2**, run the following command to test connectivity to port 5000 on agent1: {{<codeblock>}}nc -vz 100.64.0.1 5000{{</codeblock>}}
You should see a "Connection refused" message indicating the TCP connection was blocked by the deny all policy. {{<tgimg src="tcp-refused.png" width="75%" caption="Netcat output showing 'Connection refused'">}}
1. Return to **agent1** and enter Ctrl+C to terminate the netcat listener.

## Step 3 - Allow Desired Traffic
Now that we've confirmed the virtual network is blocking traffic we can add rules to allow the specific traffic we want. 

### Allow all ICMP
First, we will add a rule to allow ICMP between all addresses on the network. This is frequently used as a troubleshooting tool with very little risk.
1. In the Trustgrid portal, navigate to the default virtual network's access policies page.
1. Click the "Add Rule" button.
1. Configure the rule as below:
    1. Action: Allow
    1. Protocol: ICMP
    1. Source: 0.0.0.0/0  
    1. Destination: 0.0.0.0/0
    1. Line number: 90 (It is important that the rule number is lower than that on the default deny all policy, 100, so that it is evaluated first)
    1. Description: Allow all ICMP {{<tgimg src="allow-all-icmp.png" width="70%">}}
1. Click Save
1. Select the Review Changes panel, click Apply Changes and, then confirm. 

### Confirm ICMP Allowed
1. On **agent1** run: {{<codeblock>}}ping -c 3 100.64.0.2{{</codeblock>}}
2. On **agent2** run: {{<codeblock>}}ping -c 3 100.64.0.1{{</codeblock>}}
You should see ping responses from the other agent, confirming ICMP traffic is still permitted by the new policy.
{{<tgimg src="icmp-allowed.png" width="90%" caption="Successful ping from agent2 to agent1">}}

### Allow Specific TCP
Next, we will add a rule to allow more specific TCP traffic. Specifically, we will allow **agent2** to connect to port 5000 on **agent1**.
1. In the Trustgrid portal, navigate to the default virtual network's access policies page.
1. Click the "Add Rule" button.
1. Configure the rule as below:
    1. Action: Allow
    1. Protocol: TCP
    1. Source: 100.64.0.2/32 
    1. Destination: 100.64.0.1/32
    1. Port Range: 5000
    1. Line number: 80 (It is important that the rule number is lower than that on the default deny all policy, 100, so that it is evaluated first)
    1. Description: Allow TCP 5000 on Agent1 {{<tgimg src="tcp-5000-allow.png" width="70%">}}
1. Click Save
1. Select the Review Changes panel, click Apply Changes and, then confirm. 

### Confirm TCP 5000 Allowed
Now, we will confirm the specific TCP rule is allowing traffic as expected:
1. On **agent1**, run the following command to start netcat listening on TCP port 5000. {{<codeblock>}}nc -k -l 5000{{</codeblock>}}
1. On **agent2**, run the following command to test connectivity to port 5000 on agent1: {{<codeblock>}}nc -vz 100.64.0.1 5000{{</codeblock>}} {{<tgimg src="tcp-5000-success.png" width="75%" caption="Successful connection to port 5000 on agent1">}}

But, we also want to make sure other TCP traffic isn't allowed.
1. On **agent1** enter Ctrl+C to terminate the netcat listener. Then enter the command: {{<codeblock>}}nc -k -l 5001{{</codeblock>}} {{<alert color="info">}}Note that the port has changed to 5001{{</alert>}}
1. On **agent2**, run the following command to test connectivity to port 5000 on agent1: {{<codeblock>}}nc -vz 100.64.0.1 5001{{</codeblock>}}
{{<tgimg src="tcp-5001-refused.png" width="75%" caption="Connection refused on port 5001">}}

## Next Steps
So far we've only seen traffic passing between two agent IPs.  In the next tutorial, you will see how you can [route traffic to adjacent network devices]({{<relref "../subnet-routing">}}).