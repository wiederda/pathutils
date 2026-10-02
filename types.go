package pathutils

type PathType string

const (
	PathUNC      PathType = "UNC"
	PathAbsolute PathType = "ABSOLUTE"
	PathRelative PathType = "RELATIVE"
	PathUnknown  PathType = "UNKNOWN"
)

type PathInfo struct {
	Original string
	Cleaned  string
	Type     PathType
	IsAbs    bool
	Volume   string
	IsLong   bool
}
