module example.com/context_query

go 1.25.5

replace github.com/toondb/toondb-go => ../../../toondb-go

require github.com/toondb/toondb-go v0.0.0-00010101000000-000000000000

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/posthog/posthog-go v1.8.2 // indirect
)
