package model

import (
	"bitbucket.org/dargzero/k-anon/generalization"
)

func GetIntTable1() *Table {
	g := generalization.ExampleIntGeneralizer()
	t := NewTable(&Schema{
		Columns: []*Column{
			{"Col1", g},
			{"Col2", g},
			{"Col3", g},
			{"Col4", g},
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
	dim4 := generalization.NewIntRangeGeneralizer(10, 50)
	dim5 := generalization.ExampleGradeGeneralizer()
	t := NewTable(&Schema{
		Columns: []*Column{
			{Name: "Gender", Generalizer: &generalization.Suppressor{}},
			{Name: "Col 2", Generalizer: dim2},
			{Name: "Col 3", Generalizer: dim3},
			{Name: "Col 4", Generalizer: dim4},
			{Name: "Col 5", Generalizer: dim5},
		},
	})
	t.AddRow("Male", 25, 0, 35, "A")
	t.AddRow("Female", 25, 0, 45, "A+")
	t.AddRow("Male", 30, 2, 30, "B")
	t.AddRow("Female", 30, 1, 35, "B+")
	t.AddRow("Male", 28, 1, 40, "A-")
	t.AddRow("Female", 28, 1, 15, "B")
	t.AddRow("Male", 27, 0, 15, "B-")
	t.AddRow("Female", 27, 2, 30, "B")
	return t
}

func GetMixedTable1() *Table {
	t := NewTable(&Schema{
		Columns: []*Column{
			{"Score", generalization.ExampleIntGeneralizer()},
			{"Grade", generalization.ExampleGradeGeneralizer()},
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
			{"Score", generalization.ExampleIntGeneralizer()},
			{"Grade", generalization.ExampleGradeGeneralizer()},
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
			{"Score", generalization.ExampleIntGeneralizer()},
			{"Grade", generalization.ExampleGradeGeneralizer()},
			{"Motto", &generalization.PrefixGeneralizer{MaxWords: 5}},
			{"Remark", nil},
		},
	})
	t.AddRow(9, "A+", "cats are wild", "data1")
	t.AddRow(8, "A", "cats are evil", "data2")
	t.AddRow(6, "A-", "cats are fluffy", "data3")
	return t
}

// GetEmptyTable return an empty table with 3 rows
func GetEmptyTable() *Table {
	table := &Table{}
	return table
}
