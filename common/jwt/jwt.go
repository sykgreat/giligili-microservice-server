package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type Jwt struct {
	AccessTokenSecret, RefreshTokenSecret string
	AccessTokenExpire, RefreshTokenExpire int
	Issuer                                string
}

type Claims struct {
	UserId    int64
	Email     string
	TokenType uint // 0:accessToken,1:refreshToken
	jwt.RegisteredClaims
}

// GenerateAccessToken 生成access token
func (j *Jwt) GenerateAccessToken(id int64, email string) (string, error) {
	accessJwtKey := []byte(j.AccessTokenSecret)
	expirationTime := time.Now().Add(time.Duration(j.AccessTokenExpire) * time.Hour)
	accessClaims := &Claims{
		UserId:    id,
		Email:     email,
		TokenType: 0,
		RegisteredClaims: jwt.RegisteredClaims{
			//发放时间等
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.Issuer,
		},
	}
	return j.generateToken(accessJwtKey, accessClaims)
}

// GenerateRefreshToken 生成refresh token
func (j *Jwt) GenerateRefreshToken(id int64, email string) (string, error) {
	refreshJwtKey := []byte(j.RefreshTokenSecret)
	expirationTime := time.Now().Add(time.Duration(j.RefreshTokenExpire) * time.Hour)

	refreshClaims := &Claims{
		UserId:    id,
		Email:     email,
		TokenType: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			//发放时间等
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "leaf",
		},
	}
	return j.generateToken(refreshJwtKey, refreshClaims)
}

// GetTokenClaims 获取token的荷载数据
func (j *Jwt) GetTokenClaims(tokenString string) (*Claims, error) {
	claims := &Claims{}
	parser := jwt.NewParser()
	_, _, err := parser.ParseUnverified(tokenString, claims)
	return claims, err
}

// ParseToken 解析token
func (j *Jwt) ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	// 获取jwt的荷载数据
	claims, err := j.GetTokenClaims(tokenString)
	if err != nil {
		zap.L().Error("token荷载解析失败: " + err.Error())
	}
	// 判断类型 选择不同的密钥
	var secret []byte
	if claims.TokenType == 0 { // accessToken
		secret = []byte(j.AccessTokenSecret)
	} else if claims.TokenType == 1 { // refreshToken
		secret = []byte(j.RefreshTokenSecret)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		return secret, nil
	})

	return token, claims, err
}

// generateToken 生成token
func (j *Jwt) generateToken(key []byte, claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
