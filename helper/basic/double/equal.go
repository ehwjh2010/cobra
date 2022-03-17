package double

import "sort"

type Float64Slice []float64

func (x Float64Slice) Len() int           { return len(x) }
func (x Float64Slice) Less(i, j int) bool { return x[i] < x[j] }
func (x Float64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func sliceFloat64Equal(v1, v2 []float64, strict bool) bool {
	// 都为nil, 则相等
	v1Nil := v1 == nil
	v2Nil := v2 == nil
	if v1Nil != v2Nil {
		return false
	}

	if v1Nil {
		return true
	}

	// 长度不等, 则不相等
	if len(v1) != len(v2) {
		return false
	}

	if !strict {
		sort.Sort(Float64Slice(v1))
		sort.Sort(Float64Slice(v2))
	}

	// 遍历比较内部元素
	for index, item1 := range v1 {
		item2 := v2[index]
		if item2 != item1 {
			return false
		}
	}

	return true
}

func SliceFloat64EqualStrict(v1, v2 []float64) bool {
	return sliceFloat64Equal(v1, v2, true)
}

func SliceFloat64Equal(v1, v2 []float64) bool {
	return sliceFloat64Equal(v1, v2, false)
}
