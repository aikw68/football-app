package users

import (
	"football/cmd/util"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type UserInputParam struct {
	Email    string `validate:"required,custom-email"`
	Password string `validate:"required,custom-lowercase,custom-uppercase,custom-numeric"`
}

var validate *validator.Validate

func Validation(r *http.Request) error {

	validate = validator.New()
	err := validate.RegisterValidation("custom-email", customValidationEmail)
	if err != nil {
		return errors.WithStack(util.ERR_USER_SYSTEM_ERROR)
	}
	err = validate.RegisterValidation("custom-lowercase", includeLowercase)
	if err != nil {
		return errors.WithStack(util.ERR_USER_SYSTEM_ERROR)
	}
	err = validate.RegisterValidation("custom-uppercase", includeUppercase)
	if err != nil {
		return errors.WithStack(util.ERR_USER_SYSTEM_ERROR)
	}
	err = validate.RegisterValidation("custom-numeric", includeNumeric)
	if err != nil {
		return errors.WithStack(util.ERR_USER_SYSTEM_ERROR)
	}

	uip := &UserInputParam{
		r.FormValue("email"),
		r.FormValue("password")}

	var result util.AppErr
	err = validate.Struct(*uip)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) != 0 {
			for i := range errors {
				switch errors[i].StructField() {
				case "Email":
					switch errors[i].Tag() {
					case "required", "custom-email":
						result = util.ERR_USER_EMAIL_INCORRECT_FORMAT
					}
				case "Password":
					switch errors[i].Tag() {
					case "required", "custom-lowercase", "custom-uppercase", "custom-numeric":
						result = util.ERR_USER_PASSWORD_INCORRECT_FORMAT
					}
				}
			}
		}
		return result
	}
	return result

}

// バリデーション（メールアドレス形式）
func customValidationEmail(fl validator.FieldLevel) bool {
	return checkRegexp(`^([a-z0-9\+_\-]+)(\.[a-z0-9\+_\-]+)*@([a-z0-9\-]+\.)+[a-z]{2,6}$`, fl.Field().String())
}

// 小文字が含まれるかどうか
func includeLowercase(fl validator.FieldLevel) bool {
	return checkRegexp("[a-z]", fl.Field().String())
}

// 大文字が含まれるかどうか
func includeUppercase(fl validator.FieldLevel) bool {
	return checkRegexp("[A-Z]", fl.Field().String())
}

// 数値が含まれるかどうか
func includeNumeric(fl validator.FieldLevel) bool {
	return checkRegexp("[0-9]", fl.Field().String())
}

// 正規表現共通関数
func checkRegexp(reg, str string) bool {
	r := regexp.MustCompile(reg).Match([]byte(str))
	return r
}
