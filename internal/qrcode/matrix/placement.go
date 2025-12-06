package matrix

import (
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
)

func PlaceFinderPatterns(
	modules [][]bool,
	patterns [][]qrconst.FunctionPattern,
) {
	size := len(modules)

	startPos := [3][2]int{
		{0, 0},
		{0, size - 7},
		{size - 7, 0},
	}

	for _, pos := range startPos {
		for i := range 7 {
			for j := range 7 {
				isOuter := (i == 0 || i == 6) || (j == 0 || j == 6)
				isInner := (i >= 2 && i <= 4) && (j >= 2 && j <= 4)

				module := isOuter || isInner

				modules[pos[0]+i][pos[1]+j] = module
				patterns[pos[0]+i][pos[1]+j] = qrconst.FPFinder
			}
		}
	}
}

func PlaceSeparators(
	modules [][]bool,
	patterns [][]qrconst.FunctionPattern,
) {
	size := len(modules)

	startPos := [3][2]int{
		{0, 0},
		{0, size - 7},
		{size - 7, 0},
	}

	for i, j := 0, 0; i < 8 && j < 8; i, j = i+1, j+1 {
		// Top left
		modules[startPos[0][0]+i][startPos[0][1]+7] = false
		modules[startPos[0][0]+7][startPos[0][1]+j] = false
		patterns[startPos[0][0]+i][startPos[0][1]+7] = qrconst.FPSeparator
		patterns[startPos[0][0]+7][startPos[0][1]+j] = qrconst.FPSeparator

		// Top right
		modules[startPos[1][0]+i][startPos[1][1]-1] = false
		modules[startPos[1][0]+7][startPos[1][1]-1+j] = false
		patterns[startPos[1][0]+i][startPos[1][1]-1] = qrconst.FPSeparator
		patterns[startPos[1][0]+7][startPos[1][1]-1+j] = qrconst.FPSeparator

		// Bottom left
		modules[startPos[2][0]-1+i][startPos[2][1]+7] = false
		modules[startPos[2][0]-1][startPos[2][1]+j] = false
		patterns[startPos[2][0]-1+i][startPos[2][1]+7] = qrconst.FPSeparator
		patterns[startPos[2][0]-1][startPos[2][1]+j] = qrconst.FPSeparator
	}
}

func PlaceAlignmentPattern(
	modules [][]bool,
	patterns [][]qrconst.FunctionPattern,
) {
	size := len(modules)
	version := (size - 17) / 4

	alignmentRowCols := tables.AlignmentPatternLocations[version-1]
	alignmentCenterPos := make([][2]int, len(alignmentRowCols)*len(alignmentRowCols))
	for i, r := range alignmentRowCols {
		for j, c := range alignmentRowCols {
			alignmentCenterPos[i*len(alignmentRowCols)+j] = [2]int{r, c}
		}
	}

	for _, centerPos := range alignmentCenterPos {
		if patterns[centerPos[0]][centerPos[1]] != qrconst.FPUnoccupied {
			continue
		}

		for i := range 5 {
			for j := range 5 {
				isOuter := (i == 0 || i == 4) || (j == 0 || j == 4)
				isInner := i == 2 && j == 2

				module := isOuter || isInner

				modules[centerPos[0]-2+i][centerPos[1]-2+j] = module
				patterns[centerPos[0]-2+i][centerPos[1]-2+j] = qrconst.FPAlignment
			}
		}
	}
}

func PlaceTimingPattern(
	modules [][]bool,
	patterns [][]qrconst.FunctionPattern,
) {
	size := len(modules)

	startPos := [3][2]int{
		{0, 0},
		{0, size - 7},
		{size - 7, 0},
	}

	module := true
	// Place vertical timing pattern if not occupied by another function pattern
	for i := startPos[0][0] + 6; i < startPos[2][0]; i++ {
		if patterns[i][6].IsUnoccupied() {
			modules[i][6] = module
			patterns[i][6] = qrconst.FPTiming
		}
		module = !module
	}
	// Place horizontal timing pattern if not occupied by another function pattern
	for j := startPos[0][0] + 6; j < startPos[2][0]; j++ {
		if patterns[6][j].IsUnoccupied() {
			modules[6][j] = module
			patterns[6][j] = qrconst.FPTiming
		}
		module = !module
	}
}

func PlaceDarkModule(
	modules [][]bool,
	patterns [][]qrconst.FunctionPattern,
) {
	size := len(modules)
	version := (size - 17) / 4

	modules[4*version+9][8] = true
	patterns[4*version+9][8] = qrconst.FPDarkModule
}

func PlaceVersionInformation(
	modules [][]bool,
	patterns [][]qrconst.FunctionPattern,
) {
	size := len(modules)
	version := (size - 17) / 4
	if version < 7 {
		return
	}

	versionBitString := tables.VersionInfo[version-1]
	startPos := [3][2]int{
		{0, size - 7},
		{size - 7, 0},
	}

	var module bool

	for j := range 3 {
		for i := range 6 {
			module = versionBitString[len(versionBitString)-1-(i*3+j)] == '1'
			modules[startPos[0][0]+i][startPos[0][1]-4+j] = module
			patterns[startPos[0][0]+i][startPos[0][1]-4+j] = qrconst.FPVersionInfo
		}
	}

	for j := range 6 {
		for i := range 3 {
			module = versionBitString[len(versionBitString)-1-(j*3+i)] == '1'
			modules[startPos[1][0]-4+i][startPos[1][1]+j] = module
			patterns[startPos[1][0]-4+i][startPos[1][1]+j] = qrconst.FPVersionInfo
		}
	}
}

func PlaceMessageBits(
	messageBits string,
	modules [][]bool,
	patterns [][]qrconst.FunctionPattern,
) {
	size := len(modules)

	upward := true
	msgBitIdx := 0

	// Calculate both row and column based on given index,
	// upward condition, and column.
	row := func(idx int, upward bool) int {
		if upward {
			return size - 1 - idx/2
		} else {
			return idx / 2
		}
	}
	col := func(idx, j int) int {
		return j - idx%2
	}

	for j := size - 1; j > 0; j -= 2 {
		// Check if current column interferes with vertical timing pattern
		// If so, shift the current column once to the left
		if j == 6 {
			j--
		}

		for idx := range 2 * size {
			if patterns[row(idx, upward)][col(idx, j)].IsUnoccupied() {
				modules[row(idx, upward)][col(idx, j)] = messageBits[msgBitIdx] == '1'
				patterns[row(idx, upward)][col(idx, j)] = qrconst.FPMessageBit
				msgBitIdx++
			}
		}

		upward = !upward
	}
}

func ReserveFormatInformationArea(patterns [][]qrconst.FunctionPattern) {
	size := len(patterns)

	startPos := [3][2]int{
		{0, 0},
		{0, size - 7},
		{size - 7, 0},
	}

	for i, j := 0, 0; i < 9 && j < 9; i, j = i+1, j+1 {
		// Reserve two-module strip near the top-left finder pattern
		if patterns[startPos[0][0]+i][startPos[0][1]+8].IsUnoccupied() {
			patterns[startPos[0][0]+i][startPos[0][1]+8] = qrconst.FPFormatInfo
		}
		if patterns[startPos[0][0]+8][startPos[0][1]+j].IsUnoccupied() {
			patterns[startPos[0][0]+8][startPos[0][1]+j] = qrconst.FPFormatInfo
		}

		// Reserve one-module strip near the top-right finder pattern
		if patterns[startPos[1][0]+8][min(startPos[1][1]-1+j, size-1)].IsUnoccupied() {
			patterns[startPos[1][0]+8][min(startPos[1][1]-1+j, size-1)] = qrconst.FPFormatInfo
		}

		// Reserve one-module strip near the bottom-left finder pattern
		if patterns[min(startPos[2][0]-1+i, size-1)][startPos[2][1]+8].IsUnoccupied() {
			patterns[min(startPos[2][0]-1+i, size-1)][startPos[2][1]+8] = qrconst.FPFormatInfo
		}
	}
}

func PlaceFormatInformation(
	ecLevel qrconst.ErrorCorrectionLevel,
	modules [][]bool,
	patterns [][]qrconst.FunctionPattern,
	maskNumber int,
) {
	size := len(modules)
	formatBitString := tables.FormatInfo[ecLevel][maskNumber]

	i, j := size-1, 0
	for idx := range len(formatBitString) {
		if patterns[i][8].IsTiming() {
			i--
		} else if i == size-8 {
			i = 8
		}

		if patterns[8][j].IsTiming() {
			j++
		} else if j == 8 {
			j = size - 8
		}

		module := formatBitString[idx] == '1'
		if patterns[i][8].IsFormat() {
			modules[i][8] = module
			i--
		}
		if patterns[8][j].IsFormat() {
			modules[8][j] = module
			j++
		}
	}
}
