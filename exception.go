package earthshaker

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//indicate exception wrapper

type Exception interface {
	//解析exception包装 error
	Unwrap() error
	//打印error 字符信息
	Error() string
	//异常栈
	Stack() []byte
	//打印异常堆栈
	PrintStackTrace(write ...io.Writer)
}

func AsException(e interface{}, stack []byte) Exception {
	switch err := e.(type) {
	case error:
		return exception{err, stack}
	default:
		return exception{error: fmt.Errorf("%v", e), stack: stack}
	}
}

type exception struct {
	error error
	stack []byte
}

func (e exception) Unwrap() error {
	return e.error
}

func (e exception) Error() string {
	return e.error.Error()
}

func (e exception) Stack() []byte {
	return e.stack
}

func (e exception) PrintStackTrace(writer ...io.Writer) {
	if 0 == len(writer) {
		writer = append(writer, os.Stderr)
	}
	var sb strings.Builder
	sb.WriteString("exception:")
	sb.WriteString(e.Error())
	sb.WriteString("\n")
	sb.Write(e.Stack())
	_, _ = io.Copy(io.MultiWriter(writer...), strings.NewReader(sb.String()))
}
