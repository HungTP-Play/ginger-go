package constant

type PartType int

const (
	Type1 PartType = iota + 1
	Type2
	Type3
	Type4
	Type5
	Type6
	Type7
	Type8
	Type9
	Type10
	Type11
	Type12
	Type13
	Type14
	Type15
	Type16
)

// Symmetry part types
var CenterTypes = []PartType{Type1, Type5, Type9, Type16}

var AllTypes = []PartType{
	Type1, Type2, Type3, Type4,
	Type5, Type6, Type7, Type8,
	Type9, Type10, Type11, Type12,
	Type13, Type14, Type15, Type16,
}
