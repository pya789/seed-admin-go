package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// 获取当前运行路径
func CurrentDirectory() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return strings.Replace(dir, "\\", "/", -1), nil
}

// 判断目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 检查拓展名
func CheckExtension(path string) string {
	for i := len(path) - 1; i > 0; i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}
	return ""
}
