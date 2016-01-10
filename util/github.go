package util

import (
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Global variable for setting a github token
var TOKEN = GetToken()

// CreateClientConnection is a helper function for creating a connection to the
// Github API based on whether or not an auth token is supplied.
func CreateClientConnection() *github.Client {
	var client *github.Client
	config := new(Configuration)
	if config.Github.Token == "empty" {
		client = github.NewClient(nil)
		return client
	} else {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: TOKEN})
		tc := oauth2.NewClient(oauth2.NoContext, ts)
		client = github.NewClient(tc)
		return client
	}
}

// RecentTags print the last 10 releases for a single repo.
func RecentTags(user, project string) {

	var client = CreateClientConnection()

	opt := &github.ListOptions{Page: 1, PerPage: 10}
	releases, _, err := client.Repositories.ListTags(user, project, opt)

	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		for _, release := range releases {
			fmt.Printf("%+v\n", *release.Name)
		}
	}
}

// LatestTag returns the latest tag for a given user and project (VERSION WILL
// FAIL SILENTLY if rate limit has been exceeded).
func LatestTag(user, project string) (string, error) {

	var client = CreateClientConnection()

	releases, _, err := client.Repositories.ListTags(user, project, nil)
	var release github.RepositoryTag
	// Make sure there is a tag
	if len(releases) > 0 {
		release = releases[0]
	}

	if err != nil {
		return "", err
	} else if release.Name == nil {
		// Set a custom tag if there are none present
		return "NONE", nil
	} else {
		return *release.Name, nil
	}
}

// LatestRelease prints the latest release for a given user and project.
func LatestRelease(user, project string) {

	var client = CreateClientConnection()

	repo, _, err := client.Repositories.GetLatestRelease(user, project)

	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Version: %s\n", *repo.TagName)
	}
}

// GetStarredRepos returns an array of starred repos for a given user.
func GetStarredRepos(user string) []string {

	var client = CreateClientConnection()

	starredRepos, _, err := client.Activity.ListStarred(user, nil)
	var userStars []string

	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		for _, repo := range starredRepos {
			userStars = append(userStars, *repo.Repository.FullName)
		}
	}
	return userStars
}
