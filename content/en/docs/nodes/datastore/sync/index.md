---
categories: ["data store"]
tags: ["containers", "s3", "virtual machines"]
title: "File Sync"
linkTitle: "File Sync"
---

{{% pageinfo %}}
File syncing allows a user to transfer files to the data store as well as upload files from the data store.
{{% /pageinfo %}}

## Transfer Files

A user can populate the data store by transferring files from an Amazon S3 bucket or from an HTTP endpoint.

![img](file_explorer_toolbar_sync.png)

{{<alert color="info">}}
The destination where the file will be stored can be changed by selecting a particular directory from the [File Explorer]({{<ref "docs/nodes/datastore" >}}) side bar.
{{</alert>}}


#### Amazon S3 Transfer

A user can select to transfer a file from an S3 URI.

![img](amazon_s3_transfer.png)

##### Sync File

![img](s3_transfer.png)

{{<fields>}}
{{< field "S3 URI" >}}
The S3 URI for the location of the file to transfer.
{{< /field >}}
{{< field "Datastore Destination" >}}
The data store location where the file will be transferred to.
{{< /field >}}
{{</fields>}}

##### Configure Authentication

In order to be able to utilize the S3 File Sync option, valid AWS credentials need to be configured on the node for the applicable S3 buckets.

##### Configure Bandwidth Throttling

The S3 File Sync service will always try to utilize as much bandwidth as possible to perform the transfer of the file.  If the amount of bandwidth needs to be limited, a user can configure the amount of bandwidth allowed to be utilized in Kbps/Mbps/Gbps.

![img](s3_bandwidth.png)

{{<fields>}}
{{< field "Max Concurrent Requests" >}}
The maximum number of concurrent requests allowed.
{{< /field >}}
{{< field "Max Bandwidth" >}}
The maximum bandwidth that will be consumed for uploading and downloading data to and from Amazon S3.
{{< /field >}}
{{< field "Bandwidth Rate" >}}
Rate associated to the Max Bandwidth specified.  Supported values are Kbps, Mbps, and Gbps.
{{< /field >}}
{{</fields>}}

{{<alert color="info">}}
In general, it is recommended to first use Max Concurrent Requests to lower transfers to the desired bandwidth consumption. The Max Bandwidth setting would then be used to further limit bandwidth consumption if setting Max Concurrent Requests is unable to lower bandwidth consumption to the desired rate.
{{</alert>}}


#### HTTP Endpoint Transfer

A user can select to transfer a file from an HTTP endpoint.  The endpoint has to be unauthenticated since the node currently doesn't support providing credentials.

![img](http_transfer.png)

##### Sync File

![img](http_transfer_config.png)

{{<fields>}}
{{< field "HTTP URI" >}}
The HTTP URI for the location of the file to transfer.
{{< /field >}}
{{< field "Datastore File Name" >}}
The transferred file will be saved with this name under the datastore destination.
{{< /field >}}
{{< field "Hash (Optional)" >}}
Optional fields used to validate the integrity of the transferred file.
{{< /field >}}
{{< field "Hashing Algorithm" >}}
If a hash is specified, this specifies the hashing algorithm used during validation of the file.
{{< /field >}}
{{</fields>}}

***

Upload Files

A user can extract files from the datastore by uploading them to an Amazon S3 bucket or to an HTTP endpoint.

![img](file_upload.png)

{{<alert color="info">}}
A file needs to be selected first before the File Upload button gets enabled on the [File Explorer]({{<ref "docs/nodes/datastore" >}}) menu bar.
{{</alert>}}

#### Amazon S3 Upload

A user can select to transfer a file to an S3 URI.

![img](s3_upload.png)

##### File Upload

![img](s3_upload_config.png)

{{<fields>}}
{{< field "S3 Bucket URI" >}}
The S3 bucket URI for the location where the file will be uploaded.
{{< /field >}}
{{< field "Datastore File" >}}
The data store file to upload.
{{< /field >}}
{{</fields>}}

{{<alert color="info">}}
For configuration of AWS authentication and Bandwidth Throttling refer to [Configure Authentication]({{<ref "#configure-authentication" >}}) and [Configure Bandwidth Throttling]({{<ref "#configure-bandwidth-throttling" >}})
{{</alert>}}

#### HTTP Endpoint Upload

A user can select to upload a file to an HTTP endpoint.  The endpoint has to be unauthenticated since the node currently doesn't support providing credentials.

![img](http_upload.png)

##### File Upload

![img](http_upload_config.png)

{{<fields>}}
{{< field "HTTP URI" >}}
The HTTP URI for the location where the file will be uploaded.
{{< /field >}}
{{< field "Datastore File" >}}
The data store file to upload.
{{< /field >}}
{{< field "Use Multipart" >}}
Used to denote if multipart should be used when uploading the file. True for use multipart and false otherwise.
{{< /field >}}
{{</fields>}}

{{<alert color="info">}}
The node uses the POST method to perform the HTTP upload
{{</alert>}}