package algorithm

import (
	"gonum.org/v1/gonum/graph"
	"k-anon/model"
)

// Anonimizer operates on a given table with parameter 'k'.
// In a k-anonimized table valus are generalized or suppressed in a way,
// that given any record there are other k-1 records in the table that are identical
// to it along quasi-identifier attributes
type Anonimizer struct {
	table *model.Table
	k     int
}

func (a *Anonimizer) AnonimizeData() {
	a.computeAnonGraph()
}

func (a *Anonimizer) computeAnonGraph() graph.Undirected {
	g := BuildAnonGraph(a.table, a.k)
	d := NewDecomposer(UndirectGraph(g), a.k)
	d.Decompose()
	return d.g
}

func (a *Anonimizer) generalize(vectors []*model.Vector) {
	dims := len(vectors[0].Items)
	for i := 0; i < dims; i++ {
		//var dims []*model.Data
		for j := 0; j < len(vectors); j++ {
			//append(dims, vectors[j].Items[i])
		}
	}
}
