package util

import (
	"encoding/json"
	"os"
	"testing"
)

// TestConfigsExists test for the existence of a configuration file.
func TestConfigExists(t *testing.T) {
	homedir := os.Getenv("HOME")
	configpath := (homedir + "/.stalker.json")

	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		t.Error(configpath + " not found!")
	} else {
		t.Log("found config " + configpath)
	}
}

// TestConfigIsJSON test a configuration file to make sure it is valid JSON.
func TestConfigIsJSON(t *testing.T) {
	homedir := os.Getenv("HOME")
	configpath := (homedir + "/.stalker.json")
	file, _ := os.Open(configpath)

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		t.Error(err)
	} else {
		t.Log("config is valid JSON")
	}
}

// TestTokenIsSet checks to see if there is a token section configured.
func TestTokenIsSet(t *testing.T) {
	configuration := ReadConfig()
	token := configuration.Token

	if token == "" {
		t.Error("Token misconfigured")
	}

	// A dumb way to check if a dummy token has been used
	if len(token) < 16 {
		t.Error("Token misconfigured")
	}

	t.Log("Token set")
}

// TestUserIsSet checks to see if there is a user section configured.
func TestUserIsSet(t *testing.T) {
	configuration := ReadConfig()
	user := configuration.User

	if user == "" {
		t.Error("User misconfigured")
	} else {
		t.Log("User set")
	}
}

// TestPrintStarredRepos tries to print starred repos for a user based on custom
// configuration.
func TestPrintStarredRepos(t *testing.T) {
	PrintStarredRepos()
}

// TestPrintFromConfig tries to print custom repos base on custom configuration.
func TestPrintFromConfig(t *testing.T) {
	PrintFromConfig()
}
