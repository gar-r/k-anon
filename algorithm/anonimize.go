package algorithm

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/topo"
	"k-anon/generalization"
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
	g := a.computeAnonGraph()
	components := topo.ConnectedComponents(g)
	a.getGroups(components)
}

func (a *Anonimizer) computeAnonGraph() graph.Undirected {
	g := BuildAnonGraph(a.table, a.k)
	d := NewDecomposer(UndirectGraph(g), a.k)
	d.Decompose()
	return d.g
}

func (a *Anonimizer) getGroups(components [][]graph.Node) [][]*model.Vector {
	var groups [][]*model.Vector
	for _, component := range components {
		var rows []*model.Vector
		for _, n := range component {
			idx := int(n.ID())
			if idx < len(a.table.Rows) {
				rows = append(rows, a.table.Rows[idx])
			}
		}
		groups = append(groups, rows)
	}
	return groups
}

func anonimize(group []*model.Vector) {
	for dim := range group[0].Items {
		var data []*model.Data
		for _, v := range group {
			data = append(data, v.Items[dim])
		}
		generalize(data)
	}
}

func generalize(data []*model.Data) []*generalization.Partition {
	if data[0].IsIdentifier() {
		return generalizeIdentifier(data)
	}
	return generalizeNonIdentifier(data)
}

func generalizeIdentifier(data []*model.Data) []*generalization.Partition {
level:
	for level := 0; level < data[0].Levels(); level++ {
		var result []*generalization.Partition
		for _, d := range data {
			p := d.Generalize(level)
			if len(result) > 0 && !result[0].Equals(p) {
				continue level
			}
			result = append(result, p)
		}
		return result
	}
	panic("could not generalize items into same partition")
}

func generalizeNonIdentifier(data []*model.Data) []*generalization.Partition {
	var result []*generalization.Partition
	for _, d := range data {
		p := generalization.NewPartition(d.Value)
		result = append(result, p)
	}
	return result
}
