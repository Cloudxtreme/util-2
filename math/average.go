// Copyright 2013 Felipe Alves Cavani. All rights reserved.
// Start date:        2014-04-25
// Last modification: 2014-

package math

func AvgInt64(values []int64) float64 {
	var sum int64
	for _, val := range values {
		sum += val
	}
	return float64(sum) / float64(len(values))
}
