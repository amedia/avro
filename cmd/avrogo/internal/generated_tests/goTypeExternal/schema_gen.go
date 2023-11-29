// Code generated by avrogen. DO NOT EDIT.

package goTypeExternal

import (
	"github.com/amedia/avro/avrotypegen"
	"github.com/amedia/avro/internal/testtypes"
)

type R struct {
	F testtypes.Message
	G testtypes.CloudEvent
	H string
}

// AvroRecord implements the avro.AvroRecord interface.
func (R) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"F","type":{"fields":[{"name":"Metadata","type":{"fields":[{"name":"CloudEvent","type":{"fields":[{"name":"id","type":"string"},{"name":"source","type":"string"},{"name":"specversion","type":"string"},{"name":"time","type":{"logicalType":"timestamp-micros","type":"long"}}],"name":"CloudEvent","type":"record"}}],"name":"Metadata","type":"record"}}],"go.package":"github.com/amedia/avro/internal/testtypes","name":"com.heetch.Message","type":"record"}},{"name":"G","type":"com.heetch.CloudEvent"},{"name":"H","type":"string"}],"name":"R","type":"record"}`,
		Required: []bool{
			0: true,
			1: true,
			2: true,
		},
	}
}
