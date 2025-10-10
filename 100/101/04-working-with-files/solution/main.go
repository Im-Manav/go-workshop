package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	Address3  string `json:"address_3"`
	Address4  string `json:"address_4"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
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
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("failed to scan file: %w", err)
		}
		var p Person
		err := json.Unmarshal(scanner.Bytes(), &p)
		if err != nil {
			return fmt.Errorf("failed to unmarshal json: %w", err)
		}
		if p.Age < 30 {
			fmt.Printf("First Name: %s, Last Name: %s\n", p.FirstName, p.LastName)
		}
	}

	return nil
}
