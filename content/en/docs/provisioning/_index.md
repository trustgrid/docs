---
tags: ["provisioning", "orders"]
title: "Provisioning"
linkTitle: "Provisioning"
date: 2025-05-13
weight: 16
---

{{% pageinfo %}}
The provisioning portal allows users to manage orders, which can be used to gather information to fulfill requests for appliances.
{{% /pageinfo %}}

## Orders List View

The **Orders List View** provides a summary of all provisioning orders, offering users streamlined navigation through existing orders.

{{<tgimg src="orders-list.png" width="60%">}}

### Interface Overview

* **Search Bar**: Quickly find orders using keywords, order numbers, creator names, project names, or other relevant details.
* **Orders Table** includes the following columns:

  * **Assignee**: Individual responsible for managing or completing the order.
  * **Project**: Associated project for the order.
  * **Summary**: Brief description with order number, nodes involved, hardware types, creator, and creation timestamp.
  * **Status**: Current status of the order (e.g., "New").
  * **Actions**: Create new orders, bulk update existing orders, or download the order list as a CSV file.

## Order Edit View

The **Order Edit View** allows updating orders with new information, and validates the data's coherency before an order's information is used in the fulfillment process.

{{<tgimg src="order-edit-top.png" width="60%">}}

### General Information

On the right-hand side, you can view and edit general order workflow and assignment information.

* **Order ID**: Unique identifier for tracking.
* **Status & Priority**: Current state and urgency level. Field validation is relaxed and listed as warnings until an order is moved to the QA state. Validations must pass before an order can be moved to provisioning.
* **Due Date**: Deadline for completion.
* **Assignee & Labels**: Assign responsible individuals and categorize orders.
* **Project**: Link orders to specific projects.
* **Creator & Created Date**: User who initiated the order and timestamp.

### Networking and Technical Details

* **DNS Configuration**: Primary and secondary DNS servers.
* **Interface Configuration**:

  * IP assignment method (DHCP or Static).
  * Subnet and gateway details.
* **Cluster IP & Routed Networks**: Internal addressing and routing configurations.

### NAT Configuration

* **Virtual CIDR**: Virtual network range.
* **Local CIDR**: Local network addresses.
* **Description**: Context or purpose of each NAT rule.

### Stakeholders & Contacts

* **Contact Information**:

  * Salesforce Case #, End User Account #, Company Name, Phone, Email.
  * Technical contacts and details.

### Installation Details

* **Location & Site Type**: Physical installation location and type.
* **Clustered/HA, Nodes, Interfaces**: High-availability, node quantity, and interfaces.
* **Rack Mounts**: Inclusion of rack mounts.
* **Appliance Type & RMA**: Hardware specifics and return authorization.

### Node Information

* **Node Name & Asset ID**: Unique node identifiers and asset management tags.
* **DNS & Associated Node**: DNS configurations and related nodes.
