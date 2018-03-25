// Package main demonstrates the use of the CRDT gcounter package.
package main

import (
	crdt "github.com/russmack/crdt-go/gcounter"
)

func main() {
	// Create a cluster.
	cluster := crdt.NewCluster()

	// Add nodes to the cluster, each with an id.
	for i := 0; i < 3; i++ {
		cluster.Add(crdt.NewNode(i))
	}

	// Increment the counters of the nodes.
	cluster.Node(0).Inc()
	cluster.Node(0).Inc()
	cluster.Node(0).Inc()

	cluster.Node(1).Inc()
	cluster.Node(1).Inc()

	cluster.Node(2).Inc()
	cluster.Node(2).Inc()
	cluster.Node(2).Inc()
	cluster.Node(2).Inc()

	cluster.PrintClusterState()

	// Sync cluster nodes.
	cluster.Sync()

	cluster.PrintClusterState()
}
