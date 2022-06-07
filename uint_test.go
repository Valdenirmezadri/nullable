package nullable_test

import (
	"testing"

	"github.com/Valdenirmezadri/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestScanUint(t *testing.T) {
	nullableInt := nullable.NewUint(0)

	// uint8
	nullableInt.Scan(37)
	tests.AssertEqual(t, nullableInt.Get(), 37)

	// uint16
	nullableInt.Scan(1234)
	tests.AssertEqual(t, nullableInt.Get(), 1234)

	// uint32
	nullableInt.Scan(654321)
	tests.AssertEqual(t, nullableInt.Get(), 654321)

	// uint64
	nullableInt.Scan(50000000000)
	tests.AssertEqual(t, nullableInt.Get(), 50000000000)

	nullableInt.Scan(nil)
	tests.AssertEqual(t, nullableInt.Get(), 0)
}

func TestNewUint(t *testing.T) {
	// uint8
	var basicUint1 uint = 37
	nullableUint1 := nullable.NewUint(basicUint1)
	tests.AssertEqual(t, nullableUint1.Get(), 37)

	// uint16
	var basicUint2 uint = 1234
	nullableUint2 := nullable.NewUint(basicUint2)
	tests.AssertEqual(t, nullableUint2.Get(), 1234)

	// uint32
	var basicUint3 uint = 654321
	nullableUint3 := nullable.NewUint(basicUint3)
	tests.AssertEqual(t, nullableUint3.Get(), 654321)

	// uint64
	var basicUint4 uint = 50000000000
	nullableUint4 := nullable.NewUint(basicUint4)
	tests.AssertEqual(t, nullableUint4.Get(), 50000000000)

	nullableUint5 := nullable.NewUint(0)
	tests.AssertEqual(t, nullableUint5.Get(), 0)
}

func TestSetUint(t *testing.T) {
	nullableUint := nullable.NewUint(0)
	tests.AssertEqual(t, nullableUint.Get(), 0)

	// uint8
	var basicUint1 uint = 37
	nullableUint.Set(basicUint1)
	tests.AssertEqual(t, nullableUint.Get(), 37)
	// uint16
	var basicUint2 uint = 1234
	nullableUint.Set(basicUint2)
	tests.AssertEqual(t, nullableUint.Get(), 1234)

	// uint32
	var basicUint3 uint = 654321
	nullableUint.Set(basicUint3)
	tests.AssertEqual(t, nullableUint.Get(), 654321)

	// uint64
	var basicUint4 uint = 50000000000
	nullableUint.Set(basicUint4)
	tests.AssertEqual(t, nullableUint.Get(), 50000000000)

	nullableUint.Set(0)
	tests.AssertEqual(t, nullableUint.Get(), 0)
}

func TestJSONUint(t *testing.T) {
	var basicInt1 uint = 37
	marshalUnmarshalJSON(t, nullable.NewUint(basicInt1))

	var basicInt2 uint = 1234
	marshalUnmarshalJSON(t, nullable.NewUint(basicInt2))

	var basicInt3 uint = 654321
	marshalUnmarshalJSON(t, nullable.NewUint(basicInt3))

	var basicInt4 uint = 50000000000
	marshalUnmarshalJSON(t, nullable.NewUint(basicInt4))

	marshalUnmarshalJSON(t, nullable.NewUint(0))
}
