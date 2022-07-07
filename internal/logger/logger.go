// logging functions

package logger

import "fmt"

// Info prints a structured logging message
func Info(operationName string, msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	fmt.Printf("[+] [info] %s: %s\n", operationName, msg)
}

// Success prints a structured success message
func Success(operationName string, msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	fmt.Printf("[+] [success] %s: %s\n", operationName, msg)
}

// Debug prints a debug logging message
func Debug(operationName string, msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	fmt.Printf("[+] [debug] -- %s: %s\n", operationName, msg)
}

// Error prints a structured error message
func Error(err error, operationName string) {
	fmt.Printf("[error] %s failed -- error: %s\n", operationName, err.Error())
}
