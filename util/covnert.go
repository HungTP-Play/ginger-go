package util

import (
	"log"
	"math"
)

// Convert RGB -> HSV
func Rgb2Hlv(red float64, green float64, blue float64) (int, int, int) {
	// normalize
	red = red / 255.0
	green = green / 255.0
	blue = blue / 255.0

	cMax := math.Max(math.Max(red, green), blue)
	cMin := math.Min(math.Min(red, green), blue)
	diff := cMax - cMin

	var h float64
	var s float64
	var v float64

	// Calculate Hue
	if cMax == 0 && cMin == 0 {
		h = 0
	}

	if cMax == red {
		h = 60*((green-blue)/diff) + 360
		h = float64(int(h) % 360)
	}

	if cMax == green {
		h = 60*((blue-red)/diff) + 120
		h = float64(int(h) % 360)
	}

	if cMax == blue {
		h = 60*((red-green)/diff) + 240
		h = float64(int(h) % 360)
	}

	// Calculate saturation
	if cMax == 0 {
		s = 0
	}

	s = (diff / cMax) * 100

	// Calculate value
	v = cMax * 100

	return int(math.Round(h)), int(math.Round(s)), int(math.Round(v))
}

// Convert HSV -> RGB
func Hsv2Rgb(h float64, s float64, v float64) (int, int, int) {
	// Normallize
	s = s / 100
	v = v / 100

	c := v * s

	h60m2 := math.Mod(h/60, 2)

	x := c * (1 - math.Abs(float64(h60m2)-1))

	m := v - c
	log.Printf("M: %v", m)

	var r int
	var g int
	var b int

	var r_ float64
	var g_ float64
	var b_ float64

	if h >= 0 && h < 60 {
		r_ = c
		g_ = x
		b_ = 0
	}

	if h >= 60 && h < 120 {
		r_ = x
		g_ = c
		b_ = 0
	}

	if h >= 120 && h < 180 {
		r_ = 0
		g_ = c
		b_ = x
	}

	if h >= 180 && h < 240 {
		r_ = 0
		g_ = x
		b_ = c
	}

	if h >= 240 && h < 300 {
		r_ = x
		g_ = 0
		b_ = c
	}

	if h >= 300 && h < 360 {
		r_ = c
		g_ = 0
		b_ = x
	}

	r = int(math.Round((r_ + m) * 255))
	g = int(math.Round((g_ + m) * 255))
	b = int(math.Round((b_ + m) * 255))

	log.Printf("r, g, b: %v, %v, %v", r, g, b)
	return r, g, b
}
