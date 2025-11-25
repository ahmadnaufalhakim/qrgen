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
	ecBlockInfo := tables.ECBlockInfos[ecLevel][version-1]

	dataGroups := make([][][]string, 0, 2)

	// Fill group 1
	group1 := make([][]string, ecBlockInfo.Group1Blocks)
	start := 0
	for block := range ecBlockInfo.Group1Blocks {
		end := start + ecBlockInfo.Group1DataCodewordsPerBlock
		group1[block] = append([]string{}, dataCodewords[start:end]...)
		start = end
	}
	dataGroups = append(dataGroups, group1)

	// Fill group 2 (if present)
	if ecBlockInfo.Group2Blocks > 0 {
		group2 := make([][]string, ecBlockInfo.Group2Blocks)
		for block := range ecBlockInfo.Group2Blocks {
			end := start + ecBlockInfo.Group2DataCodewordsPerBlock
			group2[block] = append([]string{}, dataCodewords[start:end]...)
			start = end
		}
		dataGroups = append(dataGroups, group2)
	}

	return dataGroups, nil
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
	ecLevel qrconst.ErrorCorrectionLevel,
	version int,
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
