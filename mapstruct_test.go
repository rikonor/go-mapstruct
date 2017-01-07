package mapstruct

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	in := map[string]interface{}{
		"name": "Johnny",
		"age":  127,
	}

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	expectedOut := Person{
		Name: "Johnny",
		Age:  127,
	}

	var out Person
	if err := Decode(in, &out); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(out, expectedOut) {
		t.Fatalf("failed to convert correctly. got %v but expected %v", out, expectedOut)
	}
}
