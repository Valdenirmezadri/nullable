package nullable_test

import (
	"testing"

	"github.com/Valdenirmezadri/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestScanInt64(t *testing.T) {
	nullableInt := nullable.NewInt64(nil)

	// uint8
	nullableInt.Scan(37)
	tests.AssertEqual(t, nullableInt.Get(), 37)

	// int8
	nullableInt.Scan(-37)
	tests.AssertEqual(t, nullableInt.Get(), -37)

	// uint16
	nullableInt.Scan(1234)
	tests.AssertEqual(t, nullableInt.Get(), 1234)

	// int16
	nullableInt.Scan(-1234)
	tests.AssertEqual(t, nullableInt.Get(), -1234)

	// uint32
	nullableInt.Scan(654321)
	tests.AssertEqual(t, nullableInt.Get(), 654321)

	// int32
	nullableInt.Scan(-654321)
	tests.AssertEqual(t, nullableInt.Get(), -654321)

	// uint64
	nullableInt.Scan(50000000000)
	tests.AssertEqual(t, nullableInt.Get(), 50000000000)

	// int64
	nullableInt.Scan(-50000000000)
	tests.AssertEqual(t, nullableInt.Get(), -50000000000)

	nullableInt.Scan(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)
}

func TestNewInt64(t *testing.T) {
	// uint8
	var basicInt1 int64 = 37
	nullableInt1 := nullable.NewInt64(&basicInt1)
	tests.AssertEqual(t, nullableInt1.Get(), 37)

	// int8
	var basicInt2 int64 = -37
	nullableInt2 := nullable.NewInt64(&basicInt2)
	tests.AssertEqual(t, nullableInt2.Get(), -37)

	// uint16
	var basicInt3 int64 = 1234
	nullableInt3 := nullable.NewInt64(&basicInt3)
	tests.AssertEqual(t, nullableInt3.Get(), 1234)

	// int16
	var basicInt4 int64 = -1234
	nullableInt4 := nullable.NewInt64(&basicInt4)
	tests.AssertEqual(t, nullableInt4.Get(), -1234)

	// uint32
	var basicInt5 int64 = 654321
	nullableInt5 := nullable.NewInt64(&basicInt5)
	tests.AssertEqual(t, nullableInt5.Get(), 654321)

	// int32
	var basicInt6 int64 = -654321
	nullableInt6 := nullable.NewInt64(&basicInt6)
	tests.AssertEqual(t, nullableInt6.Get(), -654321)

	// uint64
	var basicInt7 int64 = 50000000000
	nullableInt7 := nullable.NewInt64(&basicInt7)
	tests.AssertEqual(t, nullableInt7.Get(), 50000000000)

	// int64
	var basicInt8 int64 = -50000000000
	nullableInt8 := nullable.NewInt64(&basicInt8)
	tests.AssertEqual(t, nullableInt8.Get(), -50000000000)
}

func TestSetInt64(t *testing.T) {
	nullableInt := nullable.NewInt64(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)

	// uint8
	var basicInt1 int64 = 37
	nullableInt.Set(&basicInt1)
	tests.AssertEqual(t, nullableInt.Get(), 37)

	// int8
	var basicInt2 int64 = -37
	nullableInt.Set(&basicInt2)
	tests.AssertEqual(t, nullableInt.Get(), -37)

	// uint16
	var basicInt3 int64 = 1234
	nullableInt.Set(&basicInt3)
	tests.AssertEqual(t, nullableInt.Get(), 1234)

	// int16
	var basicInt4 int64 = -1234
	nullableInt.Set(&basicInt4)
	tests.AssertEqual(t, nullableInt.Get(), -1234)

	// uint32
	var basicInt5 int64 = 654321
	nullableInt.Set(&basicInt5)
	tests.AssertEqual(t, nullableInt.Get(), 654321)

	// int32
	var basicInt6 int64 = -654321
	nullableInt.Set(&basicInt6)
	tests.AssertEqual(t, nullableInt.Get(), -654321)

	// uint64
	var basicInt7 int64 = 50000000000
	nullableInt.Set(&basicInt7)
	tests.AssertEqual(t, nullableInt.Get(), 50000000000)

	// int64
	var basicInt8 int64 = -50000000000
	nullableInt.Set(&basicInt8)
	tests.AssertEqual(t, nullableInt.Get(), -50000000000)

	nullableInt.Set(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)
}

func TestJSONInt64(t *testing.T) {
	var basicInt1 int64 = 37
	marshalUnmarshalJSON(t, nullable.NewInt64(&basicInt1))

	var basicInt2 int64 = -37
	marshalUnmarshalJSON(t, nullable.NewInt64(&basicInt2))

	var basicInt3 int64 = 1234
	marshalUnmarshalJSON(t, nullable.NewInt64(&basicInt3))

	var basicInt4 int64 = -1234
	marshalUnmarshalJSON(t, nullable.NewInt64(&basicInt4))

	var basicInt5 int64 = 654321
	marshalUnmarshalJSON(t, nullable.NewInt64(&basicInt5))

	var basicInt6 int64 = -654321
	marshalUnmarshalJSON(t, nullable.NewInt64(&basicInt6))

	var basicInt7 int64 = 50000000000
	marshalUnmarshalJSON(t, nullable.NewInt64(&basicInt7))

	var basicInt8 int64 = -50000000000
	marshalUnmarshalJSON(t, nullable.NewInt64(&basicInt8))

	marshalUnmarshalJSON(t, nullable.NewInt64(nil))
}

func TestInt64(t *testing.T) {
	type TestNullableInt64 struct {
		ID    uint
		Name  string
		Value nullable.Int64
		Unit  string
	}

	DB.Migrator().DropTable(&TestNullableInt64{})
	if err := DB.Migrator().AutoMigrate(&TestNullableInt64{}); err != nil {
		t.Errorf("failed to migrate nullable int64, got error: %v", err)
	}

	var matterEnergy int64 = 50000000000
	matter := TestNullableInt64{
		Name:  "matter",
		Value: nullable.NewInt64(&matterEnergy),
		Unit:  "Joule",
	}
	DB.Create(&matter)

	var antimatterEnergy int64 = -50000000000
	antimatter := TestNullableInt64{
		Name:  "antimatter",
		Value: nullable.NewInt64(&antimatterEnergy),
		Unit:  "Joule",
	}
	DB.Create(&antimatter)

	neutron := TestNullableInt64{
		Name:  "neutron",
		Value: nullable.NewInt64(nil),
		Unit:  "Joule",
	}
	DB.Create(&neutron)

	var result1 TestNullableInt64
	if err := DB.First(&result1, "name = ?", "matter").Error; err != nil {
		t.Fatal("Cannot read int64 test record of \"matter\"")
	}
	tests.AssertEqual(t, result1, matter)

	var result2 TestNullableInt64
	if err := DB.First(&result2, "name = ?", "antimatter").Error; err != nil {
		t.Fatal("Cannot read int64 test record of \"antimatter\"")
	}
	tests.AssertEqual(t, result2, antimatter)

	var result3 TestNullableInt64
	if err := DB.First(&result3, "name = ?", "neutron").Error; err != nil {
		t.Fatal("Cannot read int64 test record of \"neutron\"")
	}
	tests.AssertEqual(t, result3, neutron)
}
