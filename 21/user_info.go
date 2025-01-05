package main

import (
	"io"
	"net/http"
)

// 发送HTTP GET请求并返回响应的函数
func httpGetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// 依赖httpGetRequest函数获取用户信息的函数
func fetchUserInfo(userID string) (string, error) {
	url := "https://example.com/api/user/" + userID
	data, err := httpGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
