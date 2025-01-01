package main

import (
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
)

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		//return nil, err
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()
	buf, err := io.ReadAll(f)
	if err != nil {
		//return nil, err
		return nil, errors.Wrap(err, "read failed")
	}
	return buf, nil
}

func main() {
	filePath := "config.xml"
	_, err := ReadFile(filePath)

	if err != nil {
		log.Printf("read file failed, file_path:%s, due to %+v\n", filePath, err)
	}
}
