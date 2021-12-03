package ginext

import (
	"github.com/ehwjh2010/cobra/config"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTran "github.com/go-playground/validator/v10/translations/en"
	zhTran "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
	"strings"
	"sync"
)

var flag = false
var lock = sync.Mutex{}
var trans unTrans.Translator

// RegisterTrans 定义翻译的方法
func RegisterTrans(language string) (err error) {
	if flag {
		return nil
	}

	lock.Lock()
	defer lock.Unlock()
	if flag {
		return nil
	}
	language = strings.ToLower(language)
	zhT := zh.New() //中文翻译器
	enT := en.New() //英文翻译器
	//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
	uni := unTrans.New(enT, zhT, enT)

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
		case config.EN:
			trans, _ = uni.GetTranslator(enT.Locale())
		case config.CN:
			trans, _ = uni.GetTranslator(zhT.Locale())
		default:
			trans, _ = uni.GetTranslator(enT.Locale())
		}

		switch language {
		case config.EN:
			err = enTran.RegisterDefaultTranslations(v, trans)
		case config.CN:
			err = zhTran.RegisterDefaultTranslations(v, trans)
		default:
			err = enTran.RegisterDefaultTranslations(v, trans)
		}

		if err != nil {
			return errors.Wrap(err, "init translator failed")
		}
	}
	flag = true
	return nil
}

//Translate 翻译错误信息
func Translate(err error) (errMsg string) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		// validator.ValidationErrors类型错误则进行翻译
		for _, err := range errs {
			errMsg = err.Translate(trans)
			break
		}
	} else {
		// 非validator.ValidationErrors类型错误直接返回
		errMsg = err.Error()
	}
	return
}
