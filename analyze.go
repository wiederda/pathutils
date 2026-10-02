package pathutils

import (
	"path/filepath"
	"strings"
)

func Analyze(p string) PathInfo {
	clean := filepath.Clean(p)
	vol := filepath.VolumeName(clean)
	isAbs := filepath.IsAbs(clean)

	return PathInfo{
		Original: p,
		Cleaned:  clean,
		Type:     classify(clean, vol, isAbs),
		IsAbs:    isAbs,
		Volume:   vol,
		IsLong:   IsLongPath(clean),
	}
}

func classify(p, vol string, isAbs bool) PathType {
	if len(vol) > 1 && strings.HasPrefix(vol, `\\`) {
		return PathUNC
	}

	if isAbs {
		return PathAbsolute
	}

	if p != "" {
		return PathRelative
	}

	return PathUnknown
}
