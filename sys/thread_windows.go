package sys

import (
	"syscall"
)

func GetCurrentThreadId() (int, error) {
	var user32 *syscall.DLL
	var GetCurrentThreadId *syscall.Proc
	var err error

	user32, err = syscall.LoadDLL("Kernel32.dll")
	if err != nil {
		return 0, err
	}
	GetCurrentThreadId, err = user32.FindProc("GetCurrentThreadId")
	if err != nil {
		return 0, err
	}

	var pid uintptr
	pid, _, err = GetCurrentThreadId.Call()

	return int(pid), err
}
