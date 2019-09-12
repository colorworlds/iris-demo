package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	"reflect"
	"regexp"
)

var validate *validator.Validate
var trans  ut.Translator

func init() {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})

	// 这里举了个英文的特例，多语言的话，可以根据情况自己实现
	_en := en.New()
	trans, _ = ut.New(_en, _en).GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	registerTagValidator("phone", "{0} is a invalid phone.", phoneValidator)
}

// 检查错误
func Check(v interface{}) (error, string) {
	if err := validate.Struct(v); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, _err := range errs {
				return err, _err.Translate(trans)
			}
		}
		return err, ""
	}
	return nil, "OK"
}

// 注册自定义的tag验证器
func registerTagValidator(tagName, message string, fn validator.Func) {
	validate.RegisterValidation(tagName, fn)
	validate.RegisterTranslation(tagName, trans, func(ut ut.Translator) error {
		return ut.Add(tagName, message, false)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T(fe.Tag(), fe.Field())
		if err != nil {
			return fe.(error).Error()
		}
		return t
	})
}

// 手机号码验证器
func phoneValidator(fl validator.FieldLevel) bool {
	b, _ := regexp.MatchString(`^[\d]{11}$`, fl.Field().String())
	return b
}

