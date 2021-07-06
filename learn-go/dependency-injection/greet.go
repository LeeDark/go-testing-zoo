package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet injects io.Writer interface
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
