package model

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type UserToken struct {
	Id   int64
	Name string
	jwt.RegisteredClaims
}

// 签名密钥
const signKey = "路边有野花"

func GetJwt(id int64, name string) (string, error) {
	if id < 0 || name == "" {
		return "", errors.New("参数错误！！！")
	}
	token := &UserToken{
		Id:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "路边有野花",                                             // 签发者
			Subject:   "宝石兽",                                               // 签发对象
			Audience:  jwt.ClaimStrings{"Android", "IOS", "H5", "Web"},     //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),       //过期时间 1小时
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second * 0)), //最早使用时间 10秒之后
			IssuedAt:  jwt.NewNumericDate(time.Now()),                      //签发时间 当前时间
			ID:        "Test-1",                                            // jwt ID,类似于盐值 最好是每次都随机
		},
	}
	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, token).SignedString([]byte(signKey))
	return tokenStr, err
}

func CheckJwt(tokenStr string) (*UserToken, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil //返回签名密钥
	})

	if err != nil || !token.Valid {
		return nil, errors.New("校验失败，TOKEN不合格")
	}

	claims, ok := token.Claims.(*UserToken)
	if !ok {
		return nil, errors.New("TOKEN转义失败！")
	}

	return claims, nil
}

var JWTMap map[string]int

func GetJWTMap(name string) bool {
	if JWTMap == nil {
		JWTMap = make(map[string]int)
	}
	_, flag := JWTMap[name]
	return flag
}

func ClearJWTMap(name string) {
	delete(JWTMap, name)
}

func GetToken(c *gin.Context) {
	// 从请求中获取state参数
	state := c.Query("state")

	// 如果state参数为空，则返回错误
	if state == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数state缺失",
		})
		return
	}

	// 从Redis中获取token
	token, err := Redis.Get(c, "token_"+state).Result()
	if err != nil {
		// 如果在Redis中找不到对应的token，返回错误
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取token失败",
		})
		return
	}

	// 如果从Redis中成功获取到token，则返回token
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
