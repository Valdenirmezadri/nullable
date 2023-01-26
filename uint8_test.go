package nullable_test

import (
	"testing"

	"github.com/Valdenirmezadri/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestScanUint8(t *testing.T) {
	nullableInt := nullable.NewUint8(0)

	// uint8
	nullableInt.Scan(37)
	tests.AssertEqual(t, nullableInt.Get(), 37)

	nullableInt.Scan(nil)
	tests.AssertEqual(t, nullableInt.Get(), 0)
}

func TestNewUint8(t *testing.T) {
	// uint8
	var basicInt1 uint8 = 37
	nullableInt1 := nullable.NewUint8(basicInt1)
	tests.AssertEqual(t, nullableInt1.Get(), 37)

}

func TestSetUint8(t *testing.T) {
	nullableInt := nullable.NewUint8(0)
	tests.AssertEqual(t, nullableInt.Get(), 0)

	// uint8
	var basicInt1 uint8 = 37
	nullableInt.Set(basicInt1)
	tests.AssertEqual(t, nullableInt.Get(), 37)

	nullableInt.Set(0)
	tests.AssertEqual(t, nullableInt.Get(), 0)
}

func TestJSONUint8(t *testing.T) {
	var basicInt1 uint8 = 37
	marshalUnmarshalJSON(t, nullable.NewUint8(basicInt1))

	marshalUnmarshalJSON(t, nullable.NewUint8(0))
}

func TestUint8(t *testing.T) {
	type TestNullableInt8 struct {
		ID    uint
		Name  string
		Value nullable.Uint8
		Unit  string
	}

	DB.Migrator().DropTable(&TestNullableInt8{})
	if err := DB.Migrator().AutoMigrate(&TestNullableInt8{}); err != nil {
		t.Errorf("failed to migrate nullable int8, got error: %v", err)
	}

	var matterEnergy uint8 = 117
	matter := TestNullableInt8{
		Name:  "matter",
		Value: nullable.NewUint8(matterEnergy),
		Unit:  "Joule",
	}
	DB.Create(&matter)

	neutron := TestNullableInt8{
		Name:  "neutron",
		Value: nullable.NewUint8(0),
		Unit:  "Joule",
	}
	DB.Create(&neutron)

	var result1 TestNullableInt8
	if err := DB.First(&result1, "name = ?", "matter").Error; err != nil {
		t.Fatal("Cannot read int8 test record of \"matter\"")
	}
	tests.AssertEqual(t, result1, matter)

	var result3 TestNullableInt8
	if err := DB.First(&result3, "name = ?", "neutron").Error; err != nil {
		t.Fatal("Cannot read int8 test record of \"neutron\"")
	}
	tests.AssertEqual(t, result3, neutron)
}
