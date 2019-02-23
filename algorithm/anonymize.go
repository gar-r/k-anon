package algorithm

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/model"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/topo"
)

// Anonymizer operates on a given table with parameter 'k'.
// In a k-anonymized table values are generalized or suppressed in a way,
// that given any record there are other k-1 records in the table that are identical
// to it along quasi-identifier attributes
type Anonymizer struct {
	table *model.Table
	k     int
}

func (a *Anonymizer) anonymizeData() [][]*generalization.Partition {
	g := a.computeAnonGraph()
	components := topo.ConnectedComponents(g)
	groups := a.getGroups(components)
	var results [][]*generalization.Partition
	for _, group := range groups {
		rows := anonymize(group)
		for _, r := range rows {
			results = append(results, r)
		}
	}
	return results
}

func (a *Anonymizer) computeAnonGraph() graph.Undirected {
	g := BuildAnonGraph(a.table, a.k)
	d := NewDecomposer(UndirectGraph(g), a.k)
	d.Decompose()
	return d.g
}

func (a *Anonymizer) getGroups(components [][]graph.Node) [][]*model.Vector {
	var groups [][]*model.Vector
	for _, component := range components {
		var rows []*model.Vector
		for _, n := range component {
			idx := int(n.ID())
			if idx < len(a.table.Rows) { // skip Steiner's vertices
				rows = append(rows, a.table.Rows[idx])
			}
		}
		groups = append(groups, rows)
	}
	return groups
}

func anonymize(group []*model.Vector) [][]*generalization.Partition {
	results := makeRows(len(group), len(group[0].Items))
	for col := range group[0].Items {
		var data []*model.Data
		for _, v := range group {
			data = append(data, v.Items[col])
		}
		partitions := generalize(data)
		for row, p := range partitions {
			results[row][col] = p
		}
	}
	return results
}

func makeRows(rows, cols int) [][]*generalization.Partition {
	anonymizedRows := make([][]*generalization.Partition, rows)
	for i := range anonymizedRows {
		anonymizedRows[i] = make([]*generalization.Partition, cols)
	}
	return anonymizedRows
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
