---
title: Viewing Agent Debug Logs
linkTitle: Debug Logs
description: Details where to view the Trustgrid agent's debug logs for troubleshooting purposes
---

The Trustgrid agent generates debug logs that can be useful for troubleshooting issues. Many log entries are fairly self explanatory, but some may require the assistance to Trustgrid support to fully understand their context.

## Viewing Debug Logs
### Viewing Log Files
The agent log files are located in `/var/log/trustgrid/agent`. The active log file is `tg-agent.log`. Logs rotate every day, with the last 9 days kept in separate files appended with a number. For example, the file `tg-agent.log.1` contains the logs from the previous day.

### Viewing Logs with `journalctl` on Ubuntu
You can also view log entries with `journalctl -u tg-agent.service`. This command shows the logs for the Trustgrid agent service and allows you to filter and search through the logs easily. 

For example, you can use `journalctl -u tg-agent.service --since "2023-01-01"` to view logs since a specific date.

Or, you can use `journalctl -u tg-agent.service -f` to view logs in real-time.