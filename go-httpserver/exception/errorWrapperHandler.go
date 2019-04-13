package exception

import (
	"net/http"
	"log"
	"os"
)

// 定义一个 function 类型的 type, 返回值是 error
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 输入 appHandler 是一个函数，输出也是一个函数
func ErrWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 1.处理逻辑业务
		err := handler(writer, request)
		if err != nil {
			log.Print("error occured, ", err)

			// 2. 处理可以抛给用户的错误
			if err, ok := err.(UserError); ok {
				// 将错误写回到 http.ResponseWriter
				http.Error(writer, err.Message(), http.StatusBadRequest)
			}

			// 3. 处理不可以抛给用户的错误
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusOK
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}