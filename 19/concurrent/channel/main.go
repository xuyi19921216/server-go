package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Obj struct {
}

func handle(timeout time.Duration) *Obj {
	ch := make(chan *Obj)
	//ch := make(chan *Obj, 1)

	go func() {
		result := fn() // 逻辑处理
		ch <- result   // block
	}()
	select {
	case result := <-ch:
		return result
	case <-time.After(timeout):
		return nil
	}
}

func fn() *Obj {
	time.Sleep(10 * time.Second)
	return &Obj{}
}

func call() (string, error) {
	resp, err := http.Get("http://example.com")
	if err != nil {
		return "", err
	}
	// 不读Body也必须Close，否则会协程泄漏
	// defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status_code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
