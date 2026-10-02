package pathutils

import (
	"path/filepath"
	"runtime"
	"strings"
)

func Normalize(p string) string {
	p = filepath.Clean(strings.TrimSpace(p))

	if runtime.GOOS != "windows" {
		return p
	}

	if !filepath.IsAbs(p) {
		return p
	}

	if strings.HasPrefix(p, longPathPrefix) {
		return p
	}

	if !needsLongPath(p) {
		return p
	}

	if strings.HasPrefix(p, `\\`) {
		return `\\?\UNC\` + strings.TrimPrefix(p, `\\`)
	}

	return longPathPrefix + p
}

func needsLongPath(p string) bool {
	return len(p) >= 248
}

func NormalizeForCompare(p string) string {
	return filepath.ToSlash(filepath.Clean(p))
}
