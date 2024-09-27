// Package utils @Author hubo 2024/9/26 16:20:00
package utils

import "os"

func DirExists(dirPath string) bool {
	fi, err := os.Stat(dirPath)
	if err == nil {
		return fi.IsDir()
	}
	return false
}
