
---
linkTitle: "Trustgrid Support Access"
title: "Trustgrid Support's Node Access and Data Plane Security"
---


Trustgrid takes security and privacy seriously. **Trustgrid teams do not have access to any data that traverses the Data Plane.** All customer data moving through the Data Plane is encrypted and inaccessible to Trustgrid staff, including support and engineering teams.

For a clear understanding of the difference between the Data Plane and Control Plane, see the [Basic Architecture]({{<relref "getting-started/basic-architecture" >}}) section.

Node access is strictly limited to maintenance and troubleshooting activities, and only applies to the Control Plane. Trustgrid staff may access node management interfaces for support purposes, but this does not grant access to customer data or applications running within the Data Plane.


Trustgrid customers may restrict Trustgrid support access to their nodes via the Control Plane by [disabling remote support]({{<relref "docs/support#remote-support" >}}). Disabling support access does not interfere with automated management features such as patching, updating, logging, or authentication. Changes to the support flag are audited.
