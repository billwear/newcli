package main

import (
	"fmt"
	"os"
	"os/user"
	"log"
)

func main() {
	var username string

	// Check if a username is provided in the command line arguments
	if len(os.Args) > 1 {
		username = os.Args[1]
	} else {
		// Get current user if no username is provided
		currUser, err := user.Current()
		if err != nil {
			log.Fatalf("Failed to get current user: %s", err)
		}
		username = currUser.Username
	}

	// Lookup the user
	u, err := user.Lookup(username)
	if err != nil {
		log.Fatalf("Failed to find user: %s", err)
	}

	// Print user details
	fmt.Println("User Details:")
	fmt.Println("        Username:", u.Username)
	fmt.Println("             UID:", u.Uid)
	fmt.Println("             GID:", u.Gid)
	fmt.Println("            Name:", u.Name)
	fmt.Println("  Home Directory:", u.HomeDir)
}

