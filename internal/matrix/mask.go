package matrix

import (
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
)

func DetermineBestMaskPattern(
	modules [][]bool,
	patterns [][]qrconst.FunctionPattern,
	ecLevel qrconst.ErrorCorrectionLevel,
) func(r, c int) bool {
	bestMaskNumber := 0
	bestPenalty := int(^uint(0) >> 1)
	for maskNum, maskPattern := range tables.MaskPatterns {
		maskedModules := copyModules(modules)
		ApplyMaskPattern(maskedModules, patterns, maskPattern)
		PlaceFormatInformation(maskedModules, patterns, ecLevel, maskNum)

		penalty := TotalPenalty(maskedModules)
		if penalty < bestPenalty {
			bestMaskNumber = maskNum
			bestPenalty = penalty
		}
	}

	return tables.MaskPatterns[bestMaskNumber]
}

func ApplyMaskPattern(
	modules [][]bool,
	patterns [][]qrconst.FunctionPattern,
	maskPattern func(r, c int) bool,
) {
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
