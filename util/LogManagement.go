package util

import (
	"fmt"
	"os"
)

var LogManage *Logger

// Logger 结构体
type Logger struct {
	LogChan chan string
	Done    chan bool
}

// NewLogger 创建一个新的Logger实例
func NewLogger() *Logger {
	l := &Logger{
		LogChan: make(chan string, 100), // 缓冲区大小为100
		Done:    make(chan bool),
	}
	//go l.start()
	return l
}

// Log 记录日志
func LogSend(msg string) {
	LogManage.LogChan <- msg
}

func LogError(err error) {
	LogManage.LogChan <- err.Error()
}

// start 启动日志记录协程
func (l *Logger) start() {
	for {
		select {
		case msg := <-l.LogChan:
			_, err := fmt.Fprintln(os.Stdout, msg)
			if err != nil {
				return
			}
		case <-l.Done:
			return
		}
	}
}

// Stop 停止日志记录
func Stop() {
	LogManage.Done <- true
}

func init() {
	LogManage = NewLogger()
}
