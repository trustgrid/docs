---
Title: "Flow Logs"
linkTitle: "Flow Logs"
---

{{% pageinfo %}}
Flow logs are a record of all traffic that passes through a node. They can be used to troubleshoot connectivity issues, or to monitor traffic patterns.
{{% /pageinfo %}}

Flow logs for an individual node can be found under the `History` section of the node view. Click the `Flow Logs` tab to view the logs. For the entire organization across all nodes, navigate to `Operations`->`Flow Logs`.

By default, newest flows will be shown first.

![img](flow-logs-table.png)

{{<fields>}}
{{<field "Start Time" >}}
The time the flow started
{{</field >}}
{{<field "End Time" >}}
The time the flow ended
{{</field >}}
{{<field "Protocol" >}}
The protocol used
{{</field >}}
{{<field "Source Node" >}}
The node that initiated the flow
{{</field >}}
{{<field "Source IP" >}}
The IP address through which the node initiated the flow
{{</field >}}
{{<field "Source Port" >}}
The port through which the node initiated the flow
{{</field >}}
{{<field "Dest Node" >}}
The node that received the flow
{{</field >}}
{{<field "Dest IP" >}}
The IP address to which traffic was sent
{{</field >}}
{{<field "Dest Port" >}}
The port to which traffic was sent
{{</field >}}
{{<field "Recv Bytes" >}}
Bytes received at the source node
{{</field >}}
{{<field "Sent Bytes" >}}
Bytes sent from the source node
{{</field >}}
{{<field "TCP Flags" >}}
TCP Flags set during the flow:

- SYN - sync packet
- PSH - push packet
- ACK - ack packet
- URG - urgent packet
- FIN - finish packet
- RST - reset packet

{{</field >}}

{{</fields>}}

### Advanced Search

Flow logs can be filtered by any of the fields listed above. To filter by a field, click the `Advanced Search` button at the top right of the flow logs table.

In addition, the ordering can be changed so that oldest flows (or flows closest to the time range specified) can be shown first. There is a limit of 10,000 flows returned per search.

![img](advanced-search.png)

### Exporting Flow Logs

Trustgrid can export flow logs to any S3 compatible bucket. Exports are run every 24 hours.

On the Operations->Flow Logs page, configure credentials to export flows.

![img](s3-export.png)

{{<fields>}}
{{<field "S3 Access Key" >}}
The access key to use when making S3 API calls to the target bucket
{{</field >}}

{{<field "S3 Secret Key" >}}
The secret key to use when making S3 API calls to the target bucket
{{</field >}}

{{<field "S3 Bucket" >}}
The bucket name. You must own the S3 bucket and the access credentials must be allowed to write to the bucket.
{{</field >}}

{{<field "S3 Region" >}}
The region of the S3 bucket. This is used to determine the endpoint to use when making S3 API calls.
{{</field >}}

{{<field "S3 Prefix" >}}
The prefix to use when writing objects. If the prefix is, "prefix", then files will be written to s3://your-bucket/prefix/tg-flow-logs-YYYY-MM-DD.csv.gz
{{</field >}}
{{</fields>}}

After saving credentials, Trustgrid will attempt to write an empty file named `validation` to the bucket and prefix provided. Any errors will be displayed on the page.

#### Example IAM Policy for S3 Export

Below is an example IAM policy with the required permissions to push flow logs to your S3 bucket. Be sure you replace `example-flowlogs` in lines 13 & 14 with the name of your bucket.

<pre class="line-numbers language-json" data-line="13-14">
<code>{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": [
                "s3:PutObject",
                "s3:ListBucket",
                "s3:ListMultipartUploadParts"
            ],
            "Resource": [
                "arn:aws:s3:::example-flowlogs/*",
                "arn:aws:s3:::example-flowlogs"
            ]
        }
    ]
}
</code></pre>
