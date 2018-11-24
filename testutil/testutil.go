package testutil

import (
	"gonum.org/v1/gonum/graph"
	"testing"
)

func AssertContains(t *testing.T, component []graph.Node, ids ...int64) {
	for _, id := range ids {
		for _, n := range component {
			if n.ID() == id {
				return
			}
		}
		t.Errorf("component %v does not contain node %v", component, id)
	}
}
