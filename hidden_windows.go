//go:build windows

package pathutils

import (
	"path/filepath"
	"syscall"
)

func IsHidden(p string) bool {
	ptr, err := syscall.UTF16PtrFromString(filepath.Clean(p))
	if err != nil {
		return false
	}
	attrs, err := syscall.GetFileAttributes(ptr)
	if err != nil {
		return false
	}
	return attrs&syscall.FILE_ATTRIBUTE_HIDDEN != 0
}
