package main

import "net/http"

func setupRoutes() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
}

func main() {
	setupRoutes()

	panic(http.ListenAndServe(":8080", nil))
}
