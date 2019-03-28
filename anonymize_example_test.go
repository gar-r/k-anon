package k_anon

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/model"
	"fmt"
)

// This is a testable example to demonstrate the usage of the anonymizer.
// The outline of the steps you need to take is the following:
// 1) Define a table
//    a) define the schema (friendly name and generalizer for each column)
//          * use a 'nil' generalizer for non-identifier (skipped) columns
//          * see the Generalizer interface to implement a custom generalizer
//    b) define the rows, conforming to the above schema:
//          * number of columns should match
//          * data type should be compatible with the assigned generalizer for the column
// 2) create an Anonymizer instance, and supply the table and k parameters
// 3) call the Anonymize() function
// The supplied table will be anonymized in-place. Note, that items in partitions
// in the resulting table are not ordered (always treat them as sets).
func ExampleAnonymizer_AnonymizeData() {

	// define the schema & table
	table := &model.Table{
		Schema: &model.Schema{
			Columns: []*model.Column{
				{"Name", &generalization.Suppressor{}},
				{"Status", nil},
				{"Gender", &generalization.Suppressor{}},
				{"Age", generalization.NewIntGeneralizer(0, 120, 1)},
				{"Kids", generalization.NewIntGeneralizerFromItems(0, 1, 2, 3, 4, 5)},
				{"Income", generalization.NewIntGeneralizer(10000, 40001, 1)},
				{"Grade", generalization.GetGradeGeneralizer()},
				{"Motto", &generalization.PrefixGeneralizer{MaxWords: 100}},
			},
		},
		Rows: []*model.Row{
			model.NewRow("Joe", "employee", "male", 25, 0, 10000, "B", "cats are wonderful little beings"),
			model.NewRow("Jane", "client", "female", 25, 0, 10000, "A", "cats are my favorite kind of animals "),
			model.NewRow("Jack", "employee", "male", 30, 2, 30000, "B", "cats are very unique"),
			model.NewRow("Janet", "employee", "female", 30, 1, 35000, "A+", "cats are interesting"),
			model.NewRow("Steve", "client", "male", 28, 1, 40000, "A-", "cats are my only pets"),
			model.NewRow("Sarah", "client", "female", 27, 1, 15000, "B", "cats are my favorite!"),
			model.NewRow("Ben", "employee", "male", 25, 0, 15000, "B-", "cats are interesting, but sometimes also egoistic"),
			model.NewRow("Anne", "client", "female", 30, 2, 30000, "B+", "cats are my favorite kind of animals"),
		},
	}

	// create an anonymizer and run the anonymization
	anon := &Anonymizer{
		table: table,
		k:     2,
	}
	anon.Anonymize()

	// access & print the data
	fmt.Printf("%v", anon.table)

	// the above will produce a similar output:

	// Name	Status	Gender	Age			Kids	Income			Grade					Motto
	//[*]	[employee]	[*]	[25]		[0]		[10000]			[B+, B, B-, A+, A, A-]	[cats are]
	//[*]	[client]	[*]	[25]		[0]		[10000]			[A+, A, A-, B+, B, B-]	[cats are]
	//[*]	[employee]	[*]	[30]		[1..2]	[25000..40000]	[A-, B+, B, B-, A+, A]	[cats are]
	//[*]	[employee]	[*]	[30]		[1..2]	[25000..40000]	[A-, B+, B, B-, A+, A]	[cats are]
	//[*]	[client]	[*]	[22..29]	[0..2]	[10000..40000]	[A+, A, A-, B+, B, B-]	[cats are]
	//[*]	[client]	[*]	[22..29]	[0..2]	[10000..40000]	[B, B-, A+, A, A-, B+]	[cats are]
	//[*]	[employee]	[*]	[22..29]	[0..2]	[10000..40000]	[A+, A, A-, B+, B, B-]	[cats are]
	//[*]	[client]	[*]	[30]		[1..2]	[25000..40000]	[B-, A+, A, A-, B+, B]	[cats are]
}
