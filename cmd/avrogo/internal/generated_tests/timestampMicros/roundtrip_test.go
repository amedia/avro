// Code generated by generatetestcode.go; DO NOT EDIT.

package timestampMicros

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
                        "name": "T",
                        "type": {
                            "type": "long",
                            "logicalType": "timestamp-micros"
                        }
                    }
                ]
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "T": 1579176162000001
                    }`,
		OutDataJSON: `{
                        "T": 1579176162000001
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
