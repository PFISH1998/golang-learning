package exception

// 自定义用户异常接口
type UserError interface {
	error
	Message() string
}