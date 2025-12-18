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
/>
`}
	},
	qrconst.Circle: func(lookahead qrconst.Lookahead) []string {
		var paths []string

		paths = append(paths, `<path
	d="M 0.5 0.5 m -0.5 0 a 0.5 0.5 0 1 0 1 0 a 0.5 0.5 0 1 0 -1 0"
/>
`)

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

		if R2UR {
			paths = append(paths, `<path
	d="M 1 .5 v -.5 Z"
	stroke-width=".025"
/>
`)
		}
		if UR2U {
			paths = append(paths, `<path
	d="M 1 0 h -.5 Z"
	stroke-width=".025"
/>
`)
		}
		if U2UL {
			paths = append(paths, `<path
	d="M .5 0 h -.5 Z"
	stroke-width=".025"
/>
`)
		}
		if UL2L {
			paths = append(paths, `<path
	d="M 0 0 v .5 Z"
	stroke-width=".025"
/>
`)
		}
		if L2DL {
			paths = append(paths, `<path
	d="M 0 .5 v .5 Z"
	stroke-width=".025"
/>
`)
		}
		if DL2D {
			paths = append(paths, `<path
	d="M 0 1 h .5 Z"
	stroke-width=".025"
/>
`)
		}
		if D2DR {
			paths = append(paths, `<path
	d="M .5 1 h .5 Z"
	stroke-width=".025"
/>
`)
		}
		if DR2R {
			paths = append(paths, `<path
	d="M 1 1 v -.5 Z"
	stroke-width=".025"
/>
`)
		}

		return paths
	},
	qrconst.HorizontalBlob: func(lookahead qrconst.Lookahead) []string {
		var paths []string

		paths = append(paths, `<path
	d="M 0.5 0.5 m -0.5 0 a 0.5 0.5 0 1 0 1 0 a 0.5 0.5 0 1 0 -1 0"
/>
`)

		R := lookahead.Has(qrconst.LookR)
		L := lookahead.Has(qrconst.LookL)

		UR := lookahead.Has(qrconst.LookUR) &&
			lookahead.Lacks(qrconst.LookR, qrconst.LookU)
		UL := lookahead.Has(qrconst.LookUL) &&
			lookahead.Lacks(qrconst.LookU, qrconst.LookL)
		DL := lookahead.Has(qrconst.LookDL) &&
			lookahead.Lacks(qrconst.LookL, qrconst.LookD)
		DR := lookahead.Has(qrconst.LookDR) &&
			lookahead.Lacks(qrconst.LookD, qrconst.LookR)

		if R {
			paths = append(paths, `<path
	d="M 1 0.5 V 0 H 0.5 A 0.5 0.5 0 0 1 1 0.5"
/>
`)
			paths = append(paths, `<path
	d="M 0.5 1 H 1 V 0.5 A 0.5 0.5 0 0 1 0.5 1"
/>
`)
		}
		if L {
			paths = append(paths, `<path
	d="M 0.5 0 H 0 V 0.5 A 0.5 0.5 0 0 1 0.5 0"
/>
`)
			paths = append(paths, `<path
	d="M 0 0.5 V 1 H 0.5 A 0.5 0.5 0 0 1 0 0.5"
/>
`)
		}

		if UR {
			paths = append(paths, `<path
	d="M 1 0.5 V 0 H 0.5 A 0.5 0.5 0 0 1 1 0.5"
/>
`)
		}
		if UL {
			paths = append(paths, `<path
	d="M 0.5 0 H 0 V 0.5 A 0.5 0.5 0 0 1 0.5 0"
/>
`)
		}
		if DL {
			paths = append(paths, `<path
	d="M 0 0.5 V 1 H 0.5 A 0.5 0.5 0 0 1 0 0.5"
/>
`)
		}
		if DR {
			paths = append(paths, `<path
	d="M 0.5 1 H 1 V 0.5 A 0.5 0.5 0 0 1 0.5 1"
/>
`)
		}

		return paths
	},
	qrconst.VerticalBlob: func(lookahead qrconst.Lookahead) []string {
		var paths []string

		paths = append(paths, `<path
	d="M 0.5 0.5 m -0.5 0 a 0.5 0.5 0 1 0 1 0 a 0.5 0.5 0 1 0 -1 0"
/>
`)

		U := lookahead.Has(qrconst.LookU)
		D := lookahead.Has(qrconst.LookD)

		UR := lookahead.Has(qrconst.LookUR) &&
			lookahead.Lacks(qrconst.LookR, qrconst.LookU)
		UL := lookahead.Has(qrconst.LookUL) &&
			lookahead.Lacks(qrconst.LookU, qrconst.LookL)
		DL := lookahead.Has(qrconst.LookDL) &&
			lookahead.Lacks(qrconst.LookL, qrconst.LookD)
		DR := lookahead.Has(qrconst.LookDR) &&
			lookahead.Lacks(qrconst.LookD, qrconst.LookR)

		if U {
			paths = append(paths, `<path
	d="M 1 0.5 V 0 H 0.5 A 0.5 0.5 0 0 1 1 0.5"
/>
`)
			paths = append(paths, `<path
	d="M 0.5 0 H 0 V 0.5 A 0.5 0.5 0 0 1 0.5 0"
/>
`)
		}
		if D {
			paths = append(paths, `<path
	d="M 0 0.5 V 1 H 0.5 A 0.5 0.5 0 0 1 0 0.5"
/>
`)
			paths = append(paths, `<path
	d="M 0.5 1 H 1 V 0.5 A 0.5 0.5 0 0 1 0.5 1"
/>
`)
		}

		if UR {
			paths = append(paths, `<path
	d="M 1 0.5 V 0 H 0.5 A 0.5 0.5 0 0 1 1 0.5"
/>
`)
		}
		if UL {
			paths = append(paths, `<path
	d="M 0.5 0 H 0 V 0.5 A 0.5 0.5 0 0 1 0.5 0"
/>
`)
		}
		if DL {
			paths = append(paths, `<path
	d="M 0 0.5 V 1 H 0.5 A 0.5 0.5 0 0 1 0 0.5"
/>
`)
		}
		if DR {
			paths = append(paths, `<path
	d="M 0.5 1 H 1 V 0.5 A 0.5 0.5 0 0 1 0.5 1"
/>
`)
		}

		return paths
	},
	qrconst.Blob: func(lookahead qrconst.Lookahead) []string {
		var paths []string

		paths = append(paths, `<path
	d="M 0.5 0.5 m -0.5 0 a 0.5 0.5 0 1 0 1 0 a 0.5 0.5 0 1 0 -1 0"
/>
`)

		R := lookahead.Has(qrconst.LookR)
		U := lookahead.Has(qrconst.LookU)
		L := lookahead.Has(qrconst.LookL)
		D := lookahead.Has(qrconst.LookD)

		UR := lookahead.Has(qrconst.LookUR)
		UL := lookahead.Has(qrconst.LookUL)
		DL := lookahead.Has(qrconst.LookDL)
		DR := lookahead.Has(qrconst.LookDR)

		if R {
			paths = append(paths, `<path
	d="M 1 0 H 0.5 A 0.5 0.5 0 0 1 0.5 1 H 1 Z"
/>
`)
		}
		if U {
			paths = append(paths, `<path
	d="M 0 0 V 0.5 A 0.5 0.5 0 0 1 1 .5 V 0 Z"
/>
`)
		}
		if L {
			paths = append(paths, `<path
	d="M 0 1 H 0.5 A 0.5 0.5 0 0 1 0.5 0 H 0 Z"
/>
`)
		}
		if D {
			paths = append(paths, `<path
	d="M 1 1 V 0.5 A 0.5 0.5 0 0 1 0 .5 V 1 Z"
/>
`)
		}

		if UR {
			paths = append(paths, `<path
	d="M 1 0.5 V 0 H 0.5 A 0.5 0.5 0 0 1 1 0.5"
/>
`)
		}
		if UL {
			paths = append(paths, `<path
	d="M 0.5 0 H 0 V 0.5 A 0.5 0.5 0 0 1 0.5 0"
/>
`)
		}
		if DL {
			paths = append(paths, `<path
	d="M 0 0.5 V 1 H 0.5 A 0.5 0.5 0 0 1 0 0.5"
/>
`)
		}
		if DR {
			paths = append(paths, `<path
	d="M 0.5 1 H 1 V 0.5 A 0.5 0.5 0 0 1 0.5 1"
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
	qrconst.Star4: func(lookahead qrconst.Lookahead) []string {
		return []string{`<path
	d="M 1 0.5 L .67 .33 L 0.5 0 L .33 .33 L 0 .5 L .33 .67 L .5 1 L .67 .67 L 1 .5"
	fill="rgba(0,0,0,1.)"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="round"
	stroke-linejoin="round"
	stroke-width="0.125"
/>
`}
	},
	qrconst.Star5: func(lookahead qrconst.Lookahead) []string {
		return []string{`<path
	d="M 1 0.5 L 0.65450849 0.38774301 L 0.6545085 0.02447174 L 0.44098301 0.31836437 L 0.0954915 0.20610737 L 0.309017 0.5 L 0.0954915 0.79389263 L 0.44098301 0.68163563 L 0.6545085 0.97552826 L 0.65450849 0.61225699 L 1 .5"
	fill="rgba(0,0,0,1.)"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="round"
	stroke-linejoin="round"
	stroke-width="0.125"
/>
`}
	},
	qrconst.Star6: func(lookahead qrconst.Lookahead) []string {
		return []string{`<path
	d="M 1 .5 L 0.74999988 0.3556625 L .75 .0669873 L 0.5 0.211325 L .25 .0669873 L 0.25000012 0.3556625 L 0 .5 L 0.25000012 0.6443375 L .25 .9330127 L 0.5 0.788675 L .75 .9330127 L 0.74999988 0.6443375 L 1 .5"
	fill="rgba(0,0,0,1.)"
	stroke="rgba(0,0,0,1.)"
	stroke-linecap="round"
	stroke-linejoin="round"
	stroke-width="0.125"
/>
`}
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
		var paths []string

		cornerUR := lookahead.Has(qrconst.LookR, qrconst.LookUR, qrconst.LookU)
		cornerUL := lookahead.Has(qrconst.LookU, qrconst.LookUL, qrconst.LookL)
		cornerDL := lookahead.Has(qrconst.LookL, qrconst.LookDL, qrconst.LookD)
		cornerDR := lookahead.Has(qrconst.LookD, qrconst.LookDR, qrconst.LookR)

		if cornerUR {
			paths = append(paths, `<path
	d="M 1 .5 A .5 .5 0 0 0 .5 0"
	fill="none"
	stroke-width=".025"
/>
`)
		}
		if cornerUL {
			paths = append(paths, `<path
	d="M .5 0 A .5 .5 0 0 0 0 .5"
	fill="none"
	stroke-width=".025"
/>
`)
		}
		if cornerDL {
			paths = append(paths, `<path
	d="M 0 .5 A .5 .5 0 0 0 .5 1"
	fill="none"
	stroke-width=".025"
/>
`)
		}
		if cornerDR {
			paths = append(paths, `<path
	d="M .5 1 A .5 .5 0 0 0 1 .5"
	fill="none"
	stroke-width=".025"
/>
`)
		}

		return paths
	},
	qrconst.HorizontalBlob: func(lookahead qrconst.Lookahead) []string {
		var paths []string

		UR := lookahead.Lacks(qrconst.LookUR) &&
			lookahead.Has(qrconst.LookR, qrconst.LookU)
		UL := lookahead.Lacks(qrconst.LookUL) &&
			lookahead.Has(qrconst.LookU, qrconst.LookL)
		DL := lookahead.Lacks(qrconst.LookDL) &&
			lookahead.Has(qrconst.LookL, qrconst.LookD)
		DR := lookahead.Lacks(qrconst.LookDR) &&
			lookahead.Has(qrconst.LookD, qrconst.LookR)

		if UR {
			paths = append(paths, `<path
	d="M 1 0.5 V 0 H 0.5 A 0.5 0.5 0 0 1 1 0.5"
/>
`)
		}
		if UL {
			paths = append(paths, `<path
	d="M 0.5 0 H 0 V 0.5 A 0.5 0.5 0 0 1 0.5 0"
/>
`)
		}
		if DL {
			paths = append(paths, `<path
	d="M 0 0.5 V 1 H 0.5 A 0.5 0.5 0 0 1 0 0.5"
/>
`)
		}
		if DR {
			paths = append(paths, `<path
	d="M 0.5 1 H 1 V 0.5 A 0.5 0.5 0 0 1 0.5 1"
/>
`)
		}

		return paths
	},
	qrconst.VerticalBlob: func(lookahead qrconst.Lookahead) []string {
		var paths []string

		UR := lookahead.Lacks(qrconst.LookUR) &&
			lookahead.Has(qrconst.LookR, qrconst.LookU)
		UL := lookahead.Lacks(qrconst.LookUL) &&
			lookahead.Has(qrconst.LookU, qrconst.LookL)
		DL := lookahead.Lacks(qrconst.LookDL) &&
			lookahead.Has(qrconst.LookL, qrconst.LookD)
		DR := lookahead.Lacks(qrconst.LookDR) &&
			lookahead.Has(qrconst.LookD, qrconst.LookR)

		if UR {
			paths = append(paths, `<path
	d="M 1 0.5 V 0 H 0.5 A 0.5 0.5 0 0 1 1 0.5"
/>
`)
		}
		if UL {
			paths = append(paths, `<path
	d="M 0.5 0 H 0 V 0.5 A 0.5 0.5 0 0 1 0.5 0"
/>
`)
		}
		if DL {
			paths = append(paths, `<path
	d="M 0 0.5 V 1 H 0.5 A 0.5 0.5 0 0 1 0 0.5"
/>
`)
		}
		if DR {
			paths = append(paths, `<path
	d="M 0.5 1 H 1 V 0.5 A 0.5 0.5 0 0 1 0.5 1"
/>
`)
		}

		return paths
	},
	qrconst.Blob: func(lookahead qrconst.Lookahead) []string {
		var paths []string

		UR := lookahead.Has(qrconst.LookR, qrconst.LookU)
		UL := lookahead.Has(qrconst.LookU, qrconst.LookL)
		DL := lookahead.Has(qrconst.LookL, qrconst.LookD)
		DR := lookahead.Has(qrconst.LookD, qrconst.LookR)

		if UR {
			paths = append(paths, `<path
	d="M 1 0.5 V 0 H 0.5 A 0.5 0.5 0 0 1 1 0.5"
/>
`)
		}
		if UL {
			paths = append(paths, `<path
	d="M 0.5 0 H 0 V 0.5 A 0.5 0.5 0 0 1 0.5 0"
/>
`)
		}
		if DL {
			paths = append(paths, `<path
	d="M 0 0.5 V 1 H 0.5 A 0.5 0.5 0 0 1 0 0.5"
/>
`)
		}
		if DR {
			paths = append(paths, `<path
	d="M 0.5 1 H 1 V 0.5 A 0.5 0.5 0 0 1 0.5 1"
/>
`)
		}

		return paths
	},
	qrconst.Diamond: func(lookahead qrconst.Lookahead) []string {
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
