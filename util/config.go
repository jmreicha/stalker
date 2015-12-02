package util

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

// Configuration is a stuct used to read in repos and software projects from an
// external configuration file.
type Configuration struct {
	Repos []string
	Token string
	User  string
}

// ReadConfig is a helper function for reading in a configuration file.
func ReadConfig() *Configuration {
	homedir := os.Getenv("HOME")
	configpath := (homedir + "/.stalker.json")
	file, _ := os.Open(configpath)

	// Stop execution if config isn't found
	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		fmt.Println("Config " + configpath + " not found!")
		os.Exit(0)
	}

	// Decode json config
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return &configuration
}

// GetToken is a helper function for determining if a Github auth token has been
// set.
func GetToken() string {

	configuration := ReadConfig()
	token := configuration.Token

	if token == "" {
		return "empty"
	}
	return token
}

// IsTokenSet is a helper function that tells you if you have a github auth token set.
func IsTokenSet() {

	configuration := ReadConfig()
	token := configuration.Token

	warn := color.New(color.FgYellow).PrintFunc()
	tokenSet := color.New(color.FgGreen).PrintFunc()

	if token == "" {
		warn("GITHUB AUTH TOKEN NOT SET\n\n")
		fmt.Println("Skipping authenticaiton may create rate limiting issues")
	} else {
		tokenSet("GitHub auth token has been set\n\n")
	}
}

// PrintStarredRepos tries to print tags of repos that have been starred
// according to the "user" configuration setting that is read from the config
// file.
func PrintStarredRepos() {

	IsTokenSet()
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

// PrintFromConfig tries to print versions based on repos that have been read
// in from a config file.
func PrintFromConfig() {

	IsTokenSet()
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
