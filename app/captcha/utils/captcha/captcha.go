package captcha

import (
	"math/rand"
	"sync"
	"time"
)

var (
	once    sync.Once
	Captcha *captcha
	// 随机因子
	factor = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type captcha struct {
	Length int8
	Chars  string
	Expire int
}

func NewCaptcha(length int8, chars string, expire int) {
	once.Do(func() {
		Captcha = &captcha{
			Length: length,
			Chars:  chars,
			Expire: expire,
		}
	})
}

func (c *captcha) Generate() (string, int) {
	// 生成随机验证码
	code := make([]byte, c.Length)
	for i := 0; i < int(c.Length); i++ {
		code[i] = c.Chars[factor.Intn(len(c.Chars))]
	}

	return string(code), c.Expire
}
