package log

import (
	"os"
	"syscall"
	"unsafe"

	"github.com/pkg/errors"
)

func setEnableVirtualTerminalProcessing() error {
	var placeholder syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, os.Stdout.Fd(), syscall.TIOCGETA, uintptr(unsafe.Pointer(&placeholder)), 0, 0, 0)
	if err != 0 {
		return errors.Errorf("ioctl error: %d", err)
	}
	return nil
}
