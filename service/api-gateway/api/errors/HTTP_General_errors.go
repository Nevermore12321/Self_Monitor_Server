package errors

// BadRequest 400
func BadRequest(reason, message string) *HTTPError {
	return Newf(400, reason, message)
}

// IsBadRequest 判断是否是 400 错误
func IsBadRequest(err error) bool {
	return Code(err) == 400
}

// Unauthorized 401 未授权
func Unauthorized(reason, message string) *HTTPError {
	return Newf(401, reason, message)
}

// IsUnauthorized 判断是否是401未授权
func IsUnauthorized(err error) bool {
	return Code(err) == 401
}

// Forbidden 403 禁止
func Forbidden(reason, message string) *HTTPError {
	return Newf(403, reason, message)
}

// IsForbidden 判断是否是403
func IsForbidden(err error) bool {
	return Code(err) == 403
}

// NotFound 404 not found
func NotFound(reason, message string) *HTTPError {
	return Newf(404, reason, message)
}

// IsNotFound 判断是否是 404 not found
func IsNotFound(err error) bool {
	return Code(err) == 404
}


// Conflict 409 服务器完成客户端的PUT请求是可能返回此代码，服务器处理请求时发生了冲突
func Conflict(reason, message string) *HTTPError {
	return Newf(409, reason, message)
}

// IsConflict 判断是否是 409
func IsConflict(err error) bool {
	return Code(err) == 409
}

// InternalServer 500 服务器内部错误，无法完成请求
func InternalServer(reason, message string) *HTTPError {
	return Newf(500, reason, message)
}

// IsInternalServer 判断是否是500
func IsInternalServer(err error) bool {
	return Code(err) == 500
}

// ServiceUnavailable 503 由于超载或系统维护，服务器暂时的无法处理客户端的请求
func ServiceUnavailable(reason, message string) *HTTPError {
	return Newf(503, reason, message)
}

func IsServiceUnavailable(err error) bool {
	return Code(err) == 503
}

// GatewayTimeout 504 bad gateway
func GatewayTimeout(reason, message string) *HTTPError {
	return Newf(504, reason, message)
}

func IsGatewayTimeout(err error) bool {
	return Code(err) == 504
}

// ClientClosed 499
func ClientClosed(reason, message string) *HTTPError {
	return Newf(499, reason, message)
}

func IsClientClosed(err error) bool {
	return Code(err) == 499
}
