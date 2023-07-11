package enum

// UserModule 用户模块
const (
	UserModule = "user_module:" // UserModule 用户模块

	Token = "token:" // Token 用户token

	AccessToken = "access" // AccessToken 用户登录token

	RefreshToken = "refresh" // RefreshToken 用户刷新token
)

// CaptchaModule 验证码模块
const (
	CaptchaModule = "captcha_module:" // CaptchaModule 验证码模块

	Captcha = "captcha:" // Captcha 验证码

	CaptchaLogin = "login:" // CaptchaLogin 登录验证码

	CaptchaRegister = "register:" // CaptchaRegister 注册验证码

	CaptchaResetPassword = "reset_password:" // CaptchaResetPassword 重置密码验证码
)
