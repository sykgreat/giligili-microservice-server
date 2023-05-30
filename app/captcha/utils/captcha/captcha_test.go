package captcha

import "testing"

func TestCaptcha_Generate(t *testing.T) {
	NewCaptcha(6, "1234567890", 60)
	code, expire := Captcha.Generate()
	t.Log(code, expire)
}
