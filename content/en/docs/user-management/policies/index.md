---
Title: "Policies"
Date: 2023-01-12
---

{{% pageinfo %}}
User Management Policies define what an authenticated user is allowed to do within the Trustgrid system. Policies either allow or explicitly deny specific actions by the user.
{{% /pageinfo %}}

## Types of Policies

### Trustgrid Managed Policies

These policies start with the prefix `builtin-tg` and are created and managed by Trustgrid to provide a starting level of permissions based on common user types. Descriptions of the user types for each role are listed below.

As Trustgrid adds new features or actions these policies will automatically be updated to include appropriate permissions. These policies cannot be edited by customers.

### Customer Managed Policies

These policies can be defined by customers to meet the specific needs of their teams.

### Best Practices

- Start by attaching the Trustgrid Managed Policies that best fit the user's minimum, required needs following the principle of “least privilege.”
- If additional permissions are required, create a Customer Managed Policy that includes those specific rights rights. Then attach that policy to the user.

### Trustgrid Managed Policy Descriptions

Below are the basic descriptions of the Trustgrid Managed Policies including an example role description. Trustgrid will use this description to guide if it is appropriate to add new permissions to the policies.

The policies are listed from lowest permissions to highest.

#### builtin-tg-apps-user

This policy provides minimal permission for a user to connect to the Trustgrid applications portal for using Remote access apps. The user has no rights within the Trustgrid management portal and should likely be given a landing page of `/apps` to be redirected straight to the applications portal.

Target User: End user for remote access applications with no knowledge of the Trustgrid system.

#### builtin-tg-monitor

This policy provides read-only access to non-sensitive resources in the Trustgrid management portal.

Target User: Project manager or executive sponsor that needs to view information in the portal, but does not require tools required to troubleshoot or resolve issues.

#### builtin-tg-operator

This policy provides mostly read-only access but also has the ability to view and use tools within the portal and perform actions to resolve common issues (e.g. restarting/rebooting nodes).

Target User: Entry-level technical resources responsible for initial triage of support issues.

#### builtin-tg-node-admin

This policy provides the ability to make changes to most resources but is restricted from making changes that can have a wide impact such as changing a gateway node server settings.

Target User: Mid-level technical resources that need to make changes to things like node VPN configuration and virtual network routes, but should be restricted from making changes that can impact an entire organization when possible.

#### builtin-tg-admin

This policy has permission to change almost all resources with the exceptions of Identity Providers and User Management.

Target User: Senior technical resources needing near full access.

#### builtin-tg-access-admin

This policy grants the ability to manage permissions and identity-related resources including configuring Identity Providers, Adding/Inviting Users, and creating and assigning policies

Target User: Management or access control-related resources responsible for granting access to the Trustgrid system.

#### builtin-tg-upgrade-admin

This policy grants all the permissions required to upgrade nodes using the [Upgrade Manager]({{<relref "/docs/upgrade-manager">}}) system. 

Target User: Users responsible for upgrading Trustgrid nodes.

#### builtin-tg-provisioning-admin

This policy grants the ability to create and manage orders in the provisioning section of the portal.

Target User: Project Managers and other team members involved in managing orders of new Trustgrid nodes. 


## Resource-Scoped Policies

Policies have a Resources section that allows the permissions granted to be limited to specific Trustgrid resources, such as nodes and clusters. This can be used to present a user with an extremely limited view in the portal. For example, an edge site technical contact could be given a login that was only permitted to see the nodes at their site.

![img](resources.png)

{{<alert color="warning">}} In order to grant users access to the portal a second policy is required granting the `portal::access` permission to the `*` resource.  {{</alert>}}

### Trustgrid Resource Names or TGRN

Trustgrid has adopted a resource naming convention that follows the basic format: `tgrn:tg::<service>:<resource-type>/<resource-id>`

#### Wildcard TGRN
By default, the Resources field is populated with `*` which grants permissions to all resources. This is included by default for new policies and would need to be removed to provide TGRN restricted permissions.

#### Node TGRN

A node’s TGRN will start with the below prefix:
`tgrn:tg::nodes:node/`
Followed by the UID of the node. The node TGRN is accessible from the [infovisor]({{<relref "docs/nodes/shared/infovisor" >}}) for a node by clicking on the "Copy" button to the right of the TGRN entry.
{{<tgimg src="node-tgrn.png" width="70%" caption="TRGN entry in the Inforvisor panel">}}

A complete [node]({{<ref "docs/nodes" >}}) TGRN looks like: `tgrn:tg::nodes:node/ab3550cf-f3da-4d90-b4b1-b870fc85c9c7`

#### Cluster TGRN

Cluster TGRN start with the prefix:
`tgrn:tg::nodes:cluster/`
This is followed by the cluster’s fully qualified domain name (fqdn) which is the combination of the cluster name and your organization's Trustgrid domain. 

You can copy the Cluster TGRN from a cluster's Overview page by clicking the "Copy" button to the right of TGRN.
{{<tgimg src="cluster-tgrn.png" width="60%" caption="Cluster TGRN on Overview page">}}

A complete [cluster]({{<ref "docs/clusters" >}}) TGRN looks like: `tgrn:tg::nodes:cluster/awscluster.training.trustgrid.io`

## Condition-Scope Policies
Conditions allow a policy to be scoped to objects that match specific conditions. This could be used in addition to or in place of resource-scoped policies. 
{{<tgimg src="policy-conditions.png" width="70%" caption="Policy Conditions">}}
### Tag Conditions
Tag conditions allow a policy to be scoped to objects that have a specific tag or combination of tags. This is useful for allowing users to see and manage objects that have a specific tag. 

There are three sections to specify tag conditions:
- `All` - When this section is used all conditions specified in this section must be present on the object.
- `Any` - When this section is used any conditions specified in this section must be present on the object.
- `None` - When this section is used no conditions specified in this section can be present on the object.

Each section has two fields:
- `Tag Matches` requires the tag to be present with the matching value.
- `Tag does NOT Match` requires the tag is NOT present or has a different value.

{{<alert color="warning" title="Be Advised:">}} Using tag based policies means any user with the permissions to change the tag of an object can change what policies apply to that object. {{</alert>}}

## Policy Management

### Creating/Editing Policies

1. Navigate to `User Management` → `Policies` {{<tgimg src="policy-management.png" width="40%" caption="Policy Management Page">}}
2. For a new policy click `Create Policy`. To edit an existing find it in the list and click on its name.
3. The policy wizard has two sections. The left section groups permissions by category and is where you allow or deny permissions. The right section shows the resulting JSON of the policy and is read only.
As an example, we will create a role to grant users the ability to see, edit, and comment on Provisioning orders.
   1. Under Permissions, search for `orders`
   1. Click the green check box next to each of the desired permissions. (Clicking the X would explicitly deny that permission which will take precedent over any other policy that allows that permission)
   1. You will notice the Policy JSON updated as you select permissions
   1. Click `Save`{{<tgimg src="new-policy.png" caption="Policy updated notification" >}}

### Clone Existing Policy
To clone an existing policy click the copy button on the far right of the policy row. This will create a new policy with the same permissions pre-selected which can then be edited and named as desired. 
{{<tgimg src="clone-policy.png" width="80%" caption="Policy table showing the clone/copy policy button">}}

### Attaching Policies and Viewing Effective Permissions

Policies can attached at either:
- The [User level]({{<relref "/docs/user-management/users#attaching-policies">}}) which applies the permission to only the individual user
- The [Group Level]({{<relref "">}})
