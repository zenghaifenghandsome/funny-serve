package moddleware

import (
	"fmt"
	errormessages "funny-serve/utils/errorMessages"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("www.zzhh.asia1998")

const Passtime = time.Minute * 1
const Refreshtime = time.Hour * 24

type MyClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// 生成token
func GenToken(username string, passtime time.Duration, role string) (string, int) {
	claims := MyClaims{
		username,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(passtime).Unix(),
			Issuer:    "zzhh",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errormessages.ERROR_TOKEN_GENTOKEN_ERROR
	}
	return signedToken, errormessages.SUCCESS
}

// 验证token
func ParseToken(tokenString string) (*MyClaims, int) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
	if err != nil {
		return nil, errormessages.ERROR_TOKEN_PARSETOKEN_ERROR
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, errormessages.SUCCESS
	}
	return nil, errormessages.ERROR_TOKEN_INVALIDTOKEN
}

// 中间件
func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg": errormessages.GetErrMsg(errormessages.ERROR_TOKEN_HEADER_ERROR),
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"msg": errormessages.GetErrMsg(errormessages.ERROR_TOKEN_HEADER_ERROR),
			})
			c.Abort()
			return
		}
		myClaims, err := ParseToken(parts[1])
		if err != errormessages.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"msg": errormessages.GetErrMsg(errormessages.ERROR_TOKEN_INVALIDTOKEN),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > myClaims.ExpiresAt {
			c.JSON(http.StatusOK, gin.H{
				"msg": errormessages.GetErrMsg(errormessages.ERROR_TOKEN_TIMA_PASS),
			})
			c.Abort()
			return
		}
		c.Set("username", myClaims.Username)
		c.Set("role", myClaims.Role)
		c.Next()
	}
}

func CheckAuth(role string) func(c *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println(c.GetString("role") + "@@@@@@@@@")
		if c.GetString("role") != role {
			c.JSON(http.StatusOK, gin.H{
				"msg": errormessages.GetErrMsg(errormessages.ERROR_TOKEN_ROLE_ERROR),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
