package enum

// UserModule 用户模块
const (
	// UserModule 用户模块
	UserModule = "user_module:"
	// Token 用户token
	Token = "token:"
	// AccessToken 用户登录token
	AccessToken = "access"
	// RefreshToken 用户刷新token
	RefreshToken = "refresh"
)

// CaptchaModule 验证码模块
const (
	// CaptchaModule 验证码模块
	CaptchaModule = "captcha_module:"
	// Captcha 验证码
	Captcha = "captcha:"
	// CaptchaLogin 登录验证码
	CaptchaLogin = "login:"
	// CaptchaRegister 注册验证码
	CaptchaRegister = "register:"
	// CaptchaResetPassword 重置密码验证码
	CaptchaResetPassword = "reset_password:"
)
