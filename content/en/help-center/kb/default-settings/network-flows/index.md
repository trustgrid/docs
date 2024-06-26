---
title: "Network Flow Default Settings"
linkTitle: Network Flow
description: >
  Default settings impacting network flows traversing Trustgrid devices
---

## Timeouts
The table below details the time the Trustgrid network process waits before terminating a flow. Flows are matched based on the IP tuple (source ip, source port, destination ip, destination port).

| Protocol | States | Time(seconds) | Description |
|---|---|---|---|
| TCP | SYN_SENT, SYN_RECEIVED | 15 | Covers flows that have not yet completed a 3-way TCP handshake. e.g. (SYN, SYN/ACK, ACK) |
| TCP | CLOSE_WAIT, TIME_WAIT, FIN_WAIT  | 60 | Covers flows that have started the TCP close process.  |
| TCP | ESTABLISHED | 900 | Covers established flows. The connection must be idle for the entire 15 minutes (900 s) for this timeout to be applied. |
| UDP | not applicable | 60 | Covers all UDP flows as the protocol is stateless | 
| ICMP | not applicable | 5 | Covers all ICMP flows as the protocol is stateless | 

[Virtual Network Routes]({{<relref "/docs/domain/virtual-networks/routes">}}) are only evaluated when a flow is first created.  If you modify VPN routes after a flow is established, it will not impact existing flows. If you have an application that reuses the same IP tuple then you will need to make sure it is inactive long enough for the appropriate timeout to occur.


{{<alert color="info">}}Contact Trustgrid support if you need to adjust these settings{{</alert>}}