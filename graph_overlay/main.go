package main

import (
	"fmt"
	"log"

	toondb "github.com/toondb/toondb-go"
)

func main() {
	fmt.Println("=== ToonDB Go SDK - Graph Overlay Example ===\n")

	// Open database
	db, err := toondb.Open("./test_graph_db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()
	fmt.Println("✓ Database opened")

	// Create graph overlay
	graph := toondb.NewGraphOverlay(db, "demo")
	fmt.Println("✓ Graph overlay created")

	// Add nodes
	alice := map[string]interface{}{
		"name": "Alice",
		"role": "developer",
	}
	_, err = graph.AddNode("alice", "person", alice)
	if err != nil {
		log.Fatalf("Failed to add node: %v", err)
	}

	bob := map[string]interface{}{
		"name": "Bob",
		"role": "engineer",
	}
	_, err = graph.AddNode("bob", "person", bob)
	if err != nil {
		log.Fatalf("Failed to add node: %v", err)
	}

	toondbNode := map[string]interface{}{
		"name":        "ToonDB",
		"description": "AI Database",
	}
	_, err = graph.AddNode("toondb", "project", toondbNode)
	if err != nil {
		log.Fatalf("Failed to add node: %v", err)
	}
	fmt.Println("✓ Added 3 nodes (alice, bob, toondb)")

	// Add edges
	_, err = graph.AddEdge("alice", "knows", "bob", nil)
	if err != nil {
		log.Fatalf("Failed to add edge: %v", err)
	}

	_, err = graph.AddEdge("bob", "works_on", "toondb", nil)
	if err != nil {
		log.Fatalf("Failed to add edge: %v", err)
	}

	_, err = graph.AddEdge("alice", "contributes_to", "toondb", nil)
	if err != nil {
		log.Fatalf("Failed to add edge: %v", err)
	}
	fmt.Println("✓ Added 3 edges")

	// Get node
	node, err := graph.GetNode("alice")
	if err != nil {
		log.Fatalf("Failed to get node: %v", err)
	}
	fmt.Printf("✓ Retrieved node: %s (%s)\n", node.Properties["name"], node.Type)

	// Get edges
	edges, err := graph.GetEdges("alice", "")
	if err != nil {
		log.Fatalf("Failed to get edges: %v", err)
	}
	fmt.Printf("✓ Alice has %d edges\n", len(edges))

	// BFS traversal
	visited, err := graph.BFS("alice", 10, nil, nil)
	if err != nil {
		log.Fatalf("Failed BFS traversal: %v", err)
	}
	fmt.Printf("✓ BFS visited %d nodes: %v\n", len(visited), visited)

	fmt.Println("\n✓✓✓ SUCCESS: Graph Overlay works perfectly! ✓✓✓")
}
