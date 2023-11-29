// Code generated by generatetestcode.go; DO NOT EDIT.

package multiSchemaMutualRecursive

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
                        "name": "F",
                        "type": {
                            "type": "array",
                            "items": {
                                "name": "S",
                                "type": "record",
                                "fields": [
                                    {
                                        "name": "Data",
                                        "type": "string",
                                        "default": ""
                                    },
                                    {
                                        "name": "Child",
                                        "type": [
                                            "null",
                                            "R"
                                        ],
                                        "default": null
                                    }
                                ]
                            }
                        },
                        "default": []
                    }
                ]
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "F": [
                            {
                                "Data": "hello",
                                "Child": {
                                    "R": {
                                        "F": [
                                            {
                                                "Data": "goodbye",
                                                "Child": null
                                            }
                                        ]
                                    }
                                }
                            },
                            {
                                "Data": "whee",
                                "Child": null
                            }
                        ]
                    }`,
		OutDataJSON: `{
                        "F": [
                            {
                                "Data": "hello",
                                "Child": {
                                    "R": {
                                        "F": [
                                            {
                                                "Data": "goodbye",
                                                "Child": null
                                            }
                                        ]
                                    }
                                }
                            },
                            {
                                "Data": "whee",
                                "Child": null
                            }
                        ]
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
