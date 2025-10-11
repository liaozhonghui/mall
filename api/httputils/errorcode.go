package httputils

var (
	OK            = response{Code: 0, Message: "OK"}
	InterNalError = response{Code: 500, Message: "服务异常"}

	AuthError          = response{Code: 401, Message: "认证失败"}
	NotFound           = response{Code: 404, Message: "资源未找到"}
	Forbidden          = response{Code: 403, Message: "没有权限访问"}
	ParamError         = response{Code: 422, Message: "参数错误"}
	ServiceUnavailable = response{Code: 503, Message: "服务不可用"}
	TimeOut            = response{Code: 504, Message: "请求超时"}

	UserNotFound = response{Code: USER_NOT_FOUND, Message: "用户不存在"}
	UserError    = response{Code: USER_ERROR, Message: "用户已存在"}
)

const (
	USER_NOT_FOUND = 600 + iota
	USER_ERROR
)
