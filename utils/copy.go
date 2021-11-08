package utils

import "github.com/jinzhu/copier"

func CopyProperty(source interface{}, dst interface{}) {
	copier.Copy(dst, source)
}
