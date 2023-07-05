package utils

import (
	"log"
	"runtime"
)

func GetStack() string {
	buf := make([]byte, 1<<12)
	runtime.Stack(buf, false)
	return string(buf)
}

func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		log.Println("err:", p, GetStack())
	}
}

func RunSafe(fn func()) {
	defer Recover()
	fn()
}
