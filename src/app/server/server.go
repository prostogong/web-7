package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var count int = 0

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello, web!!!"))
}

func handleCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte(strconv.Itoa(count)))
	case http.MethodPost:
		number, err := strconv.Atoi(r.FormValue("count"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Некорректное число"))
		} else {
			count += number
			w.Write([]byte(fmt.Sprintf("Текущее значение: %d", count)))
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("name"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHello)
	mux.HandleFunc("/count", handleCount)
	mux.HandleFunc("/api/user", handleUser)

	go func() {
		err := http.ListenAndServe(":8082", mux)
		if err != nil {
			fmt.Println("Ошибка запуска сервера на порту 8082:", err)
		}
	}()

	go func() {
		err := http.ListenAndServe(":8081", mux)
		if err != nil {
			fmt.Println("Ошибка запуска сервера на порту 8081:", err)
		}
	}()

	go func() {
		err := http.ListenAndServe(":8083", mux)
		if err != nil {
			fmt.Println("Ошибка запуска сервера на порту 8083:", err)
		}
	}()

	select {}
}
