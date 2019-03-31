package algorithm

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/model"
	"bitbucket.org/dargzero/k-anon/testutil"
	"fmt"
	"testing"
)

func TestCalculateCost(t *testing.T) {

	t.Run("calculate cost for row pair", func(t *testing.T) {
		tests := []struct {
			items1, items2 []interface{}
			expectedCost   float64
		}{
			{[]interface{}{8}, []interface{}{8}, 0},
			{[]interface{}{1}, []interface{}{1}, 0},
			{[]interface{}{1}, []interface{}{2}, 0.5},
			{[]interface{}{8}, []interface{}{9}, 0.25},
			{[]interface{}{1}, []interface{}{4}, 0.75},
			{[]interface{}{1}, []interface{}{5}, 1},
			{[]interface{}{1, 1}, []interface{}{3, 6}, 1.75},
			{[]interface{}{1, 1}, []interface{}{1, 1}, 0},
			{[]interface{}{1, 1, 1, 1}, []interface{}{4, 5, 1, 2}, 2.25},
		}
		for i, test := range tests {
			t.Run(fmt.Sprintf("test #%d", i), func(t *testing.T) {
				schema := getSchema(len(test.items1))
				table := model.NewTable(schema)
				table.AddRow(test.items1...)
				table.AddRow(test.items2...)
				r1 := table.GetRows()[0]
				r2 := table.GetRows()[1]
				actualCost := CalculateCost(r1, r2, schema)
				testutil.AssertEquals(test.expectedCost, actualCost, t)
			})
		}
	})

	t.Run("calculate with non identifier attributes", func(t *testing.T) {
		gen := generalization.ExampleIntGeneralizer()
		schema := &model.Schema{
			Columns: []*model.Column{
				{"Col1", gen},
				{"Col2", gen},
				{"Col3", nil},
			},
		}
		table := model.NewTable(schema)
		table.AddRow(5, 1, "Test1")
		table.AddRow(6, 9, "Test2")
		r1 := table.GetRows()[0]
		r2 := table.GetRows()[1]
		cost := CalculateCost(r1, r2, schema)
		testutil.AssertEquals(1.5, cost, t)
	})

	t.Run("calculate with prefix attributes", func(t *testing.T) {
		gen := &generalization.PrefixGeneralizer{MaxWords: 5}
		schema := &model.Schema{
			Columns: []*model.Column{
				{"Col1", gen},
			},
		}
		table := model.NewTable(schema)
		table.AddRow("cats are wonderful little beings")
		table.AddRow("dogs are my pets")
		r1 := table.GetRows()[0]
		r2 := table.GetRows()[1]
		cost := CalculateCost(r1, r2, schema)
		testutil.AssertEquals(1.0, cost, t)
	})

	t.Run("cannot generalize into same partition", func(t *testing.T) {

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic, got none")
			}
		}()

		schema := getSchema(1)
		table := model.NewTable(schema)
		table.AddRow(5)
		table.AddRow(100)
		r1 := table.GetRows()[0]
		r2 := table.GetRows()[1]
		CalculateCost(r1, r2, schema)
	})

}

func getSchema(cols int) *model.Schema {
	g := generalization.ExampleIntGeneralizer()
	schema := &model.Schema{}
	for i := 1; i <= cols; i++ {
		schema.Columns = append(schema.Columns, &model.Column{
			Name:        fmt.Sprintf("Col%d", i),
			Generalizer: g,
		})
	}
	return schema
}
