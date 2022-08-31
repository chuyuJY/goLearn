package main

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// 自定义claim
// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 自行添加字段
	Username             string `json:"username,omitempty"`
	jwt.RegisteredClaims        // 内嵌标准声明
}

// 定义Jwt过期时间
const TokenExpire = time.Hour * 1

var Secret = []byte("签名")

func GenToken(username string) (string, error) {
	ownClaims := CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "my-project",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpire)),
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, ownClaims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Secret)
}

// 解析Token
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

//func main() {
//	tokenString, err := GenToken("tiam")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println(tokenString)
//	claims, err := ParseToken(tokenString)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println(claims)
//}
