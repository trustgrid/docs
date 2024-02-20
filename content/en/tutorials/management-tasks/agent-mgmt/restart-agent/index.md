---
title: Restart Agent Node Service
linkTitle: Restart Agent
description: Learn how to restart the Trustgrid Agent service
---

Restarting the agent service is sometimes necessary to reload configuration changes or restart the agent process. The steps to restart the agent service depends on the operating system.

{{<alert color="warning">}}Restarting the agent service will disrupt traffic across the node.{{</alert>}}

## Restarting Linux Agent
### Restart Ubuntu Agents
1. Gain console access to the agent host system via SSH or local console with a user with sudo access to run the below commands.
1. Run the command {{<codeblock>}}sudo systemctl restart tg-agent{{</codeblock>}}