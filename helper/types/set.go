package types

import (
	"fmt"

	"github.com/ehwjh2010/viper/helper/basic/double"
	"github.com/ehwjh2010/viper/helper/basic/integer"
	"github.com/ehwjh2010/viper/helper/basic/str"
)

type EmptyStruct struct{}

func NewEmptyStruct() EmptyStruct {
	return EmptyStruct{}
}

// Set 非线程安全Set
type Set struct {
	data map[interface{}]EmptyStruct
}

func (set *Set) String() string {
	return "Set<" + fmt.Sprintf("%v", set.Values()) + ">"
}

func NewSimpleSet() *Set {
	return &Set{
		data: make(map[interface{}]EmptyStruct, 10),
	}
}

// Add 添加元素
func (set *Set) Add(v interface{}) {
	if set == nil {
		return
	}
	set.data[v] = NewEmptyStruct()
}

// Del 删除元素
func (set *Set) Del(v interface{}) {
	if set == nil {
		return
	}
	delete(set.data, v)
}

// Update 批量添加元素
func (set *Set) Update(vs ...interface{}) {
	if vs == nil || set == nil {
		return
	}

	for _, v := range vs {
		set.Add(v)
	}
}

// UpdateInt64s 批量添加元素
func (set *Set) UpdateInt64s(vs ...int64) {
	if vs == nil || set == nil {
		return
	}

	for _, v := range vs {
		set.Add(v)
	}
}

// UpdateStrings 批量添加元素
func (set *Set) UpdateStrings(vs ...string) {
	if vs == nil || set == nil {
		return
	}

	for _, v := range vs {
		set.Add(v)
	}
}

// UpdateFloat64s 批量添加元素
func (set *Set) UpdateFloat64s(vs ...float64) {
	if vs == nil || set == nil {
		return
	}

	for _, v := range vs {
		set.Add(v)
	}
}

// UpdateFloat32s 批量添加元素
func (set *Set) UpdateFloat32s(vs ...float32) {
	if vs == nil || set == nil {
		return
	}

	for _, v := range vs {
		set.Add(v)
	}
}

// UpdateInt32s 批量添加元素
func (set *Set) UpdateInt32s(vs ...int32) {
	if vs == nil || set == nil {
		return
	}

	for _, v := range vs {
		set.Add(v)
	}
}

// UpdateInts 批量添加元素
func (set *Set) UpdateInts(vs ...int) {
	if vs == nil || set == nil {
		return
	}

	for _, v := range vs {
		set.Add(v)
	}
}

// Size 获取长度
func (set *Set) Size() int {
	if set == nil {
		return 0
	}
	return len(set.data)
}

// Has 是否已包含
func (set *Set) Has(v interface{}) bool {
	if set.IsEmpty() {
		return false
	}
	_, exists := set.data[v]
	return exists
}

// NotHas 是否不包含
func (set *Set) NotHas(v interface{}) bool {
	return !set.Has(v)
}

// Values 获取所有的值
func (set *Set) Values() []interface{} {
	if set.IsEmpty() {
		return make([]interface{}, 0)
	}

	j := 0
	keys := make([]interface{}, len(set.data))
	for k := range set.data {
		keys[j] = k
		j++
	}
	return keys
}

// IntValues 获取所有的值
func (set *Set) IntValues() ([]int, error) {

	if set.IsEmpty() {
		return make([]int, 0), nil
	}

	j := 0
	values := make([]int, len(set.data))
	for k := range set.data {
		v, err := integer.Any2Int(k)
		if err != nil {
			return nil, err
		}
		values[j] = v
		j++
	}
	return values, nil
}

// MustIntValues 获取所有的值
func (set *Set) MustIntValues() []int {

	if set.IsEmpty() {
		return make([]int, 0)
	}

	j := 0
	values := make([]int, len(set.data))
	for k := range set.data {
		v := integer.MustAny2Int(k)
		values[j] = v
		j++
	}
	return values
}

// Int64Values 获取所有的值
func (set *Set) Int64Values() ([]int64, error) {
	if set.IsEmpty() {
		return make([]int64, 0), nil
	}

	j := 0
	values := make([]int64, len(set.data))
	for k := range set.data {
		v, err := integer.Any2Int64(k)
		if err != nil {
			return nil, err
		}
		values[j] = v
		j++
	}
	return values, nil
}

// MustInt64Values 获取所有的值
func (set *Set) MustInt64Values() []int64 {
	if set.IsEmpty() {
		return make([]int64, 0)
	}

	j := 0
	values := make([]int64, len(set.data))
	for k := range set.data {
		v := integer.MustAny2Int64(k)
		values[j] = v
		j++
	}
	return values
}

// Int32Values 获取所有的值
func (set *Set) Int32Values() ([]int32, error) {
	if set.IsEmpty() {
		return make([]int32, 0), nil
	}

	j := 0
	values := make([]int32, len(set.data))
	for k := range set.data {
		v, err := integer.Any2Int32(k)
		if err != nil {
			return nil, err
		}
		values[j] = v
		j++
	}
	return values, nil
}

// MustInt32Values 获取所有的值
func (set *Set) MustInt32Values() []int32 {
	if set.IsEmpty() {
		return make([]int32, 0)
	}

	j := 0
	values := make([]int32, len(set.data))
	for k := range set.data {
		v := integer.MustAny2Int32(k)
		values[j] = v
		j++
	}
	return values
}

// StrValues 获取所有的值
func (set *Set) StrValues() ([]string, error) {
	if set.IsEmpty() {
		return make([]string, 0), nil
	}

	j := 0
	keys := make([]string, len(set.data))
	for k := range set.data {
		v, err := str.Any2String(k)
		if err != nil {
			return nil, err
		}
		keys[j] = v
		j++
	}
	return keys, nil
}

// MustStrValues 获取所有的值
func (set *Set) MustStrValues() []string {
	if set.IsEmpty() {
		return make([]string, 0)
	}

	j := 0
	keys := make([]string, len(set.data))
	for k := range set.data {
		v := str.MustAny2String(k)
		keys[j] = v
		j++
	}
	return keys
}

// Float64Values 获取所有的值
func (set *Set) Float64Values() ([]float64, error) {
	if set.IsEmpty() {
		return make([]float64, 0), nil
	}

	j := 0
	values := make([]float64, len(set.data))
	for k := range set.data {
		v, err := double.Any2Double(k)
		if err != nil {
			return nil, err
		}
		values[j] = v
		j++
	}
	return values, nil
}

// MustFloat64Values 获取所有的值
func (set *Set) MustFloat64Values() []float64 {
	if set.IsEmpty() {
		return make([]float64, 0)
	}

	j := 0
	values := make([]float64, len(set.data))
	for k := range set.data {
		v := double.MustAny2Double(k)
		values[j] = v
		j++
	}
	return values
}

// Union 并集
func (set *Set) Union(s *Set) *Set {
	r := set.Copy()
	r.Update(s.Values()...)
	return r
}

// Common 交集
func (set *Set) Common(s *Set) *Set {
	if set.IsEmpty() || s.IsEmpty() {
		return NewSimpleSet()
	}

	r := NewSimpleSet()
	for v := range set.data {
		if s.Has(v) {
			r.Add(v)
		}
	}
	return r
}

// Diff 差集
func (set *Set) Diff(s *Set) *Set {
	if set.IsEmpty() {
		return NewSimpleSet()
	}

	if s.IsEmpty() {
		return set.Copy()
	}

	r := NewSimpleSet()
	for v := range set.data {
		if s.NotHas(v) {
			r.Add(v)
		}
	}
	return r
}

// Copy copy自身
func (set *Set) Copy() *Set {
	if set.IsEmpty() {
		return NewSimpleSet()
	}

	copySet := NewSimpleSet()

	copyData := make(map[interface{}]EmptyStruct, len(set.data))
	for v := range set.data {
		copyData[v] = NewEmptyStruct()
	}

	copySet.data = copyData

	return copySet
}

// IsEmpty 判断集合是否为空
func (set *Set) IsEmpty() bool {
	if set == nil || len(set.data) <= 0 {
		return true
	}

	return false
}

// IsNotEmpty 判断集合不为空
func (set *Set) IsNotEmpty() bool {
	return !set.IsEmpty()
}
