package errors

import "net/http"

//  全局的错误处理函数，在 api-gateway 中使用 自定义错误
func ErrHandler(err error)  (int, interface{}) {
	if se := FromError(err); err != nil {
		return se.Code, se.ErrData()
	} else{
		return http.StatusInternalServerError, nil
	}
}

