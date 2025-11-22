package main

import (
	"net/http"

	"simple-webpage-go/api"
)

func main() {
	http.ListenAndServe(":8080", handler.Handler)
}
