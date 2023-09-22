package response

import "strconv"

type Code int

// 通用错误码
const (
	CodeOk                 Code = 200
	CodeUnknownError       Code = -1
	CodeTokenNotFoundError Code = 401
	CodeParamError         Code = 10001
	CodeTokenExpiredError  Code = 10002
	CodeLengthOfParamError Code = 10003
	CodeUpdateError        Code = 10004
	CodeUnlawfulParamError Code = 10005
	CodeSIteNameError      Code = 10006
)

// 用户相关的错误
const (
	// 注册错误
	CodeUserExistError       Code = 20001
	CodePasswordConfirmError Code = 20002

	// 登录错误
	CodeUserNotExistError Code = 21001
	CodePasswordError     Code = 21002
	CodeLoginError        Code = 21003

	//校验错误
	CodeCheckPassError Code = 22001
)

// 数据库操作相关错误
const (
	//插入数据错误
	CodeDBInsertError Code = 30001
	//查询数据库错误
	CodeDBSelectError Code = 30002
	//更新数据库错误
	CodeDBUpdateError Code = 30003
	//删除数据库数据错误
	CodeDBDeleteError Code = 30004
)

// 状态码与描述信息map
var msgMap = map[Code]string{
	CodeOk:                 "执行成功",
	CodeUnknownError:       "未知错误",
	CodeParamError:         "请求参数错误",
	CodeUnlawfulParamError: "请求参数非法",
	CodeUpdateError:        "修改失败！",
	CodeTokenNotFoundError: "需要权限",
	CodeTokenExpiredError:  "身份凭证过期或不正确",
	CodeLengthOfParamError: "请求参数长度不符合要求",

	CodeUserExistError:       "注册失败，用户已存在",
	CodePasswordConfirmError: "注册失败，两次输入密码不一致",

	CodeUserNotExistError: "登录失败，用户名不存在",
	CodePasswordError:     "登录失败，密码错误",
	CodeLoginError:        "登录失败",
	CodeCheckPassError:    "原密码错误",

	//数据库操作相关
	CodeDBInsertError: "插入数据错误，请检查数据格式或是否已存在该记录",
	CodeDBSelectError: "查询数据错误",
	CodeDBUpdateError: "更新数据错误",
	CodeDBDeleteError: "删除数据错误",
	CodeSIteNameError: "输入数据或网站名称错误",
}

// 根据状态码得到对应说明
func GetMsg(code Code) string {
	msg, ok := msgMap[code]
	if !ok {
		msg = "未知错误，错误码：" + strconv.Itoa(int(code))
	}
	return msg
}
