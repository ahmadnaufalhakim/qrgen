package matrix

import (
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
)

func DetermineBestMaskNum(
	ecLevel qrconst.ErrorCorrectionLevel,
	modules [][]bool,
	patterns [][]qrconst.FunctionPattern,
) int {
	bestMaskNum := 0
	bestPenalty := int(^uint(0) >> 1)
	for maskNum := range tables.MaskPatterns {
		maskedModules := copyModules(modules)
		ApplyMaskPattern(maskNum, maskedModules, patterns)
		PlaceFormatInformation(ecLevel, maskedModules, patterns, maskNum)

		penalty := TotalPenalty(maskedModules)
		if penalty < bestPenalty {
			bestMaskNum = maskNum
			bestPenalty = penalty
		}
	}

	return bestMaskNum
}

func ApplyMaskPattern(
	maskNum int,
	modules [][]bool,
	patterns [][]qrconst.FunctionPattern,
) {
	maskPattern := tables.MaskPatterns[maskNum]

	for i := range modules {
		for j := range modules[i] {
			if patterns[i][j].IsMessage() && maskPattern(i, j) {
				modules[i][j] = !modules[i][j]
			}
		}
	}
}

func copyModules(src [][]bool) [][]bool {
	size := len(src)
	dst := make([][]bool, size)
	for i := range src {
		dst[i] = append([]bool(nil), src[i]...)
	}
	return dst
}
