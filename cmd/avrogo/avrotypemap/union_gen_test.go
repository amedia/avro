// Code generated by avrogen. DO NOT EDIT.

package avrotypemap_test

import (
	"github.com/heetch/avro/avrotypegen"
)

type U struct {
	// Allowed types for interface{} value:
	// 	UR1
	// 	UR2
	F interface{}
}

// AvroRecord implements the avro.AvroRecord interface.
func (U) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"F","type":[{"fields":[{"name":"A","type":"int"}],"name":"UR1","type":"record"},{"fields":[{"name":"B","type":"int"}],"name":"UR2","type":"record"}]}],"name":"U","type":"record"}`,
		Required: []bool{
			0: true,
		},
		Unions: []avrotypegen.UnionInfo{
			0: {
				Type: new(interface{}),
				Union: []avrotypegen.UnionInfo{{
					Type: new(UR1),
				}, {
					Type: new(UR2),
				}},
			},
		},
	}
}

type UR1 struct {
	A int
}

// AvroRecord implements the avro.AvroRecord interface.
func (UR1) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"A","type":"int"}],"name":"UR1","type":"record"}`,
		Required: []bool{
			0: true,
		},
	}
}

type UR2 struct {
	B int
}

// AvroRecord implements the avro.AvroRecord interface.
func (UR2) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"B","type":"int"}],"name":"UR2","type":"record"}`,
		Required: []bool{
			0: true,
		},
	}
}
