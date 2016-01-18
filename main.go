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
			Usage: "Specify an alternate config file (not working yet)",
		},
		cli.BoolFlag{
			Name:  "d, db",
			Usage: "Specify an alternate DB location (not working yet)",
		},
	}

	// Commands
	app.Commands = []cli.Command{
		{
			Name:  "update",
			Usage: "Update project repos in BoltDB and email new tags",
			Subcommands: []cli.Command{
				{
					Name:  "starred",
					Usage: "Update and email BoltDB starred repos",
					Action: func(c *cli.Context) {
						util.UpdateStarredRepos(DBName)
						// Test that DB gets updated
						//util.IterateStarredRepos()
					},
				},
				{
					Name:  "custom",
					Usage: "Update and email BoltDB custom repos",
					Action: func(c *cli.Context) {
						util.UpdateCustomRepos(DBName)
						// Test that DB gets updated
						//util.IterateCustomRepos(DBName)
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
					Usage: "Print previously stored custom repos from Bolt DB",
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
}
