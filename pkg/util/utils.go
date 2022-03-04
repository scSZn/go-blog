package util

import (
	"fmt"
	"runtime"
)

func GetCallerFileAndLine(depth int) string {
	_, file, line, ok := runtime.Caller(depth)
	if ok {
		return fmt.Sprintf("%s: %d", file, line)
	}
	return ""
}
