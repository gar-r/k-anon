package algorithm

import (
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

// CreateNodesUndirected creates an undirected graph with nodeCount unconnected nodes.
func CreateNodesUndirected(nodeCount int) *simple.UndirectedGraph {
	g := simple.NewUndirectedGraph()
	batchAddNodes(nodeCount, g)
	return g
}

// CreateNodesWeightedUndirected creates a weighted undirected graph with nodeCount unconnected nodes.
func CreateNodesWeightedUndirected(nodeCount int) *simple.WeightedUndirectedGraph {
	g := simple.NewWeightedUndirectedGraph(0, math.MaxFloat64)
	batchAddNodes(nodeCount, g)
	return g
}

// CreateNodesDirected creates a directed graph with nodeCount unconnected nodes.
func CreateNodesDirected(nodeCount int) *simple.DirectedGraph {
	g := simple.NewDirectedGraph()
	batchAddNodes(nodeCount, g)
	return g
}

func batchAddNodes(nodeCount int, g graph.NodeAdder) {
	for i := 0; i < nodeCount; i++ {
		g.AddNode(g.NewNode())
	}
}

// AddEdge sets an edge between u and v in g.
func AddEdge(g interface {
	graph.EdgeAdder
	graph.Graph
}, u, v int64) {
	g.SetEdge(g.NewEdge(g.Node(u), g.Node(v)))
}

// AddWeightedEdge sets an edge with w weight between u and v in g.
func AddWeightedEdge(g interface {
	graph.WeightedEdgeAdder
	graph.Graph
}, u, v int64, w float64) {
	g.SetWeightedEdge(g.NewWeightedEdge(g.Node(u), g.Node(v), w))
}

// UndirectedConnectedComponents gets the connected components by treating the directed graph as undirected.
func UndirectedConnectedComponents(g graph.Directed) [][]graph.Node {
	var components [][]graph.Node
	if !isEmpty(g) {
		components = topo.ConnectedComponents(graph.Undirect{G: g})
	}
	return components
}

// UndirectGraph converts a directed graph to a simple undirected graph implementation.
func UndirectGraph(g graph.Directed) *simple.UndirectedGraph {
	undirected := simple.NewUndirectedGraph()
	graph.Copy(undirected, graph.Undirect{G: g})
	return undirected
}

func isEmpty(g graph.Directed) bool {
	return g.Nodes() == nil || g.Nodes().Len() < 1
}

func pickRandomVertex(component []graph.Node) graph.Node {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(component))
	return component[i]
}

func containsNode(component []graph.Node, u graph.Node) bool {
	for _, n := range component {
		if n == u {
			return true
		}
	}
	return false
}
