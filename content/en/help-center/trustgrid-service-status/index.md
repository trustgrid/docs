---
Title: "Check Trustgrid Services Status"
Date: 2023-01-04
---
View the status of all Trustgrid services on our [Statuspage Account](https://status.trustgrid.io/) 

Within the Trustgrid Portal, hover over the username in the upper right corner and then select `Status`.

![img](status.png)

## Components

Trustgrid components on the status page are broad and meant to communicate areas where there may be elevated error rates.

### Node

Node components affect services that nodes use to manage status, updates, and enroll.

* Cloud Controller - signals nodes when to update their config and receives from nodes: events, flow logs, shadow updates
* Software Repository - provides both OS and Trustgrid package updates
* License Manager - responsible for enrolling new devices with Trustgrid

### Management

* Alerts - event collection and notification services
* Management Portal - the UI used to manage your devices
* User Management - the ability to sign in using Trustgrid native authentication
* Management Gateways - the ability to invoke interactive services on nodes