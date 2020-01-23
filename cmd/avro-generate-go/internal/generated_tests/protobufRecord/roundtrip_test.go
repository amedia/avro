// Code generated by generatetestcode.go; DO NOT EDIT.

package protobufRecord

import (
	"testing"

	"github.com/heetch/avro/cmd/avro-generate-go/internal/testutil"
)

var tests = testutil.RoundTripTest{
	InSchema: `{
                "name": "MessageB",
                "type": "record",
                "fields": []
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName:   "main",
		InDataJSON: `{}`,
		OutDataJSON: `{
                        "arble": null,
                        "selected": false
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
