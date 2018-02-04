package database

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

var bucketName = []byte("database")

func init() {
	open("database.db")

}

//Open opens our key-value store instance hosted by BoltDB.
//path is the full path to the database file to be opened. The method will create a file if it doesn't exist.
func open(path string) {
	var err error
	db, err = bolt.Open(path, 0640, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(bucketName)
		return err

	})
	if err != nil {
		log.Fatal(err)
	}

}

//Put puts stores a specified value along with the specified key in the BoltDB bucket.
func Put(key string, value string) error {
	if key == "" || value == "" {
		log.Fatal("Values cannot be null.")
	}

	//Encoding the value using a buffer.
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(value)
	if err != nil {
		return nil
	}

	//Putting the encoded value next to the key in the bucket.
	return db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket(bucketName).Put([]byte(key), buf.Bytes())
		return err

	})

}

//Get gets the value from the specified key.
//Returns an error if one occurs, returns the decoded value otherwise.
func Get(key string) error {
	return db.View(func(tx *bolt.Tx) error {
		cursor := tx.Bucket(bucketName).Cursor()
		//Seeks the Bucket with a cursor.
		k, v := cursor.Seek([]byte(key))
		if k == nil || string(k) != key {
			log.Fatal("Key not found in the database.")
			return nil
		}

		decoder := gob.NewDecoder(bytes.NewReader(v))
		var value string
		//Stores the decoded value in a string.
		decoder.Decode(&value)
		fmt.Println(value)
		return nil

	})

}

//Delete deletes the entry with the specified key, if it exists.
func Delete(key string) error {
	return db.Update(func(tx *bolt.Tx) error {
		cursor := tx.Bucket(bucketName).Cursor()
		//Seeks the Bucket with a cursor.
		k, _ := cursor.Seek([]byte(key))
		if k == nil || string(k) != key {
			log.Fatal("Key not found in the database.")
			return nil
		}

		return cursor.Delete()

	})

}

//Close closes the database file.
func Close() error {
	return db.Close()
}
