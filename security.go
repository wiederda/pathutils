package pathutils

import (
	"path/filepath"
	"runtime"
	"strings"
)

// IsSubPath prüft, ob child innerhalb von parent liegt (strict sandbox check)
func IsSubPath(parent, child string) bool {
	p := filepath.Clean(parent)
	c := filepath.Clean(child)

	if runtime.GOOS == "windows" {
		p = strings.ToLower(p)
		c = strings.ToLower(c)
	}

	rel, err := filepath.Rel(p, c)
	if err != nil {
		return false
	}

	return !strings.HasPrefix(rel, "..")
}