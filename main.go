package main

import (
	"github.com/jmreicha/stalker/util"
)

func main() {

	util.PrintStarredRepos()
	//util.PrintFromConfig()

	//util.UpdateStarredRepos()
	//util.IterateStarredRepos()
	/* Testing functions
	GetToken()							// Helper function to list api token
	RecentTags()						// Function to list paged results of recent tags
	GetStarredRepos("jmreicha")			// List starred repo for a given user
	LatestTag("rancher", "rancher")		// Get the latest github tag for a project
	LatestRelease("rancher", "rancher")	// Get the latest github release for a project
	TestLatestTag("lukasz-madon", "awesome-remote-job")  //Case where there is no tag
	*/

}
