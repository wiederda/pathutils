//go:build windows

package pathutils

import (
	"os"
	"path/filepath"
	"strings"
)

func Exists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

func IsDir(p string) bool {
	info, err := os.Stat(p)
	return err == nil && info.IsDir()
}

func NameWithoutExt(p string) string {
	base := filepath.Base(p)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(base, ext)
}

func ReplaceExt(p, ext string) string {
	if ext != "" && !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	return strings.TrimSuffix(p, filepath.Ext(p)) + ext
}

func EnsureExt(p, ext string) string {
	if ext != "" && !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	if filepath.Ext(p) == ext {
		return p
	}
	return ReplaceExt(p, ext)
}
