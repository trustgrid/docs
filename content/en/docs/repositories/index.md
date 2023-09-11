---
title: "Repositories"
date: 1-27-2023
weight: 17
---

## Summary 

Trustgrid provides a fully managed private container registry for each customer account. This allows customers to deploy [containerized applications securely to their Trustgrid node appliances.]({{<ref "/docs/nodes/containers">}}) 

## Container Namespace

Each customer is isolated to their own container namespace that is listed at the bottom of the Repositories page. 

To push a container to your private registry it must be named with your namespace as a prefix when tagged. 

{{<tgimg src="docker-namespace.png" caption="Example Trustgrid repository namespace" width="40%">}}


## List View
 Uploaded container images may be viewed by navigating to Repositories in the portal. 

{{<tgimg src="docker-repo-list.png" caption="Repositories list view" width="80%">}}

From this view you can se all containers uploaded to your namespace. 

### Container Repository View

Each repository can be managed. Clicking into a container will show its uploaded tags and digests, as well as the URI that can be used to pull the container.
{{<tgimg src="docker-repo-example.png" width="80%" caption="Example repository for nginx container with two tag versions">}}

### Deleting Container Repositories
Next to each container repository name is a box that can be selected. This allows deleting the entire repository and all its tags with a single action.
1. Select the desired container repository.
1. From the Actions dropdown select **Delete**. 
1. When prompted enter the repository name and click **Confirm**. {{<tgimg src="docker-repo-delete-prompt.png" caption="Prompt to delete an example nginx repository" width="50%">}}

### Deleting Specific Container Tags
1. Navigate into the desired repository.
1. Select the desired tag version.
1. From the Actions dropdown select **Delete**. 
1. When prompted enter the tag name and click **Confirm**. {{<tgimg src="docker-tag-delete-prompt.png" caption="Prompt to delete an example nginx tag version" width="50%">}}



## Authentication

### Docker Login
In order to push to or pull containers from the private registry, your local docker client must first authenticate with your credentials. 

To authenticate with the registry, use the command provided in the **Docker Login** section at the bottom of the repositories page. 
{{<tgimg src="docker-login-button.png" caption="Docker Login with copy button" width="80%">}}

Use the copy button to copy the docker login command and paste it into your terminal to authenticate your docker client. This will cache your credentials locally for approximately 24 hours.

{{<tgimg src="docker-login-example.png" caption="Example docker login command" width="80%">}}

### Node Container Authentication

Trustgrid node appliances authenticate automatically with the Trustgrid container registry and can pull any image uploaded to the registry. All communication between the node appliance and the registry occurs using the [Trustgrid control plane networks and ports]({{<ref "/help-center/kb/site-requirements#trustgrid-control-plane">}})

## Example Usage
In the below example we will show how to pull down a container image (Alpine Linux) from the public hub.docker.com registry and then push it to the Trustgrid private registry under our namespace.

1. First pull the Alpine image from Docker Hub with the command `docker pull alpine` {{<tgimg src="docker-example-pull.png" caption="Pull Alpine image from Docker Hub" width="75%">}}. {{<alert color="info">}} Note the tag is automatically set to `latest` by Docker. If desired you can manually specify a specific tag to get a specific version. {{</alert>}}
1. Tag the image with your namespace prefix. `docker tag alpine:latest docker.trustgrid.io/namespace.trustgrid.io/alpine:latest` make sure you replace `namespace.trustgrid.io` with your actual namespace. {{<tgimg src="docker-example-tag.png" caption="Tag image with namespace prefix" width="90%">}}
1. If you haven't already, Authenticate your Docker client with the Trustgrid registry using the [Docker login command provided on the Repositories page]({{<ref "#docker-login">}})
1. Push the tagged image to the private registry. `docker push docker.trustgrid.io/namespace.trustgrid.io/alpine:latest` {{<tgimg src="docker-example-push.png" caption="Push tagged image to private registry" width="75%">}}
1. Back on the Trustgrid portal, navigate to Repositories and you should see the pushed image listed. {{<tgimg src="docker-example-repo.png" caption="Pushed image listed in portal" width="80%">}}