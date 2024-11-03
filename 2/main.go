package main

import (
	"log"
	"os"
	"runtime/pprof"
)

// http触发的方式
/* func main() {
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
} */

// 埋点触发的方式
func main() {
	// 创建CPU性能分析文件
	/* cpuProfile, err := os.Create("profile.pprof")
	if err != nil {
		log.Fatal("创建CPU性能分析文件失败:", err)
	}
	defer cpuProfile.Close()

	// 开始CPU性能分析
	if err := pprof.StartCPUProfile(cpuProfile); err != nil {
		log.Fatal("开始CPU性能分析失败:", err)
	}
	defer pprof.StopCPUProfile() */
	// 创建heap性能分析文件
	heapProfile, err := os.Create("heap.pprof")
	if err != nil {
		log.Fatal("创建内存性能分析文件失败:", err)
	}
	defer heapProfile.Close()

	// 开始heap内存性能分析
	if err := pprof.WriteHeapProfile(heapProfile); err != nil {
		log.Fatal("开始内存性能分析失败:", err)
	}
}
