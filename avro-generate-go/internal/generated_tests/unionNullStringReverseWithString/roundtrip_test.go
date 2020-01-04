// Code generated by generatetestcode.go; DO NOT EDIT.

package unionNullStringReverseWithString

import (
	"testing"

	"github.com/rogpeppe/avro/avro-generate-go/internal/testutil"
)

var test = testutil.RoundTripTest{
	InDataJSON: `{
                "OptionalString": {
                    "string": "hello"
                }
            }`,
	OutDataJSON: `{
                "OptionalString": {
                    "string": "hello"
                }
            }`,
	InSchema: `{
                "name": "R",
                "type": "record",
                "fields": [
                    {
                        "name": "OptionalString",
                        "type": [
                            "string",
                            "null"
                        ]
                    }
                ]
            }`,
	OutSchema: `{
                "name": "R",
                "type": "record",
                "fields": [
                    {
                        "name": "OptionalString",
                        "type": [
                            "string",
                            "null"
                        ]
                    }
                ]
            }`,
	GoType: new(R),
}

func TestGeneratedCode(t *testing.T) {
	test.Test(t)
}