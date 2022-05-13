package ginext

import (
	ut "github.com/go-playground/universal-translator"

	"github.com/ehwjh2010/viper/component/validator"
)

var translator ut.Translator

// RegisterTrans 定义翻译的方法
func RegisterTrans(language string) error {
	trans, err := validator.RegisterTrans(language)
	if err != nil {
		return err
	}

	translator = trans
	return nil
}

// Translate 翻译错误信息
func Translate(err error) (errMsg string) {
	return validator.Translate(err, translator)
}
