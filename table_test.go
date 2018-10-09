package main

import (
	"testing"
)

func TestRowCost(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		items    []interface{}
	}{
		{name: "zero cost", expected: 0, items: []interface{}{1, 2, 3, 4}},
		{name: "1 cost", expected: 1, items: []interface{}{1, nil, 3, 4}},
		{name: "2 cost", expected: 2, items: []interface{}{1, nil, 3, nil}},
		{name: "3 cost", expected: 3, items: []interface{}{nil, nil, 3, nil}},
		{name: "4 cost", expected: 4, items: []interface{}{nil, nil, nil, nil}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := Row{Items: test.items}
			if r.Cost() != test.expected {
				t.Errorf("Expected %d but was %d", test.expected, r.Cost())
			}
		})
	}
}

func TestTableCost(t *testing.T) {
	table := Table{Rows: []Row{
		{Items: []interface{}{1, 2, 3, 4}},
		{Items: []interface{}{1, nil, 3, 4}},
		{Items: []interface{}{1, 2, nil, 4}},
		{Items: []interface{}{nil, 2, nil, 4}},
	}}
	expected := 4
	if expected != table.Cost() {
		t.Errorf("Expected %d but was %d", expected, table.Cost())
	}
}
