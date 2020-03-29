package schema

// LoginCaptcha 登录验证码
type LoginCaptcha struct {
	CaptchaID string `json:"captcha_id"` // 验证码ID
}
