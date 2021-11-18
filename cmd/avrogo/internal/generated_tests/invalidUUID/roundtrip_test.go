// Code generated by generatetestcode.go; DO NOT EDIT.

package invalidUUID

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
                        "name": "T",
                        "type": {
                            "type": "string",
                            "logicalType": "uuid"
                        }
                    }
                ]
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "T": "invalid_uuid"
                    }`,
		OutDataJSON: `null`,
		ExpectError: map[testutil.ErrorType]string{`unmarshal`: `invalid UUID in Avro encoding: invalid UUID length: 12`},
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}