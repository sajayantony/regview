package store

import (
	"log"
	"os"
	"strings"

	"context"

	"github.com/boltdb/bolt"
)

type contextKey string

var (
	contextKeyDB = contextKey("db")
)

// Initializes the datastore if the file does not exist
func Initialize(f string) error {

	log.Printf("Initializing datastore: %s", f)
	f = normalizeFile(f)
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

//Open the data store and add to context
func Open(ctx context.Context, f string) (context.Context, *bolt.DB) {
	f = normalizeFile(f)
	db, err := bolt.Open(f, 0600, nil)
	if err != nil {
		log.Fatalf("Unable to open the data store %s", f)
	}

	ctx = context.WithValue(ctx, contextKeyDB, db)
	return ctx, db
}

// GetDB from context
func GetDB(ctx context.Context) *bolt.DB {
	return ctx.Value(contextKeyDB).(*bolt.DB)
}

func normalizeFile(f string) string {
	const storeSuffix = ".db"
	if !strings.HasSuffix(f, storeSuffix) {
		f = f + storeSuffix
	}
	return f
}
