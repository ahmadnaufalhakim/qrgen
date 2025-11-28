package qrencode

import (
	"fmt"

	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
)

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
		dividend = dividend[1:]

		divisor = make([]uint8, len(b)-i-1)
		copy(divisor, b[:len(b)-i-1])
	}

	return dividend
}

func polynomialsProduct(polynomials ...[]uint8) ([]uint8, error) {
	if len(polynomials) == 0 {
		return nil, fmt.Errorf("must provide at least 1 polynomial")
	}

	prod := []uint8{1}
	for _, p := range polynomials {
		prod = multiplyTwoPolynomials(prod, p)
	}

	return prod, nil
}

func discardLeadingZeros(polynomial []uint8) []uint8 {
	for i, coef := range polynomial {
		if coef != 0 {
			return polynomial[i:]
		}
	}

	return []uint8{}
}
