package db

import (
	"context"
	"database/sql"
	"log"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	const maxRetries = 3 
	var err error
	for i := 0; i < maxRetries; i++ {
			tx, err := store.db.BeginTx(ctx, nil)
			if err != nil {
					log.Printf("Failed to begin transaction: %v", err)
					continue 
			}
			log.Println("Transaction started.")

			q := New(tx)
			shouldCommit := true
			defer func() {
					if p := recover(); p != nil {
							log.Println("Panic occurred, rolling back transaction.")
							tx.Rollback()
							panic(p) 
					} else if err != nil || !shouldCommit {
							log.Printf("Error occurred or rollback flag set: %v, rolling back transaction.", err)
							tx.Rollback()
					} else {
							log.Println("Committing transaction.")
							if commitErr := tx.Commit(); commitErr != nil {
									log.Printf("Failed to commit transaction: %v", commitErr)
									err = commitErr
							} else {
									log.Println("Transaction committed successfully.")
							}
					}
			}()
			err = fn(q)
			if err != nil {
					log.Printf("Transaction function failed: %v", err)
					shouldCommit = false
					continue
			}
			break
	}
	return err
}