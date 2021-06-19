package builder

import "testing"

func TestPickColor(t *testing.T) {
	var first3Bytes = [16]byte{255, 255, 255}
	color := PickColor(first3Bytes)
	r, g, b, _ := color.RGBA()
	if r != 255 {
		t.Fatalf("Expect red 255, actual %v", r)
	}
	if g != 255 {
		t.Fatalf("Expect green 255, actual %v", g)
	}
	if b != 255 {
		t.Fatalf("Expect blue 255, actual %v", b)
	}
}
