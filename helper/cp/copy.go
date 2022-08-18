package cp

import "github.com/jinzhu/copier"

// CopyProperties 拷贝属性, 支持struct, slice, map等.
func CopyProperties(source interface{}, dst interface{}) {
	err := copier.Copy(dst, source)
	if err != nil {
		panic(err)
	}
}
