---
linkTitle: "ThousandEyes Agent"
title: "ThousandEyes Enterprise Agent Container"
---

## Description
Running the ThousandEyes Enterprise Agent from remote locations on Trustgrid Nodes enables organizations to monitor network latency and performance across different regions, ensuring a consistent user experience. It helps identify issues and proactively address network problems before they impact users. This monitoring can go beyond [Trustgrid's native hop monitoring]({{<ref "tutorials/gateway-tools/monitoring-network-hops-to-peers">}}) to perform additional tests.

This guide walks through how to download the latest published container, push it to the private Trustgrid container repository for your organization and then configure an Trustgrid node or cluster to run the container. 

## Prerequisites
- You will need an existing ThousandEyes account
    - Within this account you will need to generate an account token to be used in the deployment of the containers.
- The host node will need sufficient resources to run the container including:
    - 1.5-2GB of free RAM under normal operating conditions.
    - 3-5GB of free Disk Space.
    - CPU use will depend on the number of test being performed but as a general recommendation a 2-core machine that operates under 40% CPU utilization under normal conditions should be suitable.
- The container will use the host node’s WAN IP information to connect to the ThousandEyes cloud.  If behind a firewall ensure you follow [ThousandEyes recommended firewall configuration.](https://docs.thousandeyes.com/product-documentation/global-vantage-points/enterprise-agents/configuring/firewall-configuration-for-enterprise-agents)
- Additionally the container regularly attempt to upgrade its base Ubuntu packages via the public Canoncal APT repositories via port 80. There does not appear to be a static list of IPs for these repositories.  The container will run without but will log frequent warnings about failed upgrades.
- Trustgrid recommends running this container on node running versions newer than package version 1.5.20230302-1595.
- 

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
- Limited ability to control the actual image being run. If you utilize thousandeyes/enterprise-agent:latest, then any time a new version is pushed up with that tag the nodes will pull an updated version that you may not have tested.

### Pull from private Trustgrid repository
Every customer organization is given a namespace in a private container repository hosted in the Trustgrid control plane.


## Process