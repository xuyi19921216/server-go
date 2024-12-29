package main

import "fmt"

type CustomError struct {
	code    int
	message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("code %d,msg %s", e.code, e.message)
}

var ErrInvalidParam = &CustomError{code: 10000, message: "invalid param"}

func handle(req string) error {
	var p *CustomError = nil
	if len(req) == 0 {
		p = ErrInvalidParam
	}
	return p
	/* if len(req) == 0 {
		return ErrInvalidParam
	}
	return nil */
}

func main() {
	req := ""
	err := handle(req)
	if err != nil {
		fmt.Printf("err1 req %s\n", req)
	}
	req = "aa" // 函数传入的参数不为空
	err = handle(req)
	if err != nil {
		fmt.Printf("err2 req %s\n", req)
	}

}
