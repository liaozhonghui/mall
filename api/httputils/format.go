package httputils

import "fmt"

var emptyStruct struct{}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success() interface{} {
	return response{Code: OK.Code, Msg: OK.Msg, Data: emptyStruct}
}
func SuccessWithData(v interface{}) interface{} {
	return response{Code: OK.Code, Msg: OK.Msg, Data: v}
}

func Error(err error) interface{} {
	switch t := err.(type) {
	case response:
		return response{Code: t.Code, Msg: t.Msg, Data: emptyStruct}
	default:
		return response{Code: InternalError.Code, Msg: err.Error(), Data: emptyStruct}
	}
}

func (r response) Error() string {
	return fmt.Sprintf("%d|%s", r.Code, r.Msg)
}
