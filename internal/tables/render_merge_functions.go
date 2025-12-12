package tables

import (
	"math"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
)

var ModuleRenderFunctions = map[qrconst.ModuleShape]func(x, y, scale int, lookahead qrconst.Lookahead) bool{
	qrconst.Square: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return true
	},
	qrconst.Circle: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		edgeR2UR :=
			has(lookahead, qrconst.LookU) &&
				lacks(lookahead, qrconst.LookR, qrconst.LookUR) &&
				(x == scale-1 && dy >= 0)
		edgeUR2U :=
			has(lookahead, qrconst.LookR) &&
				lacks(lookahead, qrconst.LookUR, qrconst.LookU) &&
				(y == 0 && dx >= 0)
		edgeU2UL :=
			has(lookahead, qrconst.LookL) &&
				lacks(lookahead, qrconst.LookU, qrconst.LookUL) &&
				(y == 0 && dx <= 0)
		edgeUL2L :=
			has(lookahead, qrconst.LookU) &&
				lacks(lookahead, qrconst.LookUL, qrconst.LookL) &&
				(x == 0 && dy >= 0)
		edgeL2DL :=
			has(lookahead, qrconst.LookD) &&
				lacks(lookahead, qrconst.LookL, qrconst.LookDL) &&
				(x == 0 && dy <= 0)
		edgeDL2D :=
			has(lookahead, qrconst.LookL) &&
				lacks(lookahead, qrconst.LookDL, qrconst.LookD) &&
				(y == scale-1 && dx <= 0)
		edgeD2DR :=
			has(lookahead, qrconst.LookR) &&
				lacks(lookahead, qrconst.LookD, qrconst.LookDR) &&
				(y == scale-1 && dx >= 0)
		edgeDR2R :=
			has(lookahead, qrconst.LookD) &&
				lacks(lookahead, qrconst.LookDR, qrconst.LookR) &&
				(x == scale-1 && dy <= 0)

		if edgeR2UR || edgeUR2U || edgeU2UL || edgeUL2L ||
			edgeL2DL || edgeDL2D || edgeD2DR || edgeDR2R {
			return true
		}

		return euclideanDist(x, y, cx, cy) <= r*r+4
	},
	qrconst.HorizontalBlob: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		right := has(lookahead, qrconst.LookR) && (dx >= 0)
		left := has(lookahead, qrconst.LookL) && (dx <= 0)

		upperRight :=
			has(lookahead, qrconst.LookUR) &&
				lacks(lookahead, qrconst.LookR, qrconst.LookU) &&
				(dx >= 0 && dy >= 0)
		upperLeft :=
			has(lookahead, qrconst.LookUL) &&
				lacks(lookahead, qrconst.LookU, qrconst.LookL) &&
				(dx <= 0 && dy >= 0)
		lowerLeft :=
			has(lookahead, qrconst.LookDL) &&
				lacks(lookahead, qrconst.LookL, qrconst.LookD) &&
				(dx <= 0 && dy <= 0)
		lowerRight :=
			has(lookahead, qrconst.LookDR) &&
				lacks(lookahead, qrconst.LookD, qrconst.LookR) &&
				(dx >= 0 && dy <= 0)

		if right || upperRight || upperLeft ||
			left || lowerLeft || lowerRight {
			return true
		}

		return euclideanDist(x, y, cx, cy) <= r*r+4
	},
	qrconst.VerticalBlob: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		up := has(lookahead, qrconst.LookU) && (dy >= 0)
		down := has(lookahead, qrconst.LookD) && (dy <= 0)

		upperRight :=
			has(lookahead, qrconst.LookUR) &&
				lacks(lookahead, qrconst.LookR, qrconst.LookU) &&
				(dx >= 0 && dy >= 0)
		upperLeft :=
			has(lookahead, qrconst.LookUL) &&
				lacks(lookahead, qrconst.LookU, qrconst.LookL) &&
				(dx <= 0 && dy >= 0)
		lowerLeft :=
			has(lookahead, qrconst.LookDL) &&
				lacks(lookahead, qrconst.LookL, qrconst.LookD) &&
				(dx <= 0 && dy <= 0)
		lowerRight :=
			has(lookahead, qrconst.LookDR) &&
				lacks(lookahead, qrconst.LookD, qrconst.LookR) &&
				(dx >= 0 && dy <= 0)

		if upperRight || up || upperLeft ||
			lowerLeft || down || lowerRight {
			return true
		}

		return euclideanDist(x, y, cx, cy) <= r*r+4
	},
	qrconst.Blob: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		right := has(lookahead, qrconst.LookR) && (dx >= 0)
		up := has(lookahead, qrconst.LookU) && (dy >= 0)
		left := has(lookahead, qrconst.LookL) && (dx <= 0)
		down := has(lookahead, qrconst.LookD) && (dy <= 0)

		upperRight := has(lookahead, qrconst.LookUR) && (dx >= 0 && dy >= 0)
		upperLeft := has(lookahead, qrconst.LookUL) && (dx <= 0 && dy >= 0)
		lowerLeft := has(lookahead, qrconst.LookDL) && (dx <= 0 && dy <= 0)
		lowerRight := has(lookahead, qrconst.LookDR) && (dx >= 0 && dy <= 0)

		if right || upperRight || up || upperLeft ||
			left || lowerLeft || down || lowerRight {
			return true
		}

		return euclideanDist(x, y, cx, cy) <= r*r+4
	},
	qrconst.LeftLeaf: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		r := float64(scale)
		distFromUR := euclideanDist(x, y, r, 0)
		distFromDL := euclideanDist(x, y, 0, r)

		edgeR2UR :=
			has(lookahead, qrconst.LookU) &&
				lacks(lookahead, qrconst.LookR) &&
				(x == scale-1 && distFromDL > r*r+4*r+4)
		edgeUR2U :=
			has(lookahead, qrconst.LookR) &&
				lacks(lookahead, qrconst.LookU) &&
				(y == 0 && distFromDL > r*r+4*r+4)
		edgeL2DL :=
			has(lookahead, qrconst.LookD) &&
				lacks(lookahead, qrconst.LookL) &&
				(x == 0 && distFromUR > r*r+4*r+4)
		edgeDL2D :=
			has(lookahead, qrconst.LookL) &&
				lacks(lookahead, qrconst.LookD) &&
				(y == scale-1 && distFromUR > r*r+4*r+4)

		if edgeR2UR || edgeUR2U || edgeL2DL || edgeDL2D {
			return true
		}

		// upper-right center's POV
		condUR := distFromUR < r*r+4*r+4

		// lower-left center's POV
		condDL := distFromDL < r*r+4*r+4

		// leaf veins patterns
		condLeafVeins := ((x == y) ||
			(x == 1*scale/5 && y < 1*scale/5) || (y == 1*scale/5 && x < 1*scale/5) ||
			(x == 2*scale/5 && y < 2*scale/5) || (y == 2*scale/5 && x < 2*scale/5) ||
			(x == 3*scale/5 && y < 3*scale/5) || (y == 3*scale/5 && x < 3*scale/5) ||
			(x == 4*scale/5 && y < 4*scale/5) || (y == 4*scale/5 && x < 4*scale/5))
		// (x != 0 && y != 0) && (x != scale-1 && y != scale-1)

		return condUR && condDL && !condLeafVeins
	},
	qrconst.RightLeaf: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		r := float64(scale)

		distFromUL := euclideanDist(x, y, 0, 0)
		distFromDR := euclideanDist(x, y, r, r)

		edgeU2UL :=
			has(lookahead, qrconst.LookL) &&
				lacks(lookahead, qrconst.LookU) &&
				(y == 0 && distFromDR > r*r+4*r+4)
		edgeUL2L :=
			has(lookahead, qrconst.LookU) &&
				lacks(lookahead, qrconst.LookL) &&
				(x == 0 && distFromDR > r*r+4*r+4)
		edgeD2DR :=
			has(lookahead, qrconst.LookR) &&
				lacks(lookahead, qrconst.LookD) &&
				(y == scale-1 && distFromUL > r*r+4*r+4)
		edgeDR2R :=
			has(lookahead, qrconst.LookD) &&
				lacks(lookahead, qrconst.LookR) &&
				(x == scale-1 && distFromUL > r*r+4*r+4)

		if edgeU2UL || edgeUL2L || edgeD2DR || edgeDR2R {
			return true
		}

		// upper-left center's POV
		condUL := distFromUL < r*r+4*r+4

		// lower-right center's POV
		condDR := distFromDR < r*r+4*r+4

		// leaf veins patterns
		condLeafVeins := ((x+y == scale-1) ||
			(scale-x == 1*scale/5 && y < 1*scale/5) || (scale-y == 1*scale/5 && x > 1*scale/5) ||
			(scale-x == 2*scale/5 && y < 2*scale/5) || (scale-y == 2*scale/5 && x > 2*scale/5) ||
			(scale-x == 3*scale/5 && y < 3*scale/5) || (scale-y == 3*scale/5 && x > 3*scale/5) ||
			(scale-x == 4*scale/5 && y < 4*scale/5) || (scale-y == 4*scale/5 && x > 4*scale/5))
		// (x != 0 && y != scale-1) && (x != scale-1 && y != 0)

		return condUL && condDR && !condLeafVeins
	},
	qrconst.Diamond: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx

		return manhattanDist(x, y, cx, cy) <= r
	},
	qrconst.WaterDroplet: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		// Parametric function
		// x(t) = at^3 + bt^2 + ct + d
		// y(t) = pt^3 + qt^2
		a := .7 * (.66 * float64(scale))
		b := -.4 * (.66 * float64(scale))
		c := -1 * (.66 * float64(scale))
		d := .72 * (.66 * float64(scale))
		p := 1.125 * (.66 * float64(scale))
		q := 2.8 * (.66 * float64(scale))

		// Convert to cartesian function
		A := (b*p - a*q) / p
		B := (float64(x) - 6.25) - a*(float64(scale-y)-1)/p - d
		M := A*B*p - A*c*q + c*c*p
		N := B * (A*q - c*q)

		f := A*math.Pow(A*A*(float64(scale-y)-1)-N, 2) + c*M*(A*A*(float64(scale-y)-1)-N) - B*M*M

		return f > 0
	},
	qrconst.Octagon: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := math.Abs(float64(x) - cx)
		dy := math.Abs(float64(y) - cy)

		// A square intersect a diamond → octagon
		cond1 := math.Max(dx, dy) <= r
		cond2 := dx+dy <= r*math.Sqrt2

		return cond1 && cond2
	},
	qrconst.SmileyFace: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)
		cxLeftEye := cx - 4
		cyLeftEye := cy - 3
		cxRightEye := cx + 4
		cyRightEye := cy - 3

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		edgeR2UR :=
			has(lookahead, qrconst.LookU) &&
				lacks(lookahead, qrconst.LookR, qrconst.LookUR) &&
				(x == scale-1 && dy >= 0)
		edgeUR2U :=
			has(lookahead, qrconst.LookR) &&
				lacks(lookahead, qrconst.LookUR, qrconst.LookU) &&
				(y == 0 && dx >= 0)
		edgeU2UL :=
			has(lookahead, qrconst.LookL) &&
				lacks(lookahead, qrconst.LookU, qrconst.LookUL) &&
				(y == 0 && dx <= 0)
		edgeUL2L :=
			has(lookahead, qrconst.LookU) &&
				lacks(lookahead, qrconst.LookUL, qrconst.LookL) &&
				(x == 0 && dy >= 0)
		edgeL2DL :=
			has(lookahead, qrconst.LookD) &&
				lacks(lookahead, qrconst.LookL, qrconst.LookDL) &&
				(x == 0 && dy <= 0)
		edgeDL2D :=
			has(lookahead, qrconst.LookL) &&
				lacks(lookahead, qrconst.LookDL, qrconst.LookD) &&
				(y == scale-1 && dx <= 0)
		edgeD2DR :=
			has(lookahead, qrconst.LookR) &&
				lacks(lookahead, qrconst.LookD, qrconst.LookDR) &&
				(y == scale-1 && dx >= 0)
		edgeDR2R :=
			has(lookahead, qrconst.LookD) &&
				lacks(lookahead, qrconst.LookDR, qrconst.LookR) &&
				(x == scale-1 && dy <= 0)

		if edgeR2UR || edgeUR2U || edgeU2UL || edgeUL2L ||
			edgeL2DL || edgeDL2D || edgeD2DR || edgeDR2R {
			return true
		}

		distFromOrigin := euclideanDist(x, y, cx, cy)
		distFromLeftEye := euclideanDist(x, y, cxLeftEye, cyLeftEye)
		distFromRightEye := euclideanDist(x, y, cxRightEye, cyRightEye)
		mouthArcY := cy + 5 - .08*float64(dx*dx)

		face := distFromOrigin < r*r+4
		leftEye := distFromLeftEye < .025*r*r
		rightEye := distFromRightEye < .025*r*r
		nose := (x != int(cx) || (math.Abs(dy) > 1))
		mouth := float64(y) >= mouthArcY-.75 && float64(y) <= mouthArcY+.75 && math.Abs(dx) < 6

		return face &&
			!leftEye &&
			!rightEye &&
			nose &&
			!mouth
	},
	qrconst.Pointillism: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		if has(lookahead, qrconst.LookStructural) {
			return euclideanDist(x, y, cx, cy) <= r*r+9
		}

		neighbors := 0
		mask := qrconst.LookR
		for range 8 {
			if lookahead&mask != 0 {
				neighbors++
			}
			mask <<= 1
		}

		minR := 0.24 * float64(scale)
		maxR := 0.60 * float64(scale)

		k := .18
		t := 1 - math.Exp(-k*float64(neighbors))
		if t > 1 {
			t = 1
		}
		r = minR + (maxR-minR)*t

		return euclideanDist(x, y, cx, cy) <= r*r+9
	},
}

var ModuleMergeFunctions = map[qrconst.ModuleShape]func(x, y, scale int, lookahead qrconst.Lookahead) bool{
	qrconst.Square: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.Circle: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		cornerUR := has(lookahead, qrconst.LookR, qrconst.LookUR, qrconst.LookU) && (dx >= 0 && dy >= 0)
		cornerUL := has(lookahead, qrconst.LookU, qrconst.LookUL, qrconst.LookL) && (dx <= 0 && dy >= 0)
		cornerDL := has(lookahead, qrconst.LookL, qrconst.LookDL, qrconst.LookD) && (dx <= 0 && dy <= 0)
		cornerDR := has(lookahead, qrconst.LookD, qrconst.LookDR, qrconst.LookR) && (dx >= 0 && dy <= 0)

		if cornerUR || cornerUL || cornerDL || cornerDR {
			d2 := euclideanDist(x, y, cx, cy)
			return d2 >= r*r-1.25*r && d2 <= r*r+1.25*r
		}

		return false
	},
	qrconst.HorizontalBlob: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		upperRight :=
			has(lookahead, qrconst.LookU, qrconst.LookR) &&
				lacks(lookahead, qrconst.LookUR) &&
				(dx >= 0 && dy >= 0)
		upperLeft :=
			has(lookahead, qrconst.LookU, qrconst.LookL) &&
				lacks(lookahead, qrconst.LookUL) &&
				(dx <= 0 && dy >= 0)
		lowerLeft :=
			has(lookahead, qrconst.LookL, qrconst.LookD) &&
				lacks(lookahead, qrconst.LookDL) &&
				(dx <= 0 && dy <= 0)
		lowerRight :=
			has(lookahead, qrconst.LookD, qrconst.LookR) &&
				lacks(lookahead, qrconst.LookDR) &&
				(dx >= 0 && dy <= 0)

		if upperRight || upperLeft || lowerLeft || lowerRight {
			return euclideanDist(x, y, cx, cy) > r*r+4
		}

		return false
	},
	qrconst.VerticalBlob: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		upperRight :=
			has(lookahead, qrconst.LookU, qrconst.LookR) &&
				lacks(lookahead, qrconst.LookUR) &&
				(dx >= 0 && dy >= 0)
		upperLeft :=
			has(lookahead, qrconst.LookU, qrconst.LookL) &&
				lacks(lookahead, qrconst.LookUL) &&
				(dx <= 0 && dy >= 0)
		lowerLeft :=
			has(lookahead, qrconst.LookL, qrconst.LookD) &&
				lacks(lookahead, qrconst.LookDL) &&
				(dx <= 0 && dy <= 0)
		lowerRight :=
			has(lookahead, qrconst.LookD, qrconst.LookR) &&
				lacks(lookahead, qrconst.LookDR) &&
				(dx >= 0 && dy <= 0)

		if upperRight || upperLeft || lowerLeft || lowerRight {
			return euclideanDist(x, y, cx, cy) > r*r+4
		}

		return false
	},
	qrconst.Blob: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		upperRight := has(lookahead, qrconst.LookR, qrconst.LookU) && (dx >= 0 && dy >= 0)
		upperLeft := has(lookahead, qrconst.LookU, qrconst.LookL) && (dx <= 0 && dy >= 0)
		lowerLeft := has(lookahead, qrconst.LookL, qrconst.LookD) && (dx <= 0 && dy <= 0)
		lowerRight := has(lookahead, qrconst.LookD, qrconst.LookR) && (dx >= 0 && dy <= 0)

		if upperRight || upperLeft || lowerLeft || lowerRight {
			return euclideanDist(x, y, cx, cy) > r*r+4
		}

		return false
	},
	qrconst.LeftLeaf: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		cornerUR := has(lookahead, qrconst.LookR, qrconst.LookUR, qrconst.LookU) && (dx >= 0 && dy >= 0)
		cornerDL := has(lookahead, qrconst.LookL, qrconst.LookDL, qrconst.LookD) && (dx <= 0 && dy <= 0)

		if cornerUR || cornerDL {
			d2 := euclideanDist(x, y, cx, cy)
			return d2 >= r*r-r+15 && d2 <= r*r+r+15
		}

		return false
	},
	qrconst.RightLeaf: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		cornerUL := has(lookahead, qrconst.LookU, qrconst.LookUL, qrconst.LookL) && (dx <= 0 && dy >= 0)
		cornerDR := has(lookahead, qrconst.LookD, qrconst.LookDR, qrconst.LookR) && (dx >= 0 && dy <= 0)

		if cornerUL || cornerDR {
			d2 := euclideanDist(x, y, cx, cy)
			return d2 >= r*r-r+15 && d2 <= r*r+r+15
		}

		return false
	},
	qrconst.Diamond: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.WaterDroplet: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.Octagon: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.SmileyFace: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		cornerUR := has(lookahead, qrconst.LookR, qrconst.LookUR, qrconst.LookU) && (dx >= 0 && dy >= 0)
		cornerUL := has(lookahead, qrconst.LookU, qrconst.LookUL, qrconst.LookL) && (dx <= 0 && dy >= 0)
		cornerDL := has(lookahead, qrconst.LookL, qrconst.LookDL, qrconst.LookD) && (dx <= 0 && dy <= 0)
		cornerDR := has(lookahead, qrconst.LookD, qrconst.LookDR, qrconst.LookR) && (dx >= 0 && dy <= 0)

		if cornerUR || cornerUL || cornerDL || cornerDR {
			d2 := euclideanDist(x, y, cx, cy)
			return d2 >= r*r-1.25*r && d2 <= r*r+1.25*r
		}

		return false
	},
	qrconst.Pointillism: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		if has(lookahead, qrconst.LookStructural) {
			return false
		}

		neighbors := 0
		mask := qrconst.LookR
		for range 8 {
			if lookahead&mask != 0 {
				neighbors++
			}
			mask <<= 1
		}

		minR := .06 * float64(scale)
		maxR := .18 * float64(scale)

		k := .06
		t := 1 - math.Exp(-k*float64(neighbors))
		r := minR + (maxR-minR)*t

		return euclideanDist(x, y, cx, cy) <= r*r
	},
}

func mid(scale int) float64 {
	return float64(scale-1) / 2
}

func euclideanDist(x, y int, cx, cy float64) float64 {
	dx := float64(x) - cx
	dy := float64(y) - cy
	return dx*dx + dy*dy
}

func manhattanDist(x, y int, cx, cy float64) float64 {
	return math.Abs(float64(x)-cx) + math.Abs(float64(y)-cy)
}

func has(l qrconst.Lookahead, dirs ...qrconst.Lookahead) bool {
	for _, dir := range dirs {
		if l&dir == 0 {
			return false
		}
	}

	return true
}

func lacks(l qrconst.Lookahead, dirs ...qrconst.Lookahead) bool {
	for _, dir := range dirs {
		if l&dir != 0 {
			return false
		}
	}

	return true
}
