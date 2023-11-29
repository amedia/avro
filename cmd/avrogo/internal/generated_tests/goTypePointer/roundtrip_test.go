// Code generated by generatetestcode.go; DO NOT EDIT.

package goTypePointer

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
                        "name": "A",
                        "type": [
                            "null",
                            "long"
                        ]
                    }
                ]
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "non_null",
		InDataJSON: `{
                        "A": {
                            "long": 99
                        }
                    }`,
		OutDataJSON: `{
                        "A": {
                            "long": 99
                        }
                    }`,
	}, {
		TestName: "null",
		InDataJSON: `{
                        "A": null
                    }`,
		OutDataJSON: `{
                        "A": null
                    }`,
	}},
}

type R struct {
	A *int
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
