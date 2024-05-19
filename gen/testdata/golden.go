package something

// Code generated by avro/gen. DO NOT EDIT MANUALLY.

import (
	"math/big"
	"time"

	"github.com/hamba/avro/v2"
)

type InnerRecord struct {
	InnerJustBytes                   []byte    `avro:"innerJustBytes"`
	InnerPrimitiveNullableArrayUnion *[]string `avro:"innerPrimitiveNullableArrayUnion"`
}

type RecordInMap struct {
	Name string `avro:"name"`
}

type RecordInArray struct {
	AString string `avro:"aString"`
}

type RecordInNullableUnion struct {
	AString string `avro:"aString"`
}

type Record1InNonNullableUnion struct {
	AString string `avro:"aString"`
}

type Record2InNonNullableUnion struct {
	AString string `avro:"aString"`
}

type Record1InNullableUnion struct {
	AString string `avro:"aString"`
}

type Record2InNullableUnion struct {
	AString string `avro:"aString"`
}

// Test represents a golden record
type Test struct {
	// aString is just a string
	AString                         string                 `avro:"aString"`
	ABoolean                        bool                   `avro:"aBoolean"`
	AnInt                           int                    `avro:"anInt"`
	AFloat                          float32                `avro:"aFloat"`
	ADouble                         float64                `avro:"aDouble"`
	ALong                           int64                  `avro:"aLong"`
	JustBytes                       []byte                 `avro:"justBytes"`
	PrimitiveNullableArrayUnion     *[]string              `avro:"primitiveNullableArrayUnion"`
	InnerRecord                     InnerRecord            `avro:"innerRecord"`
	AnEnum                          string                 `avro:"anEnum"`
	AFixed                          [7]byte                `avro:"aFixed"`
	ALogicalFixed                   avro.LogicalDuration   `avro:"aLogicalFixed"`
	AnotherLogicalFixed             avro.LogicalDuration   `avro:"anotherLogicalFixed"`
	MapOfStrings                    map[string]string      `avro:"mapOfStrings"`
	MapOfRecords                    map[string]RecordInMap `avro:"mapOfRecords"`
	ADate                           time.Time              `avro:"aDate"`
	ADuration                       time.Duration          `avro:"aDuration"`
	ALongTimeMicros                 time.Duration          `avro:"aLongTimeMicros"`
	ALongTimestampMillis            time.Time              `avro:"aLongTimestampMillis"`
	ALongTimestampMicro             time.Time              `avro:"aLongTimestampMicro"`
	ABytesDecimal                   *big.Rat               `avro:"aBytesDecimal"`
	ARecordArray                    []RecordInArray        `avro:"aRecordArray"`
	NullableRecordUnion             *RecordInNullableUnion `avro:"nullableRecordUnion"`
	NonNullableRecordUnion          any                    `avro:"nonNullableRecordUnion"`
	NullableRecordUnionWith3Options any                    `avro:"nullableRecordUnionWith3Options"`
	Ref                             Record2InNullableUnion `avro:"ref"`
	UUID                            string                 `avro:"uuid"`
}
