package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := `{"name": "killianxu", "age": 18}`
	var mp map[string]interface{}

	json.Unmarshal([]byte(str), &mp)
	age := mp["age"].(int) // 报错：panic: interface conversion: interface {} is float64, not int
	fmt.Println(age)
}
