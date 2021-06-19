package builder

import "testing"

func TestPickSpriteIndex(t *testing.T) {
	var first4Bytes = [16]byte{0, 0, 0, 123, 99}
	// 123 -> 01111011
	// 99  -> 01100011
	//
	// Corner: 2 bit from 3rd byte                            -> 01   -> 1
	// Side: next 4 bit from 3rd byte                         -> 1110 -> 14
	// Corner: last 2 bit from 3rd + firt 2 bit from 4th byte -> 1101 -> 13
	centerIndex, sideIndex, cornerIndex := PickSpriteIndex(first4Bytes)

	if centerIndex != 1 {
		t.Fatalf("Expect 1 receive %v", centerIndex)
	}

	if sideIndex != 14 {
		t.Fatalf("Expect 14 receive %v", sideIndex)
	}

	if cornerIndex != 13 {
		t.Fatalf("Expect 13 receive %v", cornerIndex)
	}
}

func TestPickStartIndices(t *testing.T) {
	var first4Bytes = [16]byte{0, 0, 0, 0, 99}
	// 123 -> 01111011
	// 99  -> 01100011
	//
	// Start Side Index: 2nd and 3rd bit of 4th byte          -> 10   -> 2
	// Start Corner Index: 4-5th bit of 4th byte	             -> 00   -> 0
	startSideIndex, startCornerIndex := PickStartIndices(first4Bytes)

	if startSideIndex != 2 {
		t.Fatalf("Expect 3 receive %v", startSideIndex)
	}

	if startCornerIndex != 0 {
		t.Fatalf("Expect 0 receive %v", startCornerIndex)
	}
}
