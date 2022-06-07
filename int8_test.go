package nullable_test

import (
	"testing"

	"github.com/Valdenirmezadri/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestScanInt8(t *testing.T) {
	nullableInt := nullable.NewInt8(nil)

	// uint8
	nullableInt.Scan(37)
	tests.AssertEqual(t, nullableInt.Get(), 37)

	// int8
	nullableInt.Scan(-37)
	tests.AssertEqual(t, nullableInt.Get(), -37)

	nullableInt.Scan(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)
}

func TestNewInt8(t *testing.T) {
	// uint8
	var basicInt1 int8 = 37
	nullableInt1 := nullable.NewInt8(&basicInt1)
	tests.AssertEqual(t, nullableInt1.Get(), 37)

	// int8
	var basicInt2 int8 = -37
	nullableInt2 := nullable.NewInt8(&basicInt2)
	tests.AssertEqual(t, nullableInt2.Get(), -37)
}

func TestSetInt8(t *testing.T) {
	nullableInt := nullable.NewInt8(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)

	// uint8
	var basicInt1 int8 = 37
	nullableInt.Set(&basicInt1)
	tests.AssertEqual(t, nullableInt.Get(), 37)

	// int8
	var basicInt2 int8 = -37
	nullableInt.Set(&basicInt2)
	tests.AssertEqual(t, nullableInt.Get(), -37)

	nullableInt.Set(nil)
	tests.AssertEqual(t, nullableInt.Get(), nil)
}

func TestJSONInt8(t *testing.T) {
	var basicInt1 int8 = 37
	marshalUnmarshalJSON(t, nullable.NewInt8(&basicInt1))

	var basicInt2 int8 = -37
	marshalUnmarshalJSON(t, nullable.NewInt8(&basicInt2))

	marshalUnmarshalJSON(t, nullable.NewInt8(nil))
}

func TestInt8(t *testing.T) {
	type TestNullableInt8 struct {
		ID    uint
		Name  string
		Value nullable.Int8
		Unit  string
	}

	DB.Migrator().DropTable(&TestNullableInt8{})
	if err := DB.Migrator().AutoMigrate(&TestNullableInt8{}); err != nil {
		t.Errorf("failed to migrate nullable int8, got error: %v", err)
	}

	var matterEnergy int8 = 117
	matter := TestNullableInt8{
		Name:  "matter",
		Value: nullable.NewInt8(&matterEnergy),
		Unit:  "Joule",
	}
	DB.Create(&matter)

	var antimatterEnergy int8 = -117
	antimatter := TestNullableInt8{
		Name:  "antimatter",
		Value: nullable.NewInt8(&antimatterEnergy),
		Unit:  "Joule",
	}
	DB.Create(&antimatter)

	neutron := TestNullableInt8{
		Name:  "neutron",
		Value: nullable.NewInt8(nil),
		Unit:  "Joule",
	}
	DB.Create(&neutron)

	var result1 TestNullableInt8
	if err := DB.First(&result1, "name = ?", "matter").Error; err != nil {
		t.Fatal("Cannot read int8 test record of \"matter\"")
	}
	tests.AssertEqual(t, result1, matter)

	var result2 TestNullableInt8
	if err := DB.First(&result2, "name = ?", "antimatter").Error; err != nil {
		t.Fatal("Cannot read int8 test record of \"antimatter\"")
	}
	tests.AssertEqual(t, result2, antimatter)

	var result3 TestNullableInt8
	if err := DB.First(&result3, "name = ?", "neutron").Error; err != nil {
		t.Fatal("Cannot read int8 test record of \"neutron\"")
	}
	tests.AssertEqual(t, result3, neutron)
}
