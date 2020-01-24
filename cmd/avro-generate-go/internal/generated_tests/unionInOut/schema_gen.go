// Code generated by avrogen. DO NOT EDIT.

package unionInOut

import (
	"github.com/heetch/avro/avrotypegen"
)

type PrimitiveUnionTestRecord struct {
	UnionField interface{}
}

// AvroRecord implements the avro.AvroRecord interface.
func (PrimitiveUnionTestRecord) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"default":1234,"name":"UnionField","type":["int","long","float","double","string","boolean","null"]}],"name":"PrimitiveUnionTestRecord","type":"record"}`,
		Defaults: []func() interface{}{
			0: func() interface{} {
				return 1234
			},
		},
		Unions: []avrotypegen.UnionInfo{
			0: {
				Type: new(interface{}),
				Union: []avrotypegen.UnionInfo{{
					Type: new(int),
				}, {
					Type: new(int64),
				}, {
					Type: new(float32),
				}, {
					Type: new(float64),
				}, {
					Type: new(string),
				}, {
					Type: new(bool),
				}, {
					Type: nil,
				}},
			},
		},
	}
}

// TODO implement MarshalBinary and UnmarshalBinary methods?