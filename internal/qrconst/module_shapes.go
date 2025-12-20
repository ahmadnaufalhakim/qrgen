package qrconst

type ModuleShape int

const (
	Square ModuleShape = iota
	Circle
	TiedCircle
	HorizontalBlob
	VerticalBlob
	Blob
	LeftLeaf
	RightLeaf
	Diamond
	WaterDroplet
	Star4
	Star5
	Star6
	Star8
	Xs
	Octagon
	SmileyFace
	Pointillism
)

func (ms ModuleShape) String() string {
	switch ms {
	case Square:
		return "square"
	case Circle:
		return "circle"
	case TiedCircle:
		return "tiedCircle"
	case HorizontalBlob:
		return "horizontalBlob"
	case VerticalBlob:
		return "verticalBlob"
	case Blob:
		return "blob"
	case LeftLeaf:
		return "leftLeaf"
	case RightLeaf:
		return "rightLeaf"
	case Diamond:
		return "diamond"
	case WaterDroplet:
		return "waterDroplet"
	case Star4:
		return "star4"
	case Star5:
		return "star5"
	case Star6:
		return "star6"
	case Star8:
		return "star8"
	case Xs:
		return "xs"
	case Octagon:
		return "octagon"
	case SmileyFace:
		return "smileyFace"
	case Pointillism:
		return "pointillism"
	}
	return "unknown"
}
