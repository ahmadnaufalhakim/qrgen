package qrencode

import (
	"fmt"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
)

func GroupDataCodewords(
	ecLevel qrconst.ErrorCorrectionLevel,
	version int,
	dataCodewords []string,
) ([][][]string, error) {
	ecBlock := tables.ECBlocks[ecLevel][version-1]

	groups := make([][][]string, 0, 2)

	// Fill group 1
	group1 := make([][]string, ecBlock.Group1Blocks)
	start := 0
	for block := range ecBlock.Group1Blocks {
		end := start + ecBlock.Group1DataCodewords
		group1[block] = append([]string{}, dataCodewords[start:end]...)
		start = end
	}
	groups = append(groups, group1)

	// Fill group 2 (if present)
	if ecBlock.Group2Blocks > 0 {
		group2 := make([][]string, ecBlock.Group2Blocks)
		for block := range ecBlock.Group2Blocks {
			end := start + ecBlock.Group2DataCodewords
			group2[block] = append([]string{}, dataCodewords[start:end]...)
			start = end
		}
		groups = append(groups, group2)
	}

	return groups, nil
}

func GeneratorPolynomial(n int) ([]uint8, error) {
	if n <= 0 {
		return nil, fmt.Errorf("n must be positive")
	}

	if g, ok := tables.GeneratorPolynomial[n]; ok {
		return g, nil
	}

	var polynomials [][]uint8
	for i := n; i >= 0; i-- {
		if g, ok := tables.GeneratorPolynomial[i]; ok {
			polynomials = append(polynomials, g)
			break
		} else {
			polynomials = append(polynomials, []uint8{1, tables.AntilogGF256[i]})
		}
	}

	prod, err := PolynomialsProduct(polynomials...)
	if err != nil {
		return nil, err
	}

	tables.GeneratorPolynomial[n] = prod

	return prod, nil
}

func PolynomialsProduct(polynomials ...[]uint8) ([]uint8, error) {
	if len(polynomials) == 0 {
		return nil, fmt.Errorf("must provide at least 1 polynomial")
	}

	prod := []uint8{1}
	for _, p := range polynomials {
		prod = multiplyTwoPolynomials(prod, p)
	}

	return prod, nil
}

func multiplyTwoPolynomials(a, b []uint8) []uint8 {
	m := len(a)
	n := len(b)

	result := make([]uint8, m+n-1)

	for i := range m {
		for j := range n {
			result[i+j] = addGF256(result[i+j], mulGF256(a[i], b[j]))
		}
	}

	return result
}

func addGF256(x, y uint8) uint8 {
	return x ^ y
}

func mulGF256(x, y uint8) uint8 {
	if x == 0 || y == 0 {
		return 0
	}

	logX := uint16(tables.LogGF256[x])
	logY := uint16(tables.LogGF256[y])

	return tables.AntilogGF256[(logX+logY)%255]
}
