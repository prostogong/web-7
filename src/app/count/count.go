package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	var count int = 0
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
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
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера!", err)
	}
}
