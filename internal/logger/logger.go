// logging functions

package logger

import "fmt"

func Info(operationName string, msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	fmt.Printf("[+] %s: %s", operationName, msg)
}
func Debug(operationName string, msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	fmt.Printf("[+] DEBUG -- %s: %s", operationName, msg)
}
