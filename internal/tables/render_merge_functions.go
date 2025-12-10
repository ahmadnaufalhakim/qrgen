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

		// upper-right center's POV
		condUR := euclideanDist(x, y, r, 0) < r*r+2*r+1

		// lower-left center's POV
		condDL := euclideanDist(x, y, 0, r) < r*r+2*r+1

		// leaf veins patterns
		condLeafVeins := ((x == y) ||
			(x == 1*scale/5 && y < 1*scale/5) || (y == 1*scale/5 && x < 1*scale/5) ||
			(x == 2*scale/5 && y < 2*scale/5) || (y == 2*scale/5 && x < 2*scale/5) ||
			(x == 3*scale/5 && y < 3*scale/5) || (y == 3*scale/5 && x < 3*scale/5) ||
			(x == 4*scale/5 && y < 4*scale/5) || (y == 4*scale/5 && x < 4*scale/5)) &&
			(x != 0 && y != 0) && (x != scale-1 && y != scale-1)

		return condUR && condDL && !condLeafVeins
	},
	qrconst.RightLeaf: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		r := float64(scale)

		// upper-left center's POV
		condUL := euclideanDist(x, y, 0, 0) < r*r+2*r+1

		// lower-right center's POV
		condDR := euclideanDist(x, y, r, r) < r*r+2*r+1

		// leaf veins patterns
		condLeafVeins := ((x+y == scale-1) ||
			(scale-x == 1*scale/5 && y < 1*scale/5) || (scale-y == 1*scale/5 && x > 1*scale/5) ||
			(scale-x == 2*scale/5 && y < 2*scale/5) || (scale-y == 2*scale/5 && x > 2*scale/5) ||
			(scale-x == 3*scale/5 && y < 3*scale/5) || (scale-y == 3*scale/5 && x > 3*scale/5) ||
			(scale-x == 4*scale/5 && y < 4*scale/5) || (scale-y == 4*scale/5 && x > 4*scale/5)) &&
			(x != 0 && y != scale-1) && (x != scale-1 && y != 0)

		return condUL && condDR && !condLeafVeins
	},
	qrconst.Diamond: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx

		return manhattanDist(x, y, cx, cy) <= r
	},
	qrconst.DropletDown: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := float64(y) - cy

		// top half: circle cap
		if dy < 0 {
			return dx*dx+dy*dy <= r*r*0.80
		}

		// bottom half: more curved
		dy *= 1.3
		return dx*dx+dy*dy <= r*r
	},
	qrconst.DropletUp: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := float64(y) - cy

		// bottom half: circle cap
		if dy > 0 {
			return dx*dx+dy*dy <= r*r*0.80
		}

		// top: curved
		dy *= 1.3
		return dx*dx+dy*dy <= r*r
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
			return d2 >= r*r-r+.25 && d2 <= r*r+r+.25
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
		return false
	},
	qrconst.RightLeaf: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.Diamond: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.DropletDown: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.DropletUp: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.Octagon: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
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
