package algorithm

import (
	"gonum.org/v1/gonum/graph/simple"
	"testing"
)

// k = 1
// 0	1
func TestIsComplete_1(t *testing.T) {
	g := simple.NewDirectedGraph()
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	if !isComplete(g, 1) {
		t.Error()
	}
}

// k = 2
// 0	1
func TestIsComplete_2(t *testing.T) {
	g := simple.NewDirectedGraph()
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	if isComplete(g, 2) {
		t.Error()
	}
}

// k = 2
// 0 --- 1
func TestIsComplete_3(t *testing.T) {
	g := simple.NewDirectedGraph()
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.SetEdge(g.NewEdge(g.Node(0), g.Node(1)))
	if !isComplete(g, 2) {
		t.Error()
	}
}

// k = 3
// 0 --- 1 --- 2
// 3 --- 4
func TestIsComplete_4(t *testing.T) {
	g := simple.NewDirectedGraph()
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.SetEdge(g.NewEdge(g.Node(0), g.Node(1)))
	g.SetEdge(g.NewEdge(g.Node(1), g.Node(2)))
	g.SetEdge(g.NewEdge(g.Node(3), g.Node(4)))
	if isComplete(g, 3) {
		t.Error()
	}
}

// k = 3
// 0 --- 1 --- 2
// 3 --- 4 --- 5
func TestIsComplete_5(t *testing.T) {
	g := simple.NewDirectedGraph()
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.SetEdge(g.NewEdge(g.Node(0), g.Node(1)))
	g.SetEdge(g.NewEdge(g.Node(1), g.Node(2)))
	g.SetEdge(g.NewEdge(g.Node(3), g.Node(4)))
	g.SetEdge(g.NewEdge(g.Node(4), g.Node(5)))
	if !isComplete(g, 3) {
		t.Error()
	}
}
