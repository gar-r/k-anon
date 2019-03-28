package algorithm

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/model"
	"bitbucket.org/dargzero/k-anon/testutil"
	"fmt"
	"testing"
)

func TestCalculateCost(t *testing.T) {

	t.Run("calculate cost for single column", func(t *testing.T) {
		tests := []struct {
			r1, r2       *model.Row
			expectedCost float64
		}{
			{model.NewRow(8), model.NewRow(8), 0},
			{model.NewRow(1), model.NewRow(1), 0},
			{model.NewRow(1), model.NewRow(2), 0.5},
			{model.NewRow(8), model.NewRow(9), 0.25},
			{model.NewRow(1), model.NewRow(4), 0.75},
			{model.NewRow(1), model.NewRow(5), 1},
			{model.NewRow(1, 1), model.NewRow(3, 6), 1.75},
			{model.NewRow(1, 1), model.NewRow(1, 1), 0},
			{model.NewRow(1, 1, 1, 1), model.NewRow(4, 5, 1, 2), 2.25},
		}
		for i, test := range tests {
			t.Run(fmt.Sprintf("test #%d", i), func(t *testing.T) {
				schema := getSchema(len(test.r1.Data))
				actualCost := CalculateCost(test.r1, test.r2, schema)
				testutil.AssertEquals(test.expectedCost, actualCost, t)
			})
		}
	})

	t.Run("calculate with non identifier attributes", func(t *testing.T) {
		gen := generalization.GetIntGeneralizer()
		schema := &model.Schema{
			Columns: []*model.Column{
				{"Col1", gen},
				{"Col2", gen},
				{"Col3", nil},
			},
		}
		r1 := model.NewRow(5, 1, "Test1")
		r2 := model.NewRow(6, 9, "Test2")
		cost := CalculateCost(r1, r2, schema)
		testutil.AssertEquals(1.5, cost, t)
	})

}

func getSchema(cols int) *model.Schema {
	g := generalization.GetIntGeneralizer()
	schema := &model.Schema{}
	for i := 1; i <= cols; i++ {
		schema.Columns = append(schema.Columns, &model.Column{
			Name:        fmt.Sprintf("Col%d", i),
			Generalizer: g,
		})
	}
	return schema
}
