# Transactions

This example demonstrates how to use transactions in SochDB for atomic, consistent operations.

## Features Demonstrated

- Beginning transactions with `db.Begin()`
- Batch operations within a transaction
- Committing transactions (`Commit()`)
- Rolling back transactions (`Abort()`)
- Atomic multi-key updates
- Consistent read snapshots
- Scanning within transactions

## Running the Example

```bash
go run main.go
```

## Expected Output

The example will:
1. Perform a successful batch update and commit
2. Demonstrate rollback with aborted transaction
3. Execute an atomic account transfer
4. Read a consistent snapshot across multiple keys
5. Scan and aggregate values within a transaction

## Key Concepts

- **Atomicity**: All operations in a transaction succeed or fail together
- **Consistency**: Transactions maintain database invariants
- **Isolation**: Concurrent transactions don't interfere
- **Durability**: Committed changes persist across crashes
- **Snapshot Isolation**: Read transactions see a consistent point-in-time view
- **Cleanup**: Always call `Abort()` on read-only transactions or use defer

## Use Cases

- Financial transfers
- Multi-record updates
- Consistent reads across multiple keys
- Rollback on validation failure
