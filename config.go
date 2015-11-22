package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

// Use this stuct to read in repos and software projects
type Configuration struct {
	Repos []string
	Token string
	User  string
}

// Read a config file
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

// Manage the github auth token
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

// This function will try to print versions of repo's that have been starred,
// according to the configuration read in from the config file
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
			// Store latest version retrieved here in BoltDB so we can see if there is a new version
			// Notify us if there is a new version email the repo, version, link to the changelog and the changelog notes
		}
	}
}
