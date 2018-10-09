package main

type Row struct {
	Items []interface{}
}

func (r *Row) Cost() int {
	cost := 0
	for _, item := range r.Items {
		if suppressed(item) {
			cost++
		}
	}
	return cost
}

type Table struct {
	Rows []Row
}

func (t *Table) Cost() int {
	cost := 0
	for _, r := range t.Rows {
		cost += r.Cost()
	}
	return cost
}

func suppressed(item interface{}) bool {
	return item == nil
}
