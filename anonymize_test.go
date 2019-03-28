package k_anon

import (
	"bitbucket.org/dargzero/k-anon/model"
	"fmt"
	"testing"
)

func TestAnonymizer_Anonymize(t *testing.T) {

	t.Run("test k-anonymity", func(t *testing.T) {
		tables := []*model.Table{
			model.GetIntTable1(),
			model.GetMixedTable1(),
			model.GetMixedTable2(),
			model.GetMixedTable3(),
			model.GetStudentTable(),
		}
		for i, table := range tables {
			t.Run(fmt.Sprintf("table %d", i), func(t *testing.T) {
				anon := &Anonymizer{
					table: table,
					k:     2,
				}
				anon.Anonymize()
				t.Log(fmt.Sprintf("%v", anon.table))
				assertKAnonymity(table, 2, t)
			})
		}
	})
}

func assertKAnonymity(table *model.Table, k int, t *testing.T) {
	for i, r1 := range table.Rows {
		count := 0
		for _, r2 := range table.Rows {
			if inSamePartition(r1, r2, table.Schema) {
				count++
			}
		}
		if count < k {
			t.Errorf("k-anonimity violated in row %v", i)
		}
	}
}

func inSamePartition(r1, r2 *model.Row, schema *model.Schema) bool {
	for c, col := range schema.Columns {
		if col.IsIdentifier() {
			p1 := r1.Data[c]
			p2 := r2.Data[c]
			if !p1.Equals(p2) {
				return false
			}
		}
	}
	return true
}
