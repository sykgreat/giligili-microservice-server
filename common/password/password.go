package password

import "golang.org/x/crypto/bcrypt"

// GeneratePassword 生成密码
func GeneratePassword(pwd string) (string, error) {
	password := []byte(pwd)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// ComparePassword 比较密码
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
