package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
)

func main() {
	for _, url:=range os.Args[1:] {
		resp, err := http.Get(url)
		if err!=nil {
			fmt.Fprintf(os.Stderr, "test: %v\n", err)
			os.Exit(1)
		}
		data, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err!=nil {
			fmt.Fprintf(os.Stderr, "test: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", data)
	}
}
