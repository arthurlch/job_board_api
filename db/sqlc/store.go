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
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// Transaction options, setting the isolation level
	txOptions := &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	}
	tx, err := store.db.BeginTx(ctx, txOptions)
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return err
	}

	// Make sure to rollback if function or commit fails
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			log.Printf("Panic occured, rolling back: %v", p)
		} else if err != nil {
			tx.Rollback()
			log.Printf("Transaction error, rolling back: %v", err)
		}
	}()

	// Perform operations within the transaction
	q := New(tx)
	err = fn(q)
	if err != nil {
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return err
	}

	log.Println("Transaction successfully committed")
	return nil
}
