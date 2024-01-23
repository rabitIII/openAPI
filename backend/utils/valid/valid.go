package valid

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var (
	trans ut.Translator
)

func init() {
	err := InitTrans("zh")
	if err != nil {
		return
	}
}

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		//// 注册一个获取json tag的自定义ff
		//v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		//	name := strings.SplitN(fld.Tag.Get("label"), ",", 2)[0]
		//	if name == "" {
		//		// 没有label就json
		//		name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		//	}
		//	if name == "-" {
		//		return ""
		//	}
		//	return name
		//})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := utNew(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok := uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func Error(err error) (ret string) {
	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if !ok {
		return err.Error()
	}

	for _, e := range validationErrors {
		msg := e.Translate(trans)
		ret += msg + ";"
	}
	return ret
}

func ValidError(err error, obj any) (ret string, data map[string]string) {
	data = map[string]string{}
	getObj := reflect.TypeOf(obj)
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error(), data
	}

	for _, e := range validationErrors {
		msg := e.Translate(trans)
		fieldName := e.Field()
		filed, ok := getObj.Elem().FieldByName(fieldName)
		if ok {
			alias := fieldName
			label, labelOk := filed.Tag.Lookup("label")
			jsonLabel, jsonOK := filed.Tag.Lookup("json")
			if labelOk {
				alias = label
			} else {
				if jsonOK {
					alias = jsonLabel
				}
			}
			strings.ReplaceAll(msg, fieldName, alias)
		}
		ret += msg + ";"
	}
	return ret, data
}
