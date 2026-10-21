package nullable_test

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"
	"time"

	"github.com/Valdenirmezadri/nullable"
	"gorm.io/gorm/utils/tests"
)

func TestNewTime(t *testing.T) {
	now := time.Date(2026, 12, 20, 18, 0, 0, 0, time.UTC)

	tests.AssertEqual(t, nullable.NewTime(now).Get(), now)
	tests.AssertEqual(t, nullable.NewTime(now).IsNull(), false)
	tests.AssertEqual(t, nullable.NewTime(time.Time{}).IsNull(), true)
}

func TestSetTime(t *testing.T) {
	now := time.Date(2026, 12, 20, 18, 0, 0, 0, time.UTC)
	n := nullable.NewTime(time.Time{})

	n.Set(now)
	tests.AssertEqual(t, n.Get(), now)
	tests.AssertEqual(t, n.IsNull(), false)

	n.Set(time.Time{})
	tests.AssertEqual(t, n.Get().IsZero(), true)
	tests.AssertEqual(t, n.IsNull(), true)
}

func TestScanTime(t *testing.T) {
	now := time.Date(2026, 12, 20, 18, 0, 0, 0, time.UTC)
	n := nullable.NewTime(time.Time{})

	tests.AssertEqual(t, n.Scan(now), nil)
	tests.AssertEqual(t, n.Get(), now)

	tests.AssertEqual(t, n.Scan("2026-12-20 18:00:00+00:00"), nil)
	tests.AssertEqual(t, n.Get().Equal(now), true)

	tests.AssertEqual(t, n.Scan(nil), nil)
	tests.AssertEqual(t, n.IsNull(), true)

	if n.Scan(42) == nil {
		t.Fatal("scan de tipo não suportado deve falhar")
	}
}

func TestValueTime(t *testing.T) {
	now := time.Date(2026, 12, 20, 18, 0, 0, 0, time.UTC)

	v, _ := nullable.NewTime(now).Value()
	tests.AssertEqual(t, v, now)

	saoPaulo := now.In(time.FixedZone("-03", -3*3600))
	v, _ = nullable.NewTime(saoPaulo).Value()
	tests.AssertEqual(t, v, now)

	v, _ = nullable.NewTime(time.Time{}).Value()
	tests.AssertEqual(t, v, nil)
}

func TestJSONTime(t *testing.T) {
	marshalUnmarshalJSON(t, nullable.NewTime(time.Date(2026, 12, 20, 18, 0, 0, 0, time.UTC)))
	marshalUnmarshalJSON(t, nullable.NewTime(time.Time{}))

	empty, _ := json.Marshal(nullable.NewTime(time.Time{}))
	tests.AssertEqual(t, string(empty), `"0001-01-01T00:00:00Z"`)

	var fromNull nullable.Time
	tests.AssertEqual(t, json.Unmarshal([]byte("null"), &fromNull), nil)
	tests.AssertEqual(t, fromNull.IsNull(), true)
}

func TestGobTime(t *testing.T) {
	for _, in := range []nullable.Time{
		nullable.NewTime(time.Date(2026, 12, 20, 18, 0, 0, 0, time.UTC)),
		nullable.NewTime(time.Time{}),
	} {
		var buf bytes.Buffer
		if err := gob.NewEncoder(&buf).Encode(in); err != nil {
			t.Fatal(err)
		}
		var out nullable.Time
		if err := gob.NewDecoder(&buf).Decode(&out); err != nil {
			t.Fatal(err)
		}
		tests.AssertEqual(t, out.Get().Equal(in.Get()), true)
		tests.AssertEqual(t, out.IsNull(), in.IsNull())
	}
}

func TestGormTime(t *testing.T) {
	type Event struct {
		ID uint
		At nullable.Time
	}

	DB.Migrator().DropTable(&Event{})
	if err := DB.AutoMigrate(&Event{}); err != nil {
		t.Fatal(err)
	}

	now := time.Date(2026, 12, 20, 18, 0, 0, 0, time.UTC)
	with, without := Event{At: nullable.NewTime(now)}, Event{}
	DB.Create(&with)
	DB.Create(&without)

	var got Event
	DB.First(&got, with.ID)
	tests.AssertEqual(t, got.At.Get().Equal(now), true)

	var gotNull Event
	DB.First(&gotNull, without.ID)
	tests.AssertEqual(t, gotNull.At.IsNull(), true)
}
