package logUtil

import (
	"fmt"
	"log"
)

// LogHandler 日志处理类
type LogHandler struct {
	h    *log.Logger
	lv   string
	path string
}

// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *LogHandler) Println(v ...interface{}) { l.h.Output(2, fmt.Sprintln(v...)) }

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *LogHandler) Printf(format string, v ...interface{}) {
	l.h.Output(2, fmt.Sprintf(format, v...))
}

func (l *LogHandler) Path() string {
	return l.path
}

func New(logger *log.Logger, path, lv string) *LogHandler {
	return &LogHandler{
		h:    logger,
		lv:   lv,
		path: path}
}
