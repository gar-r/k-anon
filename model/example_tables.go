package model

import "k-anon/generalization"

// Table data:
// 		(each column  -> int generalizer)
// 1 1 1 1
// 1 1 1 2
// 4 5 1 1
// 1 3 5 7
func GetIntTable1() *Table {
	g := generalization.GetIntGeneralizer1()
	return &Table{
		Rows: []*Vector{
			CreateVector([]int{1, 1, 1, 1}, g),
			CreateVector([]int{1, 1, 1, 2}, g),
			CreateVector([]int{4, 5, 1, 1}, g),
			CreateVector([]int{1, 3, 5, 7}, g),
		},
	}
}

// Table data:
// 		(col1 -> suppress)
// 		(col2 -> int)
// 		(col3 -> int)
// 		(col4 -> int)
// 		(col5 -> grade)
// Male		25 	0 	35 	A
// Female	25	0	45	A+
// Male		30	2	30	B
// Female	30	1	35	B+
// Male		28	1	40	A-
// Female	28	1	15	B
// Male		27	0	15	B-
// Female	27	2	30	B
func GetStudentTable() *Table {
	dim1 := &generalization.Suppressor{}
	dim2 := generalization.NewHierarchyGeneralizer((&generalization.IntegerHierarchyBuilder{Items: []int{25, 27, 28, 30}}).NewIntegerHierarchy())
	dim3 := generalization.NewHierarchyGeneralizer((&generalization.IntegerHierarchyBuilder{Items: []int{0, 1, 2}}).NewIntegerHierarchy())
	dim4 := generalization.NewHierarchyGeneralizer((&generalization.IntegerHierarchyBuilder{Items: []int{10, 15, 30, 35, 40, 45}}).NewIntegerHierarchy())
	dim5 := generalization.GetGradeGeneralizer1()
	return &Table{
		Rows: []*Vector{
			{Items: []*Data{NewData("Male", dim1), NewData(25, dim2), NewData(0, dim3), NewData(35, dim4), NewData("A", dim5)}},
			{Items: []*Data{NewData("Female", dim1), NewData(25, dim2), NewData(0, dim3), NewData(45, dim4), NewData("A+", dim5)}},
			{Items: []*Data{NewData("Male", dim1), NewData(30, dim2), NewData(2, dim3), NewData(30, dim4), NewData("B", dim5)}},
			{Items: []*Data{NewData("Female", dim1), NewData(30, dim2), NewData(1, dim3), NewData(35, dim4), NewData("B+", dim5)}},
			{Items: []*Data{NewData("Male", dim1), NewData(28, dim2), NewData(1, dim3), NewData(40, dim4), NewData("A-", dim5)}},
			{Items: []*Data{NewData("Female", dim1), NewData(28, dim2), NewData(1, dim3), NewData(15, dim4), NewData("B", dim5)}},
			{Items: []*Data{NewData("Male", dim1), NewData(27, dim2), NewData(0, dim3), NewData(15, dim4), NewData("B-", dim5)}},
			{Items: []*Data{NewData("Female", dim1), NewData(27, dim2), NewData(2, dim3), NewData(30, dim4), NewData("B", dim5)}},
		},
	}
}

// Table data:
// 		(col1 -> int)
// 		(col2 -> grade)
// 9 	A+
// 8	A
// 5	B-
func GetMixedTable1() *Table {
	generalizer1 := generalization.GetIntGeneralizer1()
	generalizer2 := generalization.GetGradeGeneralizer1()
	return &Table{
		Rows: []*Vector{
			{
				Items: []*Data{
					NewData(9, generalizer1),
					NewData("A+", generalizer2),
				},
			},
			{
				Items: []*Data{
					NewData(8, generalizer1),
					NewData("A", generalizer2),
				},
			},
			{
				Items: []*Data{
					NewData(5, generalizer1),
					NewData("B-", generalizer2),
				},
			},
		},
	}
}

// GetEmptyTable return an empty table with 3 rows
func GetEmptyTable() *Table {
	table := &Table{
		Rows: []*Vector{
			{},
			{},
			{},
		},
	}
	return table
}

func CreateVector(items []int, g generalization.Generalizer) *Vector {
	v := &Vector{}
	for _, item := range items {
		v.Items = append(v.Items, NewData(item, g))
	}
	return v
}
