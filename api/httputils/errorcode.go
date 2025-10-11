package httputils

var (
	OK         = response{Code: 0, Msg: "OK"}
	ParamError = response{Code: 400, Msg: "参数错误"}
	AuthError  = response{Code: 401, Msg: "认证失败"}
	Forbidden  = response{Code: 403, Msg: "没有权限"}
	NotFound   = response{Code: 404, Msg: "资源不存在"}

	InternalError = response{Code: 500, Msg: "服务器异常"}
)
