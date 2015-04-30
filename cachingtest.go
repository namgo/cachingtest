package main

import (
	"io"
	"net/http"
	"strconv"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, strconv.Itoa(int(time.Now().UnixNano())))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
