# stalker

Get notified when your favorite projects and software get updated.

To get started, either move or copy `config.json.example` to `config.json` and
update the file with values to test.

To avoid throttling issues you will need to add your own github auth token in
the `Token:` section.  Likewise if there are repo's you'd like to follow you
can add them in the config file.

NOTE: The Github auth token isn't a requirement, if none if presented it will
be skipped.  However, Github will throttle unauthenticted requests which can
cause problems running this tool.  GitHub allow 60 unauthed requests per hour
and 5000 authed requests per hour.
