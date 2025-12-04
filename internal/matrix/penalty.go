package matrix

import "math"

func TotalPenalty(modules [][]bool) int {
	return PenaltyRunLength(modules) +
		PenaltyBlockPattern(modules) +
		PenaltyFinderPattern(modules) +
		PenaltyDarkAndLightModules(modules)
}

// First penalty rule gives the QR code a penalty for each group
// of five or more same-colored modules in a row and column.
func PenaltyRunLength(modules [][]bool) int {
	totalPenalty := 0

	// Calculate horizontal penalty
	for i := range len(modules) {
		sameConsecutiveModules := 0
		lastModule := !modules[i][0]
		for j := range len(modules[0]) {
			if lastModule != modules[i][j] {
				lastModule = modules[i][j]
				sameConsecutiveModules = 1
			} else {
				sameConsecutiveModules++
			}

			if sameConsecutiveModules == 5 {
				totalPenalty += 3
			} else if sameConsecutiveModules > 5 {
				totalPenalty++
			}
		}
	}

	// Calculate vertical penalty
	for j := range len(modules[0]) {
		sameConsecutiveModules := 0
		lastModule := !modules[0][j]
		for i := range len(modules) {
			if lastModule != modules[i][j] {
				lastModule = modules[i][j]
				sameConsecutiveModules = 1
			} else {
				sameConsecutiveModules++
			}

			if sameConsecutiveModules == 5 {
				totalPenalty += 3
			} else if sameConsecutiveModules > 5 {
				totalPenalty++
			}
		}
	}

	return totalPenalty
}

// Second penalty rule gives the QR code a penalty for each
// 2x2 area of same-colored modules in the matrix.
func PenaltyBlockPattern(modules [][]bool) int {
	totalPenalty := 0

	for i := 0; i < len(modules)-1; i++ {
		j := 0
		for j < len(modules[0])-1 {
			colSkip := 1

			if modules[i][j+1] == modules[i+1][j+1] &&
				modules[i][j] == modules[i][j+1] &&
				modules[i][j] == modules[i+1][j] {
				totalPenalty += 3
			} else if modules[i][j+1] != modules[i+1][j+1] {
				colSkip = 2
			}

			j += colSkip
		}
	}

	return totalPenalty
}

// Third penalty rule gives the QR code a large penalty
// if there are patterns that look similar to the finder pattern.
func PenaltyFinderPattern(modules [][]bool) int {
	totalPenalty := 0

	horizontalFixedModulesCheck := func(r, c int) bool {
		return (!modules[r][c+1] && !modules[r][c+5] && !modules[r][c+9]) &&
			(modules[r][c+4] && modules[r][c+6])
	}
	horizontalNonFixedModulesCheck := func(r, c int) bool {
		return (modules[r][c] == !modules[r][c+10]) &&
			(modules[r][c] == modules[r][c+2]) && (modules[r][c+10] == modules[r][c+8]) &&
			(modules[r][c+2] == modules[r][c+3]) && (modules[r][c+8] == modules[r][c+7])
	}

	verticalFixedModulesCheck := func(r, c int) bool {
		return (!modules[r+1][c] && !modules[r+5][c] && !modules[r+9][c]) &&
			(modules[r+4][c] && modules[r+6][c])
	}
	verticalNonFixedModulesCheck := func(r, c int) bool {
		return (modules[r][c] == !modules[r+10][c]) &&
			(modules[r][c] == modules[r+2][c]) && (modules[r+10][c] == modules[r+8][c]) &&
			(modules[r+2][c] == modules[r+3][c]) && (modules[r+8][c] == modules[r+7][c])
	}

	for i := range len(modules) {
		for j := range len(modules[0]) - 10 {
			if horizontalFixedModulesCheck(i, j) && horizontalNonFixedModulesCheck(i, j) {
				totalPenalty += 40
			}
		}
	}

	for j := range len(modules[0]) {
		for i := range len(modules) - 10 {
			if verticalFixedModulesCheck(i, j) && verticalNonFixedModulesCheck(i, j) {
				totalPenalty += 40
			}
		}
	}

	return totalPenalty
}

// Fourth penalty rule gives the QR code a penalty if more
// than half of the modules are dark or light. Larger penalty
// is given for a larger difference.
func PenaltyDarkAndLightModules(modules [][]bool) int {
	totalPenalty := 0

	darkModules := 0
	for i := range len(modules) {
		for j := range len(modules[0]) {
			if modules[i][j] {
				darkModules++
			}
		}
	}

	totalModules := len(modules) * len(modules[0])
	darkModulesPercentage := 100 * float64(darkModules) / float64(totalModules)
	prev := int(darkModulesPercentage) - int(darkModulesPercentage)%5
	next := prev + 5

	prev = int(math.Abs(float64(prev) - 50))
	next = int(math.Abs(float64(next) - 50))

	prev /= 5
	next /= 5

	totalPenalty += 10 * min(prev, next)

	return totalPenalty
}
