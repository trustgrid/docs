---
title: "Password Complexity and MFA Requirements"
---

Invited users utilize Trustgrid's native authentication provider for password complexity and multi-factor authentication. 

> Users authenticating via a custom [identity provider]({{<ref "/docs/idps">}}) will use that system's password complexity and MFA requirements.

## Password Complexity

* At least 10 characters in length
* Lower case (a-z), upper case (A-Z) and numbers (0-9)
* Special characters (-!@#$%^&*)
* No more than 2 identical characters in a row
* Cannot reuse the last 5 passwords
* Cannot be in the list of the [10,000](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/10k-most-common.txt) most common passwords
* Cannot include the user’s name, username, or email address

## Multi-factor Authentication Requirements

The Trustgrid native provider utilizes a One-time Password via applications such as Google Authenticator or similar to provide multi-factor authentication during login.

