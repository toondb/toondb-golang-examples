// Package main demonstrates basic key-value operations with SochDB Go SDK.
package main

import (
	"fmt"
	"log"

	"github.com/sochdb/sochdb-go/embedded"
)

func main() {
	fmt.Println("=== SochDB Go SDK - Basic Key-Value Operations ===\n")

	db, err := embedded.Open("./data/basic_kv_db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()
	fmt.Println("✅ Database opened successfully")

	// PUT
	key := []byte("user:1001")
	value := []byte(`{"name": "Alice"}`)
	
	if err := db.Put(key, value); err != nil {
		log.Fatalf("Failed to put: %v", err)
	}
	fmt.Printf("Stored: %s\n", key)

	// GET
	retrieved, err := db.Get(key)
	if err != nil {
		log.Fatalf("Failed to get: %v", err)
	}
	fmt.Printf("Retrieved: %s = %s\n", key, retrieved)

	fmt.Println("\n=== Example Complete ===")
}
