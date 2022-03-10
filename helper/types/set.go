package types

import "github.com/ehwjh2010/viper/helper/cast"

type EmptyStruct struct{}

func NewEmptyStruct() EmptyStruct {
	return EmptyStruct{}
}

// SimpleSet 非线程安全Set
type SimpleSet struct {
	data map[interface{}]EmptyStruct
}

func NewSimpleSet() *SimpleSet {
	return &SimpleSet{
		data: make(map[interface{}]EmptyStruct),
	}
}

// Add 添加元素
func (set *SimpleSet) Add(v interface{}) {
	if set == nil {
		return
	}
	set.data[v] = NewEmptyStruct()
}

// Del 删除元素
func (set *SimpleSet) Del(v interface{}) {
	if set == nil {
		return
	}
	delete(set.data, v)
}

// Update 批量添加元素
func (set *SimpleSet) Update(vs ...interface{}) {
	if vs == nil || set == nil {
		return
	}

	for _, v := range vs {
		set.Add(v)
	}
}

// Size 获取长度
func (set *SimpleSet) Size() int {
	if set == nil {
		return 0
	}
	return len(set.data)
}

// Has 是否已包含
func (set *SimpleSet) Has(v interface{}) bool {
	if set.IsEmpty() {
		return false
	}
	_, exists := set.data[v]
	return exists
}

// NotHas 是否不包含
func (set *SimpleSet) NotHas(v interface{}) bool {
	return !set.Has(v)
}

// Values 获取所有的值
func (set *SimpleSet) Values() []interface{} {
	if set.IsEmpty() {
		return make([]interface{}, 0)
	}

	j := 0
	keys := make([]interface{}, len(set.data))
	for k, _ := range set.data {
		keys[j] = k
		j++
	}
	return keys
}

// IntValues 获取所有的值
func (set *SimpleSet) IntValues() ([]int, error) {

	if set.IsEmpty() {
		return make([]int, 0), nil
	}

	j := 0
	values := make([]int, len(set.data))
	for k := range set.data {
		v, err := cast.Any2Int(k)
		if err != nil {
			return nil, err
		}
		values[j] = v
		j++
	}
	return values, nil
}

// Int64Values 获取所有的值
func (set *SimpleSet) Int64Values() ([]int64, error) {
	if set.IsEmpty() {
		return make([]int64, 0), nil
	}

	j := 0
	values := make([]int64, len(set.data))
	for k := range set.data {
		v, err := cast.Any2Int64(k)
		if err != nil {
			return nil, err
		}
		values[j] = v
		j++
	}
	return values, nil
}

// Int32Values 获取所有的值
func (set *SimpleSet) Int32Values() ([]int32, error) {
	if set.IsEmpty() {
		return make([]int32, 0), nil
	}

	j := 0
	values := make([]int32, len(set.data))
	for k := range set.data {
		v, err := cast.Any2Int32(k)
		if err != nil {
			return nil, err
		}
		values[j] = v
		j++
	}
	return values, nil
}

// StrValues 获取所有的值
func (set *SimpleSet) StrValues() ([]string, error) {
	if set.IsEmpty() {
		return make([]string, 0), nil
	}

	j := 0
	keys := make([]string, len(set.data))
	for k := range set.data {
		v, err := cast.Any2String(k)
		if err != nil {
			return nil, err
		}
		keys[j] = v
		j++
	}
	return keys, nil
}

// Float64Values 获取所有的值
func (set *SimpleSet) Float64Values() ([]float64, error) {
	if set.IsEmpty() {
		return make([]float64, 0), nil
	}

	j := 0
	values := make([]float64, len(set.data))
	for k := range set.data {
		v, err := cast.Any2Double(k)
		if err != nil {
			return nil, err
		}
		values[j] = v
		j++
	}
	return values, nil
}

// Union 并集
func (set *SimpleSet) Union(s *SimpleSet) *SimpleSet {
	r := set.Copy()
	r.Update(s.Values()...)
	return r
}

// Common 交集
func (set *SimpleSet) Common(s *SimpleSet) *SimpleSet {
	r := NewSimpleSet()
	for v, _ := range set.data {
		if s.Has(v) {
			r.Add(v)
		}
	}
	return r
}

// Diff 差集
func (set *SimpleSet) Diff(s *SimpleSet) *SimpleSet {
	r := NewSimpleSet()
	for v, _ := range set.data {
		if s.NotHas(v) {
			r.Add(v)
		}
	}
	return r
}

// Copy copy自身
func (set *SimpleSet) Copy() *SimpleSet {
	if set.IsEmpty() {
		return nil
	}

	copySet := NewSimpleSet()

	copyData := make(map[interface{}]EmptyStruct, len(set.data))
	for v, _ := range set.data {
		copyData[v] = NewEmptyStruct()
	}

	copySet.data = copyData

	return copySet
}

// IsEmpty 空集合
func (set *SimpleSet) IsEmpty() bool {
	if set == nil || len(set.data) <= 0 {
		return true
	}

	return false
}
