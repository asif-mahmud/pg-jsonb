package pgjsonb_test

import (
	"testing"

	"github.com/asif-mahmud/pgjsonb"
)

func TestJsonNullValue(t *testing.T) {
	var json pgjsonb.JSONB
	json.Set(nil)
	value, err := json.Value()
	if err != nil {
		t.Error(err)
	}
	v := value.([]byte)
	if string(v) != `null` {
		t.Error("expected", nil, "found", value)
	}
}

func TestJsonNumberValue(t *testing.T) {
	var json pgjsonb.JSONB
	json.Set(1.5)
	value, err := json.Value()
	if err != nil {
		t.Error(err)
	}
	v := value.([]byte)
	if string(v) != `1.5` {
		t.Error("expected", 1.5, "found", value)
	}
}

func TestJsonStringValue(t *testing.T) {
	var json pgjsonb.JSONB
	json.Set("data")
	value, err := json.Value()
	if err != nil {
		t.Error(err)
	}
	v := value.([]byte)
	if string(v) != `"data"` {
		t.Error("expected", "data", "found", value)
	}
}

func TestJsonArrayValue(t *testing.T) {
	var json pgjsonb.JSONB
	json.Set([]interface{}{1, "2"})
	value, err := json.Value()
	if err != nil {
		t.Error(err)
	}
	v := value.([]byte)
	expected := `[1,"2"]`
	if string(v) != expected {
		t.Error("expected", expected, "found", value)
	}
}

func TestJsonObjectValue(t *testing.T) {
	var json pgjsonb.JSONB
	json.Set(map[string]any{
		"a": 1,
		"b": []any{1, "2"},
	})
	value, err := json.Value()
	if err != nil {
		t.Error(err)
	}
	v := value.([]byte)
	expected := `{"a":1,"b":[1,"2"]}`
	if string(v) != expected {
		t.Error("expected", expected, "found", value)
	}
}
