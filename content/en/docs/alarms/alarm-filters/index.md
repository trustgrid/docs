---
Title: "Alarm Filters"
Tags: ["alarms", "alarm filters"]
---

{{% pageinfo %}}
Alarm Filters are used to determine which [events]({{<ref "events" >}}) trigger notifications and to define which [channels]({{<ref "channels" >}}) should receive those notifications.
{{% /pageinfo %}}

Alarm filters have the following fields:

{{<fields>}}
{{<field "Name">}}The name should be unique.{{</field>}}
{{<field "Description">}}The description is displayed in the alarm filters table.{{</field>}}
{{<field "Enabled">}}An alarm filter must be enabled for matching [event]({{<ref "events" >}}) to be sent to the selected [channel]({{<ref "channels" >}}). Deselecting the check box can be handy if you wish to suppress a specific type of [alarm]({{<ref "docs/alarms" >}}).{{</field>}}
{{<field "Channels">}}This section determines which [channels]({{<ref "channels" >}}) matching [alarms]({{<ref "docs/alarms" >}}) will be sent to.{{</field>}}
{{<field "Criteria">}}
The criteria determine which [events]({{<ref "events" >}}) will match the filter. These conditions can be set as:
* All (default) - All specified criteria must be true to match. Equivalent to a boolean AND condition.
* Any - Only one criteria must be true to match. Equivalent to a boolean OR condition.
* None - The specified criteria must be false to match. Equivalent to a boolean NOT of the criteria ANDed together.
{{</field>}}
{{<field "Node Name">}}
The "Node Name" criteria llows you to select one or more specific [node]({{<ref "docs/nodes" >}}) names. Note, even if the filter is set to `All`, the filter will match any of the selected [node]({{<ref "docs/nodes" >}}) names is associated with the [event]({{<ref "events" >}}).
{{</field>}}
{{<field "Event Type">}}
The ["Event Type"]({{<ref "event-types" >}}) criteria determines which [events]({{<ref "events" >}}) will match the filter. Note, even if the filter is set to `All`, the filter will match any of the selected [event types]({{<ref "event-types" >}}).
{{</field>}}
{{<field "Tag Matches">}}
The "Tag Matches" criteria allows you to use [tag]({{<ref "/docs/nodes/shared/tags" >}}) name/value pairs to determine if the filter should match [events]({{<ref "events" >}}). For examples, you may what production devices to send to a high priority [channel]({{<ref "channels" >}}) such as PagerDuty or OpsGenie. If your [nodes]({{<ref "docs/nodes" >}}) have a tag to indicating “prod_status=production”, you can select that name/value pair from the list to properly filter your [alarms]({{<ref "docs/alarms" >}}).

{{<tgimg src="tag-matches.png" width="50%">}}
{{</field>}}
{{<field "Tag Match Any/All">}} You can choose if multiple tags must match ALL or ANY of the selected tag criteria for the filter to match.  For example: 
- `any` would cover a scenario where you want to match say `Environment=Prod` **OR** `Environment=Test`.  
- `all` would cover if you wanted a filter to match something like `Environment=Prod` **AND** `Region=EAST`

{{<tgimg src="tag-matches-option.png" width="60%">}}
{{</field>}}
{{<field "Severity">}}
Each [event]({{<ref "events" >}}) type has a severity level associated with it. This filter will match any [event]({{<ref "events" >}}) with the selected severity type or higher. This is the only mandatory criteria.

The severity levels are:

1. INFO
1. WARNING
1. ERROR
1. CRITICAL

For example, if you select the severity level of WARNING the filter will match WARNING, ERROR and CRITICAL [events]({{<ref "events" >}}).

Some events have a corresponding [event]({{<ref "events" >}}) that will automatically resolve the alert in the portal and in some [channels]({{<ref "channels" >}}) such as PagerDuty. The corresponding event may have a different severity level, so make sure you select the lower severity for the criteria. e.g. [Node]({{<ref "docs/nodes" >}}) Disconnect is a WARNING but [Node]({{<ref "docs/nodes" >}}) Connect which resolves it is only INFO. So you’d need to select both [Event Types]({{<ref "event-types" >}}) and set the severity to INFO.

{{</field>}}
{{<field "Contains Text">}}
This field will accept any single string of text to match to the contents of an [event]({{<ref "events" >}}). For example, if all your gateways include `-gw` in the name you could enter that without quotes in the field and it would match any [event]({{<ref "events" >}}) that includes that text in the event payload. This criteria can also be used if there is another aspect of the node included in the [event]({{<ref "events" >}}) payload that doesn’t match the criteria above. To see the entire payload of an [event]({{<ref "events" >}}) configure a less specific payload and send to an email [channel]({{<ref "channels" >}}) to see the JSON.

The [event]({{<ref "events" >}}) payload includes the [node’s]({{<ref "docs/nodes" >}}) unique identifier (UID) which is a string of generated text and numbers. If your "Contains Text" criteria is too short, there is a chance a [node]({{<ref "docs/nodes" >}}) UID will also match unexpectedly.
{{</field>}}
{{<field "CEL Expression">}}
CEL expressions allow logical expressions that evaluate to `true` or `false` to determine if a filter should match an alarm. See [here]({{<ref "alarm-filters#cel-expressions">}}) for a detailed explanation.
{{</field>}}
{{</fields>}}

## CEL Expressions

CEL (Common Expression Language) is a simple expression language that allows more complex tests than simple equality checks.

When an alarm filter has a CEL expression set, it will be compiled when saved. If the compilation fails, a validation error will appear at the top of the page.

CEL expressions allow numerical comparisons, arithmetic, boolean operators, regular expressions, string matching, presence testing and list evaluation. 

Events are provided inside a `ctx` object. 

{{< highlight json >}}
{
    "details": {
        "alertId": "40e4a030-440a-4703-b33a-172416da4be2"
    },
    "domain": "demo.dev.trustgrid.io",
    "eventType": "Data Plane Disruption",
    "expires": 1699234335,
    "level": "WARNING",
    "message": "Node demo-node via Internet path abnormally disconnected",
    "nodeId": "ccd5a29e-fdc0-43d6-9408-b4184100287e",
    "nodeName": "demo-node",
    "node": {
        "uid": "59838ae6-a2b2-4c45-b7be-9378f0b265f5",
        "org": "aad89024-5927-4ebd-97e2-3cc605c1da5f",
        "domain": "dev.dev.trustgrid.io",
        "fqdn": "demo-node.demo.dev.trustgrid.io",
        "lastip": "64.17.3.164",
        "last_connect": 1699158287000,
        "name": "demo-node",
        "state": "ACTIVE",
        "cluster": "",
        "tags": {
            "autoupdate": "true",
        },
        "online": true,
        "shadow": {
            "reported": {
            "nic.ens160.duplex": "full",
            "node-core.version": "20231103-171711.d16963a",
            "node.upgrade.state": "COMPLETED",
            "repoConnectivity": "true",
            "dnsResolution": "healthy",
            "nic.ens160.mac": "00:50:56:8e:8a:03",
            "ztna-enabled": "true",
            "nic.ens192.mtu": "1500",
            "profile.name": "default",
            "nic.ens192.speed": "10000",
            "ssh.local": "false",
            "os.distro.id": "ubuntu",
            "nic.ens192.dhcp": "false",
            "netplan.saved": "true",
            "nic.ens192.ip": "10.20.10.50/24",
            "nic.ens192.duplex": "full",
            "nic.ens160.speed": "10000",
            "publishTime": 1699210399513,
            "package.version": "1.5.20231103-1880",
            "os.distro.version": "18.04.3 LTS (Bionic Beaver)",
            "domain.info.lastUpdate": 1699145917,
            "nic.ens192.gateway": "10.20.10.1",
            "os.arch": "amd64",
            "updateTime.enabled": "true",
            "nic.ens160.dns1": "172.16.11.4",
            "nic.ens160.mtu": "1500",
            "nic.ens160.ip": "172.16.22.50/24",
            "nic.ens160.gateway": "172.16.22.1",
            "nic.ens192.mac": "00:50:56:8e:c9:74",
            "version": 1699145622,
            "kvm-enabled": "false",
            "tpm.enabled": "false",
            "node.upgrade.completed.tstamp": 1699034170,
            "startup.error": "true",
            "nic.ens160.dhcp": "false"
            }
        },
        "device": {
            "mac": "00:50:56:8E:AA:28",
            "model": "esx",
            "vendor": "vmware"
        },
        "location": {
            "continent_name": "North America",
            "zip": "80301",
            "calling_code": null,
            "city": "Boulder",
            "ip": "64.17.3.164",
            "latitude": 40.04801940917969,
            "continent_code": "NA",
            "type": "ipv4",
            "country_code": "US",
            "country_flag_emoji_unicode": null,
            "country_name": "United States",
            "is_eu": false,
            "connection": {
                "asn": 27325,
                "isp": "Zcolo"
            },
            "country_flag_emoji": null,
            "location": {
                "Languages": [
                    {
                    "name": "English",
                    "native": "English",
                    "code": "en"
                    }
                ],
                "capital": "Washington D.C.",
                "geoname_id": 5574991
            },
            "region_name": "Colorado",
            "country_flag": null,
            "longitude": -105.20680236816406,
            "region_code": "CO"
        },
        "type": "Node",
        "tgrn": "tgrn:tg::nodes:node/59838ae6-a2b2-4c45-b7be-9378f0b265f5",
        "created_at": 1552940922,
        "config": {
            "cluster": {
                "master": true
            },
            "gateway": {
                "clients": [],
                "monitorHops": true,
                "maxmbps": 1001,
                "cert": "proxy.dev.trustgrid.io",
                "type": "private",
                "connectToPublic": true,
                "udpPort": 8443,
                "enabled": false,
                "master": false,
                "port": 8442,
                "paths": [],
                "udpEnabled": true,
                "host": "12.244.52.245",
                "maxClientWriteMbps": 1000
            },
            "snmp": {
                "port": 161,
                "interface": "ens160",
                "authProtocol": "SHA",
                "enabled": true,
                "privacyProtocol": "DES",
                "engineId": "7779cf92165b42f380fc9c93c",
                "username": "myuser"
            }
        }
    },
    "orgId": "aad89024-5927-4ebd-97e2-3cc605c1da5",
    "receivedTime": 1699147935,
    "subject": "Node",
    "timestamp": 1699147935,
    "_ct": "2023-11-05T01:32:15.471Z",
    "_md": "2023-11-05T01:32:15.471Z"
}
{{< /highlight >}}

Nested values can be referenced using a `.`. 

Some common tests include:

* Check if a node is a gateway: `has(ctx.node.config.gateway) && ctx.node.config.gateway.enabled`
* Check if a node has production in the name: `ctx.node.name.contains("prod")`
* Check if a node is not clustered: `!has(ctx.node.cluster) || ctx.node.cluster == ""`
* Check if a node is in Texas or Colorado: `ctx.node.location.region_code == "TX" || ctx.node.location.region_code == "CO"`
* Check if a node is a virtual machine: `ctx.node.device.vendor == "vmware"`
* Check if a node is up to date: `ctx.node.shadow.reported["package.version"] >= "1.5.20231103-1880"`

The full CEL definition can be found at [GitHub](https://github.com/google/cel-spec/blob/master/doc/langdef.md). 

You can use this [CEL playground](https://playcel.undistro.io/?content=H4sIAAAAAAAAA51YWY%2FbNhD%2BK4IfegARo%2Fvw28a7aAOk20WzaYDWhUFTtK1GIlWSctYt9r93JEqWbNNHuxusHfGb4XCOb4b6Z5JhhSfTyXxC1Mt8MrX%2BmTMLfuaTjCqcF3L0TD%2FHBRXqfdY8n08ChwbY8R07CBxsB7Hj20vfx7Ybe4EbZThYUm8%2B0eKvb%2FaqeYlzpjVktOQoo1ukRC3VWuQZyvl8ssfSLWXqeVdRDb8Hc62nAjNq3edS1JXKORvDX6pc0NZqN0pTzw98P9yvFqCt0Io%2B3%2F3y%2BP7xh5FoSaXE626fR55Rq7HNZs23bY6t90xRwaiyKqw2Fl4yLkpcFDsryyXhjFGiaDbS1wj2fiIkC7GXUnuVEXCWn0V2GjiJvQzcJHAdx0tieiT6iEs6uKg14whxEpo677YL08RPMI1s7C09OyBBaC%2FjJbVTP05WztKLwlW4V6ZluVhrWYyzJHW8wA5TL7YDuszsNKae7RMSOSFxMxyujmQP47k9H04NX%2F2VsaOToctp0EUPS5VXWjIKkBsjH7lRYEAtunj0aeCGCXjYcZwDJLvgYY2QCqsOcjd7fv%2Frw9E6KcBYKjTiaE3h9XHxdAVUK15X2V4znPhw39fDwLAiZy20AR6at8EZ%2F2rcRNCKizYfDYssJ4gy6UYOyuqqoC%2FakFVdFAeGDKlmEy4o2lIh23Jr0J7j%2Ba4LBe%2FG8OuiDBwd%2BfiMAlRXa4Hhc%2BTS2c8%2FPX14eH64N8g09s90EPNtrnZnXKXBGZO%2FUMmLWu2t21BcqM3OZM1w%2BBITjXacaQj%2FomlCpwmeOr5B7m%2FFsE0ZXhY0u2hNv0HqoVLVGuqGjmOAVoKv8oKicSaucF2oy2plRXsbgDuMmqXcoIIT3LHdChfSZCuXCNhLCY566qiXNWusvrR%2FtiHVZb1AkxWQNJJ4e7u3%2Btp2HeQ5CP6GzlsvuGLKDfk7BPwGx1X1ssjl5jnXEWmbCGR5moaufwrG5Au0jMPCcFGIhuJIEtMug98PRRPkBMi3Pjx%2FtL57B49zYr2j4EPxvSntW9pFOVtx1LDepz2ptJwXhKkbX%2FLdGtBf8e7I6a7ZXCzIpmsQZRaZoqI5rXEc%2Bi9l0pAQfHRWxB5QOgI6uRT3pnSvVNYIuk8rrdvzruYVCB36Zi9pcs643k2EQtJpbNptFPkuXpHnnaC%2BbMtD1jlXcqoq0U3AAzomvIQCgkaBFDBzWfXWOH7gxs4prSgsVF0hKgQXN4f3mC4Osa%2FnOh%2BMAzk5nnL0ksHRD9O7u6mXGEwp4cAdD1L5YowEy%2FrjbMuvWBzYeGhUQ6p9mzkxC8YOBc2aqcXA6I%2FQhzfWXUlFTkzd8e8%2BPxOYo035BSQOE8AaZho98jGguVPQvke%2B43WRNTPJCeba7NQdEI6nar1V4AAfJY4LsypwSRqlp%2FvuT9ybBye%2BM6XnfoTPq61pX8Kh8YjdSM%2BnjxdgqwKvFzC2%2FZkvasi1i77pRIaYfGI5pLz1sZlFpMlTckFbfmkT1nTodjAxZoGGYNmuebHvhW9MgFx20fiN8IKf1MQN5z573gspqgEfMFvX0Lfa%2BfT3U0TzYxDUwoMXH9i6aZUn%2FhtjYXq7GT1EnrJjh%2Bwdc%2Fr4D6N%2FCa5y1Y8%2Fn7HcQAkpzqx7NENGG%2BaTNeXN4RZ6FArDOABWvhYXQdfg6lFuzSCcQK38Su5eiB5b7wvQdp1mmIgSmCeiBArWic5ZMHhv9vN5%2BhrK8PH0ugP3rm4Qab5N1Xo6bZqFbP%2B%2B%2FR%2BXSiIoVFi2wPomFoZewyQHXa6tplW%2BNvPpcL8yJnKJ%2B%2BW2C10J1ainG7WRIgci00VhTqoSJjLFxY%2B8koYL2diul3KpMTBnuuYEpUJpZ8Mt4GV3%2BfI7itE%2BgJXIt%2B28Z1avOeqZPzUDLblobp1VT1xbkwTB8ZirMaPR4sKx%2B3CYeFNjqmGj42GnA2C1uRQDMPbhBls2XHbudT3kBQEKm4%2FwjLcgYLM2%2Bp8FdIafhtg515JKMj00GTOqP6wbmXMgb14prTDZ014zL50xEddq8yS44tAvNPzjj6eNVmNvCVabPmR3qPL%2B4bTn9irX0Of791lxHKdkBTejKFwG3spPnBVJSeqTM9K1pGKgyHLX%2FP%2FSHKi%2FDm8MuVj3W9%2Fwfmr0nkxQQqEFZeP7nBvE6eiloKyXf3bvio5YEaoNxA7m4iPZRS%2FXXPhs17Wd8Nlxp743dUMUxO5vI12LMruGnbPXyZvmHaagsrkbNC9m2du31mxDyRcrX1kdhVm5tLoYI%2Bt5Q60NlhbRIMm%2BVZagf9W5oNkbq5Vf1spSAJs9fLAG7dbXvCigVPOiUQ3rgn4rLcb3u2hmRnMG6r8j6gW1V4fuaQf63vrmG%2BvMWn8fmbz%2BC9%2BxVwpoFgAA) to test out expressions.