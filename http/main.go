package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Size of input: ", len(bs))
	return len(bs), nil
}

func main() {
	lw := logWriter{}
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Errorf("HTTP Request error: %v", err)
		os.Exit(1)
	}

	//fmt.Printf("%v\n", resp)

	// body := make([]byte, 99999)
	// resp.Body.Read(body)
	// fmt.Printf("%v\n", string(body))

	io.Copy(lw, resp.Body)
}
