package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lukegrn/days/pkg/hash"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <str_to_hash>\n", os.Args[0])
	}

	pw := os.Args[1]
	hash, err := hash.GenHash(pw)
	if err != nil {
		log.Fatalf("Error hashing string: %s\n", err.Error())
	}

	fmt.Printf("Hashed result for %s: %s\n", pw, hash)
}
