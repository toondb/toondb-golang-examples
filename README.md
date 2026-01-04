# ToonDB Go Examples

Official examples for using ToonDB with Go. This repository demonstrates best practices for integrating ToonDB into your Go applications.

## 📂 Repository Structure

```
examples/
├── basic_kv/       # Basic key-value operations
├── sql_check/      # SQL operations and data integrity
└── rag/            # Complete RAG (Retrieval-Augmented Generation) system
```

## 🚀 Quick Start

### Prerequisites

- Go 1.21+ installed
- ToonDB server binary (for IPC mode) or embedded mode support

### Installation

```bash
go get github.com/toondb/toondb-go
```

## 📚 Examples

### 1. Basic Key-Value Operations

**Location**: `examples/basic_kv/`

Demonstrates fundamental ToonDB operations:
- Opening a database connection
- PUT operations (storing key-value pairs)
- GET operations (retrieving values)
- SCAN operations (prefix-based queries)
- Proper database cleanup

```bash
cd examples/basic_kv
go run main.go
```

**What you'll learn**:
- Database initialization
- Basic CRUD operations
- Error handling patterns
- Resource management

---

### 2. SQL Operations

**Location**: `examples/sql_check/`

Shows how to use ToonDB's SQL interface:
- Creating tables with schema
- Inserting structured data
- Querying with SQL
- Data integrity verification

```bash
cd examples/sql_check
go run main.go
```

**What you'll learn**:
- SQL table creation
- Structured data storage
- Query execution
- Type handling in Go

---

### 3. RAG System (Production-Ready)

**Location**: `examples/rag/`

A complete, production-ready Retrieval-Augmented Generation system using:
- **ToonDB** as the vector store
- **Azure OpenAI** for embeddings and generation
- **Text chunking** for document processing
- **Semantic search** for context retrieval

#### Architecture

```
RAG Pipeline:
1. Document Ingestion → Chunking → Embedding → ToonDB Storage
2. Query → Embedding → Vector Search → Context Retrieval
3. Context + Query → LLM → Generated Answer
```

#### Components

- **Vector Store** (`internal/vectorstore/`): ToonDB-backed vector database with HNSW index
- **Embeddings** (`internal/embeddings/`): Azure OpenAI embeddings client
- **Chunking** (`internal/chunking/`): Text splitting and preprocessing
- **Generation** (`internal/generation/`): LLM integration for answer generation
- **RAG Pipeline** (`internal/rag/`): End-to-end orchestration

#### Setup

1. **Configure Environment**:
   ```bash
   cd examples/rag
   cp .env.example .env
   # Edit .env with your Azure OpenAI credentials
   ```

2. **Install Dependencies**:
   ```bash
   go mod download
   ```

3. **Run the Demo**:
   ```bash
   go run cmd/demo/main.go
   ```

#### Environment Variables

```env
AZURE_OPENAI_ENDPOINT=https://your-resource.openai.azure.com/
AZURE_OPENAI_API_KEY=your-api-key
AZURE_OPENAI_DEPLOYMENT=gpt-4
AZURE_OPENAI_EMBEDDING_DEPLOYMENT=text-embedding-ada-002
```

#### Features Demonstrated

- **Vector Search**: HNSW index for fast similarity search
- **Persistent Storage**: Embeddings and metadata stored in ToonDB
- **Chunking Strategies**: Configurable chunk size and overlap
- **Context Assembly**: Token-aware context building
- **Azure Integration**: Production-ready OpenAI client usage

---

## 🔑 Key Features of ToonDB Go SDK

- **Embedded Mode**: Run ToonDB directly in your Go application (FFI)
- **IPC Mode**: Connect to ToonDB server via Unix sockets
- **SQL Support**: Full SQL interface via IPC mode
- **Vector Search**: Built-in HNSW index for AI applications
- **ACID Transactions**: Group commits with Snapshot Isolation (SSI)
- **Path Operations**: Hierarchical key organization

## 📖 Documentation

- [ToonDB Go SDK Documentation](https://pkg.go.dev/github.com/toondb/toondb-go)
- [ToonDB Main Documentation](https://toondb.io)
- [API Reference](https://pkg.go.dev/github.com/toondb/toondb-go)

## 🤝 Contributing

We welcome contributions! Please feel free to submit Pull Requests with:
- New example implementations
- Improvements to existing examples
- Documentation enhancements
- Bug fixes

## 📄 License

Apache License 2.0 - see the [LICENSE](../LICENSE) file for details.

## 🔗 Related Repositories

- [toondb/toondb](https://github.com/toondb/toondb) - Main ToonDB repository
- [toondb/toondb-examples](https://github.com/toondb/toondb-examples) - Python, Node.js, and Rust examples
- [toondb/toondb-go](https://github.com/toondb/toondb-go) - Go SDK source code
