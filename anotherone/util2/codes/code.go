package codes

const (
	ErrUnknown       = 50000
	ErrServerGetFail = 50001

	ErrGinBindingQuery = 40001

	OK = 20000
)

var CodeErrorMap = map[int]string{
	ErrUnknown:       "未知错误",
	ErrServerGetFail: "服务器发送GET请求失败",

	ErrGinBindingQuery: "查询参数绑定出错",
}
