---
Title: "Support Responsibility Matrix"
Date: 2023-1-9
Weight: 15
---

{{% pageinfo %}}
The matrix below defines the primary responsible party for troubleshooting and resolving different types of issues. Trustgrid support can be engaged to assist in resolving issues not directly in scope of their responsibility.
{{% /pageinfo %}}

### Organization and Roles

#### Trustgrid

Trustgrid is a software and services provider that operates a multi-tenant cloud management system and assists customers and end-users with the deployment and management of a data transfer plane.

#### Customer

Customers have a direct, contractual relationship with Trustgrid to utilize Trustgrid’s software and services. Depending on the contractual relationship with Trustgrid, Customers may also assume the role and responsibilities of End Users.

#### End-User

End Users own and operate a physical, virtual or cloud environment in which the Trustgrid software may be installed.

### Terminology

{{<fields>}}
{{<field "Trustgrid Management Portal" >}}
Cloud-based system used to remotely manage, monitor, and configure Trustgrid Nodes.
{{</field >}}

{{<field "Trustgrid Node" >}}
Appliance running Trustgrid’s software to facilitate Trustgrid Connect, EdgeCompute, and RemoteAccess services. These appliances could be deployed on physical devices or as virtual machines. The appliance consists of the Trustgrid software and the underlying operating system (OS).
{{</field >}}

{{<field "Customer Site" >}}
Site with Trustgrid node(s) deployed and managed by the Customer.
{{</field >}}

{{<field "End-User Site" >}}
Site with Trustgrid node(s) deployed and managed by the End-User.
{{</field >}}
{{</fields>}}

### Support Matrix

| Issue Type                                                      | Trustgrid           | Customer               | End-User               |
| --------------------------------------------------------------- | ------------------- | ---------------------- | ---------------------- |
| Trustgrid software or operating system issues                   | Full                | None                   | None                   |
| Hardware appliance | Shared | Shared | None |
| Control Plane Connectivity                                      | Shared              | Shared - Customer Site | Shared - End-User Site |
| Data Plane Connectivity                                         | Limited<sup>1</sup> | Shared                 | Shared                 |
| Data Plane Performance                                          | Shared              | Shared                 | Shared                 |
| Power and physical network connectivity for Trustgrid Appliance | None                | Full - Customer Site   | Full - End-User Site   |
| Internet Service                                                | None                | Full - Customer Site   | Full - End-User Site   |
| Internet Side Firewall                                          | None                | Full - Customer Site   | Full - End-User Site   |
| Local networking including switching and firewall               | None                | Full - Customer Site   | Full - End-User Site   |

---

<sup>1</sup>Trustgrid can work with the Customer and End-User to confirm that the Trustgrid system is working as expected and provide additional information to aid in troubleshooting.

### Tier One Support

The customer’s support team should be the end-user’s first point of contact. Trustgrid’s standard support contract makes the customer responsible for performing tier one support tasks prior to escalating Trustgrid Support. This includes contacting end-user technical resources as needed.

Example tasks include:

- [Triage]({{<relref "/tutorials/operations-runbook/node-down-response">}}) and Initial [troubleshooting offline Trustgrid nodes]({{<relref "/tutorials/operations-runbook/control-plane-disconnect">}})
  - Before escalating to Trustgrid support please establish contact with an [end-user](#end-user) technical resource with [console access]({{<relref "/tutorials/local-console-utility#connecting-to-trustgrid-local-console">}})
- [Changing Trustgrid node IP addresses]({{<relref "/tutorials/wan-interface-ip">}})
- [Adding, modifying]({{<relref "/docs/nodes/appliances/vpn/nats">}}) and testing NATs on Edge devices
- Modifying [interface]({{<relref "/docs/nodes/appliances/interfaces#interface-routes">}}) and [domain routes]({{<relref "/docs/domain/virtual-networks/routes#managing-virtual-network-routes">}}), including facilitating [failover between sites]({{<relref "/tutorials/operations-runbook/site-failover">}})
- [Disabling]({{<relref "/tutorials/management-tasks/changing-node-status">}}) and removing [nodes]({{<relref "/tutorials/management-tasks/deleting-nodes">}}), [clusters]({{<relref "/tutorials/management-tasks/deleting-clusters">}}), and related configurations (e.g. routes) when no longer needed
- For customers with Enhanced Support SLA, Trustgrid provides advanced replacement for failed hardware appliances

{{<alert>}} Note: Some of the above services are managed by Trustgrid Professional Services during initial deployment {{</alert>}}
