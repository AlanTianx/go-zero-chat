package errorm

const ErrServerMsg = "服务器开小差啦，稍后再来试一试"

var msg = map[int]string{
	ErrRequestLimit:       "Number of requests exceeded the limit",
	ErrRequestParam:       "请求错误",
	ErrRequestParamAllNil: "请求错误：至少需要一项内容",
	ErrSendSms:            "发送短信失败",
	ErrCaptcha:            "验证码验证失败",
	ErrUpload:             "上传失败~请稍后再试",
	ErrUploadRType:        "上传失败-错误的资料类型",
	ErrUploadFileType:     "上传失败-文件类型错误",
	ErrUploadSize:         "上传失败-文件大小不适合",
}

func MapErrMsg(errCode int) string {
	if r, ok := msg[errCode]; ok {
		return r
	} else {
		return ErrServerMsg
	}
}
