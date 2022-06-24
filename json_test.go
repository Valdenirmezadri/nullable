package nullable_test

import (
	"encoding/json"
	"testing"

	"github.com/Valdenirmezadri/nullable"
	"gorm.io/gorm/utils/tests"
)

func marshalUnmarshalJSON(t *testing.T, target interface{}) {
	serialized, err := json.Marshal(target)
	if err != nil {
		t.Fatalf("Failed to marshal %T", target)
		return
	}

	switch target.(type) {
	case nullable.Uint:
		var unserialized nullable.Uint
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	case nullable.String:
		var unserialized nullable.String
		if err := json.Unmarshal(serialized, &unserialized); err != nil {
			t.Fatalf("Failed to unmarshal %T because: %s", target, err)
			return
		}
		tests.AssertEqual(t, unserialized, target)
	default:
		t.Fatalf("%T is not registered at json_test.go", target)
	}
}
