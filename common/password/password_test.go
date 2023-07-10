package password

import (
	"testing"
)

func TestPassword_GeneratePassword(t *testing.T) {
	s, err := GeneratePassword("qwer1234")
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}

func TestPassword_ComparePassword(t *testing.T) {
	err := ComparePassword("$2a$10$8KXVl2xzjkD1LzfOHQCeeuS.6J4Z0n4CEo16ArJBcNh9WkYxB7H/u", "qwer1234")
	if err != nil {
		t.Error(err)
	}
	t.Log("success")
}
