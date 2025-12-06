package qrencode

import (
	"fmt"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
)

// AssembleDataBlocks splits the data codewords into the exact block structure
// required by the QR Code specification for the given version and error
// correction level.
//
// A QR version may have:
// - Group 1 blocks, each containing N1 data codewords; and
// - Group 2 blocks, each containing N2 data codewords (usually N1 + 1)
//
// The function returns a flat slice of data blocks, where each block is
// a slice of data codewords (8-bit strings).
func AssembleDataBlocks(
	version int,
	ecLevel qrconst.ErrorCorrectionLevel,
	dataCodewords []string,
) ([][]string, error) {
	ecBlockInfo := tables.ECBlockInfos[ecLevel][version-1]
	group1Blocks := ecBlockInfo.Group1Blocks
	group2Blocks := ecBlockInfo.Group2Blocks
	n1 := ecBlockInfo.Group1DataCodewordsPerBlock
	n2 := ecBlockInfo.Group1DataCodewordsPerBlock

	dataBlocks := make([][]string, group1Blocks+group2Blocks)

	start := 0

	// Fill with group 1 blocks
	for i := range group1Blocks {
		end := start + n1
		if end > len(dataCodewords) {
			return nil, fmt.Errorf("insufficient data codewords for Group 1 blocks")
		}

		dataBlocks[i] = append([]string{}, dataCodewords[start:end]...)
		start = end
	}

	// Fill with group 2 blocks (if any)
	for i := range group2Blocks {
		end := start + n2
		if end > len(dataCodewords) {
			return nil, fmt.Errorf("insufficient data codewords for Group 2 blocks")
		}

		dataBlocks[group1Blocks+i] = append([]string{}, dataCodewords[start:end]...)
		start = end
	}

	// Safety check (should always match)
	if start != len(dataCodewords) {
		return nil, fmt.Errorf("data codeword count mismatch: leftover = %d", len(dataCodewords)-start)
	}

	return dataBlocks, nil
}

func MessagePolynomial(dataBlock []string) ([]uint8, error) {
	var m []uint8
	for _, dataCodeword := range dataBlock {
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
	for i := n; i > 0; i-- {
		if g, ok := tables.GeneratorPolynomial[i]; ok {
			polynomials = append(polynomials, g)
			break
		} else {
			polynomials = append(polynomials, []uint8{1, tables.AntilogGF256[i-1]})
		}
	}

	prod, err := polynomialsProduct(polynomials...)
	if err != nil {
		return nil, err
	}

	tables.GeneratorPolynomial[n] = prod

	return prod, nil
}

func GenerateErrorCorrectionBlocks(
	version int,
	ecLevel qrconst.ErrorCorrectionLevel,
	dataBlocks [][]string,
) ([][]uint8, error) {
	ecBlockInfo := tables.ECBlockInfos[ecLevel][version-1]
	n := ecBlockInfo.ECCodewordsPerBlock
	g, err := GeneratorPolynomial(n)
	if err != nil {
		return nil, err
	}

	ecBlocks := make([][]uint8, ecBlockInfo.Group1Blocks+ecBlockInfo.Group2Blocks)
	for i, dataBlock := range dataBlocks {
		m, err := MessagePolynomial(dataBlock)
		if err != nil {
			return nil, err
		}

		ecCodewords := divideTwoPolynomials(m, g)
		ecBlocks[i] = ecCodewords
	}

	return ecBlocks, nil
}
