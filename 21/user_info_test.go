package main

import (
	"fmt"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMock(t *testing.T) {
	// 使用gomonkey mock函数httpGetRequest的返回
	mockData := []byte(`{"name":"killianxu","age":32}`)
	patch := gomonkey.ApplyFunc(httpGetRequest, func(url string) ([]byte, error) {
		return mockData, nil
	})
	defer patch.Reset()

	// 模拟调用
	mockUserInfo, _ := fetchUserInfo("123")

	fmt.Printf("mocked user info: %s\n", mockUserInfo)
}

func TestFetchUserInfo(t *testing.T) {
	Convey("获取用户信息", t, func() {

		// 使用gomonkey mock函数httpGetRequest的返回
		mockData := []byte(`{"name":"killianxu","age":32}`)
		patch := gomonkey.ApplyFunc(httpGetRequest, func(url string) ([]byte, error) {
			return mockData, nil
		})
		defer patch.Reset()

		// 模拟调用
		result, _ := fetchUserInfo("123")
		So(result, ShouldEqual, string(mockData))

	})

}
