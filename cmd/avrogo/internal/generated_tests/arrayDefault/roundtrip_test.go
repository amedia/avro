// Code generated by generatetestcode.go; DO NOT EDIT.

package arrayDefault

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
                        "arrayOfInt": [
                            2,
                            3,
                            4
                        ]
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
