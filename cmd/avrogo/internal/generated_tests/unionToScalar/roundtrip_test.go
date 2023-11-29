// Code generated by generatetestcode.go; DO NOT EDIT.

package unionToScalar

import (
	"testing"

	"github.com/heetch/avro/cmd/avrogo/internal/testutil"
)

var tests = testutil.RoundTripTest{
	InSchema: `{
                "name": "PrimitiveUnionTestRecord",
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
	GoType: new(PrimitiveUnionTestRecord),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "UnionField": {
                            "int": 999
                        }
                    }`,
		OutDataJSON: `{
                        "UnionField": 999
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
