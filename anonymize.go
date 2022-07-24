package kanon

import (
	"git.okki.hu/garric/k-anon/algorithm"
	"git.okki.hu/garric/k-anon/model"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/topo"
)

// Anonymizer is a graph based data anonymizer which operates on Table data.
// The anonymizer is characterized by its K value. In a K-anonymized Table
// all records are samePartition or suppressed in a way, that given any record
// there are other K-1 records in the Table that are identical
// to it along quasi-identifier attributes.
type Anonymizer struct {
	K     int
	Table *model.Table
}

// Anonymize creates a K-anonymized Table from the input Table.
func (a *Anonymizer) Anonymize() error {
	g, err := a.computeAnonGraph()
	if err != nil {
		return err
	}
	components := topo.ConnectedComponents(g)
	groups := a.getRowGroups(components)
	a.generalize(groups)
	return nil
}

func (a *Anonymizer) computeAnonGraph() (graph.Undirected, error) {
	g, err := algorithm.BuildAnonGraph(a.Table, a.K)
	if err != nil {
		return nil, err
	}
	undirected := algorithm.UndirectGraph(g)
	d := algorithm.NewDecomposer(undirected, a.K)
	d.Decompose()
	return undirected, nil
}

func (a *Anonymizer) getRowGroups(components [][]graph.Node) [][]*model.Row {
	var groups [][]*model.Row
	for _, component := range components {
		var group []*model.Row
		for _, n := range component {
			id := int(n.ID())
			if id < len(a.Table.GetRows()) { // skip Steiner's vertices
				group = append(group, a.Table.GetRows()[id])
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
	for colIdx := 0; colIdx < len(a.Table.GetSchema().Columns); colIdx++ {
		colDef := a.Table.GetSchema().Columns[colIdx]
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
