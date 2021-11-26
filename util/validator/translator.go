package validator

import (
	"errors"
	"github.com/ehwjh2010/cobra/config"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTran "github.com/go-playground/validator/v10/translations/en"
	zhTran "github.com/go-playground/validator/v10/translations/zh"
	"strings"
	"sync"
)

var flag = false
var lock = sync.Mutex{}

var NoFoundTranslator = errors.New("not found translator")

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
	flag = true
	zhT := zh_Hans_CN.New() //中文翻译器
	enT := en.New()         //英文翻译器
	//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
	uni := unTrans.New(enT, zhT, enT)

	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		var trans unTrans.Translator
		var found bool

		switch language {
		case config.EN:
			trans, found = uni.GetTranslator(enT.Locale())
		case config.CN:
			trans, found = uni.GetTranslator(zhT.Locale())
		default:
			trans, found = uni.GetTranslator(enT.Locale())
		}

		if !found {
			return NoFoundTranslator
		}
		switch language {
		case config.EN:
			err = enTran.RegisterDefaultTranslations(v, trans)
		case config.CN:
			err = zhTran.RegisterDefaultTranslations(v, trans)
		default:
			err = enTran.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}
