package main

func Suppress(t *Table, col int) {
	for _, r := range t.Rows {
		for i := range r.Items {
			if i == col {
				r.Items[i] = nil
			}
		}
	}
}
