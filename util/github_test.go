package util

import (
	"testing"
)

// Globals for testing github API
var (
	user    = "docker"
	project = "docker"
)

// TestRecentTags tests the last 10 tags for a github repo.
func TestRecentTags(t *testing.T) {
	RecentTags(user, project)
}

// TestLatestTag tests a single release tag for a github project.
func TestLatestTag(t *testing.T) {
	LatestTag(user, project)
}

// TestLatestRelease tests the latest releaase for a github project.
func TestLatestRelease(t *testing.T) {
	LatestRelease(user, project)
}

// TestStarredRepos tests for an array of starred repos for a given user.
func TestStarredRepos(t *testing.T) {

	// Override user to get test result data
	user = "jmreicha"
	GetStarredRepos(user)
}
