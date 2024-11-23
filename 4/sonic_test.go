package main

import (
	"testing"

	"github.com/bytedance/sonic"
)

type User struct {
	Name string
	Age  int
}

// sonic序列化/反序列化
func TestSonic(t *testing.T) {
	user := User{Name: "test", Age: 18}
	// sonic序列化
	data, err := sonic.Marshal(user)
	if err != nil {
		panic(err)
	}
	var newUser User
	err = sonic.Unmarshal(data, &newUser)
	if err != nil {
		panic(err)
	}

}
