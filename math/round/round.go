// Copyright 2009 Felipe Alves Cavani. All rights reserved.
// Start date:        2014-04-09
// Last modification: 2014-

package round

import (
	"math"
)

func Round64(x float64) float64 {
	i, f := math.Modf(x)
	if f >= 0.5 {
		return i + 1.0
	}
	return i
}

func Round32(x float32) float32 {
	i, f := math.Modf(float64(x))
	if f >= 0.5 {
		return float32(i) + 1.0
	}
	return float32(i)
}
