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

		return euclideanDist(x, y, cx, cy) <= r*r+4
	},
	qrconst.HorizontalBlob: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx

		if lookahead&qrconst.LookL != 0 && dx <= 0 {
			return true
		}
		if lookahead&qrconst.LookR != 0 && dx >= 0 {
			return true
		}

		return euclideanDist(x, y, cx, cy) <= r*r+4
	},
	qrconst.VerticalBlob: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dy := cy - float64(y)

		up := (lookahead&qrconst.LookU != 0) && (dy >= 0)
		down := (lookahead&qrconst.LookD != 0) && (dy <= 0)

		if up || down {
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

		right := (lookahead&qrconst.LookR != 0) && (dx >= 0)
		up := (lookahead&qrconst.LookU != 0) && (dy >= 0)
		left := (lookahead&qrconst.LookL != 0) && (dx <= 0)
		down := (lookahead&qrconst.LookD != 0) && (dy <= 0)

		upperRight := (lookahead&qrconst.LookUR != 0) && (dx >= 0 && dy >= 0)
		upperLeft := (lookahead&qrconst.LookUL != 0) && (dx <= 0 && dy >= 0)
		lowerLeft := (lookahead&qrconst.LookDL != 0) && (dx <= 0 && dy <= 0)
		lowerRight := (lookahead&qrconst.LookDR != 0) && (dx >= 0 && dy <= 0)

		if right || upperRight || up || upperLeft ||
			left || lowerLeft || down || lowerRight {
			return true
		}

		return euclideanDist(x, y, cx, cy) <= r*r+4
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
		return false
	},
	qrconst.HorizontalBlob: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.VerticalBlob: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		return false
	},
	qrconst.Blob: func(x, y, scale int, lookahead qrconst.Lookahead) bool {
		cx := mid(scale)
		cy := mid(scale)

		r := cx
		dx := float64(x) - cx
		dy := cy - float64(y)

		upperRight := lookahead&(qrconst.LookR|qrconst.LookU) == (qrconst.LookR | qrconst.LookU)
		upperLeft := lookahead&(qrconst.LookU|qrconst.LookL) == (qrconst.LookU | qrconst.LookL)
		lowerLeft := lookahead&(qrconst.LookL|qrconst.LookD) == (qrconst.LookL | qrconst.LookD)
		lowerRight := lookahead&(qrconst.LookD|qrconst.LookR) == (qrconst.LookD | qrconst.LookR)

		upperRight = upperRight && (dx >= 0 && dy >= 0)
		upperLeft = upperLeft && (dx <= 0 && dy >= 0)
		lowerLeft = lowerLeft && (dx <= 0 && dy <= 0)
		lowerRight = lowerRight && (dx >= 0 && dy <= 0)

		if upperRight || upperLeft || lowerLeft || lowerRight {
			return euclideanDist(x, y, cx, cy) > r*r+3*r+2.25
		}

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
