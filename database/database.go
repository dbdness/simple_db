package database

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

//Open opens our key-value store instance hosted by BoltDB.
//path is the full path to the database file to be opened. The method will create a file if it doesn't exist.
func Open(path string) {
	db, err := bolt.Open(path, 0640, nil)
	if err != nil {
		log.Panic()
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte("database"))
		fmt.Println("--No previous database found. A new one is created instead--")
		return err

	})
	if err != nil {
		log.Panic()
	}

	fmt.Println("--Database is successfully opened--")
}
