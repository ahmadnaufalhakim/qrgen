package main

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrcode"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrcode/render"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
)

var (
	BLACK     = color.RGBA{0, 0, 0, 255}
	WHITE     = color.RGBA{255, 255, 255, 255}
	GREEN     = color.RGBA{27, 165, 8, 255}
	BLUE      = color.RGBA{48, 129, 229, 255}
	PINK      = color.RGBA{224, 134, 200, 255}
	YELLOW    = color.RGBA{251, 208, 67, 255}
	POO_BROWN = color.RGBA{122, 89, 1, 255}
	SILVER    = color.RGBA{188, 198, 204, 255}
)

func main() {

	// text := "8675309" //Numeric
	// text := "HELLO WORLD" //Alphanumeric
	// text := "だから僕は音楽をやめた" //Kanji
	text := "Life moves pretty fast. If you don't stop and look around once in a while, you could miss it.\nFerris Bueller" //Byte

	qrBuilder := qrcode.NewQRBuilder(text)
	qrCode, err := qrBuilder.
		WithErrorCorrectionLevel(qrconst.L).
		Build()
	if err != nil {
		fmt.Println(err)
		return
	}

	bg := GREEN
	fg := YELLOW

	start := time.Now()
	qrRenderer := render.NewRenderer().
		WithModuleShape(qrconst.SmileyFace).
		WithBackgroundColor(bg).
		WithForegroundColor(fg)

	// PNG Rendering
	f, err := os.Create("main.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = qrRenderer.RenderToWriter(*qrCode, f, qrconst.RenderPNG)
	if err != nil {
		fmt.Println(err)
		return
	}

	// SVG Rendering
	f, err = os.Create("main.svg")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = qrRenderer.RenderSVG(*qrCode, f)
	if err != nil {
		fmt.Println(err)
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("Rendering took %s\n", elapsed)
}
