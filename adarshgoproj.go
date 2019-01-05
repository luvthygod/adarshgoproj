package adarshgoproj

import (
	"golang.org/x/sys/unix"
)

func HelloWorld() string {
	return "Hello"
}

// GetUserID will get the userid
func GetUserID() int {
	return unix.Getuid()

}
