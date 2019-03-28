package k_anon

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/model"
	"fmt"
)

func ExampleAnonymizer_Continuous() {

	schema := &model.Schema{
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
	}

	table := &model.Table{
		Schema: schema,
		Rows:   []*model.Row{},
	}

	// anonymize first batch of data
	table.Add(
		model.NewRow("Joe", "employee", "male", 25, 0, 10000, "B", "cats are wonderful little beings"),
		model.NewRow("Jane", "client", "female", 25, 0, 10000, "A", "cats are my favorite kind of animals "),
		model.NewRow("Jack", "employee", "male", 30, 2, 30000, "B", "cats are very unique"),
		model.NewRow("Janet", "employee", "female", 30, 1, 35000, "A+", "cats are interesting"),
		model.NewRow("Steve", "client", "male", 28, 1, 40000, "A-", "dogs are my only pets"),
		model.NewRow("Sarah", "client", "female", 27, 1, 15000, "B", "dogs are my favorite!"),
		model.NewRow("Ben", "employee", "male", 25, 0, 15000, "B-", "cats are interesting, but sometimes also egoistic"),
		model.NewRow("Anne", "client", "female", 30, 2, 30000, "B+", "cats are my favorite kind of animals"),
	)

	anon := &Anonymizer{
		table: table,
		k:     2,
	}
	anon.Anonymize()

	// anonymize second batch
	table.Add(
		model.NewRow("Michelle", "employee", "female", 27, 1, 20000, "B", "cats are secretly extraterrestrials"),
		model.NewRow("Perseus", "client", "male", 31, 2, 38000, "A+", "cats are mischievous"),
		model.NewRow("Odysseus", "client", "male", 39, 3, 39000, "A", "dogs are war bringers"),
		model.NewRow("Helene", "employee", "female", 29, 3, 21000, "B+", "cats are silky and furry"),
	)

	anon.Anonymize()

	// anonymize third batch
	table.Add(
		model.NewRow("Donald", "client", "male", 26, 3, 20000, "A", "cats are secretly extraterrestrials"),
		model.NewRow("Hillary", "client", "female", 70, 2, 38000, "B-", "dogs are loyal"),
		model.NewRow("George", "client", "male", 65, 2, 39000, "B", "dogs are war bringers"),
		model.NewRow("Victor", "employee", "male", 45, 2, 21000, "A-", "cats are silky and furry"),
	)

	anon.Anonymize()

	// process & print results
	fmt.Printf("%v", anon.table)
	// Output:
	// 1
}
