# Stalker

Get notified when your favorite projects and software get updated.

### Getting started

Clone this repo into your GOPATH and run `go install` from the root to create the stalker binary file.

Stalker is (mostly) configuration driven so you probably won't get very far without a config.  By default, Stalker expects to find its configuration in `~/.stalker.json`.  There is a `stalker.json.example` in this repo that can be moved to the correct location to test some of the basic functionality.

After you have installed Stalker and set up a configuration file, simply run the `stalker` command from your terminal to get some basic usage options (as shown below).

One useful command to help users get started is the `stalker print custom` command.  This command will print out a few sample repos and their release tags.

Additionally, you can run the help command for any subcommand to get a brief description of its functionality.

### Usage

```
NAME:
   Stalker - Get notified when your favorite projects are updated

USAGE:
   stalker [global options] command [command options] [arguments...]

VERSION:
   0.0.2

COMMANDS:
   update	Update project repos in BoltDB and email new tags
   print	Print project repos and tags
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -c, --config-file	Specify an alternate config file (not working yet)
   -d, --db		Specify an alternate DB location (not working yet)
   --help, -h		show help
   --version, -v	print the version
```

### Add some repos

To add some repos, update `~/.stalker.json` with some values to test.

```json
{
    "Repos": [
        "github.com/docker/compose",
        "github.com/docker/machine"
    ]
}
```

Add any repos you'd like to follow to this section of the config file.

### Set up starred repo's

There is an option to look at versions of starred repos, rather than discovering them from the configuration file.  To enable this option you will need to set a user in the `config.json` file.

```json
  "Github": {
      "User": "jmreicha"
   }
```

### Use an auth token

To avoid throttling issues, IT IS HIGHLTY RECOMMENDED to add your own github auth token in
the `Token:` section.

```json
  "Github": {
      "Token": "XXX"
   }
```

You can find more information about [GitHub access tokens here](https://help.github.com/articles/creating-an-access-token-for-command-line-use/).

**NOTE:** The Github auth token isn't a requirement but Github will throttle unauthenticted requests which can
cause problems running this tool.  GitHub allows 60 unauthed requests per hour
and 5000 authed requests per hour.

### Configure Stalker to send email

This is a handy feature for receiving emails when Stalker detects a new Github version.

For Gmail, update the `~/.stalker.json` configuration to use your Gmail credentials.

```json
   "Email": {
      "Server": "smtp.gmail.com",
      "Address": "your-email@gmail.com",
      "Password": "XXX"
   }
```

**NOTE:** If you use 2FA with your email account, follow [these instructions](https://support.google.com/accounts/answer/185833) to generate your "app" password.

### Troubleshooting

 * You receive `error: GET https://api.github.com/users/jmreicha/starred: 401 Bad credentials []`.

Check your `~/.stalker.json` configuration.  Either the "Github Token" is not set or it not valid.

 * You receive `ERROR attempting to send mail: 535 5.7.8 Username and Password not accepted.`.

Check your `~/.stalker.json` configuration.  The "Email Password" is invalid or if you have 2FA turned on then you are using the incorrect password.
