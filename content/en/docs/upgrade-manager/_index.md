---
tags: ["upgrade manager"]
title: "Upgrade Manager"
linkTitle: "Upgrade Manager"
date: 2024-8-30
weight: 1
---

## Overview

The Upgrade Manager is designed to manage upgrades for individual nodes and node clusters in bulk. It provides a user-friendly interface to plan, monitor, and execute upgrades while ensuring system stability and minimal downtime.

## Key Features

- Manage upgrades for individual nodes and node clusters
- Perform pre-upgrade health checks
- Monitor online status of devices
- Verify upgrade success
- Support for staged upgrades in clustered environments

## Planning an Upgrade

To plan a new upgrade:

1. Navigate to the upgrade planning interface. From the Portal, click on the "Upgrade Manager" option in the left navigation menu.
1. Click "Create Upgrade" at the top right.
2. You can set the following parameters:
   - **Name**: Enter a name for the upgrade (e.g., "aug-2024-upgrade")
   - **Include Tags**: (Optional) Specify tags to include only specific devices
   - **Exclude Tags**: (Optional) Specify tags to exclude certain devices

### Special Behaviors

- **Require approval before failover**: When checked, this option requires manual approval before proceeding with failovers during the upgrade process. This allows for additional control and verification steps.

## Nodes

The Nodes section displays a list of devices that are candidates for the upgrade. Each node entry includes:

- **Name**: The identifier for the node
- **Tags**: Associated tags that can be used for filtering or grouping
- **Version**: The current version of the node

### Example Node Data:

| Name | Tags | Version |
|------|------|---------|
| jlp-profiled | | 1.5.20240212-1978 |
| node1-profiled | Flagged : false, abe : froman, autoupdate : true, ... | 1.5.20240826-2138 |
| terry-profiled | | 1.5.20220124-1143 |
| zthage | one : boom2, prod_status : production, two : kapow | 1.5.20240203-1968 |

## Clusters

The Clusters section shows groupings of nodes that operate together. This is particularly important for managing upgrades in high-availability setups. The table includes:

- **Name**: The cluster identifier
- **Tags**: Associated tags for the cluster
- **Health**: The current health status of the cluster

### Example Cluster Data:

| Name | Tags | Health |
|------|------|--------|
| FFF.dev.regression.trustgrid.io | | offline |
| automated-permanent.dev.regression.trustgrid.io | | unknown |
| clonepc.dev.regression.trustgrid.io | | unknown |
| ecu-test.dev.regression.trustgrid.io | one : boom2, two : kapow | offline |
| profiled-node.dev.regression.trustgrid.io | MASTER_NODE : SET_TO_NODE2, pin-active : a9f1fedd-a474-455b-a123-37a035409229 | offline |

## Upgrade Process

### Initiation
1. **Planning**: Create a new upgrade plan, specifying version and any tag-based inclusions or exclusions.
1. **Scheduling**: The upgrade is scheduled with a specific start time (e.g., "8/29/2024, 2:46:45 PM").

### Execution
1. **Workflow Creation**: The system automatically creates workflows for each device or cluster that needs to be upgraded.
1. **Prechecks**: The first step for each workflow is to validate the node or cluster can be upgraded.
   - For individual nodes, the node must be active and online.
   - For clustered nodes, the cluster must have exactly one active member and the cluster must be healthy.
1. **Dry Run**: (Optional) A dry run can be performed to simulate the upgrade process without actually upgrading the devices.-
1. **Health Checks**: Every minute, each device is checked for health and readiness to move to the next state in the workflow.
1. **State Transitions**: Devices progress through various states (e.g., NEW, READY, WORKING, DONE) as the upgrade proceeds.
1. **Clustered Upgrades**: 
   - For clustered nodes, the passive member is upgraded first.
   - After successful upgrade, a failover is initiated.
   - The new passive member (previously active) is then upgraded.
1. **Monitoring**: The system monitors the online status of devices throughout the process.
1. **Verification**: Post-upgrade, the system checks device shadows to confirm successful version updates.

## Workflow Management

The Upgrade Manager provides a detailed view of the upgrade process through its Workflows and Timeline sections.

### Workflows Table

The Workflows table displays the current status of each node or cluster in the upgrade process:

- **Name**: Identifier of the node or cluster
- **Tags**: Associated tags for filtering and categorization
- **State**: Current state of the workflow (e.g., NEW, READY, WORKING, DONE, SKIPPED)
- **Note**: Additional information, such as reasons for skipping
- **Actions**: Icons for making notes (pencil icon) or skipping the workflow (x icon)

#### Workflow States
- **NEW**: Initial state for a newly created workflow
- **READY**: The device is ready for the next step in the upgrade process
- **WORKING**: The upgrade is actively being applied to the device
- **DONE**: The upgrade has been successfully completed
- **SKIPPED**: The workflow has been manually skipped or automatically skipped due to conditions (e.g., "cluster has 0 members")

### Timeline

The Timeline section provides a detailed, chronological view of events during the upgrade process. Each entry in the Timeline represents a step or action taken during the upgrade workflow.

### Timeline Entry Format

Each entry in the Timeline contains:

- **Timestamp**: Date and time of the event (e.g., "8/29/2024, 2:48:51 PM")
- **Entity**: The specific node or cluster involved (e.g., "node: zthagw (e04f0a95-cf9f-46ca-b5fa-df8a00b0b862)")
- **Message**: Description of the action or state change

### Common Timeline Messages

1. **Workflow Initiation**:
   - "creating workflows": Indicates the start of the upgrade process.

2. **Node Transitions**:
   - "transition to upgrading": The node is beginning the upgrade process.
   - "transition to ready": The node is prepared for the next step in the upgrade.
   - "transition to done: already up to date": The node was found to already be on the target version.
   - "transition to skipped: offline nodes can't be upgraded": The node was offline and couldn't be upgraded.

3. **Cluster Transitions**:
   - "transition to skipped: cluster has 0 members": The cluster was skipped due to having no active members.
   - "transition to skipped: cluster has 1 members": A cluster with only one member was skipped.
   - "transition to skipped: cluster is offline": The entire cluster was offline and couldn't be upgraded.
   - "transition to skipped: cluster upgrade skipped": The cluster upgrade was skipped for an unspecified reason.

### Interpreting Timeline Entries

- **Sequence of Events**: The Timeline shows the progression of the upgrade across all nodes and clusters, allowing users to track the order of operations.
- **Upgrade Status**: Users can quickly identify which nodes or clusters have successfully upgraded, which were skipped, and why.
- **Troubleshooting**: The Timeline is crucial for identifying issues, such as offline nodes or empty clusters, that may require attention.

### Using the Timeline

- **Real-time Monitoring**: As the upgrade progresses, new entries are added to the Timeline, allowing for real-time monitoring of the process.
- **Post-upgrade Analysis**: After the upgrade is complete, the Timeline serves as a comprehensive log for review and analysis.
- **Filtering**: Use the search functionality to filter Timeline entries for specific nodes, clusters, or message types.

Understanding the Timeline entries is crucial for effectively managing and troubleshooting the upgrade process. It provides insights into the status of each node and cluster, helping administrators identify and address any issues that arise during the upgrade.

## Additional Important Notes

### Upgrade Limitations

- **Single Active Upgrade**: Only one upgrade can be active at a time. This ensures system stability and prevents conflicts between multiple upgrade processes.
- **Upgrade Timeout**: If a node does not transition out of the upgrade state within 90 minutes, its workflow will transition to failure. 

### Dry Runs

- Dry runs can be performed from an upgrade in the planning state.
- Purpose: To run prechecks and show which nodes would match the specified tags at that moment.
- Behavior: Dry runs do not actually upgrade any nodes or clusters.
- Use Case: Ideal for validating upgrade plans and identifying potential issues before initiating a real upgrade.

### Special Behaviors

#### Require Approval Before Failover

- When this option is selected during upgrade planning, the system will not automatically fail over any clusters.
- Failovers will only occur after explicit user approval through the UI.
- This feature allows for greater control and manual verification during critical stages of the upgrade process.

#### Manual Updates

- Nodes tagged with `manualupdate` as `true` will not be upgraded and will transition to the `skipped` state immediately.

## Best Practices

1. **Always Perform Dry Runs**: Before initiating an actual upgrade, use the dry run feature to validate your upgrade plan and identify any potential issues.

2. **Use Tags Effectively**: Carefully consider your include and exclude tags to target the correct set of nodes and clusters for your upgrade.

3. **Monitor the Timeline**: Keep a close eye on the Timeline during the upgrade process to quickly identify and address any issues.

4. **Use the "Require Approval" Feature for Critical Systems**: For high-priority or sensitive clusters, use the require-approval feature to maintain manual control over the failover process.

5. **Regular Health Checks**: Ensure all nodes and clusters are in a healthy state before initiating an upgrade to minimize the risk of failures.

## Troubleshooting

While the Upgrade Manager tool itself has limited troubleshooting capabilities, here are some common issues you might encounter and initial steps to address them:

1. **Offline Nodes**: If the Timeline shows nodes are being skipped due to being offline, check the network connectivity and status of these nodes.

2. **Empty Clusters**: For clusters showing 0 members, verify the cluster configuration and ensure member nodes are properly assigned.

3. **Upgrade Failures**: If nodes fail to upgrade, check the specific error messages in the Timeline and verify that the target version is compatible with the node's current configuration.

4. **Stuck Workflows**: If a workflow seems stuck in a particular state, you may need to manually investigate the node or cluster status outside of this tool.

[Note: Additional troubleshooting steps and external resource links to be added later.]

Remember, for more complex issues or detailed troubleshooting steps, refer to the external documentation and support resources provided by your organization.