# Stalker

Get notified when your favorite projects and software get updated.

### Getting started

To get started, clone this repo into your GOPATH and run `go install` from the root to create a basic stalker binary file.

Stalker is (mostly) configuration driven so you probably won't get very far without a config.  By default, Stalker expects to find its configuration in `~/.stalker.json`.  There is a `config.json.example` in this repo that can be moved to the correct location to test some of the basic functionality.

Once you have insatlled stalker and set up a configuration file simply run `stalker` from your terminal to get some basic help.

One useful command to help get started is the `stalker print custom` command.  This command will print out a few sample repos and their tags.

Additionally, you can run the help command for any subcommand to get a brief description of its functionality.

### Add some repo's

To get started, update `~/.stalker.json` with some values to test.

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
