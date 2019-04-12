package exception

// 自定义用户异常实现
type MyCustomError string

func (e MyCustomError) Error() string {
	return e.Message()
}

func (e MyCustomError) Message() string {
	return string(e)
}