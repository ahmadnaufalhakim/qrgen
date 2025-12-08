package render

import "math"

func GaussianKernel(sigma float64) []float64 {
	r := int(math.Ceil(3 * sigma))
	size := 2*r + 1
	kernel := make([]float64, size)

	var sum float64
	for i := -r; i <= r; i++ {
		val := math.Exp(-float64(i*i) / (2 * sigma * sigma))
		kernel[i+r] = val
		sum += val
	}

	for i := range kernel {
		kernel[i] /= sum
	}

	return kernel
}
