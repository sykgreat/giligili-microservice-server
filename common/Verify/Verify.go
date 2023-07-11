package Verify

import "regexp"

// VerifyEmailFormat 验证邮箱格式
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// VerifyPhoneFormat 验证手机号格式
func VerifyPhoneFormat(phone string) bool {
	pattern := `^1[3456789]\d{9}$` //匹配手机号
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phone)
}

// VerifyPasswordFormat 验证密码格式
func VerifyPasswordFormat(password string) bool {
	if len(password) < 8 || len(password) > 20 {
		return false
	}
	// 过滤掉这四类字符以外的密码串,直接判断不合法
	re, err := regexp.Compile(`^[a-zA-Z0-9.@$!%*#_~?&^]{8,16}$`)
	if err != nil {
		return false
	}
	match := re.MatchString(password)
	if !match {
		return false
	}

	var level = 0
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[.@$!%*#_~?&^]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, password)
		if match {
			level++
		}
	}
	if level < 3 {
		return false
	}

	return true
}
