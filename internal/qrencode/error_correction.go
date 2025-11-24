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

func MessagePolynomial(block []string) ([]uint8, error) {
	var m []uint8
	for _, dataCodeword := range block {
		b, err := bitStringToByte(dataCodeword)
		if err != nil {
			return nil, err
		}
		m = append(m, b)
	}

	return m, nil
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

func bitStringToByte(s string) (uint8, error) {
	if len(s) != 8 {
		return 0, fmt.Errorf("must be 8 bits")
	}

	b := uint8(0)
	for i := range 8 {
		b <<= 1
		if s[i] == '1' {
			b |= 1
		} else if s[i] != '0' {
			return 0, fmt.Errorf("invalid bit")
		}
	}

	return b, nil
}

func addTwoPolynomials(a, b []uint8) []uint8 {
	m := len(a)
	n := len(b)
	maxDegree := max(m, n)

	result := make([]uint8, maxDegree)

	for i := range maxDegree {
		x := uint8(0)
		if i < m {
			x = a[m-i-1]
		}

		y := uint8(0)
		if i < n {
			y = b[n-i-1]
		}

		result[maxDegree-i-1] = addGF256(x, y)
	}

	return result
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

func divideTwoPolynomials(a, b []uint8) []uint8 {
	m := len(a)
	n := len(b)

	a = append(a, make([]uint8, n-1)...)
	b = append(b, make([]uint8, m-1)...)

	dividend := a
	divisor := make([]uint8, len(b))
	copy(divisor, b)

	for i := range m {
		aLeadCoefExp := tables.LogGF256[dividend[0]]
		bLeadCoefExp := tables.LogGF256[b[0]]

		multiplierExp := 255 + uint16(aLeadCoefExp) - uint16(bLeadCoefExp)
		for j := range n {
			divisor[j] = mulGF256(divisor[j], tables.AntilogGF256[multiplierExp%255])
		}

		dividend = addTwoPolynomials(dividend, divisor)
		dividend = discardLeadingZeros(dividend)

		divisor = make([]uint8, len(b)-i-1)
		copy(divisor, b[:len(b)-i-1])
	}

	return dividend
}

func discardLeadingZeros(polynomial []uint8) []uint8 {
	for i, coef := range polynomial {
		if coef != 0 {
			return polynomial[i:]
		}
	}

	return []uint8{}
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
