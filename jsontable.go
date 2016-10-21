// jsontable implements the JSON Table Schema spec http://specs.frictionlessdata.io/json-table-schema/
//
// From the Spec:
// A JSON Table Schema consists of:
// a required list of field descriptors
// optionally, a primary key description
// optionally, a foreign _key description
// A schema is described using JSON. This might exist as a standalone document or may be embedded within another JSON structure, e.g. as part of a data package description.
package jsontable

type Schema struct {
	Fields []*Field `json:"fields"`
}

type JsonTable struct {
	Fields      []*Field      `json:"fields"`
	PrimaryKey  FieldKey      `json:"primaryKey,omitempty"`
	ForeignKeys []*ForeignKey `json:"foreignKeys,omtiempty"`
}
