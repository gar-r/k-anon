package model

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/partition"
	"fmt"
	"strings"
)

// Table contains data organized in a rectangular shape. It has a fixed
// amount of columns, which is constrained by a schema.
// Each row must in the table must conform to the table schema.
type Table struct {
	schema *Schema
	rows   []*Row
}

func NewTable(schema *Schema) *Table {
	return &Table{schema: schema}
}

func (t *Table) AddRow(items ...interface{}) {
	var data []partition.Partition
	for i, item := range items {
		if i >= len(t.schema.Columns) {
			break
		}
		col := t.schema.Columns[i]
		var p partition.Partition
		if col.Generalizer != nil {
			p = t.schema.Columns[i].Generalizer.InitItem(item)
		} else {
			p = partition.NewItem(item)
		}
		data = append(data, p)
	}
	t.rows = append(t.rows, &Row{Data: data})
}

func (t *Table) GetSchema() *Schema {
	return t.schema
}

func (t *Table) GetRows() []*Row {
	return t.rows
}

func (t *Table) String() string {
	sb := &strings.Builder{}
	t.appendHeader(sb)
	for _, row := range t.rows {
		for colIdx := range t.schema.Columns {
			p := row.Data[colIdx]
			sb.WriteString(p.String())
			sb.WriteString("\t")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (t *Table) appendHeader(sb *strings.Builder) {
	sb.WriteString("\n")
	for _, col := range t.schema.Columns {
		sb.WriteString(fmt.Sprintf("%s\t", col.Name))
	}
	sb.WriteString("\n")
}

// Schema defines the number of columns, and their attributes in a table.
type Schema struct {
	Columns []*Column
}

// Column represents a column definition in a table schema.
// If the generalizer is set to nil, the column will be treated as non-identifier.
type Column struct {
	Name        string
	Generalizer generalization.Generalizer
}

func (c *Column) IsIdentifier() bool {
	return c.Generalizer != nil
}

// Row represents a row of data in a table.
type Row struct {
	Data []partition.Partition
}
