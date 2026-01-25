# Basic Key-Value Operations

This example demonstrates fundamental SochDB operations for storing and retrieving data.

## Features Demonstrated

- Opening an embedded database
- PUT: Storing key-value pairs
- GET: Retrieving values by key
- SCAN: Iterating over keys with prefix matching
- DELETE: Removing keys
- Stats: Database statistics

## Running the Example

```bash
go run main.go
```

## Expected Output

The example will:
1. Open/create a database at `./data/basic_kv_db`
2. Store multiple user and product records
3. Retrieve individual records
4. Scan all keys with "user:" prefix
5. Delete a key and verify deletion
6. Display database statistics

## Key Concepts

- **Embedded Mode**: Database runs in the same process as your application
- **Byte Arrays**: Keys and values are stored as byte slices
- **Transactions**: Read operations use transactions for consistency
- **Cleanup**: Always call `db.Close()` to properly close the database
