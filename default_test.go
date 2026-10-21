package nullable_test

import (
	"encoding/json"
	"testing"

	"github.com/Valdenirmezadri/nullable"
	"gorm.io/gorm/utils/tests"
)

// zero vindo do cache (JSON) ou do banco (Scan) precisa sair NULL no Value
func TestZeroIsNull(t *testing.T) {
	var u nullable.Uint
	var u8 nullable.Uint8
	var u16 nullable.Uint16
	var u32 nullable.Uint32
	var s nullable.String

	tests.AssertEqual(t, json.Unmarshal([]byte("0"), &u), nil)
	tests.AssertEqual(t, json.Unmarshal([]byte("0"), &u8), nil)
	tests.AssertEqual(t, json.Unmarshal([]byte("0"), &u16), nil)
	tests.AssertEqual(t, json.Unmarshal([]byte("0"), &u32), nil)
	tests.AssertEqual(t, json.Unmarshal([]byte(`""`), &s), nil)
	for _, v := range []interface{ IsNull() bool }{u, u8, u16, u32, s} {
		tests.AssertEqual(t, v.IsNull(), true)
	}

	u.Scan(int64(0))
	u8.Scan(int64(0))
	u16.Scan(int64(0))
	u32.Scan(int64(0))
	s.Scan("")
	for _, v := range []interface{ IsNull() bool }{u, u8, u16, u32, s} {
		tests.AssertEqual(t, v.IsNull(), true)
	}

	u.Set(5)
	u.Set(0)
	tests.AssertEqual(t, u.Get(), uint(0))
	value, _ := u.Value()
	tests.AssertEqual(t, value, nil)
}
