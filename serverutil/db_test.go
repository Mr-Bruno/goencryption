package serverutil

import (
	"reflect"
	"testing"
)

func TestAddElement(t *testing.T) {
	tests := []struct {
		key   string
		value []byte
		out   DB
	}{
		{"234", []byte("1st_payload"), DB{"234": []byte("1st_payload")}},
		{"123", []byte("2nd_payload"), DB{"234": []byte("1st_payload"), "123": []byte("2nd_payload")}},
	}

	database := CreateDatabase()
	for _, c := range tests {
		AddElement(database, c.key, c.value)
		eq := reflect.DeepEqual(database, c.out)
		if !eq {
			t.Error("Expected content", c.out, "and got", database)
		}
	}
}

func TestGetElement(t *testing.T) {

	database := CreateDatabase()
	got := GetElement(database, "123")
	if got != nil {
		t.Error("Expected empty, got", got)
	}

	id := "234"
	value := []byte("payload")
	AddElement(database, id, value)
	got = GetElement(database, id)

	eq := reflect.DeepEqual(got, value)
	if !eq {
		t.Error("Expected", value, " got", got)
	}
}
