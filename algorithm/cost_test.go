package algorithm

import (
	"fmt"
	"k-anon/generalization"
	"k-anon/model"
	"k-anon/testutil"
	"testing"
)

func TestCalculateCost(t *testing.T) {
	tests := []struct {
		items1, items2 []int
		expectedCost   float64
	}{
		{items1: []int{8}, items2: []int{8}, expectedCost: 0},
		{items1: []int{1}, items2: []int{1}, expectedCost: 0},
		{items1: []int{1}, items2: []int{2}, expectedCost: 0.5},
		{items1: []int{8}, items2: []int{9}, expectedCost: 0.25},
		{items1: []int{1}, items2: []int{4}, expectedCost: 0.75},
		{items1: []int{1}, items2: []int{5}, expectedCost: 1},
		{items1: []int{1, 1}, items2: []int{3, 6}, expectedCost: 1.75},
		{items1: []int{1, 1}, items2: []int{1, 1}, expectedCost: 0},
		{items1: []int{1, 1}, items2: []int{1, 1}, expectedCost: 0},
		{items1: []int{1, 1, 1, 1}, items2: []int{4, 5, 1, 2}, expectedCost: 2.25},
	}
	g := generalization.GetIntGeneralizer1()
	for i, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", i), func(t *testing.T) {
			v1 := model.CreateVector(test.items1, g)
			v2 := model.CreateVector(test.items2, g)
			actualCost := CalculateCost(v1, v2)
			testutil.AssertEquals(test.expectedCost, actualCost, t)
		})
	}
}

func TestCalculateCost_WithNonIdentifierAttributes(t *testing.T) {
	gen := generalization.GetIntGeneralizer1()
	v1 := &model.Vector{
		Items: []*model.Data{
			model.NewIdentifier(5, gen),
			model.NewIdentifier(1, gen),
			model.NewNonIdentifier("Test1"),
		},
	}
	v2 := &model.Vector{
		Items: []*model.Data{
			model.NewIdentifier(6, gen),
			model.NewIdentifier(9, gen),
			model.NewNonIdentifier("Test2"),
		},
	}
	cost := CalculateCost(v1, v2)
	testutil.AssertEquals(1.5, cost, t)
}
