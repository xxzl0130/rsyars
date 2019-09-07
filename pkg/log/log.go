package log

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
)

type Logger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Tipsf(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
}

type color int

const (
	red    color = 31
	green        = 32
	yellow       = 33
	cyan         = 36
	white        = 37
)

type logger struct {
	enabled bool
	fd      *os.File
}

func (c *logger) Debugf(format string, v ...interface{}) {
	c.printf(white, "DEBUG", format, v...)
}

func (c *logger) Infof(format string, v ...interface{}) {
	c.printf(cyan, "INFO", format, v...)
}

func (c *logger) Tipsf(format string, v ...interface{}) {
	c.printf(green, "TIPS", format, v...)
}

func (c *logger) Warningf(format string, v ...interface{}) {
	c.printf(yellow, "WARNING", format, v...)
}

func (c *logger) Errorf(format string, v ...interface{}) {
	c.printf(red, "ERROR", format, v...)
}

func (c *logger) Fatalf(format string, v ...interface{}) {
	c.printf(red, "FATAL", format, v...)
	os.Exit(1)
}

func (c *logger) printf(tag color, typ, format string, v ...interface{}) {
	times := time.Now().Format("2006/01/02 15:04:05")
	msg := fmt.Sprintf(format, v...)
	if c.enabled {
		fmt.Printf("\x1B[%dm%s %s\x1B[0m\n", tag, times, msg)
	} else {
		fmt.Printf("[%s] %s %s\n", typ, times, msg)
	}
	_, _ = c.fd.WriteString(fmt.Sprintf("[%s] %s %s\n", typ, times, msg))
	_ = c.fd.Sync()
}

func New(logFileName string) (Logger, error) {
	logger := new(logger)
	if err := setEnableVirtualTerminalProcessing(); err != nil {
		logger.enabled = false
	} else {
		logger.enabled = true
	}
	fd, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE, 0)
	if err != nil {
		return nil, errors.WithMessagef(err, "打开日志文件 %s 失败", logFileName)
	}
	logger.fd = fd
	return logger, nil
}
