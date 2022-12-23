---
Title: "Alert Suppression"
Tags: ["alert suppression", "alarms"]
Date: 2022-12-23
---

{{% pageinfo %}}
Maintenance and other planned disruptive activity can generate a large amount of alerts and notifications for no value.  The Trustgrid system allows you to define Alert Suppression windows during which no alerts are triggered.
{{% /pageinfo %}}

## Define Alert Suppression Window
1. Navigate to Alarms -> Alert Suppression.

2. Enter a message to users. During the window this will be displayed on the Dashboard when users first log in.

![img](/docs/alarms/message-to-users.png)

3. Select a duration between 1 and 4 hours for your window.

![img](/docs/alarms/duration.png)

4. From the calendar select the date that you want the window to start.

![img](/docs/alarms/calendar-time.png)


5. From the drop down select the hour that you want the window to begin. This time will be set based off your browser’s current timezone setting.   This uses the 24 hour clock (e.g. 0 = 12 am, 14 = 2pm)

![img](/docs/alarms/schedule-time1.png)

6. Click the Schedule button. You will see a notification that the window has been scheduled and a notice of when it will begin.

![img](/docs/alarms/alert-supression.png)

To update a scheduled window, make the desired changes and click the Update button.

## Cancel Alert Suppression Window
1. Navigate to Alarms -> Alert Suppression.

2. You should see a notification with the scheduled window at the top. Confirm this is the window you wish to cancel.

3. Click the Cancel Schedule button. 

![img](/docs/alarms/cancel.png)

4. You should be notified that the window has been canceled.

![img](/docs/alarms/removal.png)

## Limitations
- Only a single window can be scheduled at any one time.

- Suppression windows apply to all devices across the entire organization.

