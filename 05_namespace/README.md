# Namespaces

This example demonstrates namespace management for multi-tenant data isolation in SochDB.

## Features Demonstrated

- Creating namespaces
- Using namespaces for data storage
- Data isolation between namespaces
- Listing all namespaces
- Namespace statistics
- Scanning data within namespaces
- Deleting namespaces
- Multi-tenant query patterns

## Running the Example

```bash
go run main.go
```

## Expected Output

The example will:
1. Create multiple namespaces (tenants)
2. Store isolated data in each namespace
3. Verify data isolation
4. List all namespaces
5. Show statistics per namespace
6. Scan namespace data
7. Delete a namespace
8. Demonstrate multi-tenant queries

## Key Concepts

- **Isolation**: Each namespace has completely separate data
- **Multi-tenancy**: Store data for multiple customers/tenants
- **Organization**: Logical grouping of related data
- **Security**: Prevent cross-tenant data access
- **Efficiency**: Shared infrastructure with logical separation

## Use Cases

- **SaaS Applications**: Separate data per customer
- **Multi-tenant Databases**: Isolate tenant data
- **Environment Separation**: dev/staging/prod namespaces
- **User Workspaces**: Personal data spaces
- **Project Management**: Separate projects/teams
- **Testing**: Isolated test environments

## Architecture Patterns

### Pattern 1: Tenant-per-Namespace
```go
// Each customer gets their own namespace
db.CreateNamespace("tenant_" + customerID)
ns := db.UseNamespace("tenant_" + customerID)
ns.Put(key, value)
```

### Pattern 2: Environment-per-Namespace
```go
// Separate development, staging, production
db.CreateNamespace("env_dev")
db.CreateNamespace("env_staging")
db.CreateNamespace("env_prod")
```

### Pattern 3: Feature-per-Namespace
```go
// Organize by application feature
db.CreateNamespace("users")
db.CreateNamespace("orders")
db.CreateNamespace("inventory")
```

## Best Practices

- Use consistent naming conventions (e.g., `tenant_`, `env_`)
- Monitor namespace statistics for capacity planning
- Implement tenant-aware query routing
- Clean up unused namespaces regularly
- Document namespace purposes
- Use namespaces for logical isolation, not security boundaries alone
