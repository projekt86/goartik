package goartik

import (
	"path"
	"runtime"
)

func getRelativePath(p string) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return p
	}
	return path.Join(path.Dir(filename), p)
}
