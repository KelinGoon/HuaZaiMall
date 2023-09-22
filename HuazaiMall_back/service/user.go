package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"singo/dal/model"
	"singo/dal/query"
	"singo/dao"
	"singo/module/auth"
	"singo/module/response"
	"time"
)

// GenerateToken 这可能是一个令牌
//
//	@生成的令牌是JSON	Web Token（JWT）。JWT是一种开放标准（RFC 7519）
//	@用于在各方之间安全地传输信息
//	@它由三个部分组成：头部（Header）、负载（Payload）和签名（Signature）
func GenerateToken(userID uint) (string, error) {
	claim := auth.JwtClaim{ //创建jwt结构体

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(auth.JwtExpiresTime).Unix(), //令牌过期时间
			IssuedAt:  time.Now().Unix(),                          //开始时间
		},

		UserID: userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(auth.JwtSecretKey))
}

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var undersign model.User
	if err := c.ShouldBind(&undersign); err != nil { //进行数据校验和绑定操作
		response.ErrRespWithMsg(c, err.Error())
	} else {
		// 先给密码做一个判断
		undersign.Password = dao.SetPassword(undersign.Password)
		err := dao.Signin(undersign) //进行登录函数判断
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
		} else {
			response.OKResp(c)
		}
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var userLoginModel dao.UserLoginModel
	if err := c.ShouldBind(&userLoginModel); err != nil { //进行数据校验和绑定操作
		response.ErrRespWithMsg(c, err.Error())
	} else {
		user, err := dao.GetUserAccount(userLoginModel) //进行登录函数判断
		if err != nil {
			response.ErrRespWithMsg(c, err.Error())
		} else {
			if dao.CheckPassword(&user, userLoginModel.Password) {
				token, err := GenerateToken(uint(user.UserID))
				if err != nil {
					response.ErrRespWithMsg(c, err.Error())
				}
				response.OKRespWithData(c, map[string]interface{}{
					"token": token,
				})
			} else {
				response.ErrRespWithMsg(c, "密码错误")
			}
		}

	}
}

// Userme get数据
func Userme(c *gin.Context) {
	a := query.User
	userdemo, err := a.Where(a.UserID.Eq(1)).First()
	if err != nil {
		response.ErrRespWithMsg(c, err.Error())
	}
	var user model.User
	user = *userdemo
	// 成功返回的示例
	response.OKRespWithData(c, map[string]interface{}{
		"user_list": user,
	})

}

// UserLogout 用户登出
//func UserLogout(c *gin.Context) {
//	// 获取当前用户的 Token（假设 Token 存储在上下文中）
//	// 构建 JSON 响应
//
//}
