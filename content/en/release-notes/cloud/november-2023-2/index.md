---
title: November 2023 Second Release Notes
linkTitle: 'November 2023 Second Release'
type: docs
date: 2023-11-15
description: "Minor release to resolve issue"
---
This minor release fixes the following issue:
## Issues Fixed
* A function was deprecated in the last release as it was believed to be no longer needed but it was discovered that some methods of registration, such as the AWS CloudFormation template, still relied on it. This release removes the deprecation and restores the function. 
