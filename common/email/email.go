package email

import (
	"net/smtp"
	"strings"
	"sync"
)

var (
	Email email
	once  sync.Once
)

type email struct {
	User     string
	Name     string
	Password string
	Host     string
	MailType string
}

func NewEmail(user, name, password, host, mailType string) {
	once.Do(func() {
		Email = email{
			User:     user,     // 邮箱账号
			Name:     name,     // 发件人名称
			Password: password, // 邮箱密码
			Host:     host,     // 邮箱服务器
			MailType: mailType, // 邮件类型
		}
	})
}

func (e *email) Send(to, subject, body string) error {
	hp := strings.Split(e.Host, ":")
	auth := smtp.PlainAuth("", e.User, e.Password, hp[0])
	var contentType string
	if e.MailType == "html" {
		contentType = "Content-Type: text/" + e.MailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + e.Name + "<" + e.User + ">" + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(e.Host, auth, e.User, sendTo, msg)
	return err
}
