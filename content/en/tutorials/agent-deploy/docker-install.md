Follow the process below to run the agent in a Docker container:

{{<alert color="info">}}
Without any elevated privileges, Docker agents can demonstrate connectivity and direct layer 4 traffic. To allow the Docker agent to route IP pool traffic from the host machine, the container requires host networking and `NET_ADMIN` and `NET_RAW` capabilities. Add these flags to the Docker command: 

`-v /dev/net/tun:/dev/net/tun --network=host --cap-add=NET_ADMIN --cap-add=NET_RAW`

To persist logs and configuration locally, volumes can be attached to the container. The configuration files are stored in `/var/lib/trustgrid/agent`, and logs are stored in `/var/log/trustgrid/agent`. To attach volumes for those directories, append to the Docker command given: 

`-v /path/to/local/config/volume:/var/lib/trustgrid/agent -v /path/to/local/log/volume:/var/log/trustgrid/agent`

An example command with all flags and volumes attached would look like this:

{{< highlight bash >}}
docker run \
  -e AUTH_TOKEN=tgt-... \
  -v /path/to/local/config/volume:/var/lib/trustgrid/agent \
  -v /path/to/local/log/volume:/var/log/trustgrid/agent \
  -v /dev/net/tun:/dev/net/tun \
  --network=host \
  --cap-add=NET_ADMIN \
  --cap-add=NET_RAW \
  trustgrid/agent
{{< /highlight >}}

{{</alert>}}

1. From the [Nodes table]({{<ref "/docs/nodes">}}), click the `+ Add Agent` button to generate an agent token. {{<tgimg src="/img/agent/add-agent-button.png" width="30%" caption="Add Agent Button" >}} 

1. When prompted provide a name and choose the Virtual Network the agent will connect to. Then click `Add`. {{<tgimg src="/img/agent/name-agent.png" width="85%" caption="Name 'agent1' and assigned to 'default' virtual network">}}

1. Make sure the "Docker" tab is selected and then click the copy button to copy the install command to your clipboard. {{<tgimg src="/img/agent/docker-install-command.png" width="85%">}} {{<alert color="warning">}}The agent token and install command are only visible on this panel.  Once closed, it is not retrievable. Either keep the panel open until you've completed the below steps or copy the token to a secure location{{</alert>}}

1. Login to your Docker host as a user with permissions to run containers.

1. Paste the command from your clipboard and hit enter. The container image will download from Dockerhub and will start automatically. 

1. After it completes you can return to the portal and close the Add Agent panel. You should see the new agent listed in the Nodes table and online. {{<tgimg src="/img/agent/agent-online.png" width="65%" caption="agent1 shows as online">}}
