package logger

import (
	"log"
	"os"
	"runtime"
	"strings"
)

func Logln(msgs ...string) {
	_, file, line, _ := runtime.Caller(1)
	file = strings.Split(file, "/")[len(strings.Split(file, "/"))-1]
	logger := log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime)
	msg := strings.Join(msgs, " ")
	logger.Printf("%s:%d: %s", file, line, msg)
}

func Logf(msg string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	file = strings.Split(file, "/")[len(strings.Split(file, "/"))-1]
	logger := log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime)
	logger.Printf("%s:%d: "+msg, append([]interface{}{file, line}, args...)...)
}

func Fatalf(msg string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	file = strings.Split(file, "/")[len(strings.Split(file, "/"))-1]
	logger := log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime)
	logger.Fatalf("%s:%d: "+msg, append([]interface{}{file, line}, args...)...)
}
