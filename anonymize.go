package k_anon

import (
	"bitbucket.org/dargzero/k-anon/algorithm"
	"bitbucket.org/dargzero/k-anon/model"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/topo"
)

// Anonymizer is a graph based data anonymizer which operates on table data.
// The anonymizer is characterized by its k value. In a k-anonymized table
// all records are samePartition or suppressed in a way, that given any record
// there are other k-1 records in the table that are identical
// to it along quasi-identifier attributes.
type Anonymizer struct {
	k     int
	table *model.Table
}

// Anonymize creates a k-anonymized Table from the input Table.
func (a *Anonymizer) Anonymize() {
	g := a.computeAnonGraph()
	components := topo.ConnectedComponents(g)
	groups := a.getRowGroups(components)
	a.generalize(groups)
}

func (a *Anonymizer) computeAnonGraph() graph.Undirected {
	g := algorithm.BuildAnonGraph(a.table, a.k)
	undirected := algorithm.UndirectGraph(g)
	d := algorithm.NewDecomposer(undirected, a.k)
	d.Decompose()
	return undirected
}

func (a *Anonymizer) getRowGroups(components [][]graph.Node) [][]*model.Row {
	var groups [][]*model.Row
	for _, component := range components {
		var group []*model.Row
		for _, n := range component {
			id := int(n.ID())
			if id < len(a.table.GetRows()) { // skip Steiner's vertices
				group = append(group, a.table.GetRows()[id])
			}
		}
		groups = append(groups, group)
	}
	return groups
}

func (a *Anonymizer) generalize(groups [][]*model.Row) {
	for _, group := range groups {
		a.generalizeRowGroup(group)
	}
}

func (a *Anonymizer) generalizeRowGroup(rows []*model.Row) {
	for colIdx := 0; colIdx < len(a.table.GetSchema().Columns); colIdx++ {
		colDef := a.table.GetSchema().Columns[colIdx]
		if colDef.IsIdentifier() {
			for level := 0; level < colDef.GetGeneralizer().Levels(); level++ {
				for _, row := range rows {
					p := colDef.GetGeneralizer().Generalize(row.Data[colIdx], level)
					row.Data[colIdx] = p
				}
				if samePartition(colIdx, rows) {
					break
				}
			}
		}
	}
}

func samePartition(colIdx int, rows []*model.Row) bool {
	if len(rows) > 1 {
		first := rows[0]
		for _, row := range rows {
			if !first.Data[colIdx].Equals(row.Data[colIdx]) {
				return false
			}
		}
	}
	return true
}
