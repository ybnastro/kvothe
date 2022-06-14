package utils

import (
	"runtime"
	"strings"
)

const (
	allocatePtr = 15
	idxCallers  = 2

	dotRune   = '.'
	slashRune = '/'
)

// GetFuncName is a func to get func name from runtime
func GetFuncName() string {
	pc := make([]uintptr, allocatePtr)

	n := runtime.Callers(idxCallers, pc)

	frames := runtime.CallersFrames(pc[:n])
	_, _ = frames.Next()

	f := runtime.FuncForPC(pc[0])
	fName := f.Name()

	var lastSlash, lastDot int

	if strings.Contains(fName, string(slashRune)) {
		lastSlash = strings.LastIndexByte(fName, slashRune)
		lastDot = strings.LastIndexByte(fName[lastSlash:], dotRune) + lastSlash
		return fName[lastDot+1:]
	}
	return fName
}
