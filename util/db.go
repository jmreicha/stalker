package util

import (
	"fmt"
	"github.com/boltdb/bolt"
	"os"
	"strings"
)

// Bolt DB settings
var (
	homedir       = os.Getenv("HOME")
	DBName        = homedir + "/version.db"
	CustomBucket  = "CustomProject"
	StarredBucket = "StarredProject"
)

// OpenDB opens up a Bolt DB connection.
func OpenDB(DBName string) *bolt.DB {
	// Open DB
	db, err := bolt.Open(DBName, 0600, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return db
}

// UpdateCustomRepos reads in a configuration file, and writes projects and
// their tags to BoltDB.  A message will be printed if a new tag has been added.
func UpdateCustomRepos(DBName string) {

	// Open DB
	db := OpenDB(DBName)
	defer db.Close()

	// Create "CustomProject" bucket if needed
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(CustomBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	configuration := ReadConfig()
	IsTokenSet()
	// Count the new tags
	var updateCount int
	// Storing the  new tags
	var new_tags []string

	// Split user and project in order to parse them separately
	for _, repo := range configuration.Repos {
		repo := strings.Split(repo, "/")
		if repo[0] == "github.com" {
			user := repo[len(repo)-2]
			project := repo[len(repo)-1]
			tag, _ := LatestTag(user, project)

			// Write project to bucket if there is a new tag
			db.Update(func(tx *bolt.Tx) error {
				var err error
				b := tx.Bucket([]byte(CustomBucket))
				v := b.Get([]byte(project))
				// Convert the tag to string for camparing
				s := string(v)
				// Check if the tag in the bucket is current, update if not
				if s != tag {
					// key=project value=tag
					err = b.Put([]byte(project), []byte(tag))
					fmt.Println(project + " has new tag " + tag)
					// TODO print old project tag and also get release notes or
					// changelog info
					new_tags = append(new_tags, project+": "+tag)
					updateCount++
					return err
				}
				return err
			})
		}
	}
	// Tally number of updated repos
	fmt.Printf("%d Repos updated\n", updateCount)
	// If there is a new tag (new_tags not empty), call the alert function to
	// email the project name with the new tag
	if len(new_tags) > 0 {
		AlertNewProjectTag(new_tags)
	}
}

// IterateCustomRepos looks at what is in BoltDB and prints out the project and
// tag based on custom repos that have been configured.
func IterateCustomRepos(DBName string) {

	// Open DB
	db := OpenDB(DBName)
	defer db.Close()

	// Iterate over Projects
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(CustomBucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("project=%s, tag=%s\n", k, v)
		}

		return nil
	})
}

// UpdateStarredRepos reads starred repo's for a user and writes projects and
// their tags to BoltDB.
func UpdateStarredRepos(DBName string) {

	// Open DB
	db := OpenDB(DBName)
	defer db.Close()

	// Create "StarredProject" bucket if needed
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(StarredBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	configuration := ReadConfig()
	IsTokenSet()
	// Count the new tags
	var updateCount int
	// Storing the  new tags
	var new_tags []string

	username := configuration.Github.User
	userRepos := GetStarredRepos(username)

	// Split out user, project and tag
	for _, repo := range userRepos {
		repo := strings.Split(repo, "/")
		user := repo[len(repo)-2]
		project := repo[len(repo)-1]
		tag, _ := LatestTag(user, project)

		// Write project to bucket if there is a new tag
		db.Update(func(tx *bolt.Tx) error {
			var err error
			b := tx.Bucket([]byte(StarredBucket))
			v := b.Get([]byte(project))
			// Convert the tag to string for camparing
			s := string(v)
			// Check if the tag in the bucket is current
			if s != tag {
				// key=project value=tag
				err = b.Put([]byte(project), []byte(tag))
				fmt.Println(project + " has new tag " + tag)
				// TODO print old project tag and also get release notes or
				// changelog info
				new_tags = append(new_tags, project+": "+tag)
				updateCount++
				return err
			}
			return err
		})
	}
	// Tally number of updated repos
	fmt.Printf("%d Repos updated\n", updateCount)
	// If there is a new tag (new_tags not empty), call the alert function to
	// email the project name with the new tag
	if len(new_tags) > 0 {
		AlertNewProjectTag(new_tags)
	}
}

// IterateStarredRepos looks at what is in BoltDB and prints out the project and
// tag.
func IterateStarredRepos() {

	// Open DB
	db := OpenDB(DBName)
	defer db.Close()

	// Iterate over Projects
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(StarredBucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("project=%s, tag=%s\n", k, v)
		}

		return nil
	})
}
