package dao

import (
	"github.com/gin-gonic/gin"
	"singo/dal/model"
	"singo/dal/query"
	crypto "singo/module/util"
)

type UserLoginModel struct {
	Account  string `form:"account" binding:"required,min=3,max=10"`
	Password string `json:"password"`
}

// Signin 注册新用户
func Signin(user model.User) error {
	u := query.User
	err := u.Create(&user)
	return err
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// GetUser 用ID获取用户
func GetUser(ID int) (model.User, error) {
	var user model.User
	// result := DB.First(&user, ID)
	// 查询
	result, err := query.User.Where(query.User.UserID.Eq(int64(ID))).First()
	user = *result
	return user, err
}

// GetUserAccount  用account获取用户
func GetUserAccount(loginModel UserLoginModel) (model.User, error) {
	q := query.User
	result, err := q.Where(q.Username.Eq(loginModel.Account)).First()
	return *result, err
}

// SetPassword 设置密码
func SetPassword(password string) string {
	bytes := crypto.AESEncrypt(password)
	Password := bytes
	return Password
}

// CheckPassword 校验密码
func CheckPassword(user *model.User, password string) bool {
	dePassword, _ := crypto.AESDecrypt(user.Password)
	return dePassword == password
}
