package main

import (
	"fmt"
	"log"
	"os"

	database "./database"
)

func main() {
	usercommand := os.Args[1]
	if usercommand == "" {
		log.Panic("No arguments given.")
		return
	}

	switch usercommand {
	case "set":
		database.Put(os.Args[2], os.Args[3])
		fmt.Println("Key and value successfully stored")
	case "get":
		fmt.Println(database.Get(os.Args[2]))
	}

}
