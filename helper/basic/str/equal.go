package str

import "sort"

func sliceStrEqual(v1, v2 []string, strict bool) bool {
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
		sort.Strings(v1)
		sort.Strings(v2)
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

func SliceStrEqualStrict(v1, v2 []string) bool {
	return sliceStrEqual(v1, v2, true)
}

// SliceStrEqual 该方法会影响原来的排序.
func SliceStrEqual(v1, v2 []string) bool {
	return sliceStrEqual(v1, v2, false)
}

func SliceByteEqual(v1, v2 []byte) bool {
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

	// 遍历比较内部元素
	for index, item1 := range v1 {
		item2 := v2[index]
		if item2 != item1 {
			return false
		}
	}

	return true
}
