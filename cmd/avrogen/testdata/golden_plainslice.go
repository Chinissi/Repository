// Code generated by avro/gen. DO NOT EDIT.
package testpkg

// Test is a test struct.
type Test struct {
	// SomeString is a string.
	SomeString        string          `avro:"someString"`
	SomeInt           int             `avro:"someInt"`
	SomeNullableMap   *map[string]int `avro:"someNullableMap"`
	SomeNullableSlice []int           `avro:"someNullableSlice"`
}
