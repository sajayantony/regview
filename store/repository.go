package store

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/sajayantony/rv/models"
)

// CreateRepository creates a new user in the given account.
func CreateRepository(ctx context.Context, r *models.Repository) error {
	// Start the transaction.
	db := GetDB(ctx)
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Setup the users bucket.
	bkt, err := tx.CreateBucketIfNotExists([]byte("REPOSITORIES"))
	if err != nil {
		return err
	}

	// Generate an ID for the new user.
	userID, err := bkt.NextSequence()
	if err != nil {
		return err
	}

	r.ID = userID

	// Marshal and save the encoded user.
	if buf, err := json.Marshal(r); err != nil {
		return err
	} else if err := bkt.Put([]byte(strconv.FormatUint(r.ID, 10)), buf); err != nil {
		return err
	}

	// Commit the transaction.
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
