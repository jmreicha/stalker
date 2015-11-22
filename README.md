# Stalker

Get notified when your favorite projects and software get updated.

### Add repo's

To get started, either move `config.json.example` to `config.json` and
update with values to test.

```json
{
    "Repos": [
        "github.com/docker/compose",
        "github.com/docker/machine"
    ]
}
```

Add any repo's you'd like to follow to this section of the config file.

### Set up starred repo's

There is an option to look at version of starred repo's, rather than reading directly from the configuration file.  To enable this option you will need to set a user in the `config.json` file.

```json
{
  "User": "jmreicha"
}
```

### Use an auth token

To avoid throttling issues you will need to add your own github auth token in
the `Token:` section.

```json
{
  "Token": "xxx"
}
```

You can find more information about [GitHub access tokens here](https://help.github.com/articles/creating-an-access-token-for-command-line-use/).

**NOTE:** The Github auth token isn't a requirement but Github will throttle unauthenticted requests which can
cause problems running this tool.  GitHub allows 60 unauthed requests per hour
and 5000 authed requests per hour.
