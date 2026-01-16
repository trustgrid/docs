---
Title: "Service Users"
---

{{% pageinfo %}}
Service Users are specialized accounts designed for machine-to-machine interactions and automated integrations. Unlike standard users, Service Users cannot log in to the Trustgrid Portal UI. instead, they are used to invoke the Trustgrid API programmatically.

Service Users can be assigned specific permissions via policies, allowing granular control over what resources an automation script or external system can access.
{{% /pageinfo %}}

{{<alert color="info">}} 
Service User functionality must be enabled by Trustgrid. Contact support to use this feature. It will be generally available in an upcoming Portal release.
{{</alert>}}

## Managing Service Users
To manage Service Users, navigate to the Service Users section in the Portal.

{{< tgimg src="service-users.png" width="50%" caption="Service Users list view." >}}

### Creating a Service User
1. Click the Add Service User button.
1. Enter a unique Name. Names must consist of alphanumeric characters and hyphens.
1. (Optional) Select initial Policies to attach.
1. Click Submit.

### Enabling and Disabling
Service Users can be temporarily disabled without deleting them. This immediately revokes their API access.

* **To Disable**: Select one or more active users in the list and click Disable.
* **To Enable**: Select one or more inactive users in the list and click Enable.

### Deleting
To permanently remove a Service User and all associated API keys:

1. Select the user(s) from the list.
1. Click the Delete icon.

### Permissions & Policies
Service Users rely on [Policies]({{<relref "/docs/user-management/policies">}}) to determine what actions they can perform.

On the Service User's detail page, you can manage these policies in the Attached Policies section.

{{< tgimg src="attached-policies.png" width="50%" caption="Policies attached to the cicd service user." >}}

* **Attach Policy**: Click Attach Policy to select an existing policy to apply to this user.
* **Detach Policy**: Click Detach next to a policy to remove those permissions from the user.

## API Access
To use a Service User for API calls, you must generate an API Token. This is done on the Service User's detail page in the API Access section.

### Generating Tokens
Click Generate API keys (or Regenerate if keys already exist).
The Client ID and Client Secret will be displayed.

{{<alert color="info">}} 
Copy the Client Secret immediately. It is only shown once and cannot be retrieved later. If you lose it, you will need to regenerate the keys.
{{</alert>}}

{{< tgimg src="api-token.png" width="50%" caption="(Fake) API access credentials." >}}

### Authenticating Requests
Include the generated credentials in the HTTP headers of your API requests. The format is `trustgrid-token {ClientId}:{ClientSecret}`.

Header Example:

```
Authorization: trustgrid-token 12345-abcde:secret-key-content-here
Authorization Header Format
The UI provides a copy-pasteable example of the header:
Authorization: trustgrid-token <ClientId>:<Secret>
```

## Best Practices
* **Least Privilege**: Create specific Policies for your integrations rather than reusing broad admin policies. Assign only the permissions necessary for the specific task the Service User performs.
* **Key Rotation**: Regularly rotate API keys by using the Regenerate API keys function. This invalidates the old secret immediately.
* **Unique Identity**: Create separate Service Users for different integrations (e.g., one for "CI/CD Pipeline" and another for "Monitoring System"). This makes audit logs clearer and allows you to revoke access for one system without affecting others.
* **Disable Unused Users**: If an integration is paused or decommissioned, disable the Service User immediately.