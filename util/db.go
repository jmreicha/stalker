package util

import (
	"fmt"
	"github.com/boltdb/bolt"
	"strings"
)

// Bolt DB settings
var DBName string = "version.db"
var CustomProjectBucket string = "CustomProject"
var StarredProjectBucket string = "StarredProject"

// UpdateCustomRepos reads in a configuration file, and writes projects and
// their tags to BoltDB.
func UpdateCustomRepos() {
	// Open DB
	db, err := bolt.Open(DBName, 0600, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	// Close DB
	defer db.Close()

	// Create "CustomProject" bucket
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(CustomProjectBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	configuration := ReadConfig()

	// Split user and project in order to parse them separately
	for _, repo := range configuration.Repos {
		repo := strings.Split(repo, "/")
		if repo[0] == "github.com" {
			user := repo[len(repo)-2]
			project := repo[len(repo)-1]
			tag, _ := LatestTag(user, project)

			// Write project to bucket
			db.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(CustomProjectBucket))
				// key=project value=tag
				err := b.Put([]byte(project), []byte(tag))
				return err
			})
		}
	}
}

// IterateCustomRepos looks at what is in BoltDB and prints out the project and
// tag based on custom repo's that have been configured.
func IterateCustomRepos() {

	// Open DB
	db, err := bolt.Open(DBName, 0600, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	// Close DB
	defer db.Close()

	// Iterate over Projects
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(CustomProjectBucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("project=%s, tag=%s\n", k, v)
		}

		return nil
	})
}

// UpdateStarredRepos reads starred repo's for a user and writes projects and
// their tags to BoltDB.
func UpdateStarredRepos() {
	// Open DB
	db, err := bolt.Open(DBName, 0600, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	// Close DB
	defer db.Close()

	// Create "StarredProject" bucket
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(StarredProjectBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	configuration := ReadConfig()
	username := configuration.User
	userRepos := GetStarredRepos(username)

	// Split out user, project and tag
	for _, repo := range userRepos {
		repo := strings.Split(repo, "/")
		user := repo[len(repo)-2]
		project := repo[len(repo)-1]
		tag, _ := LatestTag(user, project)

		// Write project
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(StarredProjectBucket))
			// key=project value=tag
			err := b.Put([]byte(project), []byte(tag))
			return err
		})
	}
}

// IterateStarredRepos looks at what is in BoltDB and prints out the project and
// tag.
func IterateStarredRepos() {

	// Open DB
	db, err := bolt.Open(DBName, 0600, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	// Close DB
	defer db.Close()

	// Iterate over Projects
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(StarredProjectBucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("project=%s, tag=%s\n", k, v)
		}

		return nil
	})
}
