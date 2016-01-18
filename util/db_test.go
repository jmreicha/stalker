package util

import (
	"os"
	"testing"
)

// Test DB settings
var (
	TestDBName = "/tmp/test.db"
)

// TestDBExists checks to see if a test database has been created.
func TestDBExists(t *testing.T) {

	if _, err := os.Stat(TestDBName); os.IsNotExist(err) {
		t.Error(TestDBName + " not found!")
	} else {
		t.Log("DB found")
	}
}

// TestOpenDB tries to open a BoltDB database.
func TestOpenDB(t *testing.T) {
	OpenDB(TestDBName)
}

// TestUpdateCustomRepos tries to update Bolt with values from a custom
// configuration.
//func TestUpdateCustomRepos(t *testing.T) {
//	UpdateCustomRepos(TestDBName)
//}

// TestIterateCustomRepos tries to print custom configured values that have been
// written to BoltDB.
//func TestIterateCustomRepos(t *testing.T) {
//	IterateCustomRepos(TestDBName)
//}

// TestUpdateStarredRepos tries to update Bolt with values for a specific github
// user, based on custom configuration.
//func TestUpdateStarredRepos(t *testing.T) {
//	UpdateStarredRepos(TestDBName)
//}

// TestIterateStarredRepos trieso to print values from Bolt for a specific
// github user, based on custom configuration.
//func TestIterateStarredRepos(t *testing.T) {
//	IterateStarredRepos()
//}
