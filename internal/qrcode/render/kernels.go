package render

import (
	"math"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
)

var Kernels = map[qrconst.KernelType]struct {
	KernelFunc    func(radius int) []float64
	DefaultRadius int
}{
	qrconst.KernelLanczos2: {
		Lanczos2Kernel, 3,
	},
	qrconst.KernelCubicSmooth: {
		CubicSmoothKernel, 1,
	},
	qrconst.KernelGaussian: {
		GaussianKernel, 2,
	},
	qrconst.KernelLanczos3: {
		Lanczos3Kernel, 4,
	},
	qrconst.KernelHann: {
		HannKernel, 2,
	},
	qrconst.KernelTriangle: {
		TriangleKernel, 1,
	},
	qrconst.KernelCosine: {
		CosineKernel, 2,
	},
	qrconst.KernelEpanechnikov: {
		EpanechnikovKernel, 2,
	},
	qrconst.KernelBSpline: {
		BSplineKernel, 2,
	},
	qrconst.KernelBox: {
		BoxKernel, 1,
	},
}

// radius: 1(default)-2
func BoxKernel(radius int) []float64 {
	size := 2*radius + 1
	kernel := make([]float64, size)

	w := 1.0 / float64(size)
	for i := range kernel {
		kernel[i] = w
	}

	return kernel
}

// radius: 1(default)-2
func TriangleKernel(radius int) []float64 {
	size := 2*radius + 1
	kernel := make([]float64, size)

	var sum float64
	for i := -radius; i <= radius; i++ {
		val := float64(radius + 1 - absInt(i))
		kernel[i+radius] = val
		sum += val
	}

	for i := range kernel {
		kernel[i] /= sum
	}

	return kernel
}

// radius = 1(default)-3
func CubicSmoothKernel(radius int) []float64 {
	size := 2*radius + 1
	kernel := make([]float64, size)

	var sum float64
	for i := -radius; i <= radius; i++ {
		x := float64(absInt(i)) / float64(radius)
		if x > 1 {
			kernel[i+radius] = 0
			continue
		}

		val := (1.0 / 6.0) * (4 - 6*x*x + 3*x*x*x)
		kernel[i+radius] = val
		sum += val
	}

	for i := range kernel {
		kernel[i] /= sum
	}

	return kernel
}

// radius = 2(default), 3
func BSplineKernel(radius int) []float64 {
	size := 2*radius + 1
	kernel := make([]float64, size)

	var sum float64
	for i := -radius; i <= radius; i++ {
		x := float64(absInt(i)) / float64(radius)
		if x >= 1 {
			kernel[i+radius] = 0
		} else {
			val := 1 - x*x*x
			kernel[i+radius] = val
			sum += val
		}
	}

	for i := range kernel {
		kernel[i] /= sum
	}

	return kernel
}

// radius = 2(default)-4
func HannKernel(radius int) []float64 {
	size := 2*radius + 1
	kernel := make([]float64, size)

	var sum float64
	for i := range size {
		val := 0.5 * (1 - math.Cos(2*math.Pi*float64(i)/float64(size-1)))
		kernel[i] = val
		sum += val
	}

	for i := range kernel {
		kernel[i] /= sum
	}

	return kernel
}

// radius = 3(default), 4
func Lanczos2Kernel(radius int) []float64 {
	a := 2.0 // Lanczos parameter
	size := 2*radius + 1
	kernel := make([]float64, size)

	var sum float64
	for i := -radius; i <= radius; i++ {
		x := float64(i)

		// Lanczos window has support only inside [-a, a]
		if math.Abs(x) > a {
			kernel[i+radius] = 0
			continue
		}

		cutoff := float64(radius) / 2.0
		lp := sinc(x / cutoff) // low-pass sinc term
		window := sinc(x / a)  // Lanczos window

		val := lp * window
		kernel[i+radius] = val
		sum += val
	}

	for i := range kernel {
		kernel[i] /= sum
	}

	return kernel
}

// radius = 4(default), 5
func Lanczos3Kernel(radius int) []float64 {
	a := 3.0 // Lanczos parameter
	size := 2*radius + 1
	kernel := make([]float64, size)

	var sum float64
	for i := -radius; i <= radius; i++ {
		x := float64(i)

		// Lanczos window has support only inside [-a, a]
		if math.Abs(x) > a {
			kernel[i+radius] = 0
			continue
		}

		cutoff := float64(radius) / 2.0
		lp := sinc(x / cutoff) // low-pass sinc term
		window := sinc(x / a)  // Lanczos window

		val := lp * window
		kernel[i+radius] = val
		sum += val
	}

	for i := range kernel {
		kernel[i] /= sum
	}

	return kernel
}

// radius = 2(default)-3
func CosineKernel(radius int) []float64 {
	size := 2*radius + 1
	kernel := make([]float64, size)

	var sum float64
	for i := -radius; i <= radius; i++ {
		x := float64(i) / float64(radius)

		val := math.Cos(0.5 * math.Pi * x)
		kernel[i+radius] = val
		sum += val
	}

	for i := range kernel {
		kernel[i] /= sum
	}

	return kernel
}

// radius = 2(default), 3
func EpanechnikovKernel(radius int) []float64 {
	size := 2*radius + 1
	kernel := make([]float64, size)

	var sum float64
	for i := -radius; i <= radius; i++ {
		x := float64(i) / float64(radius)

		val := 0.75 * (1 - x*x)
		kernel[i+radius] = val
		sum += val
	}

	for i := range kernel {
		kernel[i] /= sum
	}

	return kernel
}

// radius : 2(default)-3
func GaussianKernel(radius int) []float64 {
	sigma := float64(radius) / 3.0 // sigma approximation from radius
	size := 2*radius + 1
	kernel := make([]float64, size)

	var sum float64
	for i := -radius; i <= radius; i++ {
		val := math.Exp(-float64(i*i) / (2 * sigma * sigma))
		kernel[i+radius] = val
		sum += val
	}

	for i := range kernel {
		kernel[i] /= sum
	}

	return kernel
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sinc(x float64) float64 {
	if x == 0 {
		return 1
	}
	return math.Sin(math.Pi*x) / (math.Pi * x)
}
