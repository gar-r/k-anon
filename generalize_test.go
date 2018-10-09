package main

import "testing"

func TestSuppress(t *testing.T) {
	table := &Table{Rows: []Row{
		{Items: []interface{}{1, 2, 3, 4}},
		{Items: []interface{}{5, 6, 7, 8}},
		{Items: []interface{}{9, 1, 2, 3}},
		{Items: []interface{}{4, 5, 6, 7}},
	}}
	col := 3
	Suppress(table, col)
	for i, row := range table.Rows {
		for j, item := range row.Items {
			if j == col && item != nil {
				t.Errorf("Item [%d,%d] should be nil, but was %d", i, j, item)
			}
		}
	}
}
