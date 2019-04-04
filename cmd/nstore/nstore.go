package main

import (
	"syscall"
	"fmt"
	"os"
	"nativestore"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	args := os.Args[1:]
	nArgs := len(args)
	switch cmd := args[0]; cmd {
	case "set":
		if nArgs < 4 {
			fmt.Printf("Not enough parameters to create a new secret. Got %d, need 3\n\n", nArgs-1)
			usage(os.Args[0])
		}
		fmt.Printf("Enter secret to store: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Printf("Failed to read password: %v", err)
			os.Exit(1)
		}

		err = nativestore.Set(args[1], args[2], args[3], string(bytePassword))
		if err != nil {
			fmt.Printf("Failed to create the new secret: %v", err)
			os.Exit(1)
		}
		fmt.Printf("\nCreated!\n");

	case "get":
		if nArgs < 3 {
			fmt.Printf("Not enough parameters to get a secret. Got %d, need 2\n\n", nArgs-1)
			usage(os.Args[0])
		}
		user, secret, err := nativestore.Get(args[1], args[2])
		if err != nil {
			fmt.Printf("Couldn't get secret: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(user, ":", secret)
		
	case "del":
		if nArgs < 3 {
			fmt.Printf("Not enough parameters to delete a secret. Got %d, need 2\n\n", nArgs-1)
			usage(os.Args[0])
		}		
		err := nativestore.Del(args[1], args[2])
		if err != nil {
			fmt.Printf("Couldn't delete secret: %v", err)
			os.Exit(1)
		}
		fmt.Printf("Deleted.\n")
		default:
			usage(os.Args[0])
	}

}

func usage(cmd string) {
	fmt.Printf("Usage %s: [set|get|del] params....\n", cmd)
	fmt.Printf("   set <lbl> <url> <user>\n")
	fmt.Printf("   get <lbl> <url>\n")
	fmt.Printf("   del <lbl> <url>\n")
	os.Exit(1)
}
