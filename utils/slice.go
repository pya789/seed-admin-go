package utils

import "fmt"

type cumsum interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string
}

// 切片转in()参数字符
func SliceToInStr[T cumsum](arr []T) string {
	str := ""
	for _, item := range arr {
		if len(str) == 0 {
			str += fmt.Sprint(item)
		} else {
			str += "," + fmt.Sprint(item)
		}
	}
	return str
}

// 查询切片内是否包含某个值
func SliceIncludes[T cumsum](arr []T, formIndex T) bool {
	is := false
	for _, item := range arr {
		if item == formIndex {
			is = true
		}
	}
	return is
}

// 切片删除
func SliceDelete[T cumsum](arr []T, elem T) []T {
	j := 0
	for _, item := range arr {
		if item != elem {
			arr[j] = item
			j++
		}
	}
	return arr[:j]
}
