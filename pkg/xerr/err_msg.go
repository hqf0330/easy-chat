package xerr

var codeText = map[int]string{
	SERVER_COMMON_ERROR: "服务器异常",
	REQUEST_PARAM_ERROR: "请求参数异常",
	DB_ERROR:            "数据库繁忙，请稍后重试",
}

func ErrMsg(errCode int) string {
	// todo 校验
	return codeText[errCode]
}
