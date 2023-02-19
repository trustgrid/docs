---
title: "Volumes"
---

Volumes allow containers to persist data between executions.

![Container Volumes](volumes.png)

## Create a Volume

![Add Volume Modal](add-volume.png)

Volumes have the following fields:

{{<fields>}}
{{<field "Name">}}The name of the volume.{{</field>}}
{{<field "Encrypted">}}Whether or not to encrypt the volume. Encrypted volumes The type of volume.{{</field>}}
{{<field "Location">}}Optionally, specify a location in the node [data store]({{<ref "datastore">}}).{{</field>}}
{{</fields>}}
