package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"runtime/trace"
	"sync"
)

// 埋点触发的方式
func main() {
	// 创建trace性能分析文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	// your program here
	RunMyProgram1()
}
func RunMyProgram1() {
	// create new channel of type int
	ch := make(chan int)

	// start new anonymous goroutine
	go func() {
		// send 42 to channel
		ch <- 42
	}()
	// read from channel
	<-ch
}

func RunMyProgram2() {
	f, err := os.Create("test.png")
	if err != nil {
		log.Fatal(err)
	}
	img := createSeq(500, 300)
	if err = png.Encode(f, img); err != nil {
		panic(err)
	}
}
func createSeq(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	var w sync.WaitGroup
	w.Add(width * height)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			go func(i, j int) {
				m.Set(i, j, pixel(i, j))
				w.Done()
			}(i, j)
		}
	}
	w.Wait()
	return m
}
func pixel(i, j int) color.Color {
	return color.Gray{uint8(i * j % 255)}
}
