//go:build !windows
// +build !windows

package selenium

import "syscall"

var cmdAttr = &syscall.SysProcAttr{}
