package partdrawer

import (
	"github.com/HungTP-Play/ginger-go/constant"
)

func GetPartDrawer(partType constant.PartType) IPartDrawer {
	switch partType {
	case constant.Type1:
		return &Type1Drawer{
			Name: "Type1",
		}
	case constant.Type2:
		return &Type2Drawer{
			Name: "Type2",
		}
	case constant.Type3:
		return &Type3Drawer{
			Name: "Type3",
		}
	case constant.Type4:
		return &Type4Drawer{
			Name: "Type4",
		}
	case constant.Type5:
		return &Type5Drawer{
			Name: "Type5",
		}
	case constant.Type6:
		return &Type6Drawer{
			Name: "Type6",
		}
	case constant.Type7:
		return &Type7Drawer{
			Name: "Type7",
		}
	case constant.Type8:
		return &Type8Drawer{
			Name: "Type8",
		}
	case constant.Type9:
		return &Type9Drawer{
			Name: "Type9",
		}
	case constant.Type10:
		return &Type10Drawer{
			Name: "Type10",
		}
	case constant.Type11:
		return &Type11Drawer{
			Name: "Type11",
		}
	case constant.Type12:
		return &Type12Drawer{
			Name: "Type12",
		}
	case constant.Type13:
		return &Type13Drawer{
			Name: "Type13",
		}
	case constant.Type14:
		return &Type14Drawer{
			Name: "Type14",
		}
	case constant.Type15:
		return &Type15Drawer{
			Name: "Type15",
		}
	case constant.Type16:
		return &Type16Drawer{
			Name: "Type16",
		}
	default:
		return &Type1Drawer{
			Name: "Type1",
		}
	}
}
