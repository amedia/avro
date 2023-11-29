// Code generated by avrogen. DO NOT EDIT.

package primitiveDefaults

import (
	"github.com/amedia/avro/avrotypegen"
)

type R struct {
	Int     int     `json:"int"`
	Long    int64   `json:"long"`
	String  string  `json:"string"`
	Float   float32 `json:"float"`
	Double  float64 `json:"double"`
	Boolean bool    `json:"boolean"`
}

// AvroRecord implements the avro.AvroRecord interface.
func (R) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"default":1111,"name":"int","type":"int"},{"default":2222,"name":"long","type":"long"},{"default":"hello","name":"string","type":"string"},{"default":1.5,"name":"float","type":"float"},{"default":2.75,"name":"double","type":"double"},{"default":true,"name":"boolean","type":"boolean"}],"name":"R","type":"record"}`,
		Defaults: []func() interface{}{
			0: func() interface{} {
				return 1111
			},
			1: func() interface{} {
				return int64(2222)
			},
			2: func() interface{} {
				return "hello"
			},
			3: func() interface{} {
				return float32(1.5)
			},
			4: func() interface{} {
				return 2.75
			},
			5: func() interface{} {
				return true
			},
		},
	}
}
