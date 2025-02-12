---
title: Appliance Upgrade Process
linkTitle: Appliance Node Upgrade Process
weight: 40
description: A step by step description of how a healthy node appliance upgrades.
---

## Upgrade Process

1. DNS request for A record of repo.trustgrid.io
    1. The response will be IPs in the Trustgrid AWS networks, `35.171.100.16/28` (us-east-1) or `34.223.12.192/28` (us-west-2)
1. The node will initiate an HTTPS connection to **port 443** on the ip address returned for repo.trustgrid.io. This includes validating the TLS Certificate.  Any changes to the certificate chain will stop the process from continuing. 
1. The update process will first download and install the latest version of the Trustgrid software.
1. After the Trustgrid software is installed, the node will begin downloading any available OS security packages and installing them.
1. After installation the Trustgrid node will reboot and load the latest version of the Trustgrid software.


## Troubleshooting Upgrade Failures
1. [Confirm port 443 access to repo.trustgrid.io]({{<relref "/help-center/kb/upgrade-process/443-blocked#troubleshoot-connectivity">}})
1. Check for [interference with TLS certificate validation]({{<relref "/help-center/kb/startup-process/ssl-tls-tampering#verification">}})