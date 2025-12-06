package qrcode

import (
	"github.com/ahmadnaufalhakim/qrgen/internal/encoder"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrcode/matrix"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrencode"
)

type QRBuilder struct {
	text       string
	encMode    *qrconst.EncodingMode
	minVersion int
	ecLevel    qrconst.ErrorCorrectionLevel
	maskNum    *int
}

func NewQRBuilder(text string) *QRBuilder {
	return &QRBuilder{
		text:       text,
		encMode:    nil,
		minVersion: 1,
		ecLevel:    qrconst.M,
		maskNum:    nil,
	}
}

func (b *QRBuilder) WithEncodingMode(
	encMode qrconst.EncodingMode,
) *QRBuilder {
	b.encMode = &encMode
	return b
}

func (b *QRBuilder) WithMinVersion(minVersion int) *QRBuilder {
	b.minVersion = minVersion
	return b
}

func (b *QRBuilder) WithErrorCorrectionLevel(
	ecLevel qrconst.ErrorCorrectionLevel,
) *QRBuilder {
	b.ecLevel = ecLevel
	return b
}

func (b *QRBuilder) WithMaskNum(maskNum int) *QRBuilder {
	b.maskNum = &maskNum
	return b
}

func (b *QRBuilder) Build() (*QRCode, error) {
	// 1. Determine encoding mode
	encoder, err := encoder.NewEncoder(b.text, b.encMode)
	if err != nil {
		return nil, err
	}

	// 2. Encode input string
	dataBits, err := encoder.Encode()
	if err != nil {
		return nil, err
	}

	// 3. Determine the QR Code version
	version, err := qrencode.DetermineVersion(
		encoder.Mode(),
		b.minVersion,
		b.ecLevel,
		encoder.CharCount(),
	)
	if err != nil {
		return nil, err
	}

	// 4. Construct the bit strings from the mode indicator,
	// char count indicator, and the actual data bits
	bitStrings := append(
		[]string{
			qrencode.ModeIndicator(encoder.Mode()),
			qrencode.CharCountIndicator(
				encoder.Mode(),
				version,
				encoder.CharCount(),
			),
		},
		dataBits...,
	)

	// 5. Assemble data codewords using the bit strings
	dataCodewords, err := qrencode.AssembleDataCodewords(
		version,
		b.ecLevel,
		bitStrings,
	)
	if err != nil {
		return nil, err
	}

	// 6. Assemble data blocks using the previously
	// assembled data codewords
	dataBlocks, err := qrencode.AssembleDataBlocks(
		version,
		b.ecLevel,
		dataCodewords,
	)
	if err != nil {
		return nil, err
	}

	// 7. Generate the error correction blocks for each data block
	ecBlocks, err := qrencode.GenerateErrorCorrectionBlocks(
		version,
		b.ecLevel,
		dataBlocks,
	)
	if err != nil {
		return nil, err
	}

	// 8. Interleave blocks
	messageBitString, err := qrencode.InterleaveBlocks(
		version,
		b.ecLevel,
		dataBlocks,
		ecBlocks,
	)
	if err != nil {
		return nil, err
	}

	// 9. Construct the QR Code object
	qrCode := NewQRCode(
		version,
		b.ecLevel,
		messageBitString,
	)

	// 10. Place modules in the QR Code matrix
	err = b.placeModules(qrCode)
	if err != nil {
		return nil, err
	}

	return qrCode, nil
}

func (b *QRBuilder) placeModules(qr *QRCode) error {
	// Determine the mask pattern
	if b.maskNum != nil {
		qr.MaskNum = *b.maskNum
	} else {
		qr.MaskNum = matrix.DetermineBestMaskNum(
			qr.ECLevel,
			qr.Modules,
			qr.Patterns,
		)
	}

	// Place modules and function patterns
	matrix.PlaceFinderPatterns(
		qr.Modules,
		qr.Patterns,
	)
	matrix.PlaceSeparators(
		qr.Modules,
		qr.Patterns,
	)
	matrix.PlaceAlignmentPattern(
		qr.Modules,
		qr.Patterns,
	)
	matrix.PlaceTimingPattern(
		qr.Modules,
		qr.Patterns,
	)
	matrix.PlaceDarkModule(
		qr.Modules,
		qr.Patterns,
	)
	matrix.PlaceVersionInformation(
		qr.Modules,
		qr.Patterns,
	)
	matrix.ReserveFormatInformationArea(
		qr.Patterns,
	)
	matrix.PlaceFormatInformation(
		qr.ECLevel,
		qr.Modules,
		qr.Patterns,
		qr.MaskNum,
	)
	matrix.PlaceMessageBits(
		qr.MessageBits,
		qr.Modules,
		qr.Patterns,
	)
	matrix.ApplyMaskPattern(
		qr.MaskNum,
		qr.Modules,
		qr.Patterns,
	)

	return nil
}
