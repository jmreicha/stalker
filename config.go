package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

// Configuration is a  stuct to read in repos and software projects from an
// external configuration file.
type Configuration struct {
	Repos []string
	Token string
	User  string
}

// ReadConfi is a helper for reading in a configuration file.
func ReadConfig() *Configuration {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return &configuration
}

// GetToken is a helper for determining if a Github auth token has been set.
func GetToken() string {
	configuration := ReadConfig()
	token := configuration.Token

	warn := color.New(color.FgYellow).PrintFunc()
	tokenSet := color.New(color.FgGreen).PrintFunc()

	if token == "" {
		warn("GITHUB AUTH TOKEN NOT SET\n\n")
		fmt.Println("Skipping authenticaiton may create rate limiting issues\n")
		return "empty"
	} else {
		tokenSet("GitHub auth token has been set\n\n")
		return token
	}
}

// PrintStarredRepos tries to print tags of repo's that have been starred
// according to the "user" configuration setting that is read in from the
// config file.
func PrintStarredRepos() {
	configuration := ReadConfig()

	username := configuration.User
	userRepos := GetStarredRepos(username)

	for _, repo := range userRepos {
		repo := strings.Split(repo, "/")
		user := repo[len(repo)-2]
		project := repo[len(repo)-1]
		tag, _ := LatestTag(user, project)
		fmt.Println("User: " + user + " Project: " + project + " Tag: " + tag)
	}
}

// PrintFromConfig tries to print versions based on repo's that have been read
// in from a config file.
func PrintFromConfig() {
	configuration := ReadConfig()

	// Split user and project in order to parse them separately
	for _, repo := range configuration.Repos {
		repo := strings.Split(repo, "/")
		if repo[0] == "github.com" {
			user := repo[len(repo)-2]
			project := repo[len(repo)-1]
			tag, _ := LatestTag(user, project)
			fmt.Println("User: " + user + " Project: " + project + " Tag: " + tag)

			// TODO
			// Notify if there is a new version and email the repo, version,
			// link to the changelog and the changelog notes
		}
	}
}
