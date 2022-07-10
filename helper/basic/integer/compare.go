package integer

import "math"

func MinInt(vs ...int) int {
	r := math.MaxInt
	for _, v := range vs {
		if v < r {
			r = v
		}
	}

	return r
}

func MinInt32(vs ...int32) int32 {
	var r int32 = math.MaxInt32
	for _, v := range vs {
		if v < r {
			r = v
		}
	}
	return r
}

func MinInt64(vs ...int64) int64 {
	var r int64 = math.MaxInt64
	for _, v := range vs {
		if v < r {
			r = v
		}
	}
	return r
}

func MaxInt(vs ...int) int {
	r := math.MinInt
	for _, v := range vs {
		if v > r {
			r = v
		}
	}

	return r
}

func MaxInt32(vs ...int32) int32 {
	var r int32 = math.MinInt32
	for _, v := range vs {
		if v > r {
			r = v
		}
	}

	return r
}

func MaxInt64(vs ...int64) int64 {
	var r int64 = math.MinInt64
	for _, v := range vs {
		if v > r {
			r = v
		}
	}

	return r
}
