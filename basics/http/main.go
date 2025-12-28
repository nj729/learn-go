package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(res.Body, err)
	// byteSlice := make([]byte, 99999)
	// res.Body.Read(byteSlice)
	// fmt.Println(string(byteSlice))
	lw := logWriter{}

	io.Copy(lw, res.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println("Bytes Processed : ", len(bs))
	return len(bs), nil
}
