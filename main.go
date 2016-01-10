package main

import (
	"github.com/codegangsta/cli"
	"github.com/jmreicha/stalker/util"
	"os"
)

// Global defaults
var (
	homedir = os.Getenv("HOME")
	DBName  = homedir + "/version.db"
)

func main() {

	// CLI Options
	app := cli.NewApp()
	app.Name = "Stalker"
	app.Usage = "Get notified when your favorite projects are updated"
	// This gets updated manaully
	app.Version = "0.0.2"

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
			Usage: "Update project repos and tags in BoltDB",
			Subcommands: []cli.Command{
				{
					Name:  "starred",
					Usage: "Update BoltDB starred repos",
					Action: func(c *cli.Context) {
						util.UpdateStarredRepos(DBName)
						// Test that DB gets updated
						util.IterateStarredRepos()
					},
				},
				{
					Name:  "custom",
					Usage: "Update BoltDB custom repos",
					Action: func(c *cli.Context) {
						util.UpdateCustomRepos(DBName)
						// Test that DB gets updated
						util.IterateCustomRepos(DBName)
					},
				},
			},
		},
		{
			Name:  "print",
			Usage: "Print project repos and tags",
			Subcommands: []cli.Command{
				{
					Name:  "starred",
					Usage: "Print starred repos",
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
					Usage: "Print custom repos from Bolt DB",
					Action: func(c *cli.Context) {
						util.IterateCustomRepos(DBName)
					},
				},
				{
					Name:  "starred-db",
					Usage: "Print starred repos from Bolt DB",
					Action: func(c *cli.Context) {
						util.IterateStarredRepos()
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
	/*
		util.AlertNewProjectTag()
	*/

	/* Testing functions
	RecentTags()						// Function to list paged results of recent tags
	LatestTag("rancher", "rancher")		// Get the latest github tag for a project
	LatestRelease("rancher", "rancher")	// Get the latest github release for a project
	*/
}
