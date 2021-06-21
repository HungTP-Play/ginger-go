package util

import (
	"testing"
)

func TestRgb2Hsv(t *testing.T) {
	red := 50
	green := 168
	blue := 82

	h, s, v := Rgb2Hlv(float64(red), float64(green), float64(blue))

	if h != 136 {
		t.Fatalf("Expect 136, receive %v", int(h))
	}

	if s != 70 {
		t.Fatalf("Expect 70, receive %v", int(s))
	}

	if v != 66 {
		t.Fatalf("Expect 66 recevice %v", v)
	}
}

func TestHsv2Rgb(t *testing.T) {
	h := 100
	s := 70
	v := 66

	r, g, b := Hsv2Rgb(float64(h), float64(s), float64(v))

	if r != 90 {
		t.Fatalf("Expect 50, receive %v", r)
	}

	if g != 168 {
		t.Fatalf("Expect 168, receive %v", g)
	}

	if b != 50 {
		t.Fatalf("Expect 82, receive %v", b)
	}
}
