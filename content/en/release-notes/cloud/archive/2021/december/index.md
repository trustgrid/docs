---
title: "December"
linktitle: "December"
date: 2021-12-14
type: docs
weight: 12
---

## Data Plane Status Table
Previously Trustgrid had a visual indicator on the Node detail page to show the Data Plane health of a node, but it only showed the state as `Connected`, `Degraded` or `Disconnected`.  While connected and disconnected were clear, there was little information on why the status was Degraded. 

With this release hovering over the Data Plane status icon will bring up a pop-up screen showing all the peers and and their connection status.  Additionally, if the peer is connected it will show:

* The public IP of the peer
* The mode used to connect to with the peer (TLS or UDP)
* The route trip time (RTT) between the devices when the tool was last run.  

This data can be searched and refreshed.  In a future release we will make historical RTT time viewable for each peer. 

![Data Plane Status](data-plane-hover.png)

## Flow Log Improvements
Flow logs provide valuable metadata about traffic across Trustgrid networks but our search had performance and reliability issues that made it difficult to use.   With this release we’ve rearchitected how search these logs to allow for a significant performance improvement. 

## Improved IP Geolocation Data
Trustgrid has changed our public IP geolocation data providers to improve the accuracy of the information we use to place nodes on our Network Maps.  This new provider will also allow us to present additional information such as ISP and ASN in the future, which could be handy for identifying when multiple devices from the same ISP or ASN start having issues. 

## Custom Portal Roles
Trustgrid is working on moving our permission structure beyond the four default roles (Monitor, Operator, Node-Administrator and Administrator) to allow customers to tailor roles to just the permissions their users need. For example, you could allow a user Monitor (read-only) permissions to most of the Trustgrid environment but grant them the ability to stop and start containers.  

Currently this requires working with Trustgrid Support to define these new roles but once defined they can be assigned by Node Administrators and Administrators.

## Other Improvements
* You can now search the Nodes table for Management IPs and Device types
* Fixed a broken link in the new user invite email for customers without IDP integration

