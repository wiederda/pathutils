package pathutils

import (
	"path/filepath"
	"runtime"
	"strings"
)

func IsLongPath(p string) bool {
	if runtime.GOOS != "windows" {
		return false
	}
	return strings.HasPrefix(p, longPathPrefix)
}

func LongPath(p string) string {
	if runtime.GOOS != "windows" {
		return p
	}

	p = filepath.Clean(p)

	if strings.HasPrefix(p, longPathPrefix) {
		return p
	}

	if strings.HasPrefix(p, `\\`) {
		return `\\?\UNC\` + strings.TrimPrefix(p, `\\`)
	}

	return longPathPrefix + p
}

func FromLongPath(p string) string {
	if runtime.GOOS != "windows" {
		return p
	}

	if strings.HasPrefix(p, `\\?\UNC\`) {
		return `\\` + strings.TrimPrefix(p, `\\?\UNC\`)
	}

	if strings.HasPrefix(p, longPathPrefix) {
		return strings.TrimPrefix(p, longPathPrefix)
	}

	return p
}
