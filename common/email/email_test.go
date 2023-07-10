package email

import "testing"

func TestEmail_Send(t *testing.T) {
	NewEmail("sykissky@qq.com", "水月空", "zuouwyjqnwocbfig", "smtp.qq.com:25", "html")
	err := Email.Send("663992442@qq.com", "测试邮件", "测试邮件")
	if err != nil {
		t.Error(err)
	}
	t.Log("success")
}
