package tables

import (
	"math"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
)

var PixelRenderFunctions = map[qrconst.ModuleShape]func(x, y, scale int, lookahead qrconst.Lookahead) bool{
	qrconst.Square: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return true
	},
	qrconst.Circle: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx

		return euclideanDist(x, y, cx, cy) <= r*r+4
	},
	qrconst.TiedCircle: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		edgeR2UR :=
			lookahead.Has(qrconst.LookU) &&
				lookahead.Lacks(qrconst.LookR, qrconst.LookUR) &&
				(x == scale-1 && dy >= 0)
		edgeUR2U :=
			lookahead.Has(qrconst.LookR) &&
				lookahead.Lacks(qrconst.LookUR, qrconst.LookU) &&
				(y == 0 && dx >= 0)
		edgeU2UL :=
			lookahead.Has(qrconst.LookL) &&
				lookahead.Lacks(qrconst.LookU, qrconst.LookUL) &&
				(y == 0 && dx <= 0)
		edgeUL2L :=
			lookahead.Has(qrconst.LookU) &&
				lookahead.Lacks(qrconst.LookUL, qrconst.LookL) &&
				(x == 0 && dy >= 0)
		edgeL2DL :=
			lookahead.Has(qrconst.LookD) &&
				lookahead.Lacks(qrconst.LookL, qrconst.LookDL) &&
				(x == 0 && dy <= 0)
		edgeDL2D :=
			lookahead.Has(qrconst.LookL) &&
				lookahead.Lacks(qrconst.LookDL, qrconst.LookD) &&
				(y == scale-1 && dx <= 0)
		edgeD2DR :=
			lookahead.Has(qrconst.LookR) &&
				lookahead.Lacks(qrconst.LookD, qrconst.LookDR) &&
				(y == scale-1 && dx >= 0)
		edgeDR2R :=
			lookahead.Has(qrconst.LookD) &&
				lookahead.Lacks(qrconst.LookDR, qrconst.LookR) &&
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

		right := lookahead.Has(qrconst.LookR) && (dx >= 0)
		left := lookahead.Has(qrconst.LookL) && (dx <= 0)

		upperRight :=
			lookahead.Has(qrconst.LookUR) &&
				lookahead.Lacks(qrconst.LookR, qrconst.LookU) &&
				(dx >= 0 && dy >= 0)
		upperLeft :=
			lookahead.Has(qrconst.LookUL) &&
				lookahead.Lacks(qrconst.LookU, qrconst.LookL) &&
				(dx <= 0 && dy >= 0)
		lowerLeft :=
			lookahead.Has(qrconst.LookDL) &&
				lookahead.Lacks(qrconst.LookL, qrconst.LookD) &&
				(dx <= 0 && dy <= 0)
		lowerRight :=
			lookahead.Has(qrconst.LookDR) &&
				lookahead.Lacks(qrconst.LookD, qrconst.LookR) &&
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

		up := lookahead.Has(qrconst.LookU) && (dy >= 0)
		down := lookahead.Has(qrconst.LookD) && (dy <= 0)

		upperRight :=
			lookahead.Has(qrconst.LookUR) &&
				lookahead.Lacks(qrconst.LookR, qrconst.LookU) &&
				(dx >= 0 && dy >= 0)
		upperLeft :=
			lookahead.Has(qrconst.LookUL) &&
				lookahead.Lacks(qrconst.LookU, qrconst.LookL) &&
				(dx <= 0 && dy >= 0)
		lowerLeft :=
			lookahead.Has(qrconst.LookDL) &&
				lookahead.Lacks(qrconst.LookL, qrconst.LookD) &&
				(dx <= 0 && dy <= 0)
		lowerRight :=
			lookahead.Has(qrconst.LookDR) &&
				lookahead.Lacks(qrconst.LookD, qrconst.LookR) &&
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

		right := lookahead.Has(qrconst.LookR) && (dx >= 0)
		up := lookahead.Has(qrconst.LookU) && (dy >= 0)
		left := lookahead.Has(qrconst.LookL) && (dx <= 0)
		down := lookahead.Has(qrconst.LookD) && (dy <= 0)

		upperRight := lookahead.Has(qrconst.LookUR) && (dx >= 0 && dy >= 0)
		upperLeft := lookahead.Has(qrconst.LookUL) && (dx <= 0 && dy >= 0)
		lowerLeft := lookahead.Has(qrconst.LookDL) && (dx <= 0 && dy <= 0)
		lowerRight := lookahead.Has(qrconst.LookDR) && (dx >= 0 && dy <= 0)

		if right || upperRight || up || upperLeft ||
			left || lowerLeft || down || lowerRight {
			return true
		}

		return euclideanDist(x, y, cx, cy) <= r*r+4
	},
	qrconst.LeftMandorla: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		R_U := lookahead.HasAny(qrconst.LookR, qrconst.LookU) &&
			(dx >= 0 || dy >= 0)
		L_D := lookahead.HasAny(qrconst.LookL, qrconst.LookD) &&
			(dx <= 0 || dy <= 0)

		if R_U || L_D {
			return true
		} else if (dx >= 0 && dy >= 0) || (dx <= 0 && dy <= 0) {
			return euclideanDist(x, y, cx, cy) < r*r
		} else {
			return true
		}
	},
	qrconst.RightMandorla: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		U_L := lookahead.HasAny(qrconst.LookU, qrconst.LookL) &&
			(dx <= 0 || dy >= 0)
		D_R := lookahead.HasAny(qrconst.LookD, qrconst.LookR) &&
			(dx >= 0 || dy <= 0)

		if U_L || D_R {
			return true
		} else if (dx <= 0 && dy >= 0) || (dx >= 0 && dy <= 0) {
			return euclideanDist(x, y, cx, cy) < r*r
		} else {
			return true
		}
	},
	qrconst.LeftLeaf: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		r := float64(scale)
		distFromUR := euclideanDist(x, y, r, 0)
		distFromDL := euclideanDist(x, y, 0, r)

		// edgeR2UR :=
		// 	lookahead.Has(qrconst.LookU) &&
		// 		lookahead.Lacks(qrconst.LookR) &&
		// 		(x == scale-1 && distFromDL > r*r+4*r+4)
		// edgeUR2U :=
		// 	lookahead.Has(qrconst.LookR) &&
		// 		lookahead.Lacks(qrconst.LookU) &&
		// 		(y == 0 && distFromDL > r*r+4*r+4)
		// edgeL2DL :=
		// 	lookahead.Has(qrconst.LookD) &&
		// 		lookahead.Lacks(qrconst.LookL) &&
		// 		(x == 0 && distFromUR > r*r+4*r+4)
		// edgeDL2D :=
		// 	lookahead.Has(qrconst.LookL) &&
		// 		lookahead.Lacks(qrconst.LookD) &&
		// 		(y == scale-1 && distFromUR > r*r+4*r+4)

		// if edgeR2UR || edgeUR2U || edgeL2DL || edgeDL2D {
		// 	return true
		// }

		// upper-right center's POV
		condUR := distFromUR < r*r+4*r+4

		// lower-left center's POV
		condDL := distFromDL < r*r+4*r+4

		// leaf veins patterns
		condLeafVeins := (x == y)
		if lookahead.Has(qrconst.LookStructural) {
			condLeafVeins = (condLeafVeins ||
				(x == 2*scale/5 && y < 2*scale/5) || (y == 2*scale/5 && x < 2*scale/5) ||
				(x == 7*scale/10 && y < 7*scale/10) || (y == 7*scale/10 && x < 7*scale/10))
		} else {
			condLeafVeins = (condLeafVeins ||
				(x == 1*scale/5 && y < 1*scale/5) || (y == 1*scale/5 && x < 1*scale/5) ||
				(x == 2*scale/5 && y < 2*scale/5) || (y == 2*scale/5 && x < 2*scale/5) ||
				(x == 3*scale/5 && y < 3*scale/5) || (y == 3*scale/5 && x < 3*scale/5) ||
				(x == 4*scale/5 && y < 4*scale/5) || (y == 4*scale/5 && x < 4*scale/5))
			// (x != 0 && y != 0) && (x != scale-1 && y != scale-1)
		}

		return condUR && condDL && !condLeafVeins
	},
	qrconst.RightLeaf: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		r := float64(scale)

		distFromUL := euclideanDist(x, y, 0, 0)
		distFromDR := euclideanDist(x, y, r, r)

		// edgeU2UL :=
		// 	lookahead.Has(qrconst.LookL) &&
		// 		lookahead.Lacks(qrconst.LookU) &&
		// 		(y == 0 && distFromDR > r*r+4*r+4)
		// edgeUL2L :=
		// 	lookahead.Has(qrconst.LookU) &&
		// 		lookahead.Lacks(qrconst.LookL) &&
		// 		(x == 0 && distFromDR > r*r+4*r+4)
		// edgeD2DR :=
		// 	lookahead.Has(qrconst.LookR) &&
		// 		lookahead.Lacks(qrconst.LookD) &&
		// 		(y == scale-1 && distFromUL > r*r+4*r+4)
		// edgeDR2R :=
		// 	lookahead.Has(qrconst.LookD) &&
		// 		lookahead.Lacks(qrconst.LookR) &&
		// 		(x == scale-1 && distFromUL > r*r+4*r+4)

		// if edgeU2UL || edgeUL2L || edgeD2DR || edgeDR2R {
		// 	return true
		// }

		// upper-left center's POV
		condUL := distFromUL < r*r+4*r+4

		// lower-right center's POV
		condDR := distFromDR < r*r+4*r+4

		// leaf veins patterns
		condLeafVeins := (x+y == scale-1)
		if lookahead.Has(qrconst.LookStructural) {
			condLeafVeins = (condLeafVeins ||
				(scale-x == 2*scale/5+1 && y < 2*scale/5) || (scale-y == 2*scale/5-1 && x >= 2*scale/5-1) ||
				(scale-x == 7*scale/10+1 && y < 7*scale/10) || (scale-y == 7*scale/10-1 && x >= 7*scale/10-1))
		} else {
			condLeafVeins = (condLeafVeins ||
				(scale-x == 1*scale/5+2 && y <= 1*scale/5) || (scale-y == 1*scale/5 && x >= 1*scale/5-1) ||
				(scale-x == 2*scale/5+2 && y <= 2*scale/5) || (scale-y == 2*scale/5 && x >= 2*scale/5-1) ||
				(scale-x == 3*scale/5+2 && y <= 3*scale/5) || (scale-y == 3*scale/5 && x >= 3*scale/5-1) ||
				(scale-x == 4*scale/5+2 && y <= 4*scale/5) || (scale-y == 4*scale/5 && x >= 4*scale/5-1))
			// (x != 0 && y != 0) && (x != scale-1 && y != scale-1)
		}

		return condUL && condDR && !condLeafVeins
	},
	qrconst.Diamond: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return isInPolygon(x, y, scale, 4)
	},
	qrconst.Pentagon: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return isInPolygon(x, y, scale, 5)
	},
	qrconst.Hexagon: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return isInPolygon(x, y, scale, 6)
	},
	qrconst.Octagon: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return isInPolygon(x, y, scale, 8)
	},
	qrconst.Star4: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return isInStar(x, y, scale, 4)
	},
	qrconst.Star5: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return isInStar(x, y, scale, 5)
	},
	qrconst.Star6: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return isInStar(x, y, scale, 6)
	},
	qrconst.Star8: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return isInStar(x, y, scale, 8)
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
	qrconst.Xs: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		forwardSlash :=
			(float64(x+y) >= 1.5*float64(scale)/2-1) && (float64(x+y) <= 2.5*float64(scale)/2-1)

		backwardSlash :=
			(float64(x-y) >= -.5*float64(scale)/2 && float64(x-y) <= .5*float64(scale)/2)

		right :=
			lookahead.Lacks(qrconst.LookR) &&
				lookahead.HasAny(qrconst.LookDR, qrconst.LookUR) &&
				(x == scale-1 && y == scale/2)
		up :=
			lookahead.Lacks(qrconst.LookU) &&
				lookahead.HasAny(qrconst.LookUR, qrconst.LookUL) &&
				(x == scale/2 && y == 0)
		left :=
			lookahead.Lacks(qrconst.LookL) &&
				lookahead.HasAny(qrconst.LookUL, qrconst.LookDL) &&
				(x == 0 && y == scale/2)
		down :=
			lookahead.Lacks(qrconst.LookD) &&
				lookahead.HasAny(qrconst.LookDL, qrconst.LookDR) &&
				(x == scale/2 && y == scale-1)

		if lookahead.Has(qrconst.LookStructural) {
			forwardSlash = forwardSlash ||
				(x+y == 1*scale/2) || (x+y == 3*scale/2-1)
			backwardSlash = backwardSlash ||
				(x-y == -1*scale/2) || (x-y == 1*scale/2)
		}

		return forwardSlash || backwardSlash ||
			right || up || left || down
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

		// edgeR2UR :=
		// 	lookahead.Has(qrconst.LookU) &&
		// 		lookahead.Lacks(qrconst.LookR, qrconst.LookUR) &&
		// 		(x == scale-1 && dy >= 0)
		// edgeUR2U :=
		// 	lookahead.Has(qrconst.LookR) &&
		// 		lookahead.Lacks(qrconst.LookUR, qrconst.LookU) &&
		// 		(y == 0 && dx >= 0)
		// edgeU2UL :=
		// 	lookahead.Has(qrconst.LookL) &&
		// 		lookahead.Lacks(qrconst.LookU, qrconst.LookUL) &&
		// 		(y == 0 && dx <= 0)
		// edgeUL2L :=
		// 	lookahead.Has(qrconst.LookU) &&
		// 		lookahead.Lacks(qrconst.LookUL, qrconst.LookL) &&
		// 		(x == 0 && dy >= 0)
		// edgeL2DL :=
		// 	lookahead.Has(qrconst.LookD) &&
		// 		lookahead.Lacks(qrconst.LookL, qrconst.LookDL) &&
		// 		(x == 0 && dy <= 0)
		// edgeDL2D :=
		// 	lookahead.Has(qrconst.LookL) &&
		// 		lookahead.Lacks(qrconst.LookDL, qrconst.LookD) &&
		// 		(y == scale-1 && dx <= 0)
		// edgeD2DR :=
		// 	lookahead.Has(qrconst.LookR) &&
		// 		lookahead.Lacks(qrconst.LookD, qrconst.LookDR) &&
		// 		(y == scale-1 && dx >= 0)
		// edgeDR2R :=
		// 	lookahead.Has(qrconst.LookD) &&
		// 		lookahead.Lacks(qrconst.LookDR, qrconst.LookR) &&
		// 		(x == scale-1 && dy <= 0)

		// if edgeR2UR || edgeUR2U || edgeU2UL || edgeUL2L ||
		// 	edgeL2DL || edgeDL2D || edgeD2DR || edgeDR2R {
		// 	return true
		// }

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
		if lookahead.Has(qrconst.LookStructural) {
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

		k := .16
		t := 1 - math.Exp(-k*float64(neighbors))
		if t > 1 {
			t = 1
		}
		r = minR + (maxR-minR)*t

		return euclideanDist(x, y, cx, cy) <= r*r+9
	},
}

var PixelMergeFunctions = map[qrconst.ModuleShape]func(x, y, scale int, lookahead qrconst.Lookahead) bool{
	qrconst.Square: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.Circle: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.TiedCircle: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		cornerUR := lookahead.Has(qrconst.LookR, qrconst.LookUR, qrconst.LookU) && (dx >= 0 && dy >= 0)
		cornerUL := lookahead.Has(qrconst.LookU, qrconst.LookUL, qrconst.LookL) && (dx <= 0 && dy >= 0)
		cornerDL := lookahead.Has(qrconst.LookL, qrconst.LookDL, qrconst.LookD) && (dx <= 0 && dy <= 0)
		cornerDR := lookahead.Has(qrconst.LookD, qrconst.LookDR, qrconst.LookR) && (dx >= 0 && dy <= 0)

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
			lookahead.Has(qrconst.LookU, qrconst.LookR) &&
				lookahead.Lacks(qrconst.LookUR) &&
				(dx >= 0 && dy >= 0)
		upperLeft :=
			lookahead.Has(qrconst.LookU, qrconst.LookL) &&
				lookahead.Lacks(qrconst.LookUL) &&
				(dx <= 0 && dy >= 0)
		lowerLeft :=
			lookahead.Has(qrconst.LookL, qrconst.LookD) &&
				lookahead.Lacks(qrconst.LookDL) &&
				(dx <= 0 && dy <= 0)
		lowerRight :=
			lookahead.Has(qrconst.LookD, qrconst.LookR) &&
				lookahead.Lacks(qrconst.LookDR) &&
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
			lookahead.Has(qrconst.LookU, qrconst.LookR) &&
				lookahead.Lacks(qrconst.LookUR) &&
				(dx >= 0 && dy >= 0)
		upperLeft :=
			lookahead.Has(qrconst.LookU, qrconst.LookL) &&
				lookahead.Lacks(qrconst.LookUL) &&
				(dx <= 0 && dy >= 0)
		lowerLeft :=
			lookahead.Has(qrconst.LookL, qrconst.LookD) &&
				lookahead.Lacks(qrconst.LookDL) &&
				(dx <= 0 && dy <= 0)
		lowerRight :=
			lookahead.Has(qrconst.LookD, qrconst.LookR) &&
				lookahead.Lacks(qrconst.LookDR) &&
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

		upperRight := lookahead.Has(qrconst.LookR, qrconst.LookU) && (dx >= 0 && dy >= 0)
		upperLeft := lookahead.Has(qrconst.LookU, qrconst.LookL) && (dx <= 0 && dy >= 0)
		lowerLeft := lookahead.Has(qrconst.LookL, qrconst.LookD) && (dx <= 0 && dy <= 0)
		lowerRight := lookahead.Has(qrconst.LookD, qrconst.LookR) && (dx >= 0 && dy <= 0)

		if upperRight || upperLeft || lowerLeft || lowerRight {
			return euclideanDist(x, y, cx, cy) > r*r+4
		}

		return false
	},
	qrconst.LeftMandorla: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.RightMandorla: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.LeftLeaf: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		// cx := mid(scale)
		// cy := mid(scale)

		// r := cx
		// dx := float64(x) - cx
		// dy := cy - float64(y)

		// cornerUR := lookahead.Has(qrconst.LookR, qrconst.LookUR, qrconst.LookU) && (dx >= 0 && dy >= 0)
		// cornerDL := lookahead.Has(qrconst.LookL, qrconst.LookDL, qrconst.LookD) && (dx <= 0 && dy <= 0)

		// if cornerUR || cornerDL {
		// 	d2 := euclideanDist(x, y, cx, cy)
		// 	return d2 >= r*r-r+15 && d2 <= r*r+r+15
		// }

		return false
	},
	qrconst.RightLeaf: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		// cx := mid(scale)
		// cy := mid(scale)

		// r := cx
		// dx := float64(x) - cx
		// dy := cy - float64(y)

		// cornerUL := lookahead.Has(qrconst.LookU, qrconst.LookUL, qrconst.LookL) && (dx <= 0 && dy >= 0)
		// cornerDR := lookahead.Has(qrconst.LookD, qrconst.LookDR, qrconst.LookR) && (dx >= 0 && dy <= 0)

		// if cornerUL || cornerDR {
		// 	d2 := euclideanDist(x, y, cx, cy)
		// 	return d2 >= r*r-r+15 && d2 <= r*r+r+15
		// }

		return false
	},
	qrconst.Diamond: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.Pentagon: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.Hexagon: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.Octagon: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.Star4: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.Star5: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.Star6: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.Star8: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.WaterDroplet: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.Xs: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		upperRightFill := float64(x-y) >= 1.5*float64(scale)/2 && float64(x-y) <= 2*float64(scale)/2
		upperLeftFill := float64(x+y) >= 0*float64(scale)/2 && float64(x+y) <= 0.5*float64(scale)/2-1
		lowerLeftFill := float64(x-y) >= -2*float64(scale)/2 && float64(x-y) <= -1.5*float64(scale)/2
		lowerRightFill := float64(x+y) >= 3.5*float64(scale)/2-1 && float64(x+y) <= 4*float64(scale)/2

		if !lookahead.HasAny(qrconst.LookFinder, qrconst.LookSeparator) {
			upperRightFill = upperRightFill || x-y == 1*scale/2+1
			upperLeftFill = upperLeftFill || x+y == 1*scale/2-1
			lowerLeftFill = lowerLeftFill || x-y == -1*scale/2-1
			lowerRightFill = lowerRightFill || x+y == 3*scale/2
		}

		upperRight := lookahead.Has(qrconst.LookR, qrconst.LookU) &&
			upperRightFill
		upperLeft := lookahead.Has(qrconst.LookU, qrconst.LookL) &&
			upperLeftFill
		lowerLeft := lookahead.Has(qrconst.LookL, qrconst.LookD) &&
			lowerLeftFill
		lowerRight := lookahead.Has(qrconst.LookD, qrconst.LookR) &&
			lowerRightFill

		return upperRight || upperLeft || lowerLeft || lowerRight
	},
	qrconst.SmileyFace: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		// cx := mid(scale)
		// cy := mid(scale)

		// r := cx
		// dx := float64(x) - cx
		// dy := cy - float64(y)

		// cornerUR := lookahead.Has(qrconst.LookR, qrconst.LookUR, qrconst.LookU) && (dx >= 0 && dy >= 0)
		// cornerUL := lookahead.Has(qrconst.LookU, qrconst.LookUL, qrconst.LookL) && (dx <= 0 && dy >= 0)
		// cornerDL := lookahead.Has(qrconst.LookL, qrconst.LookDL, qrconst.LookD) && (dx <= 0 && dy <= 0)
		// cornerDR := lookahead.Has(qrconst.LookD, qrconst.LookDR, qrconst.LookR) && (dx >= 0 && dy <= 0)

		// if cornerUR || cornerUL || cornerDL || cornerDR {
		// 	d2 := euclideanDist(x, y, cx, cy)
		// 	return d2 >= r*r-1.25*r && d2 <= r*r+1.25*r
		// }

		return false
	},
	qrconst.Pointillism: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		if lookahead.HasAny(qrconst.LookFinder, qrconst.LookSeparator) {
			return false
		}

		cx := mid(scale)
		cy := mid(scale)

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

func isInPolygon(x, y, scale, n int) bool {
	cx := mid(scale)
	cy := mid(scale)

	r := cx

	dx := float64(x) - cx
	dy := cy - float64(y)

	// polar coordinates
	rho := math.Sqrt(euclideanDist(x, y, cx, cy))
	if rho > r {
		return false
	}
	theta := math.Atan2(dy, dx) + math.Pi/float64(n) - math.Pi/2
	if theta < 0 {
		theta += 2 * math.Pi
	}

	// sector angle
	alpha := 2 * math.Pi / float64(n)
	// reduce angle to first sector
	thetaPrime := math.Mod(theta+alpha/2, alpha) - alpha/2
	// apothem
	apothem := (r + .0625) * math.Cos(math.Pi/float64(n))
	// max allowed radius at this angle
	rhoMax := apothem / math.Cos(thetaPrime)

	return rho <= rhoMax
}

func isInStar(x, y, scale, n int) bool {
	cx := mid(scale)
	cy := mid(scale)

	dx := float64(x) - cx
	dy := cy - float64(y)

	d2 := euclideanDist(x, y, cx, cy)

	angle := math.Atan2(dy, dx)
	rBase := float64(scale) * 0.33
	starRadius := rBase * (1 + 0.25*math.Cos(float64(n)*angle))

	return math.Sqrt(d2) <= starRadius*1.25
}
