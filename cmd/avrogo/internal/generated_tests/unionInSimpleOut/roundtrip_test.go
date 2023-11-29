// Code generated by generatetestcode.go; DO NOT EDIT.

package unionInSimpleOut

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
                        "name": "UnionField",
                        "type": [
                            "int",
                            "long",
                            "float",
                            "double",
                            "string",
                            "boolean",
                            "null"
                        ],
                        "default": 1234
                    }
                ]
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "UnionField": {
                            "string": "hello"
                        }
                    }`,
		OutDataJSON: `{
                        "UnionField": "hello"
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
