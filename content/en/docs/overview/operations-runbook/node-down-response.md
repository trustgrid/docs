---
Title: "Node Down Response"
Date: 2022-12-29
Tags: ["node", "troubleshoot", "help"]
---

{{% pageinfo %}}
This process is intended to help customer support personnel quickly identify the scope of a [node]({{< ref "docs/concepts/node" >}}) down problem and get services back online as quickly as possible. 
{{% /pageinfo %}}

### Node Down Triage

#### Triage Checklist
- Determine Production Status of the Trustgrid Node

- Confirm High Availability or Disaster Recover is Functioning

    - If clustered, is the partner active and working

    - If there is a second site, confirm if it is active

- Determine the type of outage (Control Plane, Data Plane or Both)

### Determine Production Status of the Trustgrid Node
Use the [Production Status Tags]({{< ref "docs/concepts/tag/prod-status-tag" >}}}) to determine if the [node]({{< ref "docs/concepts/node" >}}) is in use and expected to be online.

### Confirm High Availability or Disaster Recover is Functioning
Before troubleshooting why a [node]({{< ref "docs/concepts/node" >}}) is down we should determine if the services it provided can be provided in the interim by a cluster member or devices at a secondary/disaster-recovery site.

##### Is the Node a Cluster Member
- Check if partner member is online:

  - If Yes:

        - Confirm partner member is now active and operating normally. 

        - The issue is limited to this specific Node. This could be its power input, the hardware or operating system, configuration, or its internet connection (if different than partner member).

  - If No:

        - Possible site-level issue including power or internet provider

        - Proceed to Secondary / DR Site 

##### Is there a Secondary / DR Site for this Node/Cluster?
- Are the [nodes]({{< ref "docs/concepts/node" >}}) deployed at the secondary site online in the Trustgrid?

  - If No:

        - Confirm nodes for other end-user sites are not also offline which might indicate a wider spread issue. Escalate to Trustgrid Support if you suspect a major issue with the Trustgrid system.

        - If limited to a single customer it is recommended to contact that customer immediately. They may be experiencing an outage or performing maintenance.

  - If Yes:

        - If the customer is configured for Automatic Failover between sites, verify that traffic is flowing through the active member at the secondary site.

        - If the customer is configured for Manual Failover you will need to adjust the route destination to point to the secondary site.

### Determine Connectivity Status of the Trustgrid Node
Because Trustgrid provides independent control and data planes, there are a few ways it can manifest as “down”:

{{< field-def "Control Plane Down" >}}
The [node]({{< ref "docs/concepts/node" >}}) appears offline from within the Trustgrid portal. This indicates that the [node]({{< ref "docs/concepts/node" >}}) shutdown or has not sent a heartbeat notification to Trustgrid within the past 10 minutes. 
{{< /field-def >}} 

{{< field-def "Data Plane Down" >}}
The [node]({{< ref "docs/concepts/node" >}}) appears online from within the portal but reports it is unable to connect to one or more gateway nodes. This can be indicated by the Data Plane Status indicator when viewing the [node]({{< ref "docs/concepts/node" >}}) in the portal, or by receiving a “Gateway Connectivity Health Check” failure [event]({{< ref "docs/alarms/events" >}}) notification. Both of these only work if the Control Plane is currently working. 
{{< /field-def >}} 

{{< field-def "Both Control and Data Plane Down" >}} 
In this situation the [node]({{< ref "docs/concepts/node" >}}) appears down in the Trustgrid Portal and users/applications are unable to reach services across the data plane between the Gateway and Edge sites. This is the most common scenario. While the Data Plane is most critical for the services provided across the device, first priority should be restoring the Control Plane connection so that additional troubleshooting tools are available. Often this process also uncovers the reason the data plane is down.
{{< /field-def >}}
