package initdb

import (
	"log"
	"os"
	"strings"

	"github.com/boltdb/bolt"
)

// Initializes the datastore if the file does not exist
func Initialize(f string) error {

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	const storeSuffix = ".db"
	if !strings.HasSuffix(f, storeSuffix) {
		f = f + storeSuffix
	}
	log.Printf("Initializing datastore: %s", f)

	if _, err := os.Stat(f); os.IsNotExist(err) {
		db, err := bolt.Open(f, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	} else {
		log.Fatalf("%s already exists.", f)
	}

	return nil
}
