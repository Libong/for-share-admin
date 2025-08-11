package errors

import "libong/common/server/http/code"

var (
	PERMISSION_FAIL     = code.Error(2000, "权限非法")
	REQUEST_PARAM_ERROR = code.Error(20001, "请求参数错误")
)
