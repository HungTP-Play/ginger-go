package drawer

import "testing"

func TestCalRotation(t *testing.T) {
	var tableMap = map[int]int{
		0: 1,
		1: 2,
		2: 3,
		3: 0,
	}
	startFrom := 3

	for k, v := range tableMap {
		rotation := CalRotation(k, startFrom)
		if rotation != v {
			t.Fatalf("Expect %v, receive %v", v, rotation)
		}
	}
}
