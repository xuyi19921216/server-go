package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	slices := getCapacitySlices()
	getCapacitySet(slices)
	// 设置响应头，这里设置Content-Type为text/plain，表示返回纯文本内容
	w.Header().Set("Content-Type", "text/plain")
	// 向客户端写入响应内容
	fmt.Fprintln(w, "hh")
}
func main() {
	// 注册路由，当客户端访问根路径"/"时，会调用Handler函数进行处理
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

// getRawSlices 未提前指定容量
func getRawSlices() []int {
	n := 10000
	slices := make([]int, 0)
	for i := 0; i < n; i++ {
		slices = append(slices, i)
	}
	return slices
}

// getCapacitySlices 提前指定容量
func getCapacitySlices() []int {
	n := 10000
	slices := make([]int, 0, n)
	for i := 0; i < n; i++ {
		slices = append(slices, i)
	}
	return slices
}

// 用bool表示集合
func getRawSet(slices []int) map[int]bool {
	set := make(map[int]bool, 0)
	for _, item := range slices {
		set[item] = true
	}
	return set
}

// 用空结构体表示集合
func getEmptyStructSet(slices []int) map[int]struct{} {
	set := make(map[int]struct{}, 0)
	for _, item := range slices {
		set[item] = struct{}{}
	}
	return set
}

// 提前指定容量
func getCapacitySet(slices []int) map[int]struct{} {
	set := make(map[int]struct{}, len(slices))
	for _, item := range slices {
		set[item] = struct{}{}
	}
	return set
}
