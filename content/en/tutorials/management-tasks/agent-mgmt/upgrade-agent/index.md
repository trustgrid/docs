---
title: Upgrade Trustgrid Agent Nodes
linkTitle: Upgrade Agent Nodes
description: Learn how to upgrade Trustgrid Agent nodes to get the latest features and security upgrades.
---
## How an Agent Node Upgrades

The Trustgrid agent upgrade process will depend on the operating system the agent runs on. For the most part the process will include:
1. Gaining console access to the agent host system either via SSH or local console
1. Running a package management tool to update the agent software package
1. Restarting the agent service if not done automatically by the package manager

{{<alert color="warning">}}Restarting the agent service is important to ensure it loads the latest configuration and starts with the updated code but will disrupt traffic across the node.{{</alert>}}

## Agent Upgrade Process on Linux

### Upgrade Ubuntu-based Agent

1. Gain console access to the agent host system via SSH or local console with a user with sudo access to run the below commands.
1. Run the command {{<codeblock>}}sudo apt update && sudo DEBIAN_FRONTEND=noninteractive apt-get install -y tg-agent{{</codeblock>}}
1. The above command will refresh the package list and install the latest version of the tg-agent package. This will automatically restart the tg-agent service.