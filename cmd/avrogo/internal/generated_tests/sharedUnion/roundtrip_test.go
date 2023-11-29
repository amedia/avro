// Code generated by generatetestcode.go; DO NOT EDIT.

package sharedUnion

import (
	"testing"

	"github.com/heetch/avro/cmd/avrogo/internal/testutil"
)

var tests = testutil.RoundTripTest{
	InSchema: `{
                "name": "R",
                "type": "record",
                "fields": [
                    {
                        "name": "A",
                        "type": [
                            "int",
                            "string",
                            "float"
                        ]
                    },
                    {
                        "name": "B",
                        "type": [
                            "int",
                            "string",
                            "float"
                        ]
                    }
                ]
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "A": {
                            "int": 244
                        },
                        "B": {
                            "string": "hello"
                        }
                    }`,
		OutDataJSON: `{
                        "A": {
                            "int": 244
                        },
                        "B": {
                            "string": "hello"
                        }
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
