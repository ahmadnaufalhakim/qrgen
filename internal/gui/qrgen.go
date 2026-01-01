package gui

import (
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"log"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrcode"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrcode/render"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
)

type qrBuildState struct {
	text       string
	minVersion int
	ecLevel    qrconst.ErrorCorrectionLevel
	maskNum    *int
}

type QRGeneratorApp struct {
	app        fyne.App
	window     fyne.Window
	currentTab string

	// Input fields
	plainTextEntry       *widget.Entry
	urlEntry             *widget.Entry
	twitterEntry         *widget.Entry
	emailEntry           *widget.Entry
	wifiSSIDEntry        *widget.Entry
	wifiPasswordEntry    *widget.Entry
	wifiEncryptionSelect *widget.Select
	telephoneEntry       *widget.Entry

	// QR Code options
	ecLevelRadio      *widget.RadioGroup
	minVersionEntry   *widget.Entry
	maskPatternSelect *widget.Select

	// Render options
	moduleShapeSelect  *widget.Select
	backgroundColorBtn *widget.Button
	foregroundColorBtn *widget.Button
	kernelTypeSelect   *widget.Select
	radiusSlider       *widget.Slider
	radiusLabel        *widget.Label

	// Preview
	previewContainer *fyne.Container
	previewImage     *canvas.Image

	// Data
	currentQRCode  *qrcode.QRCode
	lastBuildState *qrBuildState
	renderer       *render.QRRenderer
	updateTimer    *time.Timer

	dirtyQR     bool // encoding/matrix dirty
	dirtyRender bool // visual dirty

	qrIcon    *fyne.StaticResource
	paintIcon *fyne.StaticResource
}

//go:embed qr_icon.png
var qrIconBytes []byte

//go:embed paint_icon.png
var paintIconBytes []byte

func NewQRGeneratorApp() *QRGeneratorApp {
	app := app.NewWithID("qrgen")
	window := app.NewWindow("QRGen")

	qrIconResource := &fyne.StaticResource{
		StaticName:    "qr_icon.png",
		StaticContent: qrIconBytes,
	}
	paintIconResource := &fyne.StaticResource{
		StaticName:    "paint_icon.png",
		StaticContent: paintIconBytes,
	}

	qrGenApp := &QRGeneratorApp{
		app:         app,
		window:      window,
		currentTab:  "plaintext",
		renderer:    render.NewRenderer(),
		updateTimer: time.NewTimer(500 * time.Millisecond),
		dirtyQR:     true,
		dirtyRender: true,
		qrIcon:      qrIconResource,
		paintIcon:   paintIconResource,
	}
	qrGenApp.updateTimer.Stop()

	return qrGenApp
}

func (a *QRGeneratorApp) Run() {
	a.buildUI()
	a.window.ShowAndRun()
}

func (a *QRGeneratorApp) buildUI() {
	// Create tab container for different input types
	tabs := container.NewAppTabs(
		container.NewTabItem("Plain Text", a.buildPlainTextTab()),
		container.NewTabItem("URL", a.buildURLTab()),
		container.NewTabItem("Twitter", a.buildTwitterTab()),
		container.NewTabItem("Email", a.buildEmailTab()),
		container.NewTabItem("WiFi", a.buildWifiTab()),
		container.NewTabItem("Telephone", a.buildTelephoneTab()),
	)

	tabs.OnSelected = func(ti *container.TabItem) {
		a.currentTab = strings.ToLower(strings.ReplaceAll(ti.Text, " ", ""))
		a.markQRDirty()
	}

	// Basic options
	basicOptionsPanel := a.buildBasicOptionsPanel()

	// Advanced options
	advancedOptionsPanel := a.buildAdvancedOptionsPanel()

	// Preview panel
	previewPanel := a.buildPreviewPanel()

	// Save button
	saveButton := widget.NewButton("Save QR Code", a.showFileSaveDialog)

	// Layout
	leftPanel := container.NewVBox(
		tabs,
		widget.NewSeparator(),
		container.NewPadded(basicOptionsPanel),
		widget.NewSeparator(),
		container.NewPadded(advancedOptionsPanel),
		saveButton,
	)

	mainContent := container.NewHSplit(
		container.NewVScroll(container.NewPadded(leftPanel)),
		previewPanel,
	)
	mainContent.SetOffset(.4)

	a.window.SetContent(mainContent)
	a.window.Resize(fyne.NewSize(1200, 900))
}

func (a *QRGeneratorApp) buildPlainTextTab() fyne.CanvasObject {
	a.plainTextEntry = widget.NewMultiLineEntry()
	a.plainTextEntry.SetPlaceHolder("Enter text to encode ..")
	a.plainTextEntry.OnChanged = func(_ string) {
		a.markQRDirty()
	}

	return container.NewPadded(a.plainTextEntry)
}

func (a *QRGeneratorApp) buildURLTab() fyne.CanvasObject {
	a.urlEntry = widget.NewEntry()
	a.urlEntry.SetPlaceHolder("https://example.com")
	a.urlEntry.OnChanged = func(_ string) {
		a.markQRDirty()
	}

	return container.NewPadded(a.urlEntry)
}

func (a *QRGeneratorApp) buildTwitterTab() fyne.CanvasObject {
	a.twitterEntry = widget.NewEntry()
	a.twitterEntry.SetPlaceHolder("@username or twitter.com/username")
	a.twitterEntry.OnChanged = func(_ string) {
		a.markQRDirty()
	}

	return container.NewPadded(a.twitterEntry)
}

func (a *QRGeneratorApp) buildEmailTab() fyne.CanvasObject {
	a.emailEntry = widget.NewEntry()
	a.emailEntry.SetPlaceHolder("email@example.com")
	a.emailEntry.OnChanged = func(_ string) {
		a.markQRDirty()
	}

	return container.NewPadded(a.emailEntry)
}

func (a *QRGeneratorApp) buildWifiTab() fyne.CanvasObject {
	a.wifiSSIDEntry = widget.NewEntry()
	a.wifiSSIDEntry.SetPlaceHolder("Network SSID")
	a.wifiSSIDEntry.OnChanged = func(_ string) {
		a.markQRDirty()
	}

	a.wifiPasswordEntry = widget.NewPasswordEntry()
	a.wifiPasswordEntry.SetPlaceHolder("Password (optional)")
	a.wifiPasswordEntry.OnChanged = func(_ string) {
		a.markQRDirty()
	}

	a.wifiEncryptionSelect = widget.NewSelect(
		[]string{
			"WPA",
			"WEP",
			"None",
		},
		func(s string) {
			a.markQRDirty()
		},
	)
	a.wifiEncryptionSelect.SetSelected("WPA")

	form := widget.NewForm(
		widget.NewFormItem("SSID", a.wifiSSIDEntry),
		widget.NewFormItem("Password", a.wifiPasswordEntry),
		widget.NewFormItem("Encryption", a.wifiEncryptionSelect),
	)

	return container.NewPadded(form)
}

func (a *QRGeneratorApp) buildTelephoneTab() fyne.CanvasObject {
	a.telephoneEntry = widget.NewEntry()
	a.telephoneEntry.SetPlaceHolder("+1-234-567-8900")
	a.telephoneEntry.OnChanged = func(_ string) {
		a.markQRDirty()
	}

	return container.NewPadded(a.telephoneEntry)
}

func (a *QRGeneratorApp) buildBasicOptionsPanel() fyne.CanvasObject {
	// Error correction level
	ecLevels := []string{
		"L (Low - 7%)",
		"M (Medium - 15%)",
		"Q (Quartile - 25%)",
		"H (High - 30%)",
	}
	a.ecLevelRadio = widget.NewRadioGroup(
		ecLevels,
		func(s string) {
			a.markQRDirty()
		},
	)
	a.ecLevelRadio.Required = true
	a.ecLevelRadio.SetSelected("M (Medium - 15%)")

	// Minimum version
	minVersion := binding.NewInt()
	_ = minVersion.Set(1)

	a.minVersionEntry = widget.NewEntryWithData(
		binding.IntToString(minVersion),
	)
	a.minVersionEntry.SetPlaceHolder("1-40")
	a.minVersionEntry.OnChanged = validateMinVersionEntry(
		minVersion,
		a.minVersionEntry,
	)

	minVersionSlider := widget.NewSliderWithData(
		1, 40, binding.IntToFloat(minVersion),
	)
	minVersionSlider.Step = 1

	minVersion.AddListener(binding.NewDataListener(func() {
		a.markQRDirty()
	}))

	// Module shape
	moduleShapes := []string{
		"Square",
		"Circle",
		"HorizontalBlob",
		"VerticalBlob",
		"Blob",
		"LeftLeaf",
		"RightLeaf",
		"Diamond",
		"WaterDroplet",
		"Star4",
		"Star5",
		"Star6",
		"Xs",
		"Octagon",
		"SmileyFace",
		"Pointillism",
	}
	a.moduleShapeSelect = widget.NewSelect(
		moduleShapes,
		func(s string) {
			a.markRenderDirty()
		},
	)
	a.moduleShapeSelect.SetSelected("Square")

	// Colors
	a.backgroundColorBtn = widget.NewButton(
		"Background: White",
		func() { a.chooseColor("background") },
	)
	a.foregroundColorBtn = widget.NewButton(
		"Foreground: Black",
		func() { a.chooseColor("foreground") },
	)

	// Basic Options form
	form := widget.NewForm(
		widget.NewFormItem("Error Correction", a.ecLevelRadio),
		widget.NewFormItem("Min. version (1-40)", container.NewVBox(
			a.minVersionEntry,
			minVersionSlider,
		)),
		widget.NewFormItem("Module Shape", a.moduleShapeSelect),
		widget.NewFormItem("Background Color", a.backgroundColorBtn),
		widget.NewFormItem("Foreground Color", a.foregroundColorBtn),
	)

	return widget.NewCard("Basic Options", "", container.NewPadded(form))
}

func (a *QRGeneratorApp) buildQRAdvancedPanel() fyne.CanvasObject {
	// Mask pattern
	maskPatterns := []string{
		"Auto",
		"Pattern 0", "Pattern 1", "Pattern 2", "Pattern 3",
		"Pattern 4", "Pattern 5", "Pattern 6", "Pattern 7",
	}
	a.maskPatternSelect = widget.NewSelect(
		maskPatterns,
		func(s string) {
			a.markQRDirty()
		},
	)
	a.maskPatternSelect.SetSelected("Auto")

	form := widget.NewForm(
		widget.NewFormItem("Mask Pattern", a.maskPatternSelect),
	)

	return container.NewPadded(form)
}

func (a *QRGeneratorApp) buildRenderAdvancedPanel() fyne.CanvasObject {
	// Radius slider and label
	a.radiusSlider = widget.NewSlider(1, 10)
	a.radiusSlider.Step = 1
	a.radiusSlider.Value = 3

	a.radiusLabel = widget.NewLabel("Radius: 3")

	a.radiusSlider.OnChanged = func(f float64) {
		a.radiusLabel.SetText(fmt.Sprintf("Radius: %d", int(f)))
		a.markRenderDirty()
	}

	// Kernel type
	kernelTypes := []string{
		"Lanczos2",
		"CubicSmooth",
		"Gaussian",
		"Lanczos3",
		"Hann",
		"Triangle",
		"Cosine",
		"Epanechnikov",
		"BSpline",
		"Box",
	}
	a.kernelTypeSelect = widget.NewSelect(
		kernelTypes,
		func(s string) {
			if kernel, ok := render.Kernels[s]; ok {
				minRadius := float64(kernel.DefaultRadius)

				// Update the slider minimum value
				a.radiusSlider.Min = minRadius

				// Clamp current value if needed
				if a.radiusSlider.Value < minRadius {
					a.radiusSlider.SetValue(minRadius)
				} else {
					a.radiusLabel.SetText(
						fmt.Sprintf("Radius: %d", int(a.radiusSlider.Value)),
					)
				}

				a.radiusSlider.Refresh()
			}

			a.markRenderDirty()
		},
	)
	a.kernelTypeSelect.SetSelected("Lanczos2")

	form := widget.NewForm(
		widget.NewFormItem("Kernel Type", a.kernelTypeSelect),
		widget.NewFormItem("Kernel Radius", container.NewVBox(
			a.radiusLabel,
			a.radiusSlider,
		)),
	)

	return container.NewPadded(form)
}

func (a *QRGeneratorApp) buildAdvancedOptionsPanel() fyne.CanvasObject {
	// QR Code advanced options
	qrAdvancedPanel := a.buildQRAdvancedPanel()
	// Render advanced options
	renderAdvancedPanel := a.buildRenderAdvancedPanel()

	// QR Code section
	qrExpanded := false
	qrContent := container.NewVBox(qrAdvancedPanel)
	qrContent.Hide()
	qrTitleBtn := widget.NewButtonWithIcon("QR Code Options", a.qrIcon, func() {
		qrExpanded = !qrExpanded
		if qrExpanded {
			qrContent.Show()
		} else {
			qrContent.Hide()
		}
	})
	qrTitleBtn.Importance = widget.LowImportance

	// Render section
	renderExpanded := false
	renderContent := container.NewVBox(renderAdvancedPanel)
	renderContent.Hide()
	renderTitleBtn := widget.NewButtonWithIcon("Render Options", a.paintIcon, func() {
		renderExpanded = !renderExpanded
		if renderExpanded {
			renderContent.Show()
		} else {
			renderContent.Hide()
		}
	})
	renderTitleBtn.Importance = widget.LowImportance

	advancedContent := container.NewVBox(
		qrTitleBtn,
		qrContent,
		widget.NewSeparator(),
		renderTitleBtn,
		renderContent,
	)

	return widget.NewCard("Advanced Options", "", container.NewPadded(advancedContent))
}

func (a *QRGeneratorApp) buildPreviewPanel() fyne.CanvasObject {
	a.previewImage = canvas.NewImageFromResource(nil)
	a.previewImage.FillMode = canvas.ImageFillContain
	a.previewImage.SetMinSize(
		fyne.NewSize(float32(800), float32(800)),
	)

	a.previewContainer = container.NewStack(a.previewImage)

	return container.NewPadded(a.previewContainer)
}

func (a *QRGeneratorApp) markQRDirty() {
	a.dirtyQR = true
	a.dirtyRender = true
	a.scheduleQRUpdate()
}

func (a *QRGeneratorApp) markRenderDirty() {
	a.dirtyRender = true
	a.scheduleQRUpdate()
}

func (a *QRGeneratorApp) scheduleQRUpdate() {
	// Stop the existing timer if it's still running
	if !a.updateTimer.Stop() {
		select {
		case <-a.updateTimer.C:
		default:
		}
	}

	// Start a new timer
	a.updateTimer.Reset(200 * time.Millisecond)

	go func() {
		<-a.updateTimer.C

		img := a.buildPreviewImage()
		if img == nil {
			return
		}

		fyne.Do(func() {
			a.previewImage.Image = img
			a.previewImage.Refresh()
		})
	}()
}

func (a *QRGeneratorApp) buildPreviewImage() image.Image {
	if a.dirtyQR {
		input := a.getCurrentInput()

		desired := qrBuildState{
			text:       input,
			minVersion: a.getMinVersion(),
			ecLevel:    a.getECLevel(),
			maskNum:    a.getMaskPattern(),
		}

		if a.lastBuildState == nil || !isSameBuild(a.lastBuildState, &desired) {
			builder := qrcode.NewQRBuilder(desired.text).
				WithMinVersion(desired.minVersion).
				WithErrorCorrectionLevel(desired.ecLevel).
				WithMaskNum(desired.maskNum)

			qrCode, err := builder.Build()
			if err != nil {
				dialog.ShowError(err, a.window)
				return nil
			}

			a.currentQRCode = qrCode
			a.lastBuildState = &desired
		}

		a.dirtyQR = false
	}

	if a.dirtyRender {
		a.renderer.
			WithModuleShape(a.getModuleShape()).
			WithKernelType(a.getKernelType()).
			WithRadius(a.getRadius())

		a.dirtyRender = false
	}

	return a.renderer.RenderImage(*a.currentQRCode)
}

func (a *QRGeneratorApp) showFileSaveDialog() {
	fileSaveDialog := dialog.NewFileSave(
		func(writer fyne.URIWriteCloser, err error) {
			if err != nil || writer == nil {
				return
			}
			defer writer.Close()

			a.saveQRCode(writer)
		},
		a.window,
	)

	// Set default filename
	fileSaveDialog.SetFileName("qrcode.png")

	fileSaveDialog.SetFilter(storage.NewExtensionFileFilter([]string{
		".png", ".jpg", ".jpeg",
	}))

	fileSaveDialog.Show()
}

func (a *QRGeneratorApp) saveQRCode(writer fyne.URIWriteCloser) {
	if a.currentQRCode == nil {
		writer.Close()
		return
	}

	path := writer.URI().Path()
	ext := strings.ToLower(writer.URI().Extension())

	var format qrconst.RenderFormat
	switch ext {
	case ".png":
		format = qrconst.RenderPNG
	case ".jpg", ".jpeg":
		format = qrconst.RenderJPEG
	default:
		writer.Close()

		err := storage.Delete(writer.URI())
		if err != nil {
			log.Printf("Error deleting incomplete file: %v\n", err)
		}

		dialog.ShowError(
			fmt.Errorf("unsupported file format"),
			a.window,
		)
		return
	}

	if err := a.renderer.RenderToWriter(*a.currentQRCode, writer, format); err != nil {
		writer.Close()

		err := storage.Delete(writer.URI())
		if err != nil {
			log.Printf("Error deleting incomplete file: %v\n", err)
		}

		dialog.ShowError(err, a.window)
		return
	}

	writer.Close()
	dialog.ShowInformation("Saved", fmt.Sprintf("QR Code saved to %s", path), a.window)
}

func (a *QRGeneratorApp) chooseColor(target string) {
	colorPicker := dialog.NewColorPicker("Choose color", "Select a color", func(c color.Color) {
		if c == nil {
			return
		}

		r, g, b, _ := c.RGBA()
		rgba := color.RGBA{
			R: uint8(r >> 8),
			G: uint8(g >> 8),
			B: uint8(b >> 8),
			A: 255,
		}

		switch target {
		case "background":
			a.renderer.WithBackgroundColor(rgba)
			a.backgroundColorBtn.SetText(fmt.Sprintf("Background: RGB(%d,%d,%d)",
				rgba.R, rgba.G, rgba.B))
		case "foreground":
			a.renderer.WithForegroundColor(rgba)
			a.foregroundColorBtn.SetText(fmt.Sprintf("Foreground: RGB(%d,%d,%d)",
				rgba.R, rgba.G, rgba.B))
		}

		a.markRenderDirty()
	}, a.window)
	colorPicker.Advanced = true // enable color wheel

	colorPicker.Show()
}

func (a *QRGeneratorApp) getCurrentInput() string {
	switch a.currentTab {
	case "plaintext":
		return a.plainTextEntry.Text
	case "url":
		return a.urlEntry.Text
	case "twitter":
		return a.twitterEntry.Text
	case "email":
		return a.emailEntry.Text
	case "wifi":
		ssid := a.wifiSSIDEntry.Text
		password := a.wifiPasswordEntry.Text
		encryption := a.wifiEncryptionSelect.Selected

		if (ssid != "") &&
			(password != "") &&
			(encryption != "") {
			return fmt.Sprintf("WIFI:S:%s;T:%s;P:%s;;",
				ssid, encryption, password)
		}

		return ""
	case "telephone":
		telephone := a.telephoneEntry.Text

		if telephone != "" {
			return fmt.Sprintf("tel:%s", a.telephoneEntry.Text)
		}

		return ""
	default:
		return a.plainTextEntry.Text
	}
}

func (a *QRGeneratorApp) getECLevel() qrconst.ErrorCorrectionLevel {
	selected := a.ecLevelRadio.Selected
	switch {
	case strings.Contains(selected, "L"):
		return qrconst.L
	case strings.Contains(selected, "M"):
		return qrconst.M
	case strings.Contains(selected, "Q"):
		return qrconst.Q
	case strings.Contains(selected, "H"):
		return qrconst.H
	default:
		return qrconst.M
	}
}

func (a *QRGeneratorApp) getMinVersion() int {
	version, err := strconv.Atoi(strings.TrimSpace(a.minVersionEntry.Text))
	if err != nil {
		return 1
	}

	if version < 1 {
		return 1
	}
	if version > 40 {
		return 40
	}

	return version
}

func (a *QRGeneratorApp) getMaskPattern() *int {
	selected := a.maskPatternSelect.Selected
	if selected == "Auto" {
		return nil
	}

	// Extract pattern number from "Pattern X"
	parts := strings.Split(selected, " ")
	if len(parts) != 2 {
		return nil
	}

	maskNum, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil
	}

	if maskNum < 0 || maskNum > 7 {
		return nil
	}

	return &maskNum
}

func (a *QRGeneratorApp) getModuleShape() qrconst.ModuleShape {
	switch a.moduleShapeSelect.Selected {
	case "Square":
		return qrconst.Square
	case "Circle":
		return qrconst.Circle
	case "HorizontalBlob":
		return qrconst.HorizontalBlob
	case "VerticalBlob":
		return qrconst.VerticalBlob
	case "Blob":
		return qrconst.Blob
	case "LeftLeaf":
		return qrconst.LeftLeaf
	case "RightLeaf":
		return qrconst.RightLeaf
	case "Diamond":
		return qrconst.Diamond
	case "WaterDroplet":
		return qrconst.WaterDroplet
	case "Star4":
		return qrconst.Star4
	case "Star5":
		return qrconst.Star5
	case "Star6":
		return qrconst.Star6
	case "Xs":
		return qrconst.Xs
	case "Octagon":
		return qrconst.Octagon
	case "SmileyFace":
		return qrconst.SmileyFace
	case "Pointillism":
		return qrconst.Pointillism
	default:
		return qrconst.Square
	}
}

func (a *QRGeneratorApp) getKernelType() string {
	return a.kernelTypeSelect.Selected
}

func (a *QRGeneratorApp) getRadius() int {
	if a.radiusSlider == nil {
		return 3
	}

	return int(a.radiusSlider.Value)
}

func digitsOnly(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}

	return b.String()
}

func validateMinVersionEntry(
	minVersion binding.Int,
	entry *widget.Entry,
) func(string) {
	return func(s string) {
		cleanStr := digitsOnly(s)

		// If empty string -> reset minVersion to 1
		if cleanStr == "" {
			_ = minVersion.Set(1)
			return
		}

		// If length of string is 3 or more -> do nothing
		if len(cleanStr) > 2 {
			entry.SetText(s[:len(s)-1])
			return
		}

		// If numeric but out of range -> clamp
		v, err := strconv.Atoi(cleanStr)
		if err != nil {
			return
		}

		if v < 1 {
			v = 1
		} else if v > 40 {
			v = 40
		}

		// If string is valid -> update binding
		_ = minVersion.Set(v)
		entry.SetText(strconv.Itoa(v))
	}
}

func isSameBuild(a, b *qrBuildState) bool {
	if a == nil || b == nil {
		return false
	}

	// If both are templates (empty text)
	if a.text == "" && b.text == "" {
		// Changing version changes the modules count
		if a.minVersion != b.minVersion {
			return false
		}

		// EC level and mask number don't matter for template modules
		return true
	}

	if a.text != b.text ||
		a.minVersion != b.minVersion ||
		a.ecLevel != b.ecLevel {
		return false
	}

	if (a.maskNum == nil) != (b.maskNum == nil) {
		return false
	}
	if a.maskNum != nil && *a.maskNum != *b.maskNum {
		return false
	}

	return true
}
