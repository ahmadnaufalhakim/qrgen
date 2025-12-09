package qrconst

type KernelType int

const (
	KernelLanczos2 KernelType = iota
	KernelCubicSmooth
	KernelGaussian
	KernelLanczos3
	KernelHann
	KernelTriangle
	KernelCosine
	KernelEpanechnikov
	KernelBSpline
	KernelBox
)
