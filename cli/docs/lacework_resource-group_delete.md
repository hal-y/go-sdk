---
title: "lacework resource-group delete"
slug: lacework_resource-group_delete
hide_title: true
---

## lacework resource-group delete

Delete a resource group

### Synopsis

Delete a single resource group by it's resource group ID.

```
lacework resource-group delete <resource_group_id> [flags]
```

### Options

```
  -h, --help   help for delete
```

### Options inherited from parent commands

```
  -a, --account string      account subdomain of URL (i.e. <ACCOUNT>.lacework.net)
  -k, --api_key string      access key id
  -s, --api_secret string   secret access key
      --api_token string    access token (replaces the use of api_key and api_secret)
      --debug               turn on debug logging
      --json                switch commands output from human-readable to json format
      --nocache             turn off caching
      --nocolor             turn off colors
      --noninteractive      turn off interactive mode (disable spinners, prompts, etc.)
      --organization        access organization level data sets (org admins only)
  -p, --profile string      switch between profiles configured at ~/.lacework.toml
      --subaccount string   sub-account name inside your organization (org admins only)
```

### SEE ALSO

* [lacework resource-group](lacework_resource-group.md)	 - Manage resource groups

