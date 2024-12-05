package performance

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Id   int
	Name string
}

// GenerateIdsRaw 原始待优化函数
func GenerateIdsRaw(users []*User) (string, string, []byte) {
	names := ""
	idStr := ""
	var nameByte []byte
	for index := range users {
		idStr = fmt.Sprint(users[index].Id)
		names = names + "," + users[index].Name
		nameByte = []byte(users[index].Name)
	}

	return idStr, names, nameByte
}

// GenerateIdsBuilder 使用strings.Builder拼接字符串
func GenerateIdsBuilder(users []*User) (string, string, []byte) {
	names := ""
	idStr := ""
	var nameByte []byte
	length := 0
	for index := range users {
		idStr = fmt.Sprint(users[index].Id)
		nameByte = []byte(users[index].Name)
		length += len(users[index].Name) + 1
	}
	var builder strings.Builder
	builder.Grow(length) // 预分配
	for index := range users {
		builder.WriteString(",")
		builder.WriteString(users[index].Name)
	}
	return idStr, names, nameByte
}

// GenerateIdsStrconv 使用strconv实现整型转字符串
func GenerateIdsStrconv(users []*User) (string, string, []byte) {
	names := ""
	idStr := ""
	var nameByte []byte
	length := 0
	for index := range users {
		idStr = strconv.Itoa(users[index].Id)
		nameByte = []byte(users[index].Name)
		length += len(users[index].Name) + 1
	}
	var builder strings.Builder
	builder.Grow(length) // 预分配
	for index := range users {
		builder.WriteString(",")
		builder.WriteString(users[index].Name)
	}
	return idStr, names, nameByte
}

func GenerateIdsUnsafe(users []*User) (string, string, []byte) {
	names := ""
	idStr := ""
	var nameByte []byte
	length := 0
	for index := range users {
		idStr = strconv.Itoa(users[index].Id)
		// unsafe包实现字符串转字节切片
		nameByte = Str2Bytes(users[index].Name)
		length += len(users[index].Name) + 1
	}
	var builder strings.Builder
	builder.Grow(length) // 预分配
	for index := range users {
		builder.WriteString(",")
		builder.WriteString(users[index].Name)
	}
	return idStr, names, nameByte
}
