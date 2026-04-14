---
linkTitle: "ThousandEyes Agent Container"
title: "Deploying ThousandEyes Enterprise Agent Container on Trustgrid"
---

## Description
Running the ThousandEyes Enterprise Agent from remote locations on Trustgrid Nodes enables organizations to monitor network latency and performance across different regions, ensuring a consistent user experience. It helps identify issues and proactively address network problems before they impact users. This monitoring can go beyond [Trustgrid's native hop monitoring]({{<ref "tutorials/gateway-tools/monitoring-network-hops-to-peers">}}) to perform additional tests. For organizations already using ThousandEyes this can provide a single pane of glass view of network performance.

This guide walks through how to download the latest published container, push it to the private Trustgrid container repository for your organization and then configure an Trustgrid node or cluster to run the container. 

## Prerequisites
- You will need an existing ThousandEyes account
    - Within this account you will need to generate an account token to be used in the deployment of the containers.
- The host node will need sufficient resources to run the container including:
    - 1.5-2GB of free RAM under normal operating conditions
    - 3-5GB of free Disk Space
    - CPU use will depend on the number of test being performed but as a general recommendation a 2-core machine that operates under 40% CPU utilization under normal conditions should be suitable
- The container will use the host node’s WAN IP information to connect to the ThousandEyes cloud.  If behind a firewall ensure you follow [ThousandEyes recommended firewall configuration.](https://docs.thousandeyes.com/product-documentation/global-vantage-points/enterprise-agents/configuring/firewall-configuration-for-enterprise-agents)
    - Additionally the container regularly attempt to upgrade its base Ubuntu packages via the public Canonical APT repositories via port 80. There does not appear to be a static list of IPs for these repositories.  The container will run without but will log frequent warnings about failed upgrades.
- Trustgrid recommends running this container on node running versions newer than package version 1.5.20230302-1595

{{<alert color="warning">}} Note: ThousandEyes support does not officially support running their agent container on Trustgrid nodes because Trustgrid (like many other edge compute platforms) does not utilize the Docker runtime to run and manage containers.  However, Trustgrid has successfully deployed and run this container in several environments.
{{</alert>}}



## Container Image Management
There are two ways that the [ThousandEyes Enterprise Agent container published at hub.docker.com](https://hub.docker.com/r/thousandeyes/enterprise-agent) can be pulled down by a Trustgrid node to execute. 
### Pull Directly from hub.docker.com
In this method, the Trustgrid node will attempt to connect directly to hub.docker.com. 

Pros: 
- Quick and easy to setup

Cons: 
- Requires all nodes to have port 443 access to hub.docker.com.  This can be a challenge in highly secure environments. 
- Limited ability to control the actual image being run. For example, if you utilize thousandeyes/enterprise-agent:latest, then any time a new version is pushed up with that tag the nodes will pull an updated version that you may not have tested.

### Pull from private Trustgrid repository
Every customer organization is given a namespace in a [private container repository]({{<ref "docs/repositories">}}) hosted in the Trustgrid control plane, which nodes should already have access to.  In this method, you first pull down the image to a system with docker running, log into your organization's custom repository, tag the image in the correct namespace, and push the image back up.

Pros:
- More secure.
- Provides control over what image is available and used.

Cons: 
- Requires a third device with docker running to pull/tag/push the image.


## Process
### (Optional) Push Container Image to Trustgrid Repository
The steps below need to be performed on a machine with Docker installed and running.
{{<alert color="info">}} In the below `:latest` can be replaced by a specific version tag if desired.{{</alert>}}
1. Login to your [organization's private repository]({{<ref "docs/repositories#docker-login">}}).
1. run the command `docker pull thousandeyes/enterprise-agent:latest`
1. run the command `docker tag thousandeyes/enterprise-agent:latest docker.trustgrid.io/<orgname>.trustgrid.io/thousandeyes-enterprise-agent:latest` being sure to replace `<orgname>` with your organization's Trustgrid domain.
1. run the command `docker push docker.trustgrid.io/<orgname>.trustgrid.io/thousandeyes-enterprise-agent:latest`

### Create Required Volumes
The ThousandEyes container requires three persistent volumes.  
1. In the Trustgrid portal navigate to the node or cluster where you want to run the containers.
1. Select the Container Management → Volumes panel.
1. Use the **Add Volume** button to create the three volumes listed below.
- `te-agent-lib`
- `te-agent-logs`
- `te-browserbot-lib`

### Configure the Container
Navigate to Container Management → Containers
#### Create the Container
1. Click Add Container
    1. Give the container a name
    1. Leave Execution Type as “On Demand”
    1. Select the image
        - If you followed the above process to push the image to your private repository: 
            1. From the image drop down select the image you pushed to the Trustgrid repository above. `<orgname>.trustgrid.io/thousandeyes-enterprise-agent`
            1. From the tag drop down select **latest** (or the version tag you applied)
        - If you chose to pull directly from hub.docker.com:
            1. Enter `thousandeyes/enterprise-agent` in the image section
            1. Leave `latest` in the tag field unless you want to specify a different version tag published
            {{<tgimg src="te-docker-hub.png" width="40%">}}
    1. Click Save
1. In the Containers table, click the link to open the properties of the newly defined container and proceed to the steps below to complete the configuration

#### Configure Environmental Variables
In the Environmental Variables Section
1. Add  a variable named `TEAGENT_ACCOUNT_TOKEN` The value will need to be the [account group token generated within your ThousandEyes portal](https://docs.thousandeyes.com/product-documentation/global-vantage-points/enterprise-agents/installing/where-can-i-get-the-account-group-token)
1. Add a variable named `TEAGENT_INET` with a value of 4

(Optional) the `HOSTNAME` variable can be defined to determine the name used to register the agent in the ThousandEyes system. By default it will register using the the hostname of the Trustgrid node. 

{{<tgimg src="te-env-vars.png" width="80%" caption="Environmental Variables">}}


#### Add Volume Mounts
You will need to add the following Mounts to the container.

|Type|Source|Destination|
|----|------|-----------|
|Volume|te-agent-lib|/var/lib/te-agent|
|Volume|te-agent-logs|/var/log/agent|
|Volume|te-browserbot-lib|/var/lib/te-browserbot|

1. Click Add Entry
1. Select Volume as the type
1. Select the Source Volume you created earlier
1. Copy and paste the destination from the table above
1. Click the green arrow to save the entry
1. Repeat with the other two mounts

Click the Save button

{{<tgimg src="te-vol-mounts.png" width="80%" caption="Volume Mounts">}}

#### Set Resource Limits
Resource Limits are intended to reduce the risk of the container consuming CPU, Memory or IO resources to the point they impact other services on the node such as VPN traffic. 
{{<alert color="warning">}}
The below settings are based on the resource available on a **dual core device with 4GB of total memory** with no other containers running and moderate levels of VPN traffic.  Other device types may warrant different settings
{{</alert>}}

1. Enter `40` for the CPU Max %
1. Enter `1536` for the Memory Max (granting up to 1.5G of RAM)
1. Enter `1024` for the Memory High (this will attempt to limit the container to 1G of RAM)
1. Leave the IO and Limits settings blank.
1. Click Save

{{<tgimg src="te-resource-limits.png" width="50%" caption="Resource Limits for a device with 4G of available memory">}}

#### Linux Capabilities
The ThousandEyes Agent container needs the NET_ADMIN and SYS_ADMIN capabilities to function as expected. 

1. Click Add Entry
1. From the list select NET_ADMIN and click the green checkbox to save
1. Repeat to add SYS_ADMIN
1. Click the Save button

### Test Container Functionality
On Demand containers only run when specifically started via the Trustgrid control plane. This is idea for testing the configuration before setting the containers to run as a service

1. Navigate to the new container’s Overview page
    1. If the container was defined at the Cluster level, navigate to one of the member nodes' Container Management → Containers page and click on the container.  Usually best to start with the standby member if the cluster is in production
1. Click the “Start” button
1. A new window should pop-up (if not check to see if you browser is blocking pop-ups).  The container will:
    1. The first time the container is run it will download the container image. The ThousandEyes agent container is ~1.2G at this time so it can take a while depending on available. Subsequent starts will only check to see if new layers are available. 
    1. Start the container and display the standard output
1. After the container has started, log into your ThousandEyes account and confirm you see the Enterprise Agent listed.
1. Stop the running container by clicking the Terminate button.
1. (Optionally) If working with a cluster repeat the above with the steps with the other cluster member.

### Set the Container to Service 
Once satisfied the containers are running as expected you’ll want to set the container to run as a Service so that it stays running at all times

1. Navigate to Container Management → Containers and select the ThousandEyes container
1. If dealing with clustered nodes do this from the Cluster page
1. On the Overview page change the Execution Type from “On Demand” to “Service {{<tgimg src="te-exec-type.png" width="30%">}}
1. Click Save
1. After a minute or two you should be able to click the Logs button to confirm the container started