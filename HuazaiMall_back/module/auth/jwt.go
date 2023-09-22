package auth

import (
	"github.com/dgrijalva/jwt-go"
	"singo/module/response"

	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	// JwtSecretKey jwt加密秘钥
	JwtSecretKey = "keyasdfasdf"
	// JwtExpiresTime jwt过期时间
	JwtExpiresTime = time.Hour * 24
	// auth 就是 JSON Web Token（JSON Web令牌）
)

// JwtClaim jwt编码的结构体
type JwtClaim struct {
	jwt.StandardClaims
	UserID uint
}

// JwtRequired jwt中间件
func JwtRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获得token
		// userToken := c.Request.Header.Get("Authorization")
		userToken := c.GetHeader("X-Token")
		// 判断请求头中是否有token
		if userToken == "" {
			userToken = c.Request.Header.Get("Authorization")
			if userToken == "" {
				response.ErrRespWithMsg(c, "tokon无效")
				c.Abort()
				return
			}
		}

		// userToken = strings.Split(userToken, " ")[1] // 处理token中的“Bearer”
		tokenParts := strings.Split(userToken, " ")
		if len(tokenParts) < 2 {
			// 处理错误：没有足够的部分
			response.ErrRespWithMsg(c, "tokon格式不正确")
			c.Abort()
			return
		} else {
			// 提取第二部分（索引 1）
			userToken := tokenParts[1]
			// 这里可以继续处理有效令牌

			// 解码token值
			token, err := jwt.ParseWithClaims(userToken, &JwtClaim{},
				func(token *jwt.Token) (interface{}, error) {
					return []byte(JwtSecretKey), nil
				})
			if err != nil || token.Valid != true {

				// 过期或者非正确处理
				c.JSON(http.StatusOK, response.ErrorResponse(response.CodeTokenExpiredError))
				response.ErrRespWithMsg(c, "token已过期")
				c.Abort()
				return
			}

			// 判断令牌是否在黑名单里面
			//if result, _ := cache.RedisClient.SIsMember(cache.Context, "auth:baned", token.Raw).Result(); result {
			//	c.JSON(200, serializer.ErrorResponse(serializer.CodeTokenExpiredError))
			//	c.Abort()
			//	return
			//}

			// 用户id存入上下文
			if jwtStruct, ok := token.Claims.(*JwtClaim); ok {
				c.Set("user_id", &jwtStruct.UserID)
			}

			// 将Token放入Context, 用于退出登录添加黑名单
			c.Set("token", token.Raw)
		}
	}
}
