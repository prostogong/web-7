package main

import (
	"net/http"
)

func main() {
	// Регистрируем обработчики для разных путей
	mux_hello := http.NewServeMux()
	mux_hello.HandleFunc("/", handleRequest)

	// Запускаем веб-сервер на порту 8082
	go http.ListenAndServe(":8082", mux_hello)
	select {}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	// Разрешаем доступ из любых источников (или укажите конкретные)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain")

	w.Write([]byte("Hello, web!!!"))
}
