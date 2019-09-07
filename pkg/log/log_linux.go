package log

import (
	"os"

	"golang.org/x/sys/unix"
)

func setEnableVirtualTerminalProcessing() error {
	if _, err := unix.IoctlGetTermios(int(os.Stdout.Fd()), unix.TCGETS); err != nil {
		return err
	}
	return nil
}
