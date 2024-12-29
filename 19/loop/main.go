package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}
	/* h1 := 0
	// slice长度
	hn := len(values)
	v := "" // for i,v := range 中的 v

	for ; h1 < hn; h1++ {
		v = values[h1]
		go func() {
			fmt.Println(v)
			done <- true
		}()
	} */

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}

}
