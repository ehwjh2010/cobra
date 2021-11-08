package utils

import "github.com/jinzhu/copier"

func CopyProperty(source interface{}, dst interface{}) error {
	return copier.Copy(dst, source)
}
