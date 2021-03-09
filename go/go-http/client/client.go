package main

import "fmt"
import "net/http"
import "bytes"
//import "encoding/json"

func main() {
	fmt.Println("Hello World!")
	resp, err := http.Get("https://quotes.rest/qod")
	if err!=nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	buf := bytes.NewBuffer(make([]byte, 1024))

	length,_ := buf.ReadFrom(resp.Body)

	fmt.Println(len(buf.Bytes()))
	fmt.Println(length)
	fmt.Println(string(buf.Bytes()))
}