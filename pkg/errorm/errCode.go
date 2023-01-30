package errorm

// num < 1000 在api返回时将errMsg全部拦截改为 "服务器开小差啦，稍后再来试一试"
const (
	Success       = 0   // 请求成功
	Failed        = 1   // 请求失败
	ErrDbParams   = 101 // 数据库请求参数错误
	ErrDbNotFound = 102 // 数据库请求记录未找到
	ErrRedis      = 201 // redis请求错误
	ErrSmsInit    = 301 //短信服务错误

	ErrRequestLimit = 10001 // 请求被限流

	ErrRequestParam       = 20001 // 请求参数错误
	ErrRequestParamAllNil = 20002 // 请求参数全部为空

	ErrSendSms = 30001 //发送短信失败

	ErrCaptcha = 40001 //短信验证码错误

	ErrUpload         = 50001 //上传失败
	ErrUploadRType    = 50002 // 上传失败-资料类型错误
	ErrUploadFileType = 50003 // 上传失败-文件类型错误
	ErrUploadSize     = 50004 // 上传失败-文件大小不符合
)
