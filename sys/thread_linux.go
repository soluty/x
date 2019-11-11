package sys

import "syscall"

func GetCurrentThreadId() (int, error) {
	return syscall.Gettid(), nil
}
