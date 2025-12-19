package tables

import (
	"fmt"
	"math"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
)

var PathRenderFunctions = map[qrconst.ModuleShape]func(lookahead qrconst.Lookahead) []string{
	qrconst.Square: func(lookahead qrconst.Lookahead) []string {
		return []string{`<path
	d="M 0 0 h 1 v 1 h -1 Z"
	fill="rgba(0,0,0,1.)"
/>
`}
	},
	qrconst.Circle: func(lookahead qrconst.Lookahead) []string {
		return []string{`<path
	d="M .5 .5 m -.5 0 a .5 .5 0 1 0 1 0 a .5 .5 0 1 0 -1 0"
	fill="rgba(0,0,0,1.)"
/>
`}
	},
	qrconst.TiedCircle: func(lookahead qrconst.Lookahead) []string {
		var paths []string

		paths = append(paths, `<path
	d="M .5 .5 m -.5 0 a .5 .5 0 1 0 1 0 a .5 .5 0 1 0 -1 0"
	fill="rgba(0,0,0,1.)"
/>
`)

		UR := lookahead.Lacks(qrconst.LookR, qrconst.LookU)
		UL := lookahead.Lacks(qrconst.LookU, qrconst.LookL)
		DL := lookahead.Lacks(qrconst.LookL, qrconst.LookD)
		DR := lookahead.Lacks(qrconst.LookD, qrconst.LookR)

		R2UR := lookahead.Has(qrconst.LookU) &&
			lookahead.Lacks(qrconst.LookR, qrconst.LookUR)
		UR2U := lookahead.Has(qrconst.LookR) &&
			lookahead.Lacks(qrconst.LookUR, qrconst.LookU)
		U2UL := lookahead.Has(qrconst.LookL) &&
			lookahead.Lacks(qrconst.LookU, qrconst.LookUL)
		UL2L := lookahead.Has(qrconst.LookU) &&
			lookahead.Lacks(qrconst.LookUL, qrconst.LookL)
		L2DL := lookahead.Has(qrconst.LookD) &&
			lookahead.Lacks(qrconst.LookL, qrconst.LookDL)
		DL2D := lookahead.Has(qrconst.LookL) &&
			lookahead.Lacks(qrconst.LookDL, qrconst.LookD)
		D2DR := lookahead.Has(qrconst.LookR) &&
			lookahead.Lacks(qrconst.LookD, qrconst.LookDR)
		DR2R := lookahead.Has(qrconst.LookD) &&
			lookahead.Lacks(qrconst.LookDR, qrconst.LookR)

		if UR {
			paths = append(paths, `<path
	d="M 1 .5 A .5 .5 0 0 0 .5 0"
	fill="none"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}
		if UL {
			paths = append(paths, `<path
	d="M .5 0 A .5 .5 0 0 0 0 .5"
	fill="none"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}
		if DL {
			paths = append(paths, `<path
	d="M 0 .5 A .5 .5 0 0 0 .5 1"
	fill="none"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}
		if DR {
			paths = append(paths, `<path
	d="M .5 1 A .5 .5 0 0 0 1 .5"
	fill="none"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}

		if R2UR {
			paths = append(paths, `<path
	d="M 1 .5 v -.5 Z"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}
		if UR2U {
			paths = append(paths, `<path
	d="M 1 0 h -.5 Z"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}
		if U2UL {
			paths = append(paths, `<path
	d="M .5 0 h -.5 Z"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}
		if UL2L {
			paths = append(paths, `<path
	d="M 0 0 v .5 Z"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}
		if L2DL {
			paths = append(paths, `<path
	d="M 0 .5 v .5 Z"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}
		if DL2D {
			paths = append(paths, `<path
	d="M 0 1 h .5 Z"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}
		if D2DR {
			paths = append(paths, `<path
	d="M .5 1 h .5 Z"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}
		if DR2R {
			paths = append(paths, `<path
	d="M 1 1 v -.5 Z"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}

		return paths
	},
	qrconst.HorizontalBlob: func(lookahead qrconst.Lookahead) []string {
		var paths []string

		paths = append(paths, `<path
	d="M .5 .5 m -.5 0 a .5 .5 0 1 0 1 0 a .5 .5 0 1 0 -1 0"
	fill="rgba(0,0,0,1.)"
/>
`)

		R := lookahead.Has(qrconst.LookR)

		UR := lookahead.Has(qrconst.LookUR) &&
			lookahead.Lacks(qrconst.LookR, qrconst.LookU)
		DR := lookahead.Has(qrconst.LookDR) &&
			lookahead.Lacks(qrconst.LookD, qrconst.LookR)

		if R {
			paths = append(paths, `<path
	d="M .5 0 V 1 H 1.5 V 0 Z"
	fill="rgba(0,0,0,1.)"
/>
`)
		}

		if UR {
			paths = append(paths, `<path
	d="M .5 .5 H 1 A .5 .5 0 0 1 1.5 0 V -.5 H 1 A .5 .5 0 0 1 .5 0 Z"
	fill="rgba(0,0,0,1.)"
/>
`)
		}
		if DR {
			paths = append(paths, `<path
	d="M .5 .5 V 1 A .5 .5 0 0 1 1 1.5 H 1.5 V 1 A .5 .5 0 0 1 1 .5 Z"
	fill="rgba(0,0,0,1.)"
/>
`)
		}

		return paths
	},
	qrconst.VerticalBlob: func(lookahead qrconst.Lookahead) []string {
		var paths []string

		paths = append(paths, `<path
	d="M .5 .5 m -.5 0 a .5 .5 0 1 0 1 0 a .5 .5 0 1 0 -1 0"
	fill="rgba(0,0,0,1.)"
/>
`)

		D := lookahead.Has(qrconst.LookD)

		DL := lookahead.Has(qrconst.LookDL) &&
			lookahead.Lacks(qrconst.LookL, qrconst.LookD)
		DR := lookahead.Has(qrconst.LookDR) &&
			lookahead.Lacks(qrconst.LookD, qrconst.LookR)

		if D {
			paths = append(paths, `<path
	d="M 1 .5 H 0 V 1.5 H 1 Z"
	fill="rgba(0,0,0,1.)"
/>
`)
		}

		if DL {
			paths = append(paths, `<path
	d="M .5 .5 H 0 A .5 .5 0 0 1 -.5 1 V 1.5 H 0 A .5 .5 0 0 1 .5 1 Z"
	fill="rgba(0,0,0,1.)"
/>
`)
		}
		if DR {
			paths = append(paths, `<path
	d="M .5 .5 V 1 A .5 .5 0 0 1 1 1.5 H 1.5 V 1 A .5 .5 0 0 1 1 .5 Z"
	fill="rgba(0,0,0,1.)"
/>
`)
		}

		return paths
	},
	qrconst.Blob: func(lookahead qrconst.Lookahead) []string {
		var paths []string

		paths = append(paths, `<path
	d="M .5 .5 m -.5 0 a .5 .5 0 1 0 1 0 a .5 .5 0 1 0 -1 0"
	fill="rgba(0,0,0,1.)"
/>
`)

		R := lookahead.Has(qrconst.LookR)
		UR := lookahead.Has(qrconst.LookUR)

		D := lookahead.Has(qrconst.LookD)
		DR := lookahead.Has(qrconst.LookDR)

		if R {
			paths = append(paths, `<path
	d="M .5 0 V 1 H 1.5 V 0 Z"
	fill="rgba(0,0,0,1.)"
/>
`)
		}
		if UR {
			paths = append(paths, `<path
	d="M .5 .5 H 1 A .5 .5 0 0 1 1.5 0 V -.5 H 1 A .5 .5 0 0 1 .5 0 Z"
	fill="rgba(0,0,0,1.)"
/>
`)
		}

		if D {
			paths = append(paths, `<path
	d="M 1 .5 H 0 V 1.5 H 1 Z"
	fill="rgba(0,0,0,1.)"
/>
`)
		}
		if DR {
			paths = append(paths, `<path
	d="M .5 .5 V 1 A .5 .5 0 0 1 1 1.5 H 1.5 V 1 A .5 .5 0 0 1 1 .5 Z"
	fill="rgba(0,0,0,1.)"
/>
`)
		}

		return paths
	},
	qrconst.Diamond: func(lookahead qrconst.Lookahead) []string {
		return []string{`<path
	d="M 1 .5 L .5 0 L 0 .5 L .5 1 Z"
/>
`}
	},
	qrconst.WaterDroplet: func(lookahead qrconst.Lookahead) []string {
		return []string{`<path
	d="M 0 0 C .1 .1 1 .15 1 .667 Q 1 1 .667 1 C .25 1 .1 .1 0 0"
	fill="rgba(0,0,0,1.)"
/>
`}

	},
	qrconst.Star4: func(lookahead qrconst.Lookahead) []string {
		return []string{`<path
	d="M 1 .5 L .67 .33 L .5 0 L .33 .33 L 0 .5 L .33 .67 L .5 1 L .67 .67 L 1 .5"
	fill="rgba(0,0,0,1.)"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="round"
	stroke-linejoin="round"
	stroke-width=".125"
/>
`}
	},
	qrconst.Star5: func(lookahead qrconst.Lookahead) []string {
		return []string{`<path
	d="M 1 .5 L .65450849 .38774301 L .6545085 .02447174 L .44098301 .31836437 L .0954915 .20610737 L .309017 .5 L .0954915 .79389263 L .44098301 .68163563 L .6545085 .97552826 L .65450849 .61225699 L 1 .5"
	fill="rgba(0,0,0,1.)"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="round"
	stroke-linejoin="round"
	stroke-width=".125"
/>
`}
	},
	qrconst.Star6: func(lookahead qrconst.Lookahead) []string {
		return []string{`<path
	d="M 1 .5 L .74999988 .3556625 L .75 .0669873 L .5 .211325 L .25 .0669873 L .25000012 .3556625 L 0 .5 L .25000012 .6443375 L .25 .9330127 L .5 .788675 L .75 .9330127 L .74999988 .6443375 L 1 .5"
	fill="rgba(0,0,0,1.)"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="round"
	stroke-linejoin="round"
	stroke-width=".1"
/>
`}
	},
	qrconst.Star8: func(lookahead qrconst.Lookahead) []string {
		return []string{`<path
	d="M 1 .5 L .85355299 .35355356 L .85355339 .14644661 L .64644644 .14644701 L .5 0 L .35355356 .14644701 L .14644661 .14644661 L .14644701 .35355356 L 0 .5 L .14644701 .64644644 L .14644661 .85355339 L .35355356 .85355299 L .5 1 L .64644644 .85355299 L .85355339 .85355339 L .85355299 .64644644 L 1 .5"
	fill="rgba(0,0,0,1.)"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="round"
	stroke-linejoin="round"
	stroke-width=".025"
/>
`}
	},
	qrconst.Xs: func(lookahead qrconst.Lookahead) []string {
		var paths []string

		d := "M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 "

		R := lookahead.Has(qrconst.LookR)
		UR := lookahead.Has(qrconst.LookUR)

		D := lookahead.Has(qrconst.LookD)
		DR := lookahead.Has(qrconst.LookDR)

		if D {
			d += "V 1.25 "
		} else {
			d += "V 1 H .25 "
		}
		d += "L .5 .75 "

		if DR {
			d += "L 1 1.25 H 1.25 V 1 "
		} else {
			if D {
				d += "L 1 1.25 V 1 "
			} else {
				d += "L .75 1 H 1 "
			}

			if R {
				d += "H 1.25 "
			} else {
				d += "V .75 "
			}
		}
		d += "L .75 .5 "

		if UR {
			d += "L 1.25 0 V -.25 H 1 Z"
		} else {
			if R {
				d += "L 1.25 0 H .75 Z"
			} else {
				d += "L 1 .25 V 0 H .75 Z"
			}
		}

		paths = append(paths, fmt.Sprintf(
			`<path
	d="%s"
	fill="rgba(0,0,0,1.)"
/>
`,
			d,
		))

		if lookahead.Has(qrconst.LookStructural) {
			paths = append(paths, `<path
	d="M 1 .5 L .5 0 L 0 .5 L .5 1 Z"
	fill="none"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="round"
	stroke-linejoin="round"
	stroke-width=".025"
/>
`)
		}

		return paths
	},
	qrconst.Pointillism: func(lookahead qrconst.Lookahead) []string {
		if lookahead.Has(qrconst.LookStructural) {
			return []string{`<circle
	cx=".5" cy=".5" r=".5"
/>
`}
		}

		neighbors := 0
		mask := qrconst.LookR
		for range 8 {
			if lookahead&mask != 0 {
				neighbors++
			}
			mask <<= 1
		}

		minR := .24
		maxR := .60

		k := .18
		t := 1 - math.Exp(-k*float64(neighbors))
		if t > 1 {
			t = 1
		}
		r := minR + (maxR-minR)*t

		return []string{
			fmt.Sprintf(`<circle
	cx=".5" cy=".5" r="%f"
/>
`,
				r,
			),
		}
	},
}

var PathMergeFunctions = map[qrconst.ModuleShape]func(lookahead qrconst.Lookahead) []string{
	qrconst.Square: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.Circle: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.TiedCircle: func(lookahead qrconst.Lookahead) []string {
		var paths []string

		cornerUR := lookahead.Has(qrconst.LookR, qrconst.LookUR, qrconst.LookU)
		cornerUL := lookahead.Has(qrconst.LookU, qrconst.LookUL, qrconst.LookL)
		cornerDL := lookahead.Has(qrconst.LookL, qrconst.LookDL, qrconst.LookD)
		cornerDR := lookahead.Has(qrconst.LookD, qrconst.LookDR, qrconst.LookR)

		if cornerUR {
			paths = append(paths, `<path
	d="M 1 .5 A .5 .5 0 0 0 .5 0"
	fill="none"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}
		if cornerUL {
			paths = append(paths, `<path
	d="M .5 0 A .5 .5 0 0 0 0 .5"
	fill="none"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}
		if cornerDL {
			paths = append(paths, `<path
	d="M 0 .5 A .5 .5 0 0 0 .5 1"
	fill="none"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}
		if cornerDR {
			paths = append(paths, `<path
	d="M .5 1 A .5 .5 0 0 0 1 .5"
	fill="none"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="square"
	stroke-linejoin="round"
	stroke-width=".05"
/>
`)
		}

		return paths
	},
	qrconst.HorizontalBlob: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.VerticalBlob: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.Blob: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.Diamond: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.WaterDroplet: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.Star4: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.Star5: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.Star6: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.Star8: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.Xs: func(lookahead qrconst.Lookahead) []string {
		var paths []string

		UR := lookahead.Has(qrconst.LookR, qrconst.LookU)
		UL := lookahead.Has(qrconst.LookU, qrconst.LookL)
		DL := lookahead.Has(qrconst.LookL, qrconst.LookD)
		DR := lookahead.Has(qrconst.LookD, qrconst.LookR)

		if UR {
			paths = append(paths, `<path
	d="M 1 .5 L .5 0"
	fill="none"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="round"
	stroke-linejoin="round"
	stroke-width=".025"
/>
`)
		}
		if UL {
			paths = append(paths, `<path
	d="M .5 0 L 0 .5"
	fill="none"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="round"
	stroke-linejoin="round"
	stroke-width=".025"
/>
`)
		}
		if DL {
			paths = append(paths, `<path
	d="M 0 .5 L .5 1"
	fill="none"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="round"
	stroke-linejoin="round"
	stroke-width=".025"
/>
`)
		}
		if DR {
			paths = append(paths, `<path
	d="M .5 1 L 1 .5"
	fill="none"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="round"
	stroke-linejoin="round"
	stroke-width=".025"
/>
`)
		}

		return paths
	},
	qrconst.Pointillism: func(lookahead qrconst.Lookahead) []string {
		if lookahead.Has(qrconst.LookStructural) {
			return nil
		}

		neighbors := 0
		mask := qrconst.LookR
		for range 8 {
			if lookahead&mask != 0 {
				neighbors++
			}
			mask <<= 1
		}

		minR := .06
		maxR := .18

		k := .06
		t := 1 - math.Exp(-k*float64(neighbors))
		r := minR + (maxR-minR)*t

		return []string{
			fmt.Sprintf(`<circle
	cx=".5" cy=".5" r="%f"
/>
`,
				r,
			),
		}
	},
}
