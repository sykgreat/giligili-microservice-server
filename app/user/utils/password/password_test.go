package password

import (
	"testing"
)

func TestPassword_GeneratePassword(t *testing.T) {
	s, err := GeneratePassword("123456")
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}

func TestPassword_ComparePassword(t *testing.T) {
	err := ComparePassword("$2a$10$ik0m3lEvyzFkXQHXa3TsOeXqEs/pcLCkcOIMfmMUmlewAlA1edzCq", "123456")
	if err != nil {
		t.Error(err)
	}
	t.Log("success")
}
