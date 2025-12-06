package qrcode

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
)

type QRCode struct {
	Version     int
	ECLevel     qrconst.ErrorCorrectionLevel
	Size        int
	MessageBits string
	Modules     [][]bool
	Patterns    [][]qrconst.FunctionPattern
	MaskNum     int
}

func NewQRCode(
	version int,
	ecLevel qrconst.ErrorCorrectionLevel,
	messageBitString string,
) *QRCode {
	size := (version-1)*4 + 21
	messageBits := messageBitString
	modules := make([][]bool, size)
	patterns := make([][]qrconst.FunctionPattern, size)
	for i := range size {
		modules[i] = make([]bool, size)
		patterns[i] = make([]qrconst.FunctionPattern, size)
	}
	maskNum := 0

	return &QRCode{
		Version:     version,
		ECLevel:     ecLevel,
		Size:        size,
		MessageBits: messageBits,
		Modules:     modules,
		Patterns:    patterns,
		MaskNum:     maskNum,
	}
}

func (qr *QRCode) RenderPNG(filename string, scale int) error {
	if scale < 1 {
		scale = 1
	}

	imgSize := qr.Size * scale
	img := image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))

	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	for y := range qr.Size {
		for x := range qr.Size {
			var c color.RGBA
			if qr.Modules[y][x] {
				c = black
			} else {
				c = white
			}

			// Fill the scale*scale block
			startX, startY := x*scale, y*scale
			for dy := range scale {
				for dx := range scale {
					img.Set(startX+dx, startY+dy, c)
				}
			}
		}
	}

	// Create output file
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Encode PNG
	return png.Encode(f, img)
}
