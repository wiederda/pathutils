//go:build !windows

package pathutils

import (
	"path/filepath"
	"strings"
)

func IsHidden(p string) bool {
	base := filepath.Base(p)
	return strings.HasPrefix(base, ".")
}
