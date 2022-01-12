package types

type EmptyStruct struct{}

func NewEmptyStruct() EmptyStruct {
	return EmptyStruct{}
}

//SimpleSet 非线程安全Set
type SimpleSet struct {
	data map[interface{}]EmptyStruct
}

func NewSimpleSet() *SimpleSet {
	return &SimpleSet{
		data: make(map[interface{}]EmptyStruct),
	}
}

//Add 添加元素
func (set *SimpleSet) Add(v interface{}) {
	set.data[v] = NewEmptyStruct()
}

//Del 删除元素
func (set *SimpleSet) Del(v interface{}) {
	delete(set.data, v)
}

//Update 批量添加元素
func (set *SimpleSet) Update(vs ...interface{}) {
	for v := range vs {
		set.Add(v)
	}
}

//Size 获取长度
func (set SimpleSet) Size() int {
	return len(set.data)
}

//Has 是否已包含
func (set SimpleSet) Has(v interface{}) bool {
	_, exists := set.data[v]
	return exists
}

//NotHas 是否不包含
func (set SimpleSet) NotHas(v interface{}) bool {
	_, exists := set.data[v]
	return !exists
}

//Values 获取所有的值
func (set SimpleSet) Values() []interface{} {
	j := 0
	keys := make([]interface{}, len(set.data))
	for k := range set.data {
		keys[j] = k
		j++
	}
	return keys
}

//IntValues 获取所有的值
func (set SimpleSet) IntValues() []int {
	j := 0
	keys := make([]int, len(set.data))
	for k := range set.data {
		keys[j] = k.(int)
		j++
	}
	return keys
}

//Int64Values 获取所有的值
func (set SimpleSet) Int64Values() []int64 {
	j := 0
	keys := make([]int64, len(set.data))
	for k := range set.data {
		keys[j] = k.(int64)
		j++
	}
	return keys
}

//Int32Values 获取所有的值
func (set SimpleSet) Int32Values() []int32 {
	j := 0
	keys := make([]int32, len(set.data))
	for k := range set.data {
		keys[j] = k.(int32)
		j++
	}
	return keys
}

//StrValues 获取所有的值
func (set SimpleSet) StrValues() []string {
	j := 0
	keys := make([]string, len(set.data))
	for k := range set.data {
		keys[j] = k.(string)
		j++
	}
	return keys
}

//Float64Values 获取所有的值
func (set SimpleSet) Float64Values() []float64 {
	j := 0
	keys := make([]float64, len(set.data))
	for k := range set.data {
		keys[j] = k.(float64)
		j++
	}
	return keys
}

//Float32Values 获取所有的值
func (set SimpleSet) Float32Values() []float32 {
	j := 0
	keys := make([]float32, len(set.data))
	for k := range set.data {
		keys[j] = k.(float32)
		j++
	}
	return keys
}

//Union 并集
func (set SimpleSet) Union(s SimpleSet) *SimpleSet {
	r := set.Copy()
	r.Update(s.Values())
	return r
}

//Common 交集
func (set SimpleSet) Common(s SimpleSet) *SimpleSet {
	r := NewSimpleSet()
	for v, _ := range set.data {
		if s.Has(v) {
			r.Add(v)
		}
	}
	return r
}

//Diff 差集
func (set SimpleSet) Diff(s SimpleSet) *SimpleSet {
	r := NewSimpleSet()
	for v, _ := range set.data {
		if s.NotHas(v) {
			r.Add(v)
		}
	}
	return r
}

//Copy copy自身
func (set SimpleSet) Copy() *SimpleSet {
	copySet := NewSimpleSet()

	copyData := make(map[interface{}]EmptyStruct, len(set.data))
	for v, _ := range set.data {
		copyData[v] = NewEmptyStruct()
	}

	copySet.data = copyData

	return copySet
}
