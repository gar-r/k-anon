package model

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"fmt"
	"strings"
)

// Table contains data organized in a rectangular shape. It has a fixed
// amount of columns, which is constrained by a schema.
// Each row must in the table must conform to the table schema.
type Table struct {
	Schema *Schema
	Rows   []*Row
}

func (t *Table) Add(row ...*Row) {
	t.Rows = append(t.Rows, row...)
}

func (t *Table) String() string {
	sb := &strings.Builder{}
	t.appendHeader(sb)
	for _, row := range t.Rows {
		for colIdx := range t.Schema.Columns {
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
	for _, col := range t.Schema.Columns {
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
	Data []*generalization.Partition
}

func NewRow(items ...interface{}) *Row {
	var data []*generalization.Partition
	for _, item := range items {
		data = append(data, generalization.NewPartition(item))
	}
	return &Row{Data: data}
}
