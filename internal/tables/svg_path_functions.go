package tables

import (
	"strconv"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
)

type rule struct {
	cond   bool
	suffix string
}

var PathSymbols = map[qrconst.ModuleShape][]string{
	qrconst.Square: {
		`<path
			id="` + qrconst.Square.String() + `__render"
			d="M 1 0 H 0 V 1 H 1 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Square.String() + `__render--r"
			d="M .96 0 V 1 H 1.04 V 0 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Square.String() + `__render--d"
			d="M 1 .96 H 0 V 1.04 H 1 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
	},
	qrconst.Circle: {
		`<path
			id="` + qrconst.Circle.String() + `__render"
			d="M .5 .5 m -.5 0 a .5 .5 0 1 0 1 0 a .5 .5 0 1 0 -1 0"
			fill="rgba(0,0,0,1.)"
		/>`,
	},
	qrconst.TiedCircle: {
		`<path
			id="` + qrconst.TiedCircle.String() + `__render"
			d="M .5 .5 m -.5 0 a .5 .5 0 1 0 1 0 a .5 .5 0 1 0 -1 0"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__render--ur"
			d="M 1 .5 A .5 .5 0 0 0 .5 0"
			fill="none"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__render--ul"
			d="M .5 0 A .5 .5 0 0 0 0 .5"
			fill="none"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__render--dl"
			d="M 0 .5 A .5 .5 0 0 0 .5 1"
			fill="none"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__render--dr"
			d="M .5 1 A .5 .5 0 0 0 1 .5"
			fill="none"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__render--r2ur"
			d="M 1 .5 v -.5 Z"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__render--ur2u"
			d="M 1 0 h -.5 Z"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__render--u2ul"
			d="M .5 0 h -.5 Z"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__render--ul2l"
			d="M 0 0 v .5 Z"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__render--l2dl"
			d="M 0 .5 v .5 Z"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__render--dl2d"
			d="M 0 1 h .5 Z"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__render--d2dr"
			d="M .5 1 h .5 Z"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__render--dr2r"
			d="M 1 1 v -.5 Z"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__merge--ur"
			d="M 1 .5 A .5 .5 0 0 0 .5 0"
			fill="none"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__merge--ul"
			d="M .5 0 A .5 .5 0 0 0 0 .5"
			fill="none"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__merge--dl"
			d="M 0 .5 A .5 .5 0 0 0 .5 1"
			fill="none"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
		`<path
			id="` + qrconst.TiedCircle.String() + `__merge--dr"
			d="M .5 1 A .5 .5 0 0 0 1 .5"
			fill="none"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="square"
			stroke-linejoin="round"
			stroke-width=".05"
		/>`,
	},
	qrconst.HorizontalBlob: {
		`<path
			id="` + qrconst.HorizontalBlob.String() + `__render"
			d="M .5 .5 m -.5 0 a .5 .5 0 1 0 1 0 a .5 .5 0 1 0 -1 0"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.HorizontalBlob.String() + `__render--r"
			d="M .5 0 V 1 H 1.5 V 0 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.HorizontalBlob.String() + `__render--ur"
			d="M .5 .5 H 1 A .5 .5 0 0 1 1.5 0 V -.5 H 1 A .5 .5 0 0 1 .5 0 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.HorizontalBlob.String() + `__render--dr"
			d="M .5 .5 V 1 A .5 .5 0 0 1 1 1.5 H 1.5 V 1 A .5 .5 0 0 1 1 .5 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
	},
	qrconst.VerticalBlob: {
		`<path
			id="` + qrconst.VerticalBlob.String() + `__render"
			d="M .5 .5 m -.5 0 a .5 .5 0 1 0 1 0 a .5 .5 0 1 0 -1 0"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.VerticalBlob.String() + `__render--d"
			d="M 1 .5 H 0 V 1.5 H 1 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.VerticalBlob.String() + `__render--dl"
			d="M .5 .5 H 0 A .5 .5 0 0 1 -.5 1 V 1.5 H 0 A .5 .5 0 0 1 .5 1 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.VerticalBlob.String() + `__render--dr"
			d="M .5 .5 V 1 A .5 .5 0 0 1 1 1.5 H 1.5 V 1 A .5 .5 0 0 1 1 .5 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
	},
	qrconst.Blob: {
		`<path
			id="` + qrconst.Blob.String() + `__render"
			d="M .5 .5 m -.5 0 a .5 .5 0 1 0 1 0 a .5 .5 0 1 0 -1 0"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Blob.String() + `__render--r"
			d="M .5 0 V 1 H 1.5 V 0 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Blob.String() + `__render--ur"
			d="M .5 .5 H 1 A .5 .5 0 0 1 1.5 0 V -.5 H 1 A .5 .5 0 0 1 .5 0 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Blob.String() + `__render--d"
			d="M 1 .5 H 0 V 1.5 H 1 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Blob.String() + `__render--dr"
			d="M .5 .5 V 1 A .5 .5 0 0 1 1 1.5 H 1.5 V 1 A .5 .5 0 0 1 1 .5 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
	},
	// TODO: implement LeftLeaf and RightLeaf path symbols
	qrconst.LeftLeaf:  {},
	qrconst.RightLeaf: {},
	qrconst.Diamond: {
		`<path
			id="` + qrconst.Diamond.String() + `__render"
			d="M 1 .5 L .5 0 L 0 .5 L .5 1 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
	},
	qrconst.WaterDroplet: {
		`<path
			id="` + qrconst.WaterDroplet.String() + `__render"
			d="M 0 0 C .1 .1 1 .15 1 .667 Q 1 1 .667 1 C .25 1 .1 .1 0 0"
			fill="rgba(0,0,0,1.)"
		/>`,
	},
	qrconst.Star4: {
		`<path
			id="` + qrconst.Star4.String() + `__render"
			d="M 1 .5 L .67 .33 L .5 0 L .33 .33 L 0 .5 L .33 .67 L .5 1 L .67 .67 L 1 .5"
			fill="rgba(0,0,0,1.)"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="round"
			stroke-linejoin="round"
			stroke-width=".125"
		/>`,
	},
	qrconst.Star5: {
		`<path
			id="` + qrconst.Star5.String() + `__render"
			d="M 1 .5 L .65450849 .38774301 L .6545085 .02447174 L .44098301 .31836437 L .0954915 .20610737 L .309017 .5 L .0954915 .79389263 L .44098301 .68163563 L .6545085 .97552826 L .65450849 .61225699 L 1 .5"
			fill="rgba(0,0,0,1.)"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="round"
			stroke-linejoin="round"
			stroke-width=".125"
		/>`,
	},
	qrconst.Star6: {
		`<path
			id="` + qrconst.Star6.String() + `__render"
			d="M 1 .5 L .74999988 .3556625 L .75 .0669873 L .5 .211325 L .25 .0669873 L .25000012 .3556625 L 0 .5 L .25000012 .6443375 L .25 .9330127 L .5 .788675 L .75 .9330127 L .74999988 .6443375 L 1 .5"
			fill="rgba(0,0,0,1.)"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="round"
			stroke-linejoin="round"
			stroke-width=".125"
		/>`,
	},
	qrconst.Star8: {
		`<path
			id="` + qrconst.Star8.String() + `__render"
			d="M 1 .5 L .85355299 .35355356 L .85355339 .14644661 L .64644644 .14644701 L .5 0 L .35355356 .14644701 L .14644661 .14644661 L .14644701 .35355356 L 0 .5 L .14644701 .64644644 L .14644661 .85355339 L .35355356 .85355299 L .5 1 L .64644644 .85355299 L .85355339 .85355339 L .85355299 .64644644 L 1 .5"
			fill="rgba(0,0,0,1.)"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="round"
			stroke-linejoin="round"
			stroke-width=".125"
		/>`,
	},
	qrconst.Xs: {
		`<path
			id="` + qrconst.Xs.String() + `__structural"
			d="M 1 .5 L .5 0 L 0 .5 L .5 1 Z"
			fill="none"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="round"
			stroke-linejoin="round"
			stroke-width=".025"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__render"
			d="M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 V 1 H .25 L .5 .75 L .75 1 H 1 V .75 L .75 .5 L 1 .25 V 0 H .75 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__render--ur"
			d="M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 V 1 H .25 L .5 .75 L .75 1 H 1 V .75 L .75 .5 L 1.25 0 V -.25 H 1 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__render--r"
			d="M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 V 1 H .25 L .5 .75 L .75 1 H 1.25 L .75 .5 L 1.25 0 H .75 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__render--ur-r"
			d="M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 V 1 H .25 L .5 .75 L .75 1 H 1.25 L .75 .5 L 1.25 0 V -.25 H 1 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__render--dr"
			d="M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 V 1 H .25 L .5 .75 L 1 1.25 H 1.25 V 1 L .75 .5 L 1 .25 V 0 H .75 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__render--ur-dr"
			d="M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 V 1 H .25 L .5 .75 L 1 1.25 H 1.25 V 1 L .75 .5 L 1.25 0 V -.25 H 1 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__render--r-dr"
			d="M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 V 1 H .25 L .5 .75 L 1 1.25 H 1.25 V 1 L .75 .5 L 1.25 0 H .75 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__render--d"
			d="M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 V 1.25 L .5 .75 L 1 1.25 V 1 V .75 L .75 .5 L 1 .25 V 0 H .75 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__render--ur-d"
			d="M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 V 1.25 L .5 .75 L 1 1.25 V 1 V .75 L .75 .5 L 1.25 0 V -.25 H 1 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__render--r-d"
			d="M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 V 1.25 L .5 .75 L 1 1.25 V 1 H 1.25 L .75 .5 L 1.25 0 H .75 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__render--ur-r-d"
			d="M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 V 1.25 L .5 .75 L 1 1.25 V 1 H 1.25 L .75 .5 L 1.25 0 V -.25 H 1 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__render--dr-d"
			d="M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 V 1.25 L .5 .75 L 1 1.25 H 1.25 V 1 L .75 .5 L 1 .25 V 0 H .75 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__render--ur-dr-d"
			d="M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 V 1.25 L .5 .75 L 1 1.25 H 1.25 V 1 L .75 .5 L 1.25 0 V -.25 H 1 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__render--r-dr-d"
			d="M .5 .25 L .25 0 H 0 V .25 L .25 .5 L 0 .75 V 1.25 L .5 .75 L 1 1.25 H 1.25 V 1 L .75 .5 L 1.25 0 H .75 Z"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<g id="` + qrconst.Xs.String() + `__render--structural">
			<use href="#` + qrconst.Xs.String() + `__render" />
			<use href="#` + qrconst.Xs.String() + `__structural" />
		</g>`,
		`<g id="` + qrconst.Xs.String() + `__render--structural-ur">
			<use href="#` + qrconst.Xs.String() + `__render--ur" />
			<use href="#` + qrconst.Xs.String() + `__structural" />
		</g>`,
		`<g id="` + qrconst.Xs.String() + `__render--structural-r">
			<use href="#` + qrconst.Xs.String() + `__render--r" />
			<use href="#` + qrconst.Xs.String() + `__structural" />
		</g>`,
		`<g id="` + qrconst.Xs.String() + `__render--structural-ur-r">
			<use href="#` + qrconst.Xs.String() + `__render--ur-r" />
			<use href="#` + qrconst.Xs.String() + `__structural" />
		</g>`,
		`<g id="` + qrconst.Xs.String() + `__render--structural-dr">
			<use href="#` + qrconst.Xs.String() + `__render--dr" />
			<use href="#` + qrconst.Xs.String() + `__structural" />
		</g>`,
		`<g id="` + qrconst.Xs.String() + `__render--structural-ur-dr">
			<use href="#` + qrconst.Xs.String() + `__render--ur-dr" />
			<use href="#` + qrconst.Xs.String() + `__structural" />
		</g>`,
		`<g id="` + qrconst.Xs.String() + `__render--structural-r-dr">
			<use href="#` + qrconst.Xs.String() + `__render--r-dr" />
			<use href="#` + qrconst.Xs.String() + `__structural" />
		</g>`,
		`<g id="` + qrconst.Xs.String() + `__render--structural-d">
			<use href="#` + qrconst.Xs.String() + `__render--d" />
			<use href="#` + qrconst.Xs.String() + `__structural" />
		</g>`,
		`<g id="` + qrconst.Xs.String() + `__render--structural-ur-d">
			<use href="#` + qrconst.Xs.String() + `__render--ur-d" />
			<use href="#` + qrconst.Xs.String() + `__structural" />
		</g>`,
		`<g id="` + qrconst.Xs.String() + `__render--structural-r-d">
			<use href="#` + qrconst.Xs.String() + `__render--r-d" />
			<use href="#` + qrconst.Xs.String() + `__structural" />
		</g>`,
		`<g id="` + qrconst.Xs.String() + `__render--structural-ur-r-d">
			<use href="#` + qrconst.Xs.String() + `__render--ur-r-d" />
			<use href="#` + qrconst.Xs.String() + `__structural" />
		</g>`,
		`<g id="` + qrconst.Xs.String() + `__render--structural-dr-d">
			<use href="#` + qrconst.Xs.String() + `__render--dr-d" />
			<use href="#` + qrconst.Xs.String() + `__structural" />
		</g>`,
		`<g id="` + qrconst.Xs.String() + `__render--structural-ur-dr-d">
			<use href="#` + qrconst.Xs.String() + `__render--ur-dr-d" />
			<use href="#` + qrconst.Xs.String() + `__structural" />
		</g>`,
		`<g id="` + qrconst.Xs.String() + `__render--structural-r-dr-d">
			<use href="#` + qrconst.Xs.String() + `__render--r-dr-d" />
			<use href="#` + qrconst.Xs.String() + `__structural" />
		</g>`,
		`<path
			id="` + qrconst.Xs.String() + `__merge--ur"
			d="M 1 .5 L .5 0"
			fill="none"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="round"
			stroke-linejoin="round"
			stroke-width=".025"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__merge--ul"
			d="M .5 0 L 0 .5"
			fill="none"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="round"
			stroke-linejoin="round"
			stroke-width=".025"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__merge--dl"
			d="M 0 .5 L .5 1"
			fill="none"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="round"
			stroke-linejoin="round"
			stroke-width=".025"
		/>`,
		`<path
			id="` + qrconst.Xs.String() + `__merge--dr"
			d="M .5 1 L 1 .5"
			fill="none"
			stroke="rgba(0,0,0,1.)"
			stroke-linecap="round"
			stroke-linejoin="round"
			stroke-width=".025"
		/>`,
	},
	// TODO: implement Octagon path symbols
	qrconst.Octagon: {},
	qrconst.SmileyFace: {
		`<mask id="` + qrconst.SmileyFace.String() + `__mask">
			<rect width="100%" height="100%" fill="black" />

			<!-- Face (visible) -->
			<path
				d="M 0.5 0.5 m -0.5 0
				a 0.5 0.5 0 1 0 1 0
				a 0.5 0.5 0 1 0 -1 0"
				fill="white"
			/>

			<!-- Eyes (cut out) -->
			<path
				d="M 0.775 0.333
				a 0.075 0.075 0 1 0 -0.15 0
				a 0.075 0.075 0 1 0 0.15 0"
				fill="black"
			/>
			<path
				d="M 0.225 0.333
				a 0.075 0.075 0 1 0 0.15 0
				a 0.075 0.075 0 1 0 -0.15 0"
				fill="black"
			/>

			<!-- Nose (stroke cuts) -->
			<path
				d="M 0.5 0.45 V 0.55"
				stroke="black"
				stroke-width="0.035"
				stroke-linecap="square"
			/>

			<!-- Mouth (stroke cuts) -->
			<path
				d="M 0.2625 0.65 Q 0.5 0.9 0.7375 0.65"
				stroke="black"
				stroke-width="0.05"
				stroke-linecap="round"
				fill="none"
			/>
		</mask>`,
		`<rect
			id="` + qrconst.SmileyFace.String() + `__render"
			width="1" height="1"
			fill="rgba(0,0,0,1.)"
			mask="url(#` + qrconst.SmileyFace.String() + `__mask)"
		/>`,
	},
	qrconst.Pointillism: {
		`<circle
			id="` + qrconst.Pointillism.String() + `__render--structural"
			cx=".5" cy=".5" r=".5"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__render--0"
			cx=".5" cy=".5" r=".24"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__render--1"
			cx=".5" cy=".5" r=".2932282359721639"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__render--2"
			cx=".5" cy=".5" r=".33858634665347126"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__render--3"
			cx=".5" cy=".5" r=".3772379789497893"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__render--4"
			cx=".5" cy=".5" r=".4101747273445025"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__render--5"
			cx=".5" cy=".5" r=".4382415729178002"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__render--6"
			cx=".5" cy=".5" r=".4621585610489597"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__render--7"
			cx=".5" cy=".5" r=".48253927393570584"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__render--8"
			cx=".5" cy=".5" r=".49990657183685006"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__merge--0"
			cx=".5" cy=".5" r=".06"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__merge--1"
			cx=".5" cy=".5" r=".06698825596989015"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__merge--2"
			cx=".5" cy=".5" r=".0735695475939411"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__merge--3"
			cx=".5" cy=".5" r=".07976757463064736"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__merge--4"
			cx=".5" cy=".5" r=".08560465667201358"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__merge--5"
			cx=".5" cy=".5" r=".09110181351819385"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__merge--6"
			cx=".5" cy=".5" r=".09627884087147627"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__merge--7"
			cx=".5" cy=".5" r=".10115438162219319"
			fill="rgba(0,0,0,1.)"
		/>`,
		`<circle
			id="` + qrconst.Pointillism.String() + `__merge--8"
			cx=".5" cy=".5" r=".10574599298326308"
			fill="rgba(0,0,0,1.)"
		/>`,
	},
}

var PathRenderFunctions = map[qrconst.ModuleShape]func(lookahead qrconst.Lookahead) []string{
	qrconst.Square: func(lookahead qrconst.Lookahead) []string {
		paths := []string{
			use(qrconst.Square, "render", ""),
		}

		R := lookahead.Has(qrconst.LookR)
		D := lookahead.Has(qrconst.LookD)

		paths = applyRules(
			paths, qrconst.Square, "render",
			[]rule{
				{R, "--r"},
				{D, "--d"},
			},
		)

		return paths
	},
	qrconst.Circle: func(lookahead qrconst.Lookahead) []string {
		return []string{
			use(qrconst.Circle, "render", ""),
		}
	},
	qrconst.TiedCircle: func(lookahead qrconst.Lookahead) []string {
		paths := []string{
			use(qrconst.TiedCircle, "render", ""),
		}

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

		paths = applyRules(
			paths, qrconst.TiedCircle, "render",
			[]rule{
				{UR, "--ur"},
				{UL, "--ul"},
				{DL, "--dl"},
				{DR, "--dr"},

				{R2UR, "--r2ur"},
				{UR2U, "--ur2u"},
				{U2UL, "--u2ul"},
				{UL2L, "--ul2l"},
				{L2DL, "--l2dl"},
				{DL2D, "--dl2d"},
				{D2DR, "--d2dr"},
				{DR2R, "--dr2r"},
			},
		)

		return paths
	},
	qrconst.HorizontalBlob: func(lookahead qrconst.Lookahead) []string {
		paths := []string{
			use(qrconst.HorizontalBlob, "render", ""),
		}

		R := lookahead.Has(qrconst.LookR)

		UR := lookahead.Has(qrconst.LookUR) &&
			lookahead.Lacks(qrconst.LookR, qrconst.LookU)
		DR := lookahead.Has(qrconst.LookDR) &&
			lookahead.Lacks(qrconst.LookD, qrconst.LookR)

		paths = applyRules(
			paths, qrconst.HorizontalBlob, "render",
			[]rule{
				{R, "--r"},

				{UR, "--ur"},
				{DR, "--dr"},
			},
		)

		return paths
	},
	qrconst.VerticalBlob: func(lookahead qrconst.Lookahead) []string {
		paths := []string{
			use(qrconst.VerticalBlob, "render", ""),
		}

		D := lookahead.Has(qrconst.LookD)

		DL := lookahead.Has(qrconst.LookDL) &&
			lookahead.Lacks(qrconst.LookL, qrconst.LookD)
		DR := lookahead.Has(qrconst.LookDR) &&
			lookahead.Lacks(qrconst.LookD, qrconst.LookR)

		paths = applyRules(
			paths, qrconst.VerticalBlob, "render",
			[]rule{
				{D, "--d"},

				{DL, "--dl"},
				{DR, "--dr"},
			},
		)

		return paths
	},
	qrconst.Blob: func(lookahead qrconst.Lookahead) []string {
		paths := []string{
			use(qrconst.Blob, "render", ""),
		}

		R := lookahead.Has(qrconst.LookR)
		UR := lookahead.Has(qrconst.LookUR)

		D := lookahead.Has(qrconst.LookD)
		DR := lookahead.Has(qrconst.LookDR)

		paths = applyRules(
			paths, qrconst.Blob, "render",
			[]rule{
				{R, "--r"},
				{UR, "--ur"},
				{D, "--d"},
				{DR, "--dr"},
			},
		)

		return paths
	},
	// TODO: implement LeftLeaf and RightLeaf path rendering functions
	qrconst.LeftLeaf: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.RightLeaf: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.Diamond: func(lookahead qrconst.Lookahead) []string {
		return []string{
			use(qrconst.Diamond, "render", ""),
		}
	},
	qrconst.WaterDroplet: func(lookahead qrconst.Lookahead) []string {
		return []string{
			use(qrconst.WaterDroplet, "render", ""),
		}

	},
	qrconst.Star4: func(lookahead qrconst.Lookahead) []string {
		return []string{
			use(qrconst.Star4, "render", ""),
		}
	},
	qrconst.Star5: func(lookahead qrconst.Lookahead) []string {
		return []string{
			use(qrconst.Star5, "render", ""),
		}
	},
	qrconst.Star6: func(lookahead qrconst.Lookahead) []string {
		return []string{
			use(qrconst.Star6, "render", ""),
		}
	},
	qrconst.Star8: func(lookahead qrconst.Lookahead) []string {
		return []string{
			use(qrconst.Star8, "render", ""),
		}
	},
	qrconst.Xs: func(lookahead qrconst.Lookahead) []string {
		R := lookahead.Has(qrconst.LookR)
		UR := lookahead.Has(qrconst.LookUR)

		D := lookahead.Has(qrconst.LookD)
		DR := lookahead.Has(qrconst.LookDR)

		if lookahead.Has(qrconst.LookStructural) {
			rules := []rule{
				{R && DR && D, "--structural-r-dr-d"},
				{UR && DR && D, "--structural-ur-dr-d"},
				{DR && D, "--structural-dr-d"},
				{UR && R && D, "--structural-ur-r-d"},
				{R && D, "--structural-r-d"},
				{UR && D, "--structural-ur-d"},
				{D, "--structural-d"},

				{R && DR, "--structural-r-dr"},
				{UR && DR, "--structural-ur-dr"},
				{DR, "--structural-dr"},
				{UR && R, "--structural-ur-r"},
				{R, "--structural-r"},
				{UR, "--structural-ur"},
			}

			for _, rule := range rules {
				if rule.cond {
					return []string{
						use(qrconst.Xs, "render", rule.suffix),
					}
				}
			}

			return []string{
				use(qrconst.Xs, "render", "--structural"),
			}
		}

		rules := []rule{
			{R && DR && D, "--r-dr-d"},
			{UR && DR && D, "--ur-dr-d"},
			{DR && D, "--dr-d"},
			{UR && R && D, "--ur-r-d"},
			{R && D, "--r-d"},
			{UR && D, "--ur-d"},
			{D, "--d"},

			{R && DR, "--r-dr"},
			{UR && DR, "--ur-dr"},
			{DR, "--dr"},
			{UR && R, "--ur-r"},
			{R, "--r"},
			{UR, "--ur"},
		}

		for _, rule := range rules {
			if rule.cond {
				return []string{
					use(qrconst.Xs, "render", rule.suffix),
				}
			}
		}

		return []string{
			use(qrconst.Xs, "render", ""),
		}
	},
	// TODO: implement Octagon path rendering functions
	qrconst.Octagon: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.SmileyFace: func(lookahead qrconst.Lookahead) []string {
		return []string{
			use(qrconst.SmileyFace, "render", ""),
		}
	},
	qrconst.Pointillism: func(lookahead qrconst.Lookahead) []string {
		if lookahead.Has(qrconst.LookStructural) {
			return []string{
				use(qrconst.Pointillism, "render", "--structural"),
			}
		}

		neighbors := 0
		mask := qrconst.LookR
		for range 8 {
			if lookahead&mask != 0 {
				neighbors++
			}
			mask <<= 1
		}

		return []string{
			use(qrconst.Pointillism, "render", "--"+strconv.Itoa(neighbors)),
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

		UR := lookahead.Has(qrconst.LookR, qrconst.LookUR, qrconst.LookU)
		UL := lookahead.Has(qrconst.LookU, qrconst.LookUL, qrconst.LookL)
		DL := lookahead.Has(qrconst.LookL, qrconst.LookDL, qrconst.LookD)
		DR := lookahead.Has(qrconst.LookD, qrconst.LookDR, qrconst.LookR)

		paths = applyRules(
			paths, qrconst.TiedCircle, "merge",
			[]rule{
				{UR, "--ur"},
				{UL, "--ul"},
				{DL, "--dl"},
				{DR, "--dr"},
			},
		)

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
	// TODO: implement LeftLeaf and RightLeaf path merging functions (if necessary)
	qrconst.LeftLeaf: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.RightLeaf: func(lookahead qrconst.Lookahead) []string {
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

		paths = applyRules(
			paths, qrconst.Xs, "merge",
			[]rule{
				{UR, "--ur"},
				{UL, "--ul"},
				{DL, "--dl"},
				{DR, "--dr"},
			},
		)

		return paths
	},
	// TODO: implement Octagon and SmileyFace path merging functions (if necessary)
	qrconst.Octagon: func(lookahead qrconst.Lookahead) []string {
		return nil
	},
	qrconst.SmileyFace: func(lookahead qrconst.Lookahead) []string {
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

		return []string{
			use(qrconst.Pointillism, "merge", "--"+strconv.Itoa(neighbors)),
		}
	},
}

func use(shape qrconst.ModuleShape, variant, suffix string) string {
	return `<use href="#` + shape.String() + `__` + variant + suffix + `"/>`
}

func applyRules(
	paths []string,
	shape qrconst.ModuleShape,
	variant string,
	rules []rule,
) []string {
	for _, rule := range rules {
		if rule.cond {
			paths = append(paths, use(shape, variant, rule.suffix))
		}
	}

	return paths
}
