package tables

import "github.com/ahmadnaufalhakim/qrgen/internal/qrconst"

var CharacterCountIndicatorBits = map[qrconst.EncodingMode][3]int{
	qrconst.NumericMode:      {10, 12, 14},
	qrconst.AlphanumericMode: {9, 11, 13},
	qrconst.ByteMode:         {8, 16, 16},
	qrconst.KanjiMode:        {8, 10, 12},
}
