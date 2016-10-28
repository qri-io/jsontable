// jsontable implements the JSON Table Schema spec http://specs.frictionlessdata.io/json-table-schema/
//
// From the Spec:
// A JSON Table Schema consists of:
// a required list of field descriptors
// optionally, a primary key description
// optionally, a foreign _key description
// A schema is described using JSON. This might exist as a standalone document or may be embedded within another JSON structure, e.g. as part of a data package description.
package jsontable

import (
	"fmt"
)

type Table struct {
	Fields      []*Field      `json:"fields"`
	PrimaryKey  FieldKey      `json:"primaryKey,omitempty"`
	ForeignKeys []*ForeignKey `json:"foreignKeys,omtiempty"`
}

func (t *Table) RowToStrings(row []interface{}) (strs []string, err error) {
	if len(row) != len(t.Fields) {
		err = fmt.Errorf("row is not the same length as the table's fields")
		return
	}
	strs = make([]string, len(t.Fields))
	for i, field := range t.Fields {
		str, err := field.Type.ValueToString(row[i])
		if err != nil {
			return nil, err
		}
		strs[i] = str
	}
	return
}

func (t *Table) RowToBytes(row []interface{}) (bytes [][]byte, err error) {
	if len(row) != len(t.Fields) {
		err = fmt.Errorf("row is not the same length as the table's fields")
		return
	}
	bytes = make([][]byte, len(t.Fields))
	for i, field := range t.Fields {
		val, err := field.Type.ValueToBytes(row[i])
		if err != nil {
			return nil, err
		}
		bytes[i] = val
	}
	return
}
