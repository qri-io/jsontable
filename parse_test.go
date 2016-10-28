package jsontable

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		FileName string
		err      error
	}{
		{"airport-codes.json", nil},
		{"continent-codes.json", nil},
	}

	for i, c := range cases {
		data, err := ioutil.ReadFile(filepath.Join("test_files", c.FileName))
		if err != nil {
			t.Errorf("case %d couldn't read file: %s", i, err.Error())
		}

		tbl := &Table{}
		if err := json.Unmarshal(data, tbl); err != c.err {
			t.Errorf("case %d parse error mismatch. expected: '%s', got: '%s'", i, c.err, err)
		}

	}
}
