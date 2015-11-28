package main

import (
	//"fmt"
	"github.com/codegangsta/cli"
	"github.com/jmreicha/stalker/util"
	"os"
)

func main() {

	// CLI Options
	app := cli.NewApp()
	app.Name = "Stalker"
	app.Usage = "Get notified when your favorite projects are updated"
	// This gets updated by hand
	app.Version = "0.0.1"

	// Flags
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "c, config-file",
			Usage: "Specify an alternate config file",
		},
		cli.BoolFlag{
			Name:  "d, db",
			Usage: "Specify an alternate DB location",
		},
	}

	// Commands
	app.Commands = []cli.Command{
		{
			Name:  "update",
			Usage: "Update project repo's and tags in BoltDB",
			Subcommands: []cli.Command{
				{
					Name:  "starred",
					Usage: "Update BoltDB starred repo's",
					Action: func(c *cli.Context) {
						util.UpdateStarredRepos()
						// Test that DB gets updated
						util.IterateStarredRepos()
					},
				},
				{
					Name:  "custom",
					Usage: "Update BoltDB custom repo's",
					Action: func(c *cli.Context) {
						util.UpdateCustomRepos()
						// Test that DB gets updated
						util.IterateCustomRepos()
					},
				},
			},
		},
		{
			Name:  "print",
			Usage: "Print project repo's and tags",
			Subcommands: []cli.Command{
				{
					Name:  "starred",
					Usage: "Print starred repo's",
					Action: func(c *cli.Context) {
						util.PrintStarredRepos()
					},
				},
				{
					Name:  "custom",
					Usage: "Print from configuration",
					Action: func(c *cli.Context) {
						util.PrintFromConfig()
					},
				},
				{
					Name:  "custom-db",
					Usage: "Print custom repo's from Bolt DB",
					Action: func(c *cli.Context) {
						//util.PrintFromConfig()
					},
				},
				{
					Name:  "starred-db",
					Usage: "Print starred repo's from Bolt DB",
					Action: func(c *cli.Context) {
						//util.PrintFromConfig()
					},
				},
			},
		},
	}

	// Actions
	app.Action = func(c *cli.Context) {
		// Show help if no args are passed
		cli.ShowAppHelp(c)
	}

	app.Run(os.Args)

	/* Testing functions
	RecentTags()						// Function to list paged results of recent tags
	LatestTag("rancher", "rancher")		// Get the latest github tag for a project
	LatestRelease("rancher", "rancher")	// Get the latest github release for a project
	TestLatestTag("lukasz-madon", "awesome-remote-job")  //Case where there is no tag
	*/
}
