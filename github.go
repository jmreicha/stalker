package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var TOKEN string = GetToken()

// List last 10 releases for single repo
func RecentTags(user, project string) {

	var client *github.Client
	config := new(Configuration)
	if config.Token == "empty" {
		client = github.NewClient(nil)
	} else {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: TOKEN})
		tc := oauth2.NewClient(oauth2.NoContext, ts)
		client = github.NewClient(tc)
	}

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

func TestLatestTag(user, project string) {

	var client *github.Client
	config := new(Configuration)
	if config.Token == "empty" {
		client = github.NewClient(nil)
	} else {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: TOKEN})
		tc := oauth2.NewClient(oauth2.NoContext, ts)
		client = github.NewClient(tc)
	}

	releases, _, err := client.Repositories.ListTags(user, project, nil)
	var release github.RepositoryTag
	// Make sure there is a tag
	if len(releases) > 0 {
		release = releases[0]
	}

	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else if release.Name != nil {
		fmt.Println(*release.Name)
	} else {
		fmt.Println("NONE")
	}
}

// Get latest tag (VERSION WILL FAIL SILENTLY if rate limit has been exceeded)
func LatestTag(user, project string) (string, error) {

	var client *github.Client
	config := new(Configuration)
	if config.Token == "empty" {
		client = github.NewClient(nil)
	} else {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: TOKEN})
		tc := oauth2.NewClient(oauth2.NoContext, ts)
		client = github.NewClient(tc)
	}

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

// Only get latest release
func LatestRelease(user, project string) {

	var client *github.Client
	config := new(Configuration)
	if config.Token == "empty" {
		client = github.NewClient(nil)
	} else {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: TOKEN})
		tc := oauth2.NewClient(oauth2.NoContext, ts)
		client = github.NewClient(tc)
	}

	repo, _, err := client.Repositories.GetLatestRelease(user, project)

	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Version: %s\n", *repo.TagName)
	}
	//return *repo.TagName
}

// List the name of starred repo's for a user
func GetStarredRepos(user string) []string {

	var client *github.Client
	config := new(Configuration)
	if config.Token == "empty" {
		client = github.NewClient(nil)
	} else {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: TOKEN})
		tc := oauth2.NewClient(oauth2.NoContext, ts)
		client = github.NewClient(tc)
	}

	starredRepos, _, err := client.Activity.ListStarred(user, nil)
	userStars := make([]string, 0)

	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		for _, repo := range starredRepos {
			//fmt.Printf("%+v\n", *repo.Repository.FullName)
			userStars = append(userStars, *repo.Repository.FullName)
		}
	}
	return userStars
}

// Get latest release for single repo (this function is primarily for testing)
func WriteVersion() {

	client := github.NewClient(nil)
	latest, _, err := client.Repositories.GetLatestRelease("hashicorp", "terraform")

	// Open the DB
	db, err := bolt.Open("version.db", 0600, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	// Close DB
	defer db.Close()

	// Create version bucket
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("Version"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	// Create Repo bucket
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("Repo"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	// Write repo version
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		//err := b.Put([]byte("answer"), []byte(github.Stringify(latest)))
		err := b.Put([]byte("version"), []byte(*latest.TagName))
		return err
	})

	// Print repo version
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("answer"))
		fmt.Printf("Terraform version: %s\n", v)
		return nil
	})

}
