package main

import (
	"fmt"
	"log"
	"os"

	database "./database"
)

func main() {
	//Closes the database file after all code has run.
	defer database.Close()

	usercommand := os.Args[1]
	if usercommand == "" {
		log.Panic("No arguments given.")
		return
	}

	switch usercommand {
	case "--set":
		database.Put(os.Args[2], os.Args[3])
		fmt.Println("---Key and value successfully stored---")
	case "--get":
		database.Get(os.Args[2])
	case "--delete":
		database.Delete(os.Args[2])
		fmt.Println("--Entry successfully deleted--")
	default:
		fmt.Println("unknown command: '" + usercommand + "'")
		printUsage()
	}

	database.Close()

}

func printUsage() {
	fmt.Println("usage: 	./simple_db [--set <key> <value>] [--get <value>] [--delete <key>]")
}
