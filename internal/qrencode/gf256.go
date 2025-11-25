package qrencode

import "github.com/ahmadnaufalhakim/qrgen/internal/tables"

func addGF256(x, y uint8) uint8 {
	return x ^ y
}

func mulGF256(x, y uint8) uint8 {
	if x == 0 || y == 0 {
		return 0
	}

	logX := uint16(tables.LogGF256[x])
	logY := uint16(tables.LogGF256[y])

	return tables.AntilogGF256[(logX+logY)%255]
}
