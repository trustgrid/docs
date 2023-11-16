---
linkTitle: Diagnostics
Title: Gateway Diagnostics
description: Tools and Information for troubleshooting gateway connectivity issues 
weight: 30
---

The diagnostics page includes stats on aggregate traffic across all peer tunnels regardless of overlay service as well as tools for troubleshooting connectivity.

{{<tgimg src="gateway-diags.png" width="80%" caption="Gateway Diagnostics Page">}}

## Gateway Troubleshooting Tool

The Troubleshoot Gateway Traffic tool allows you to inspect live diagnostic messages about traffic between this node and its configured peers. 
{{<tgimg src="gateway-diag-troubleshoot.png" width="60%" caption="Gateway Troubleshooting Tool">}}

Clicking the "Troubleshoot Gateway Traffic" button will open the below diaglog that allows for filtering the output seen in the Troubleshoot Gateway tool by peer, local or service. Selecting what to filter on and clicking "Apply" will update the output seen in the tool. Accepting the default will display all gateway messages.

{{<tgimg src="launch-troubleshoot-gateway.png" width="50%" caption="Troubleshoot Gateway Traffic Filter Dialog">}}

The output can be useful in troubleshooting why a node and a peer are not connecting.

{{<tgimg src="troubleshoot-gateway-traffic.png" width="80%" caption="Example output of Troubleshoot Gateway tool">}}