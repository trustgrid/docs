---
categories: ["node"]
tags: ["data store", "containers", "s3", "virtual machines"]
title: "Data Store"
linkTitle: "Data Store"
---

{{% pageinfo %}}
The data store allows artifacts (i.e. container volumes, virtual machine images) to be synced from/to a node.  The data store view is located under Compute > Data Store.
{{% /pageinfo %}}

### File Explorer

Allows a user to navigate, sync, create directories, delete artifacts from the data store.

![img](file_explorer.png)

- ![img](home.png) - Takes you to the Root of the directory hierarchy.
- ![img](back.png) - Takes you up one directory from your current location in the directory hierarchy.
- ![img](refresh.png) - Refreshes the directory hierarchy to get the latest state of the data store.
- ![img](create_dir.png) - Create a new sub directory. See [Directory Creation]({{<ref "docs/nodes/datastore/mkdir" >}}) for more information
- ![img](sync.png) - Transfer a file to the data store. See [File Sync]({{<ref "docs/nodes/datastore/sync" >}}) for more information
- ![img](upload.png) - Upload a file from the data store. See [File Sync]({{<ref "docs/nodes/datastore/sync" >}}) for more information
- ![img](delete.png) - Delete a file from the data store.


### Recent Tasks

Allows a user to view the status of any long running tasks (i.e. Transfer file from Amazon S3, Upload file to HTTP endpoint, etc).  Any task that is executing can be cancelled.  Completed tasks can be cleared manually or they will be automatically removed 6 hours after completion time.

![img](recent_tasks.png)



