package errors

import "libong/common/server/http/code"

var (
	REQUEST_PARAM_ERROR = code.Error(20001, "请求参数错误")
)
