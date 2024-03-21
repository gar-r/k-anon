package kanon

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gar-r/k-anon/generalization"
	"github.com/gar-r/k-anon/hierarchy"

	"github.com/gar-r/k-anon/model"
)

const rangeMin = 0
const rangeMax = 100

func BenchmarkAnonymizerColumns(b *testing.B) {
	gen := generalization.NewIntRangeGeneralizer(rangeMin, rangeMax)
	for col := 1; col <= 50; col++ {
		b.Run(fmt.Sprintf("columns/%d", col), func(b *testing.B) {
			table := randomTable(col, 100, gen)
			anonymizeTableBench(table, 4, b)
		})
	}
}

func BenchmarkAnonymizerRows(b *testing.B) {
	gen := generalization.NewIntRangeGeneralizer(rangeMin, rangeMax)
	for rows := 10; rows <= 100; rows += 10 {
		b.Run(fmt.Sprintf("rows/%d", rows), func(b *testing.B) {
			table := randomTable(10, rows, gen)
			anonymizeTableBench(table, 4, b)
		})
	}
}

func BenchmarkAnonymizerK(b *testing.B) {
	gen := generalization.NewIntRangeGeneralizer(rangeMin, rangeMax)
	for k := 2; k <= 100; k += 5 {
		b.Run(fmt.Sprintf("k/%d", k), func(b *testing.B) {
			table := randomTable(5, 200, gen)
			anonymizeTableBench(table, k, b)
		})
	}
}

func BenchmarkAnonymizerColumnTypes(b *testing.B) {
	items := make([]interface{}, rangeMax)
	for i := range items {
		items[i] = i
	}
	h, _ := hierarchy.AutoBuild(10, items...)
	tests := []struct {
		name        string
		generalizer generalization.Generalizer
	}{
		{"int", generalization.NewIntRangeGeneralizer(rangeMin, rangeMax)},
		{"float", generalization.NewFloatRangeGeneralizer(rangeMin, rangeMax)},
		{"suppress", &generalization.Suppressor{}},
		{"prefix", &generalization.PrefixGeneralizer{MaxWords: rangeMax}},
		{"hierarchy", &generalization.HierarchyGeneralizer{Hierarchy: h}},
	}
	for _, test := range tests {
		for i := 5; i < 100; i += 5 {
			b.Run(fmt.Sprintf("%v/%d", test.name, i), func(b *testing.B) {
				table := randomTable(1, i, test.generalizer)
				anonymizeTableBench(table, 4, b)
			})
		}
	}
}

func anonymizeTableBench(table *model.Table, k int, b *testing.B) {
	b.Helper()
	for i := 0; i < b.N; i++ {
		anon := &Anonymizer{
			Table: table,
			K:     k,
		}
		err := anon.Anonymize()
		if err != nil {
			b.Error("error while anonymizing table", err)
		}
	}
}

func randomTable(nCols, nRows int, generalizer generalization.Generalizer) *model.Table {
	cols := makeCols(nCols, generalizer)
	table := model.NewTable(&model.Schema{Columns: cols})
	addRandomRows(nRows, nCols, table)
	return table
}

func addRandomRows(nRows int, nCols int, table *model.Table) {
	for i := 0; i < nRows; i++ {
		row := make([]interface{}, nCols)
		for j := range row {
			row[j] = rand.Intn(rangeMax-rangeMin) + rangeMin
		}
		table.AddRow(row...)
	}
}

func makeCols(nCols int, generalizer generalization.Generalizer) []*model.Column {
	cols := make([]*model.Column, nCols)
	for i := range cols {
		cols[i] = model.NewColumn(fmt.Sprintf("col %d", i), generalizer)
	}
	return cols
}
