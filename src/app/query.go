package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("name"))
}

func main() {
	http.HandleFunc("/api/user", handler)

	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
