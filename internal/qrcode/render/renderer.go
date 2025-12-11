package render

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrcode"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
)

type QRRenderer struct {
	moduleShape     qrconst.ModuleShape
	backgroundColor color.RGBA
	foregroundColor color.RGBA
	kernelType      qrconst.KernelType
	kernelFunc      func(radius int) []float64
	radius          int
}

func NewRenderer() *QRRenderer {
	return &QRRenderer{
		moduleShape:     qrconst.Square,
		backgroundColor: color.RGBA{255, 255, 255, 255},
		foregroundColor: color.RGBA{0, 0, 0, 255},
		kernelType:      qrconst.KernelLanczos2,
		kernelFunc:      Lanczos2Kernel,
		radius:          3,
	}
}

func (r *QRRenderer) WithModuleShape(
	moduleShape qrconst.ModuleShape,
) *QRRenderer {
	r.moduleShape = moduleShape
	return r
}

func (r *QRRenderer) WithBackgroundColor(
	backgroundColor color.RGBA,
) *QRRenderer {
	r.backgroundColor = backgroundColor
	return r
}

func (r *QRRenderer) WithForegroundColor(
	foregroundColor color.RGBA,
) *QRRenderer {
	r.foregroundColor = foregroundColor
	return r
}

func (r *QRRenderer) WithKernelType(
	kernelType qrconst.KernelType,
) *QRRenderer {
	r.kernelType = kernelType

	kernelCfg := Kernels[r.kernelType]
	r.kernelFunc = kernelCfg.KernelFunc
	r.radius = kernelCfg.DefaultRadius
	return r
}

func (r *QRRenderer) WithRadius(
	radius int,
) *QRRenderer {
	r.radius = radius
	return r
}

func (r *QRRenderer) RenderPNG(
	qr qrcode.QRCode,
	filename string,
) error {
	version := qr.Version

	// Set scale based on the QR Code version
	var scale int
	switch {
	case version >= 30:
		scale = 15
	case version >= 20:
		scale = 17
	case version >= 10:
		scale = 19
	default:
		scale = 21
	}
	margin := scale

	// Prepare the image matrix
	imgSize := qr.Size*scale + 2*margin
	img := image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))

	// Background and module colors
	bg := r.backgroundColor
	fg := r.foregroundColor

	// Fill margins
	for x := range qr.Size + 1 {
		for dy := range margin {
			for dx := range margin {
				img.Set(x*margin+dx, dy, bg)
				img.Set(imgSize-1-x*margin-dx, imgSize-1-dy, bg)
			}
		}
	}
	for y := range qr.Size + 1 {
		for dx := range margin {
			for dy := range margin {
				img.Set(imgSize-margin+dx, y*margin+dy, bg)
				img.Set(margin-1-dx, imgSize-1-y*margin-dy, bg)
			}
		}
	}

	// Fill modules
	for y := range qr.Size {
		for x := range qr.Size {
			lookahead := buildLookahead(qr, x, y)
			startX, startY := x*scale+margin, y*scale+margin
			module := qr.Modules[y][x]

			if module {
				for dy := range scale {
					for dx := range scale {
						if tables.ModuleRenderFunctions[r.moduleShape](dx, dy, scale, lookahead) {
							img.Set(startX+dx, startY+dy, fg)
						} else {
							img.Set(startX+dx, startY+dy, bg)
						}
					}
				}
			} else {
				for dy := range scale {
					for dx := range scale {
						if tables.ModuleMergeFunctions[r.moduleShape](dx, dy, scale, lookahead) {
							img.Set(startX+dx, startY+dy, fg)
						} else {
							img.Set(startX+dx, startY+dy, bg)
						}
					}
				}
			}
		}
	}

	// Apply blurring kernel
	blurHorizontal(img, r.kernelFunc(r.radius))
	blurVertical(img, r.kernelFunc(r.radius))

	// Create output file
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Encode PNG
	return png.Encode(f, img)
}

func buildLookahead(qr qrcode.QRCode, x, y int) qrconst.Lookahead {
	size := qr.Size
	lookahead := qrconst.Lookahead(0)

	if x < size-1 && qr.Modules[y][x+1] {
		lookahead |= qrconst.LookR
	}
	if x < size-1 && y > 0 && qr.Modules[y-1][x+1] {
		lookahead |= qrconst.LookUR
	}
	if y > 0 && qr.Modules[y-1][x] {
		lookahead |= qrconst.LookU
	}
	if x > 0 && y > 0 && qr.Modules[y-1][x-1] {
		lookahead |= qrconst.LookUL
	}
	if x > 0 && qr.Modules[y][x-1] {
		lookahead |= qrconst.LookL
	}
	if x > 0 && y < size-1 && qr.Modules[y+1][x-1] {
		lookahead |= qrconst.LookDL
	}
	if y < size-1 && qr.Modules[y+1][x] {
		lookahead |= qrconst.LookD
	}
	if x < size-1 && y < size-1 && qr.Modules[y+1][x+1] {
		lookahead |= qrconst.LookDR
	}

	if qr.Patterns[y][x].IsFinder() {
		lookahead |= qrconst.LookFinder
	}
	if qr.Patterns[y][x].IsAlignment() {
		lookahead |= qrconst.LookAlignment
	}
	if qr.Patterns[y][x].IsSeparator() {
		lookahead |= qrconst.LookSeparator
	}

	return lookahead
}

func blurHorizontal(img *image.RGBA, kernel []float64) {
	bounds := img.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	r := len(kernel) / 2

	tmp := make([]color.RGBA, w)

	for y := range h {
		// copy row
		for x := range w {
			tmp[x] = img.RGBAAt(x, y)
		}

		for x := range w {
			var rSum, gSum, bSum, aSum float64

			for k := -r; k <= r; k++ {
				xx := x + k
				if xx < 0 {
					xx = 0
				} else if xx >= w {
					xx = w - 1
				}

				color := tmp[xx]
				weight := kernel[k+r]

				rSum += weight * float64(color.R)
				gSum += weight * float64(color.G)
				bSum += weight * float64(color.B)
				aSum += weight * float64(color.A)
			}

			img.SetRGBA(x, y, color.RGBA{
				uint8(rSum),
				uint8(gSum),
				uint8(bSum),
				uint8(aSum),
			})
		}
	}
}

func blurVertical(img *image.RGBA, kernel []float64) {
	bounds := img.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	r := len(kernel) / 2

	tmp := make([]color.RGBA, h)

	for x := range w {
		// copy row
		for y := range h {
			tmp[y] = img.RGBAAt(x, y)
		}

		for y := range h {
			var rSum, gSum, bSum, aSum float64

			for k := -r; k <= r; k++ {
				yy := y + k
				if yy < 0 {
					yy = 0
				} else if yy >= h {
					yy = h - 1
				}

				color := tmp[yy]
				weight := kernel[k+r]

				rSum += weight * float64(color.R)
				gSum += weight * float64(color.G)
				bSum += weight * float64(color.B)
				aSum += weight * float64(color.A)
			}

			img.SetRGBA(x, y, color.RGBA{
				uint8(rSum),
				uint8(gSum),
				uint8(bSum),
				uint8(aSum),
			})
		}
	}
}
