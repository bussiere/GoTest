package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

func encode(data []byte) *bytes.Buffer {
	bb := &bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, bb)
	encoder.Write([]byte(data))
	encoder.Close()
	return bb
}

func main() {
	data := make([]byte, 0, 25e5)
	file, err := ioutil.ReadFile("/home/bussiere/workspace/Gotest/gobook.pdf")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		data = append(data, file...)
	}
	bb := encode(data)
	fmt.Println(len(data), bb.Len())
	fmt.Println(len(data), bb)
}
