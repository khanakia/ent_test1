package filesystem

import "os"

// VisibilityConverter is an interface for converting visibility settings
// between integer representations and string representations.
type VisibilityConverter interface {
	// ForFile converts the visibility string to an integer representation for files.
	ForFile(visibility string) os.FileMode

	// ForDirectory converts the visibility string to an integer representation for directories.
	ForDirectory(visibility string) os.FileMode

	// InverseForFile converts the integer visibility back to string representation for files.
	InverseForFile(visibility os.FileMode) string

	// InverseForDirectory converts the integer visibility back to string representation for directories.
	InverseForDirectory(visibility os.FileMode) string

	// DefaultForDirectories returns the default visibility setting for directories as an integer.
	DefaultForDirectories() os.FileMode
}
