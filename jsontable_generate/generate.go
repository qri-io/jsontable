package jsontable_generate

import (
	"math/rand"
	"time"

	"github.com/qri-io/datatype"
	"github.com/qri-io/datatype/datatype_generate"
	"github.com/qri-io/jsontable"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateRandomTableOpt specifies the options for GenerateRandomTable
type RandomTableOpt struct {
	// use a provided name instead of a random one
	Name string
	// number of random fields to generate, default between 1 & 10
	NumFields int
	// constrict creation to provided types, blank means any valid datatype
	Datatypes []datatype.Type
	// set fields to get a specific set of fields back
	// overrides numfields, datatypes
	Fields []*jsontable.Field
}

// GenerateRandomTable creates a random valid table. Provide option func(s) to customize
func RandomTable(options ...func(*RandomTableOpt)) *jsontable.Table {
	opt := &RandomTableOpt{
		Name:      randString(16),
		NumFields: rand.Intn(9) + 1,
		Fields:    nil,
	}
	for _, option := range options {
		option(opt)
	}

	if opt.Fields == nil {
		opt.Fields = make([]*jsontable.Field, opt.NumFields)
		for i := 0; i < opt.NumFields; i++ {
			opt.Fields[i] = RandomField(func(o *RandomFieldOpt) {
				o.Datatypes = opt.Datatypes
			})
		}
	}

	return &jsontable.Table{
		Fields: opt.Fields,
	}
}

// RandomFieldOpt are the options for RandomField
type RandomFieldOpt struct {
	// use a provided name instead of a random one
	Name jsontable.Name
	// use a provided type instead of a random one
	Type datatype.Type
	// constrict random types to a provided set, blank means any valid datatype
	Datatypes []datatype.Type
}

// RandomField generates a random field, optionally configured
func RandomField(options ...func(*RandomFieldOpt)) *jsontable.Field {
	opt := &RandomFieldOpt{
		Name:      RandomName(16),
		Datatypes: nil,
	}
	for _, option := range options {
		option(opt)
	}

	if opt.Type == datatype.Unknown {
		if opt.Datatypes != nil {
			opt.Type = opt.Datatypes[rand.Intn((len(opt.Datatypes)-1))+1]
		} else {
			opt.Type = datatype.Type(rand.Intn(9) + 1)
		}
	}

	return &jsontable.Field{
		Name: opt.Name,
		Type: opt.Type,
	}
}

// Random Rows generates random row data
func RandomRows(tbl *jsontable.Table, numRows int) (rows [][]interface{}) {
	rows = make([][]interface{}, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]interface{}, len(tbl.Fields))
		for j, field := range tbl.Fields {
			row[j] = datatype_generate.RandomValue(field.Type)
		}
		rows[i] = row
	}

	return
}

func RandomName(maxLength int) jsontable.Name {
	return jsontable.Name(randString(rand.Intn(maxLength-1) + 1))
}

var alphaNumericRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphaNumericRunes[rand.Intn(len(alphaNumericRunes))]
	}
	return string(b)
}
