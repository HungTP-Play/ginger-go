package builder_test

import (
	"reflect"
	"testing"

	"github.com/HungTP-Play/ginger-go/module/builder"
)

var reflexArrTestTable = []struct {
	in  []byte
	out []byte
}{
	{[]byte{1, 2, 3}, []byte{1, 2, 1}},
	{[]byte{1, 2, 3, 4}, []byte{1, 2, 2, 1}},
	{[]byte{1, 2, 3, 4, 5}, []byte{1, 2, 3, 2, 1}},
	{[]byte{1, 2, 3, 4, 5, 6}, []byte{1, 2, 3, 3, 2, 1}},
	{[]byte{1, 2, 3, 4, 5, 6, 7}, []byte{1, 2, 3, 4, 3, 2, 1}},
	{[]byte{1, 2, 3, 4, 5, 6, 7, 8}, []byte{1, 2, 3, 4, 4, 3, 2, 1}},
	{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}, []byte{1, 2, 3, 4, 5, 4, 3, 2, 1}},
}

func TestReflexArr(t *testing.T) {
	t.Parallel()
	for _, val := range reflexArrTestTable {
		input := val.in
		actualOutput := builder.ReflexArr(input, len(input))
		if reflect.DeepEqual(val.out, actualOutput) == false {
			t.Fatalf("expect %v, actual %v", val.out, actualOutput)
		}
	}
}
