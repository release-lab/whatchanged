package logger

import (
	"fmt"
	"log"
	"os"
)

var printter *log.Logger

func init() {
	printter = log.New(os.Stderr, "[whatchanged]: ", log.Ltime)
}

func New(namespace string) *log.Logger {
	return log.New(os.Stderr, fmt.Sprintf("[whatchanged %s]: ", namespace), log.Ltime)
}

func Printf(format string, v ...interface{}) {
	printter.Printf(format, v...)
}

func Println(v ...interface{}) {
	printter.Println(v...)
}

func Print(v ...interface{}) {
	printter.Print(v...)
}
