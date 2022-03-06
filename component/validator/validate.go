package validator

import (
	"github.com/ehwjh2010/viper/global"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en2 "github.com/go-playground/validator/v10/translations/en"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
	"strings"
)

// RegisterTrans 定义翻译的方法
func RegisterTrans(language string) (translator ut.Translator, err error) {
	var trans ut.Translator

	language = strings.ToLower(language)
	zhT := zh.New() //中文翻译器
	enT := en.New() //英文翻译器
	//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
	uni := ut.New(enT, zhT, enT)

	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 获取json tag中的字段名
		// TODO fix 如果是表单数据, 则会有问题, 查询字符串 待查看
		// TODO 添加常见的tag解析
		//v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		//	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		//	if name == "-" {
		//		return ""
		//	}
		//	return name
		//})

		switch language {
		case global.English:
			trans, _ = uni.GetTranslator(enT.Locale())
		case global.Chinese:
			trans, _ = uni.GetTranslator(zhT.Locale())
		default:
			trans, _ = uni.GetTranslator(enT.Locale())
		}

		switch language {
		case global.English:
			err = en2.RegisterDefaultTranslations(v, trans)
		case global.Chinese:
			err = zh2.RegisterDefaultTranslations(v, trans)
		default:
			err = en2.RegisterDefaultTranslations(v, trans)
		}

		if err != nil {
			return nil, errors.Wrap(err, "init translator failed")
		}
	}
	return trans, nil
}

// Translate 翻译错误信息
func Translate(err error, tran ut.Translator) (errMsg string) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		// validator.ValidationErrors类型错误则进行翻译
		for _, err := range errs {
			errMsg = err.Translate(tran)
			break
		}
	} else {
		// 非validator.ValidationErrors类型错误直接返回
		errMsg = err.Error()
	}
	return
}
