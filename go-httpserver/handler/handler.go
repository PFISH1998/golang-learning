package handler

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
	"../exception"
	"fmt"
)

const PathPrefix = "/list/"

// 实现一个读取文件的 HttpServer 处理器
// 假设访问 http://localhost:8888/list/abc.txt
func HandleFileListing(writer http.ResponseWriter, request *http.Request) error {
	// 1. 如果 url 不是以 /list/ 开头， 自定义用户错误
	if strings.Index(request.URL.Path, PathPrefix) != 0 {
		return exception.MyCustomError("url path need startWith /list/")
	}
	fmt.Println("path", request.URL.Path)
	// 字符串切割
	path := request.URL.Path[len(PathPrefix):]
	// 2. 打开文件
	file, err := os.Open(path)
	if err != nil {
		// 遇到错误直接返回，由错误统一处理器进行处理
		return err
	}
	defer file.Close()

	// 3. 读取文件到 byte[]
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// 4. 将 byte[] all 写出到响应流
	writer.Write(all)
	return nil
}