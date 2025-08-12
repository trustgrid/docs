---
title: Running Agent Diagnostic Command
linkTitle: Agent Diagnostic
description: Details how to run the Trustgrid Agent `diag` command and understand the output
---
The `diag` command is designed to determine if an agent has the required connectivity to connect to the Trustgrid Control Plane.

## Running `diag` Command
1. Connect to the machine with the Trustgrid agent installed and open a console. If running as a container you will need to use the appropriate tools to get command line interface (cli) inside the container not the host.
1. Run the command `tg-agent diag`
1. Review the output

```
tg-agent diag
--------------------------------------------------------------------------------
-------------------------------------- DNS -------------------------------------
--------------------------------------------------------------------------------

keymaster.trustgrid.io:
	35.171.100.28

gatekeeper.trustgrid.io:
	35.171.100.26
	35.171.100.25
	35.171.100.27

zuul.trustgrid.io:
	35.171.100.19
	35.171.100.20

repo.trustgrid.io:
	35.171.100.29

--------------------------------------------------------------------------------
--------------------------------- CONNECTIVITY ---------------------------------
--------------------------------------------------------------------------------

keymaster.trustgrid.io:443:
	35.171.100.28:443:              OK

gatekeeper.trustgrid.io:8443:
	35.171.100.26:8443:             OK
	35.171.100.27:8443:             OK
	35.171.100.25:8443:             OK

zuul.trustgrid.io:8443:
	35.171.100.20:8443:             OK
	35.171.100.19:8443:             OK

repo.trustgrid.io:443:
	35.171.100.29:443:              OK

--------------------------------------------------------------------------------
-------------------------------------- TLS -------------------------------------
--------------------------------------------------------------------------------

keymaster.trustgrid.io:443:
	35.171.100.28:443:              OK

repo.trustgrid.io:443:
	35.171.100.29:443:              OK

--------------------------------------------------------------------------------
------------------------------------- MTLS -------------------------------------
--------------------------------------------------------------------------------

gatekeeper.trustgrid.io:8443:
	35.171.100.27:8443:             OK
	35.171.100.26:8443:             OK
	35.171.100.25:8443:             OK

zuul.trustgrid.io:8443:
	35.171.100.20:8443:             OK
	35.171.100.19:8443:             OK
```

## Understanding the Output
The `diag` command steps through several levels to help determine where a breakdown in connectivity may be occurring.
### DNS
First, the agent will attempt to resolve key hostnames for Trustgrid control plane services using the agent OS's configured DNS servers. Some will return a single IP, some will return several.

Most likely, all or none of these services will show OK. If, any are failing investigate the DNS settings on the host OS.

### Connectivity
Next, the agent will attempt to perform a TCP port connection, using `netcat` like test, to the appropriate ports for each control plane service. Specifically it will attempt connecting to:
- TCP Port 8443 for:
    - gatekeeper.trustgrid.io
    - zuul.trustgrid.io
- TCP Port 443 for:
    - keymaster.trustgrid.io
    - repo.trustgrid.io

If any of these are failing you will need to investigate any firewalls and routing between the agent and the internet.  If `traceroute` or `mtr` are installed, running the TCP versions of these tests to the appropriate port may help identify where the connection is being dropped.

### TLS 
Next, the agent will validate that the appropriate TLS certificate is being received without alteration. TLS integrity is critical to the security of Trustgrid communications. Any failure here indicates something is altering the TLS certificates. 

Any security appliance performing TLS decryption/re-encryption needs to be configure to exclude traffic to *.trustgrid.io and the [Trustgrid Control Plane address spaces]({{<relref "/help-center/kb/site-requirements/#trustgrid-control-plane">}})

See the document on [SSL/TLS tampering]({{<relref "/help-center/kb/startup-process/ssl-tls-tampering">}}) for related information.

### MTLS
Finally, the agent will utilize its own certificate to verify mutual TLS authentication is working to services that utilize that functionality.  If this fails, there may be an issue with the local certificate that would require working with Trustgrid support to investigate. Though, the quickest resolution might be to re-register the agent with a new token as this will refresh the certificate. 