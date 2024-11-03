package main

import (
	"net/http"
	_ "net/http/pprof"
)

// http触发的方式
func main() {
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
