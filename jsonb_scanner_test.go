package pgjsonb_test

import (
	"testing"

	pgjsonb "github.com/asif-mahmud/pg-jsonb"
)

func TestConstructor(t *testing.T) {
	data := map[string]any{
		"value": 1,
	}
	expected := `{"value":1}`
	j := pgjsonb.New(data)
	if j.String() != expected {
		t.Error("failed to construct new jsonb")
	}
}

func TestBytesScan(t *testing.T) {
	data := `{}`
	var jsonb pgjsonb.JSONB
	err := jsonb.Scan([]byte(data))
	if err != nil {
		t.Error(err)
	}
	if jsonb.String() != data {
		t.Error("failed to scan []byte")
	}
}

func TestNullScan(t *testing.T) {
	var jsonb pgjsonb.JSONB
	err := jsonb.Scan(nil)
	if err != nil {
		t.Error(err)
	}
	if jsonb.String() != "null" {
		t.Error("failed to scan []byte")
	}
}

func TestJsonNullScan(t *testing.T) {
	data := `null`
	var jsonb pgjsonb.JSONB
	err := jsonb.Scan([]byte(data))
	if err != nil {
		t.Error(err)
	}
	if jsonb.String() != data {
		t.Error("failed to scan []byte")
	}
}

func TestJsonNumberScan(t *testing.T) {
	data := `1.5`
	var jsonb pgjsonb.JSONB
	err := jsonb.Scan([]byte(data))
	if err != nil {
		t.Error(err)
	}
	if jsonb.String() != data {
		t.Error("failed to scan []byte")
	}
}

func TestJsonStringScan(t *testing.T) {
	data := `"data"`
	var jsonb pgjsonb.JSONB
	err := jsonb.Scan([]byte(data))
	if err != nil {
		t.Error(err)
	}
	if jsonb.String() != data {
		t.Error("failed to scan []byte")
	}
}

func TestJsonArrayScan(t *testing.T) {
	data := `[1,"2"]`
	var jsonb pgjsonb.JSONB
	err := jsonb.Scan([]byte(data))
	if err != nil {
		t.Error(err)
	}
	if jsonb.String() != data {
		t.Error("failed to scan []byte")
	}
}

func TestJsonObjectScan(t *testing.T) {
	data := `{"a":1,"b":[2]}`
	var jsonb pgjsonb.JSONB
	err := jsonb.Scan([]byte(data))
	if err != nil {
		t.Error(err)
	}
	if jsonb.String() != data {
		t.Error("failed to scan []byte")
	}
}
