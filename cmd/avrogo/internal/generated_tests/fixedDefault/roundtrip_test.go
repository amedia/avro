// Code generated by generatetestcode.go; DO NOT EDIT.

package fixedDefault

import (
	"testing"

	"github.com/amedia/avro/cmd/avrogo/internal/testutil"
)

var tests = testutil.RoundTripTest{
	InSchema: `{
                "name": "R",
                "type": "record",
                "fields": [
                    {
                        "name": "_",
                        "type": "int",
                        "default": 0
                    }
                ]
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "_": 0
                    }`,
		OutDataJSON: `{
                        "fixedField": "hello"
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
