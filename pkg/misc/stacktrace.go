package misc

import "runtime"

const maxStackTraceSize = 8192 * 4

// GetStackTrace returns current stacktrace
func GetStackTrace() string {
	b := make([]byte, maxStackTraceSize)
	n := runtime.Stack(b, false)
	return string(b[:n])
}
