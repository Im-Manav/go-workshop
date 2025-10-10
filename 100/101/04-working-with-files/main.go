package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address1  string
	Address2  string
	Address3  string
	Address4  string
	Postcode  string
	Country   string
}

func main() {
	err := run()
	if err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
	}
}

func run() error {
	f, err := os.Open("data.jsonl")
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	// You can read files line by line using a Scanner.
	// This is efficient, because it uses buffered I/O - only reading part of the file into memory at a time.
	// In modern computers, a lot of time is spent waiting for the disk to read data, making syscalls, or allocating memory, so buffering these operations is a big win.
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Err() != nil {
			return fmt.Errorf("failed to scan file: %w", err)
		}
		var p Person
		err := json.Unmarshal(scanner.Bytes(), &p)
		if err != nil {
			return fmt.Errorf("failed to unmarshal json: %w", err)
		}
		// #v prints out the Go syntax representation of the value.
		fmt.Printf("%#v\n", p)
		// You can also access individual fields.
		fmt.Printf("First Name: %s, Last Name: %s\n", p.FirstName, p.LastName)
	}

	return nil
}
