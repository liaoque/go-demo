package rule

import (
	"errors"
	"github.com/dchest/captcha"
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/labstack/echo/v4"
)

func CaptchaEquals(c echo.Context) validation.RuleFunc {
	return func(value  interface{}) error {
		cookie, err := c.Cookie("captchaId")
		if err != nil {
			return err
		}
		verCode,ok := value.(string)
		if !ok {
			return errors.New("必须是字符串")
		}
		if !captcha.VerifyString(cookie.Value, verCode) {
			return err
		}
		return nil
	}
}

