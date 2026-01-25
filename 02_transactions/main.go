// Package main demonstrates transaction handling with SochDB Go SDK.
//
// This example covers:
// - Beginning transactions
// - Batch operations within transactions
// - Transaction commit and rollback
// - ACID properties
// - Error handling
package main

import (
	"fmt"
	"log"

	"github.com/sochdb/sochdb-go/embedded"
)

func main() {
	fmt.Println("=== SochDB Go SDK - Transactions ===\n")

	// Open database
	db, err := embedded.Open("./data/transactions_db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()
	fmt.Println("✅ Database opened successfully")

	// Example 1: Successful Transaction (Commit)
	fmt.Println("\n1. Successful Transaction (Commit)")
	fmt.Println("-----------------------------------")

	txn := db.Begin()

	// Batch operations within transaction
	updates := map[string]string{
		"account:alice":   "1000",
		"account:bob":     "500",
		"account:charlie": "750",
	}

	for k, v := range updates {
		if err := txn.Put([]byte(k), []byte(v)); err != nil {
			txn.Abort()
			log.Fatalf("Failed to put in transaction: %v", err)
		}
		fmt.Printf("  Staged: %s = %s\n", k, v)
	}

	// Commit transaction
	if err := txn.Commit(); err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}
	fmt.Println("✅ Transaction committed successfully")

	// Verify data was persisted
	val, _ := db.Get([]byte("account:alice"))
	fmt.Printf("Verified: account:alice = %s\n", val)

	// Example 2: Failed Transaction (Rollback)
	fmt.Println("\n2. Failed Transaction (Rollback)")
	fmt.Println("--------------------------------")

	txn2 := db.Begin()

	// Stage some operations
	txn2.Put([]byte("temp:key1"), []byte("temporary_value"))
	txn2.Put([]byte("temp:key2"), []byte("will_be_rolled_back"))
	fmt.Println("  Staged temporary keys")

	// Abort transaction (rollback)
	txn2.Abort()
	fmt.Println("✅ Transaction aborted (rolled back)")

	// Verify data was NOT persisted
	_, err = db.Get([]byte("temp:key1"))
	if err != nil {
		fmt.Println("Verified: temporary keys were not persisted")
	}

	// Example 3: Transfer Transaction (Atomic)
	fmt.Println("\n3. Atomic Transfer Transaction")
	fmt.Println("------------------------------")

	// Read current balances
	aliceBalance, _ := db.Get([]byte("account:alice"))
	bobBalance, _ := db.Get([]byte("account:bob"))
	fmt.Printf("Before: Alice=%s, Bob=%s\n", aliceBalance, bobBalance)

	// Perform atomic transfer
	txn3 := db.Begin()

	// Debit Alice's account
	txn3.Put([]byte("account:alice"), []byte("800"))

	// Credit Bob's account
	txn3.Put([]byte("account:bob"), []byte("700"))

	// Record transfer
	txn3.Put([]byte("transfer:001"), []byte(`{"from":"alice","to":"bob","amount":200}`))

	if err := txn3.Commit(); err != nil {
		log.Fatalf("Transfer failed: %v", err)
	}

	// Verify transfer
	aliceBalance, _ = db.Get([]byte("account:alice"))
	bobBalance, _ = db.Get([]byte("account:bob"))
	fmt.Printf("After:  Alice=%s, Bob=%s\n", aliceBalance, bobBalance)
	fmt.Println("✅ Atomic transfer completed")

	// Example 4: Read Transaction (Consistent Snapshot)
	fmt.Println("\n4. Read Transaction (Consistent View)")
	fmt.Println("-------------------------------------")

	readTxn := db.Begin()
	defer readTxn.Abort()

	// Get consistent snapshot of all accounts
	accounts := []string{"account:alice", "account:bob", "account:charlie"}
	fmt.Println("Consistent snapshot:")

	for _, acc := range accounts {
		val, err := readTxn.Get([]byte(acc))
		if err == nil {
			fmt.Printf("  %s = %s\n", acc, val)
		}
	}
	fmt.Println("✅ Read consistent data snapshot")

	// Example 5: Scan within Transaction
	fmt.Println("\n5. Scan within Transaction")
	fmt.Println("--------------------------")

	scanTxn := db.Begin()
	defer scanTxn.Abort()

	iter := scanTxn.ScanPrefix([]byte("account:"))
	defer iter.Close()

	total := 0
	count := 0
	for {
		k, v, ok := iter.Next()
		if !ok {
			break
		}
		// Parse balance (simplified)
		var balance int
		fmt.Sscanf(string(v), "%d", &balance)
		total += balance
		count++
		fmt.Printf("  %s = %s\n", k, v)
	}
	fmt.Printf("✅ Total balance across %d accounts: %d\n", count, total)

	fmt.Println("\n=== Example Complete ===")
}
