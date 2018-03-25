// Package gcounter is a CRDT of the convergent type.
// This is a POC demonstration.
package gcounter

// TODO(rm): write a properly usable, non-demo, implementation.

import (
	"fmt"
	"strconv"
	"strings"
)

// gcounter is the CRDT of a node.
type gcounter []int

// cluster represents a collection of nodes
// each having a synchronised gcounter CRDT.
type cluster struct {
	nodes []*node
}

// NewCluster creates a new cluster of nodes.
func NewCluster() *cluster {
	return &cluster{}
}

// Add adds a node to the cluster.
func (c *cluster) Add(n *node) {
	c.nodes = append(c.nodes, n)
}

// Node returns a node specified by id from the cluster.
func (c *cluster) Node(i int) *node {
	return c.nodes[i]
}

// Sync syncs the nodes in the cluster with each other.
// IRL the nodes would do this autonomously.
func (c *cluster) Sync() {
	for nid, _ := range c.nodes {
		p := NewPayload(c.nodes[nid])
		for i, node := range c.nodes {
			if i == nid {
				continue
			}

			node.merge(p)
		}
	}
}

// PrintClusterState displays the state of the cluster on stdout.
func (c *cluster) PrintClusterState() {
	fmt.Println("\nCluster state:")
	for i := range c.nodes {
		fmt.Printf("Node %d: %+v\n", c.nodes[i].id, c.nodes[i].gcounter)
	}
	fmt.Println("")
}

// payload is really just a node, has the same data -
// we only have a specific struct to elucidate the concept.
type payload struct {
	fromNodeId int
	gcounter   gcounter
}

// NewPayload creates a payload from a given node.
func NewPayload(n *node) *payload {
	return &payload{
		fromNodeId: n.id,
		gcounter:   n.gcounter,
	}
}

// node is a node in the cluster.
type node struct {
	id       int
	gcounter gcounter
}

// NewNode creates a new node with the provided id.
func NewNode(i int) *node {
	n := &node{
		id:       i,
		gcounter: make(gcounter, 3),
	}
	return n
}

// SetId sets the id of the node.
func (n *node) SetId(i int) {
	n.id = i
}

// Inc increments the counter in the array slot associated with the node.
func (n *node) Inc() {
	n.gcounter[n.id]++
}

// merge compares the payload with the node's record of state
// and determines the update.
func (n *node) merge(p *payload) {
	fmt.Printf("sync \t\t\tnode %d:\t%v\n", n.id, n.gcounter)
	fmt.Printf("with payload from \tnode %d: %d\n", p.fromNodeId, p.gcounter)

	for i, _ := range p.gcounter {
		if p.gcounter[i] > n.gcounter[i] {
			n.gcounter[i] = p.gcounter[i]
		}
	}

	fmt.Printf("sync result\t\tnode %d: %v\n\n", n.id, n.gcounter)
}

// payloadString converts a gcounter array to a CSV string.
func payloadString(v []int) string {
	ss := make([]string, len(v))
	for i, j := range v {
		ss[i] = strconv.Itoa(j)
	}
	s := strings.Join(ss, ",")
	return s
}
