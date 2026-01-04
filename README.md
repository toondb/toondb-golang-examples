# ToonDB Go SDK Examples

This directory contains examples for using the ToonDB Go SDK.

## Prerequisites

- [Go](https://go.dev/) 1.21+ installed.
- ToonDB Go SDK: `github.com/toondb/toondb-go`

## 📂 Repository Structure

```
toondb-golang-cd basic_kv/       # Basic key-value operations
├── sql_check/      # SQL operations and data integrity
├── rag/            # Complete RAG (Retrieval-Augmented Generation) system
└── README.md
```

## Examples

### 1. Basic Key-Value Operations
Located in `basic_kv/`. Demonstrates how to open a database, put, get, and scan keys.

```bash
cd basic_kv
go run main.go
```

### 2. SQL Integrity Check
**Location**: `rag/`. Demonstrates SQL table creation and data integrity verification.

```bash
cd sql_check
go run main.go
```
