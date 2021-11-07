package utils

import "github.com/jinzhu/copier"

func CopyProperty(source interface{}, dst interface{}) error {
	return copier.CopyWithOption(dst, source, copier.Option{DeepCopy: true, IgnoreEmpty: false})
}
