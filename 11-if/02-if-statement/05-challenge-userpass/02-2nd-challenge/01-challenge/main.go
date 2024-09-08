package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOK = "Access granted to %q.\n"
)

var users = []string{"jack", "inanc"}
var passwords = []string{"1888", "1879"}

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	for i, user := range users {
		if u == user {
			if p == passwords[i] {
				fmt.Printf(accessOK, u)
				return
			}
			fmt.Printf(errPwd, u)
			return
		}
	}

	fmt.Printf(errUser, u)
}
