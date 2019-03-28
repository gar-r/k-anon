package k_anon

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/model"
	"fmt"
	"strings"
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
	suppressor := &generalization.Suppressor{}
	age := generalization.NewHierarchyGeneralizer(
		(&generalization.IntegerHierarchyBuilder{Items: makeRange(0, 100, 1)}).NewIntegerHierarchy())
	kids := generalization.NewHierarchyGeneralizer(
		(&generalization.IntegerHierarchyBuilder{Items: []int{0, 1, 2, 3, 4, 5}}).NewIntegerHierarchy())
	income := generalization.NewHierarchyGeneralizer(
		(&generalization.IntegerHierarchyBuilder{Items: makeRange(0, 500000, 5000)}).NewIntegerHierarchy())
	grade := generalization.GetGradeGeneralizer()
	motto := &generalization.PrefixGeneralizer{MaxWords: 100}

	// define input table
	table := &model.Table{
		Rows: []*model.Vector{
			{
				Items: []*model.Data{
					model.NewIdentifier("Joe", suppressor),
					model.NewNonIdentifier("employee"),
					model.NewIdentifier("Male", suppressor),
					model.NewIdentifier(25, age),
					model.NewIdentifier(0, kids),
					model.NewIdentifier(10000, income),
					model.NewIdentifier("B", grade),
					model.NewIdentifier("cats are wonderful little beings", motto),
				},
			},
			{
				Items: []*model.Data{
					model.NewIdentifier("Jane", suppressor),
					model.NewNonIdentifier("client"),
					model.NewIdentifier("Female", suppressor),
					model.NewIdentifier(25, age),
					model.NewIdentifier(0, kids),
					model.NewIdentifier(10000, income),
					model.NewIdentifier("A", grade),
					model.NewIdentifier("cats are my favorite kind of animals", motto),
				},
			},
			{
				Items: []*model.Data{
					model.NewIdentifier("Jack", suppressor),
					model.NewNonIdentifier("employee"),
					model.NewIdentifier("Male", suppressor),
					model.NewIdentifier(30, age),
					model.NewIdentifier(2, kids),
					model.NewIdentifier(30000, income),
					model.NewIdentifier("B", grade),
					model.NewIdentifier("cats are very unique", motto),
				},
			},
			{
				Items: []*model.Data{
					model.NewIdentifier("Janet", suppressor),
					model.NewNonIdentifier("employee"),
					model.NewIdentifier("Female", suppressor),
					model.NewIdentifier(30, age),
					model.NewIdentifier(1, kids),
					model.NewIdentifier(35000, income),
					model.NewIdentifier("A+", grade),
					model.NewIdentifier("cats are interesting", motto),
				},
			},
			{
				Items: []*model.Data{
					model.NewIdentifier("Steve", suppressor),
					model.NewNonIdentifier("client"),
					model.NewIdentifier("Male", suppressor),
					model.NewIdentifier(28, age),
					model.NewIdentifier(1, kids),
					model.NewIdentifier(40000, income),
					model.NewIdentifier("A-", grade),
					model.NewIdentifier("cats are my only pets", motto),
				},
			},
			{
				Items: []*model.Data{
					model.NewIdentifier("Sarah", suppressor),
					model.NewNonIdentifier("client"),
					model.NewIdentifier("Female", suppressor),
					model.NewIdentifier(27, age),
					model.NewIdentifier(1, kids),
					model.NewIdentifier(15000, income),
					model.NewIdentifier("B", grade),
					model.NewIdentifier("cats are my favorite!", motto),
				},
			},
			{
				Items: []*model.Data{
					model.NewIdentifier("Ben", suppressor),
					model.NewNonIdentifier("employee"),
					model.NewIdentifier("Male", suppressor),
					model.NewIdentifier(25, age),
					model.NewIdentifier(0, kids),
					model.NewIdentifier(15000, income),
					model.NewIdentifier("B-", grade),
					model.NewIdentifier("cats are interesting, but sometimes also egoistic", motto),
				},
			},
			{
				Items: []*model.Data{
					model.NewIdentifier("Anne", suppressor),
					model.NewNonIdentifier("client"),
					model.NewIdentifier("Female", suppressor),
					model.NewIdentifier(30, age),
					model.NewIdentifier(2, kids),
					model.NewIdentifier(30000, income),
					model.NewIdentifier("B+", grade),
					model.NewIdentifier("cats are my favorite kind of animals", motto),
				},
			},
		},
	}

	// Step 2&3: create an anonymizer and run the anonymization
	anon := &Anonymizer{
		table: table,
		k:     2,
	}
	result := anon.anonymizeData()

	// Step 4: process the data
	prettyPrintResult(result)

	// Since the partitioning has a random element, the algorithm might
	// produce slightly different (but still correct) outputs for the same input.
	// In this case for example the output might look something like this:
	//0:	[*]	[employee]	[Female]	[25..30]	[1..1]	[0..55000]		[A-, B+, B, B-, A+, A]	[cats are]
	//1:	[*]	[client]	[Female]	[25..30]	[1..1]	[0..55000]		[B+, B, B-, A+, A, A-]	[cats are]
	//2:	[*]	[client]	[Male]		[28..30]	[1..2]	[30000..40000]	[A, A-, B+, B, B-, A+]	[cats are]
	//3:	[*]	[employee]	[Male]		[28..30]	[1..2]	[30000..40000]	[A, A-, B+, B, B-, A+]	[cats are]
	//4:	[*]	[client]	[Female]	[25..30]	[0..2]	[0..55000]		[A+, A, A-, B+, B, B-]	[cats are my favorite kind of animals]
	//5:	[*]	[client]	[Female]	[25..30]	[0..2]	[0..55000]		[B, B-, A+, A, A-, B+]	[cats are my favorite kind of animals]
	//6:	[*]	[employee]	[Male]		[25..25]	[0..0]	[0..25000]		[B+, B, B-]				[cats are]
	//7:	[*]	[employee]	[Male]		[25..25]	[0..0]	[0..25000]		[B+, B, B-]				[cats are]
}

func prettyPrintResult(result [][]*generalization.Partition) {
	sb := strings.Builder{}
	for row, part := range result {
		sb.WriteString(fmt.Sprintf("%d:\t", row))
		for col, val := range part {
			switch col {
			case 0, 1, 2, 6, 7:
				sb.WriteString(val.String())
			case 3, 4, 5:
				s, err := val.IntRangeString()
				if err != nil {
					panic(err.Error())
				}
				sb.WriteString(s)
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
