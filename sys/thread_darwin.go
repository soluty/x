package sys

import (
	"fmt"
	"syscall"
)

func GetCurrentThreadId() (int, error) {
	tid, _, e := syscall.RawSyscall(syscall.SYS_THREAD_SELFID, 0, 0, 0)
	if e != 0 {
		return 0, fmt.Errorf("syscall err: ", e)
	}
	return int(tid), nil
}
