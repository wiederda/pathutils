package pathutils

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

func AbsPath(p string) (string, error) {
	clean := filepath.Clean(p)

	if runtime.GOOS != "windows" && looksLikeWindowsPath(clean) {
		return "", fmt.Errorf("windows path not allowed on %s: %s", runtime.GOOS, clean)
	}

	abs, err := filepath.Abs(clean)
	if err != nil {
		return "", fmt.Errorf("failed to resolve absolute path: %w", err)
	}

	return abs, nil
}

func IsAbs(p string) bool {
	return filepath.IsAbs(p)
}

func looksLikeWindowsPath(p string) bool {
	if len(p) >= 2 && p[1] == ':' {
		return true
	}
	if strings.HasPrefix(p, `\\`) {
		return true
	}
	return false
}
