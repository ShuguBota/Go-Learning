package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	// Simple approach
	// buffer := make([]byte, 99999)
	// size, _ := resp.Body.Read(buffer)
	// fmt.Printf("%s", buffer[:size])

	io.Copy(os.Stdout, resp.Body)
}