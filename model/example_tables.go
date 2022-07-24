package model

import (
	"git.okki.hu/garric/k-anon/generalization"
)

func GetIntTable1() *Table {
	g := generalization.ExampleIntGeneralizer()
	t := NewTable(&Schema{
		Columns: []*Column{
			NewColumn("Col1", g),
			NewColumn("Col2", g),
			NewColumn("Col3", g),
			NewColumn("Col4", g),
		},
	})
	t.AddRow(1, 1, 1, 1)
	t.AddRow(1, 1, 1, 2)
	t.AddRow(4, 5, 1, 1)
	t.AddRow(1, 3, 5, 7)
	return t
}

func GetStudentTable() *Table {
	dim2 := generalization.NewIntRangeGeneralizer(25, 35)
	dim3 := generalization.NewIntRangeGeneralizer(0, 3)
	dim4 := generalization.NewFloatRangeGeneralizer(0.0, 5.0)
	dim5 := generalization.ExampleGradeGeneralizer()
	t := NewTable(&Schema{
		Columns: []*Column{
			NewColumn("Gender", &generalization.Suppressor{}),
			NewColumn("Col2", dim2),
			NewColumn("Col3", dim3),
			NewColumn("Col4", dim4),
			NewColumn("Col5", dim5),
		},
	})
	t.AddRow("Male", 25, 0, 3.487, "A")
	t.AddRow("Female", 25, 0, 4.234, "A+")
	t.AddRow("Male", 30, 2, 2.89, "B")
	t.AddRow("Female", 30, 1, 3.34, "B+")
	t.AddRow("Male", 28, 1, 3.98, "A-")
	t.AddRow("Female", 28, 1, 2.534, "B")
	t.AddRow("Male", 27, 0, 2.06, "B-")
	t.AddRow("Female", 27, 2, 2.45, "B")
	return t
}

func GetMixedTable1() *Table {
	t := NewTable(&Schema{
		Columns: []*Column{
			NewColumn("Score", generalization.ExampleIntGeneralizer()),
			NewColumn("Grade", generalization.ExampleGradeGeneralizer()),
		},
	})
	t.AddRow(9, "A+")
	t.AddRow(8, "A")
	t.AddRow(5, "B-")
	return t
}

func GetMixedTable2() *Table {
	t := NewTable(&Schema{
		Columns: []*Column{
			NewColumn("Score", generalization.ExampleIntGeneralizer()),
			NewColumn("Grade", generalization.ExampleGradeGeneralizer()),
		},
	})
	t.AddRow(9, "A+")
	t.AddRow(2, "B-")
	t.AddRow(6, "A-")
	t.AddRow(4, "B+")
	return t
}

func GetMixedTable3() *Table {
	t := NewTable(&Schema{
		Columns: []*Column{
			NewColumn("Score", generalization.ExampleIntGeneralizer()),
			NewColumn("Grade", generalization.ExampleGradeGeneralizer()),
			NewColumn("Motto", &generalization.PrefixGeneralizer{MaxWords: 5}),
			NewColumn("Remark", nil),
		},
	})
	t.AddRow(9, "A+", "cats are wild", "data1")
	t.AddRow(8, "A", "cats are evil", "data2")
	t.AddRow(6, "A-", "cats are fluffy", "data3")
	return t
}

// GetEmptyTable return an empty table
func GetEmptyTable() *Table {
	table := &Table{}
	return table
}
