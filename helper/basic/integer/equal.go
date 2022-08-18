package integer

import "sort"

type Int32s []int32

func (x Int32s) Len() int           { return len(x) }
func (x Int32s) Less(i, j int) bool { return x[i] < x[j] }
func (x Int32s) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type Int64s []int64

func (x Int64s) Len() int           { return len(x) }
func (x Int64s) Less(i, j int) bool { return x[i] < x[j] }
func (x Int64s) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func sliceIntEqual(v1, v2 []int, strict bool) bool {
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
		sort.Ints(v1)
		sort.Ints(v2)
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

func SliceIntEqualStrict(v1, v2 []int) bool {
	return sliceIntEqual(v1, v2, true)
}

func SliceIntEqual(v1, v2 []int) bool {
	return sliceIntEqual(v1, v2, false)
}

func sliceInt32Equal(v1, v2 []int32, strict bool) bool {
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
		sort.Sort(Int32s(v1))
		sort.Sort(Int32s(v2))
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

func SliceInt32EqualStrict(v1, v2 []int32) bool {
	return sliceInt32Equal(v1, v2, true)
}

func SliceInt32Equal(v1, v2 []int32) bool {
	return sliceInt32Equal(v1, v2, false)
}

func sliceInt64Equal(v1, v2 []int64, strict bool) bool {
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
		sort.Sort(Int64s(v1))
		sort.Sort(Int64s(v2))
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

func SliceInt64EqualStrict(v1, v2 []int64) bool {
	return sliceInt64Equal(v1, v2, true)
}

func SliceInt64Equal(v1, v2 []int64) bool {
	return sliceInt64Equal(v1, v2, false)
}
