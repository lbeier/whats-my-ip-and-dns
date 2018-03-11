package main

import (
	"github.com/tutabeier/ipleak/pkg/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Handler)
	if err := http.ListenAndServe(":9999", nil); err != nil {
		panic(err)
	}
}
