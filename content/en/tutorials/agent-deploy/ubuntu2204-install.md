1. From the [Nodes table]({{<ref "/docs/nodes">}}), click the `+ Add Agent` button to generate an agent token. {{<tgimg src="/img/agent/add-agent-button.png" width="30%" caption="Add Agent Button" >}} 

1. When prompted provide a name and choose the Virtual Network the agent will connect to. Then click `Add`. {{<tgimg src="/img/agent/name-agent.png" width="85%" caption="Name 'agent1' and assigned to 'default' virtual network">}}

1. Make sure the "Ubuntu Jammy" tab is selected and then click the copy button to copy the install command to your clipboard. {{<tgimg src="/img/agent/ubuntu-install-command.png" width="85%">}} {{<alert color="warning">}}The agent token and install command are only visible on this panel.  Once closed, it is not retrievable. Either keep the panel open until you've completed the below steps or copy the token to a secure location{{</alert>}}

1. Login to your Ubuntu instances as a user with sudo permissions.

1. Paste the command from your clipboard and hit enter. The install process will run automatically. {{<tgimg src="/img/agent/tg-agent-install.gif" width="95%" alt="Terminal window showing the Trustgrid agent being installed">}}

1. After it completes you can return to the portal and close the Add Agent panel. You should see the new agent listed in the Nodes table and online. {{<tgimg src="/img/agent/agent-online.png" width="65%" caption="agent1 shows as online">}}
