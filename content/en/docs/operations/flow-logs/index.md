---
title: "Organization Flow Logs"
linkTitle: "Flow Logs"

---

{{<pageinfo>}}This page shows Flow Logs from all nodes in an organization. {{</pageinfo>}}

See this page for more info on [Flow Logs]({{<ref "/help-center/flow-logs">}})

## Flow Log Export

Trustgrid can export flow logs to any S3 bucket. Set your bucket name and apply the policy provided, and then Trustgrid will configure replication for your flow logs into your bucket.

{{<tgimg src=s3-export.png width="50%" caption="Fields for configuring S3 replication">}}

{{<fields>}}
{{<field "AWS Account ID">}}This is the AWS Account ID that contains the bucket you wish to have Flow Logs replicated to.
{{</field>}}
{{<field "S3 Bucket" >}}
The bucket name. You must own the S3 bucket and apply the policy in AWS to allow Trustgrid to replicate files to the bucket.
{{</field >}}
{{</fields>}}

{{<alert color="info">}}After submitting the above information, Trustgrid will complete their required configuration within 2 business days {{</alert>}}

### Requirements

#### Enable Bucket Versioning
You must [enable bucket versioning](https://docs.aws.amazon.com/AmazonS3/latest/userguide/manage-versioning-examples.html) on the target S3 bucket for replication to work.

#### Configure Bucket Policy

Below is an example Bucket Policy with the required permissions to push flow logs to your S3 bucket. Be sure you replace `example-flowlogs` in lines 19 & 20 with the name of your bucket and [add this bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/userguide/add-bucket-policy.html) to your S3 bucket.

<pre class="line-numbers language-json" data-line="19-20">
<code>{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "allow-tg-writes",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::079972220921:root"
      },
      "Action": [
        "s3:ReplicateObject",
        "s3:ReplicateDelete",
        "s3:ReplicateTags",
        "s3:List*",
        "s3:GetBucketVersioning",
        "s3:PutBucketVersioning"
      ],
      "Resource": [
        "arn:aws:s3:::example-flowlogs",
        "arn:aws:s3:::example-flowlogs/*"
      ]
    }
  ]
}
</code></pre>