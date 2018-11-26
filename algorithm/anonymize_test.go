package algorithm

import (
	"fmt"
	"gonum.org/v1/gonum/graph/topo"
	"k-anon/generalization"
	"k-anon/model"
	"k-anon/testutil"
	"strings"
	"testing"
)

// This is a testable example to demonstrate k-anonymization usage
// 1) supply a table to anonymize
//    * the table consists of row-vectors
//    * each row-vector should have the same 'schema'
//         (this is not enforced, but you will get an error during anonymization otherwise)
//    * each vector consists of data items
//    * there are two types of data items: identifier and non-identifier
//    * each identifier data item should have an associated generalizer: this can be based on a generalization
//      hierarchy, a custom generalizer, or something as simple as a value suppressor
//    * pre-built hierarchies can help you define your own without rolling your own custom
//    * non-identifiers only have a value, and will be ignored during the anonymization process
//    * it is recommended to assign the same generalizer for all data items in a column of the table
//         (again, this is not enforced but you will get an error during anonymization if the
//          respective data items cannot be generalized into the same partition)
// 2) create an Anonymizer instance, and supply the table and k parameters
// 3) call the AnonymizeData function on the Anonymizer
// 4) you will get the results in the form of a 2D slice of partitions
//       * each partition represents a level of generalization for a given value
//         in the respective generalization hierarchy
//       * items in partitions are not ordered (treat them as sets)
//       * you can create your own "pretty-printer" for partitions if needed. The basic cases
//         are already implemented, for example int range "[x..y]", and suppressed value "*".
func ExampleAnonymizer_AnonymizeData() {

	// Step 1:

	// define generalizers
	gender := &generalization.Suppressor{}
	age := generalization.NewHierarchyGeneralizer(
		(&generalization.IntegerHierarchyBuilder{Items: makeRange(0, 100, 1)}).NewIntegerHierarchy())
	kids := generalization.NewHierarchyGeneralizer(
		(&generalization.IntegerHierarchyBuilder{Items: []int{0, 1, 2, 3, 4, 5}}).NewIntegerHierarchy())
	income := generalization.NewHierarchyGeneralizer(
		(&generalization.IntegerHierarchyBuilder{Items: makeRange(0, 500000, 5000)}).NewIntegerHierarchy())
	grade := generalization.GetGradeGeneralizer()

	// define input table
	table := &model.Table{
		Rows: []*model.Vector{
			{
				Items: []*model.Data{
					model.NewNonIdentifier("Joe"),
					model.NewIdentifier("Male", gender),
					model.NewIdentifier(25, age),
					model.NewIdentifier(0, kids),
					model.NewIdentifier(10000, income),
					model.NewIdentifier("B", grade),
				},
			},
			{
				Items: []*model.Data{
					model.NewNonIdentifier("Jane"),
					model.NewIdentifier("Female", gender),
					model.NewIdentifier(25, age),
					model.NewIdentifier(0, kids),
					model.NewIdentifier(10000, income),
					model.NewIdentifier("A", grade),
				},
			},
			{
				Items: []*model.Data{
					model.NewNonIdentifier("Jack"),
					model.NewIdentifier("Male", gender),
					model.NewIdentifier(30, age),
					model.NewIdentifier(2, kids),
					model.NewIdentifier(30000, income),
					model.NewIdentifier("B", grade),
				},
			},
			{
				Items: []*model.Data{
					model.NewNonIdentifier("Janet"),
					model.NewIdentifier("Female", gender),
					model.NewIdentifier(30, age),
					model.NewIdentifier(1, kids),
					model.NewIdentifier(35000, income),
					model.NewIdentifier("B+", grade),
				},
			},
			{
				Items: []*model.Data{
					model.NewNonIdentifier("Steve"),
					model.NewIdentifier("Male", gender),
					model.NewIdentifier(28, age),
					model.NewIdentifier(1, kids),
					model.NewIdentifier(40000, income),
					model.NewIdentifier("A-", grade),
				},
			},
			{
				Items: []*model.Data{
					model.NewNonIdentifier("Sarah"),
					model.NewIdentifier("Female", gender),
					model.NewIdentifier(27, age),
					model.NewIdentifier(1, kids),
					model.NewIdentifier(15000, income),
					model.NewIdentifier("C", grade),
				},
			},
			{
				Items: []*model.Data{
					model.NewNonIdentifier("Ben"),
					model.NewIdentifier("Male", gender),
					model.NewIdentifier(25, age),
					model.NewIdentifier(0, kids),
					model.NewIdentifier(15000, income),
					model.NewIdentifier("C-", grade),
				},
			},
			{
				Items: []*model.Data{
					model.NewNonIdentifier("Anne"),
					model.NewIdentifier("Female", gender),
					model.NewIdentifier(30, age),
					model.NewIdentifier(2, kids),
					model.NewIdentifier(30000, income),
					model.NewIdentifier("B+", grade),
				},
			},
		},
	}

	// Step 2&3: create an anonymizer and run the anonymization
	anon := &Anonymizer{
		table: table,
		k:     3,
	}
	result := anon.anonymizeData()

	// Step 4: process the data
	sb := strings.Builder{}
	for i, row := range result {
		sb.WriteString(fmt.Sprintf("%d:\t", i))
		for j, col := range row {
			switch j {
			case 0, 5:
				sb.WriteString(col.String())
				break
			case 1, 2, 3, 4:
				s, err := col.IntRangeString()
				if err != nil {
					panic(err.Error())
				}
				sb.WriteString(s)
				break
			}
			sb.WriteString("\t")
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

func makeRange(from, to, increment int) []int {
	var r []int
	for i := from; i < to; i += increment {
		r = append(r, i)
	}
	return r
}

func TestAnonymizer_AnonymizeData(t *testing.T) {
	gen1 := generalization.GetIntGeneralizer()
	gen2 := generalization.GetGradeGeneralizer()
	table := &model.Table{
		Rows: []*model.Vector{
			{
				Items: []*model.Data{
					model.NewIdentifier(9, gen1),
					model.NewIdentifier("A+", gen2),
				},
			},
			{
				Items: []*model.Data{
					model.NewIdentifier(2, gen1),
					model.NewIdentifier("B-", gen2),
				},
			},
			{
				Items: []*model.Data{
					model.NewIdentifier(6, gen1),
					model.NewIdentifier("A-", gen2),
				},
			},
			{
				Items: []*model.Data{
					model.NewIdentifier(4, gen1),
					model.NewIdentifier("B+", gen2),
				},
			},
		},
	}
	anon := &Anonymizer{
		table: table,
		k:     2,
	}
	result := anon.anonymizeData()
	assertKAnonymity(table, result, 2, t)
}

func assertKAnonymity(table *model.Table, data [][]*generalization.Partition, k int, t *testing.T) {
	for i, r1 := range data {
		count := 0
		for _, r2 := range data {
			if inSamePartition(r1, r2, func(col int) bool {
				return table.Rows[0].Items[col].IsIdentifier()
			}) {
				count++
			}
		}
		if count < k {
			t.Errorf("k-anonimity violated in row %v", i)
		}
	}
}

func inSamePartition(r1, r2 []*generalization.Partition, isIdColumn func(int) bool) bool {
	for col := 0; col < len(r1); col++ {
		p1 := r1[col]
		p2 := r2[col]
		if isIdColumn(col) && !p1.Equals(p2) {
			return false
		}
	}
	return true
}

func TestGetGroups(t *testing.T) {
	v0 := model.CreateVector([]int{}, nil)
	v1 := model.CreateVector([]int{}, nil)
	v2 := model.CreateVector([]int{}, nil)
	v3 := model.CreateVector([]int{}, nil)
	v4 := model.CreateVector([]int{}, nil)
	table := &model.Table{Rows: []*model.Vector{v0, v1, v2, v3, v4}}
	a := &Anonymizer{
		table: table,
		k:     2,
	}
	g := CreateNodesUndirected(5)
	AddEdge(g, 0, 3)
	AddEdge(g, 1, 2)
	groups := a.getGroups(topo.ConnectedComponents(g))
	testutil.AssertEquals(3, len(groups), t)
	assertGroup(groups, v0, v3)
	assertGroup(groups, v1, v2)
	assertGroup(groups, v4)
}

func TestGeneralizeIdentifier(t *testing.T) {
	gen := generalization.GetIntGeneralizer()
	data := []*model.Data{
		model.NewIdentifier(1, gen),
		model.NewIdentifier(2, gen),
		model.NewIdentifier(7, gen),
	}
	partitions := generalize(data)
	expected := generalization.NewPartition(1, 2, 3, 4, 5, 6, 7, 8, 9)
	for _, p := range partitions {
		if !expected.Equals(p) {
			t.Errorf("incorrect partition: %v", p)
		}
	}
}

func TestGeneralizeNonIdentifier(t *testing.T) {
	data := []*model.Data{
		model.NewNonIdentifier("test1"),
		model.NewNonIdentifier("test2"),
		model.NewNonIdentifier("test3"),
	}
	partitions := generalize(data)
	p1 := generalization.NewPartition("test1")
	p2 := generalization.NewPartition("test2")
	p3 := generalization.NewPartition("test3")
	for _, p := range partitions {
		if !(p.Equals(p1) || p.Equals(p2) || p.Equals(p3)) {
			t.Errorf("incorrect partition: %v", p)
		}
	}
}

func TestAnonymize(t *testing.T) {
	gen1 := generalization.GetIntGeneralizer()
	gen2 := generalization.GetGradeGeneralizer()
	groups := []*model.Vector{
		{
			Items: []*model.Data{
				model.NewIdentifier(9, gen1),
				model.NewIdentifier("A+", gen2),
				model.NewNonIdentifier("data1"),
			},
		},
		{
			Items: []*model.Data{
				model.NewIdentifier(8, gen1),
				model.NewIdentifier("A", gen2),
				model.NewNonIdentifier("data2"),
			},
		},
		{
			Items: []*model.Data{
				model.NewIdentifier(6, gen1),
				model.NewIdentifier("A-", gen2),
				model.NewNonIdentifier("data3"),
			},
		},
	}
	partitions := anonymize(groups)
	testutil.AssertEquals(3, len(partitions), t)
	assertSamePartition([]*generalization.Partition{
		partitions[0][0],
		partitions[1][0],
		partitions[2][0]}, t)
	assertSamePartition([]*generalization.Partition{
		partitions[0][1],
		partitions[1][1],
		partitions[2][1]}, t)
	partitions[0][2].Equals(generalization.NewPartition("data1"))
	partitions[1][2].Equals(generalization.NewPartition("data2"))
	partitions[2][2].Equals(generalization.NewPartition("data3"))
}

func assertSamePartition(p []*generalization.Partition, t *testing.T) {
	first := p[0]
	for i := 1; i < len(p); i++ {
		if !first.Equals(p[i]) {
			t.Errorf("partitions are not equal: %v, %v", first, p[i])
		}
	}
}

func assertGroup(groups [][]*model.Vector, items ...*model.Vector) bool {
	for _, group := range groups {
		if composedOf(group, items...) {
			return true
		}
	}
	return false
}

func composedOf(group []*model.Vector, items ...*model.Vector) bool {
	if len(group) != len(items) {
		return false
	}
	for _, item := range items {
		if !contains(group, item) {
			return false
		}
	}
	return true
}

func contains(group []*model.Vector, item *model.Vector) bool {
	for _, i := range group {
		if i == item {
			return true
		}
	}
	return false
}
