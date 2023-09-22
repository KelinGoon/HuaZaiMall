package response

// Response 基本响应格式
type Response struct {
	Code Code        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// OkResponse 请求成功响应
func OkResponse(data interface{}) *Response {
	return &Response{
		Code: CodeOk,
		Msg:  GetMsg(CodeOk),
		Data: data,
	}
}

// ErrorResponse 请求失败响应
func ErrorResponse(code Code) *Response {
	return &Response{
		Code: code,
		Msg:  GetMsg(code),
		Data: nil,
	}
}
