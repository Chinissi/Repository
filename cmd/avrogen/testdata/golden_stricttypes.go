// Code generated by avro/gen. DO NOT EDIT.
package testpkg

// Test is a test struct.
type Test struct {
	// SomeString is a string.
	SomeString        string            `avro:"someString"`
	SomeInt           int32             `avro:"someInt"`
	SomeNullableMap   *map[string]int32 `avro:"someNullableMap"`
	SomeNullableSlice *[]int32          `avro:"someNullableSlice"`
}
