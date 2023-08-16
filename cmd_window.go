//go:build windows
// +build windows

package selenium

import "syscall"

var cmdAttr = &syscall.SysProcAttr{
	CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
}
