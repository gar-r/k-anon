package testutil

import (
	"reflect"
	"testing"

	"gonum.org/v1/gonum/graph"
)

// AssertContains ensures, that a component contains the nodes with the given ids.
func AssertContains(t *testing.T, component []graph.Node, ids ...int64) {
	t.Helper()
	for _, id := range ids {
		for _, n := range component {
			if n.ID() == id {
				return
			}
		}
		t.Errorf("component %v does not contain node %v", component, id)
	}
}

// AssertEdgeCost ensures, that there is an edge between two given nodes with the given cost in the given weighted graph.
func AssertEdgeCost(t *testing.T, graph graph.Weighted, node1, node2 int, expectedCost float64) {
	t.Helper()
	cost, exists := graph.Weight(int64(node1), int64(node2))
	if !exists {
		t.Errorf("expected edge between %d and %d, but was not found", node1, node2)
	}
	if expectedCost != cost {
		t.Errorf("expected cost %v, got %v", expectedCost, cost)
	}
}

// AssertVertexReplaced ensures, that the 'original' vertex is replaced with 'new', and 'new' is
// connected to each node in 'connections' given the g graph.
func AssertVertexReplaced(t *testing.T, g graph.Graph, original, new int64, connections ...int64) {
	t.Helper()
	for _, conn := range connections {
		if g.HasEdgeBetween(original, conn) {
			t.Errorf("unexpected edge between %v and %v", original, conn)
		}
		if !g.HasEdgeBetween(new, conn) {
			t.Errorf("expected edge between %v and %v", new, conn)
		}
	}
}

// AssertNil ensures, that the given value is nil.
func AssertNil(value interface{}, t *testing.T) {
	t.Helper()
	if !isNil(value) {
		t.Errorf("expected nil, got %v", value)
	}
}

// AssertNotNil ensures, that the given value is not nil.
func AssertNotNil(value interface{}, t *testing.T) {
	t.Helper()
	if isNil(value) {
		t.Errorf("unexpected nil value")
	}
}

func isNil(object interface{}) bool {
	if object == nil {
		return true
	}
	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}
	return false
}

// AssertEquals ensures, that the two given values are equal.
func AssertEquals(expected interface{}, actual interface{}, t *testing.T) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
