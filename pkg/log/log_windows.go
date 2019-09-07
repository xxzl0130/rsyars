package log

import (
	"syscall"
	"unsafe"

	"github.com/pkg/errors"
)

const (
	kKernel32Name       = "kernel32.dll"
	kGetStdHandleName   = "GetStdHandle"
	kGetConsoleModeName = "GetConsoleMode"
	kSetConsoleModeName = "SetConsoleMode"

	kStdOutputHandle                 = uintptr(4294967285)
	kInvalidHandleValue              = ^uintptr(0)
	kEnableVirtualTerminalProcessing = 0x0004
)

var (
	addrGetStdHandle   = mustGetProcAddress(mustLoadLibrary(kKernel32Name), kGetStdHandleName)
	addrGetConsoleMode = mustGetProcAddress(mustLoadLibrary(kKernel32Name), kGetConsoleModeName)
	addrSetConsoleMode = mustGetProcAddress(mustLoadLibrary(kKernel32Name), kSetConsoleModeName)
)

func setEnableVirtualTerminalProcessing() error {
	h, _, err := syscall.Syscall(addrGetStdHandle, 1, kStdOutputHandle, 0, 0)
	if h == kInvalidHandleValue {
		return errors.WithMessage(err, "get std output handle failed")
	}

	mode := uint32(0)
	r, _, err := syscall.Syscall(addrGetConsoleMode, 2, h, uintptr(unsafe.Pointer(&mode)), 0)
	if r == 0 {
		return errors.WithMessage(err, "get console mode failed")
	}

	mode = mode | kEnableVirtualTerminalProcessing
	r, _, err = syscall.Syscall(addrSetConsoleMode, 2, h, uintptr(mode), 0)
	if r == 0 {
		return errors.WithMessage(err, "set console mode failed")
	}

	return nil
}

func mustLoadLibrary(name string) uintptr {
	lib, err := syscall.LoadLibrary(name)
	if err != nil {
		panic(err)
	}
	return uintptr(lib)
}

func mustGetProcAddress(lib uintptr, name string) uintptr {
	addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
	if err != nil {
		panic(err)
	}
	return uintptr(addr)
}
