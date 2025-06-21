package utils

import (
	"time"

	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt" // bcrypt 库（对密码进行Hash）
)

// Secret 从环境变量中获取 JWT 的签名密钥
// 一定要设置好 .env 中的 JWT_SECRET 值，生产时不可泄漏
var Secret = []byte(os.Getenv("JWT_SECRET"))

// CustomClaims 定义我们自己的Claims（可以放需要携带到JWT中的数据）
// jwt.RegisteredClaims 是 JWT 本身定义的一系列字段（包括过期时间）
type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// HashPassword 函数对传入的纯文本密码进行bcryptHash
// 这可以保障我们在数据中只保存Hash而不是纯文本
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword 函数对输入的纯文本进行bcrypt对比，验证密码是否正确
// password：输入时提供的纯文本
// hashed：数据中保存的bcrypt Hash
func CheckPassword(password, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

// GenerateJWT 函数签发 JWT，payload 中包括Username和过期时间
// 这通常是在认证时（login时）由后端颁发给客户端
func GenerateJWT(username string) (string, error) {
	// 设置Claims，Username可以由应用自由定义，ExpiresAt为过期时间（此为72小时）
	claims := CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}

	// 创建时指定签名方式为 HMAC-SHA256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 用Secret进行签名，生产时需要设置好 JWT_SECRET 环境变量
	signed, err := token.SignedString(Secret)
	if err != nil {
		return "", err
	}

	// 一般我们都会在Header中携带时需要 "Bearer " 前缀
	return "Bearer " + signed, nil
}

// ParseJWT 函数用来验证 JWT 的合法性，并解析其中的数据（Claims）
// 一般是在需要认证时使用，如中间件中验证每一个API的Authorization头
func ParseJWT(signed string) (*CustomClaims, error) {
	if len(signed) > 7 && signed[:7] == "Bearer " {
		signed = signed[7:]
	}

	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(signed, claims, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	}, jwt.WithValidMethods([]string{"HS256"}))

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
