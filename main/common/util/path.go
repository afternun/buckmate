package util

import "strings"

func Resolve(workDir string, path string) string {
	if len(path) > 0 {
		if strings.HasPrefix(path, "/") {
			return path
		}

		return workDir + "/" + path
	}
	return workDir
}
