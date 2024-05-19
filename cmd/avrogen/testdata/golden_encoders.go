package testpkg

// Code generated by avro/gen. DO NOT EDIT MANUALLY.

import (
	"github.com/hamba/avro/v2"
)

// Test is a test struct
type Test struct {
	// SomeString is a string
	SomeString string `avro:"someString"`
	SomeInt    int    `avro:"someInt"`
}

var schemaTest = avro.MustParse(`{"name":"a.b.test","type":"record","fields":[{"name":"someString","type":"string"},{"name":"someInt","type":"int"}]}`)

// Schema returns the schema for Test.
func (o *Test) Schema() avro.Schema {
	return schemaTest
}

// Unmarshal decodes b into the receiver.
func (o *Test) Unmarshal(b []byte) error {
	return avro.Unmarshal(o.Schema(), b, o)
}

// Marshal encodes the receiver.
func (o *Test) Marshal() ([]byte, error) {
	return avro.Marshal(o.Schema(), o)
}
