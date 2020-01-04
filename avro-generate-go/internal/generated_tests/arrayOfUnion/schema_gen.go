// Code generated by avrogen. DO NOT EDIT.

package arrayOfUnion

import "github.com/rogpeppe/avro"

type R struct {
	F []interface{}
}

// AvroRecord implements the avro.AvroRecord interface.
func (R) AvroRecord() avro.RecordInfo {
	return avro.RecordInfo{
		Schema: `{"fields":[{"name":"F","type":{"items":["int","string"],"type":"array"}}],"name":"R","type":"record"}`,
		Required: []bool{
			0: true,
		},
		Unions: [][]interface{}{
			0: {new(int), new(string)},
		},
	}
}

// TODO implement MarshalBinary and UnmarshalBinary methods?