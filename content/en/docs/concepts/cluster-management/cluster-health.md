---
Title: "Cluster Site Health State"
Tags: ["concepts", "clusters", "cluster health", "cluster management"]
Date: 2022-12-23
---

{{% pageinfo %}}
A healthy cluster should have two online Trustgrid nodes in a healthy state. The Trustgrid control plane can monitor the member nodes and generate events when their state changes such that the health of the cluster is impacted. 
{{% /pageinfo %}}

## Basics
> These events are based on the information reported by each member node to the control plane.  Because of this:
>- There can be a delay between a change and the Cluster Site Health events
>- If member nodes have issues preventing control plane connections but not impacting data plane or other services the events may not accurately reflect the state of the cluster

## Cluster Health State
The Trustgrid will now display one of three states in the portal:

| State | Meaning |
|-------|---------|
| **Healthy** | Both members are online with a single member active. |
| **Unhealthy** | There is at least one online and healthy member of the cluster with the active role AND one or more of these conditions is true:
| |- One of the cluster members is offline, unhealthy or disabled
| |- There are **two** cluster members with the active role |
| **Offline** | Both nodes are offline, unhealthy or disabled. |
