package main

import (
	"io"
	"net/http"
	"sync"
	"testing"
)

var urls = []string{
	"http://www.golang.org/",
	"http://www.google.com/",
	"http://www.somestupidname.com/",
}

func TestWaitGroup(t *testing.T) {
	// 创建WaitGroup对象
	wg := sync.WaitGroup{}
	results := make([]string, len(urls))
	for index, url := range urls {
		url := url
		index := index
		// 在创建协程执行任务之前，调用Add方法
		wg.Add(1)
		go func() {
			// 任务完成后，调用Done方法
			defer wg.Done()
			// Fetch the URL.
			resp, err := http.Get(url)
			if err != nil {
				return
			}

			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return
			}
			results[index] = string(body)

		}()
	}
	// 主协程阻塞，等待所有的任务执行完成
	wg.Wait()
}

func TestGetAllErr(t *testing.T) {
	// 创建WaitGroup对象
	wg := sync.WaitGroup{}
	results := make([]string, len(urls))
	// 保存所有并发任务执行结果
	errors := make([]error, len(urls))

	for index, url := range urls {
		url := url
		index := index
		// 在创建协程执行任务之前，调用Add方法
		wg.Add(1)
		go func() {
			// 任务完成后，调用Done方法
			defer wg.Done()
			// Fetch the URL.
			resp, err := http.Get(url)
			if err != nil {
				errors[index] = err // 保存执行结果
				return
			}

			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				errors[index] = err // 保存执行结果
				return
			}
			results[index] = string(body)

		}()
	}
	// 主协程阻塞，等待所有的任务执行完成
	wg.Wait()

}
