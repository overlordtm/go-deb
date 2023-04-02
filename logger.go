package deb

import "log"

type LoggerLike interface {
	Printf(format string, args ...interface{})
	Println(args ...interface{})
	Print(args ...interface{})
}

var (
	// logger is the logger used by the deb package
	logger LoggerLike = log.Default()
)

func SetLogger(l LoggerLike) {
	logger = l
}
