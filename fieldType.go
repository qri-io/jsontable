package jsontable

import (
	"encoding/json"
	"fmt"
)

type FieldType int

const (
	UnknownFieldType FieldType = iota
	StringFieldType
	NumberFieldType
	IntegerFieldType
	BooleanFieldType
	ObjectFieldType
	ArrayFieldType
	DateFieldType
	DurationFieldType
	GeopointFieldType
	AnyFieldType
)

func (ft FieldType) String() string {
	s, ok := map[FieldType]string{
		UnknownFieldType:  "",
		StringFieldType:   "string",
		NumberFieldType:   "number",
		IntegerFieldType:  "integer",
		BooleanFieldType:  "boolean",
		ObjectFieldType:   "object",
		ArrayFieldType:    "array",
		DateFieldType:     "date",
		DurationFieldType: "duration",
		GeopointFieldType: "geopoint",
		AnyFieldType:      "any",
	}[ft]

	if !ok {
		return ""
	}

	return s
}

func (ft FieldType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, ft.String())), nil
}

func (ft *FieldType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Filed type should be a string, got %s", data)
	}

	got, ok := map[string]FieldType{
		"":         UnknownFieldType,
		"string":   StringFieldType,
		"number":   NumberFieldType,
		"integer":  IntegerFieldType,
		"boolean":  BooleanFieldType,
		"object":   ObjectFieldType,
		"array":    ArrayFieldType,
		"date":     DateFieldType,
		"duration": DurationFieldType,
		"geopoint": GeopointFieldType,
		"any":      AnyFieldType,
	}[s]
	if !ok {
		return fmt.Errorf("invalid FieldType %q", s)
	}

	*ft = got
	return nil
}
