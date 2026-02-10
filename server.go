package main

import (
	"fmt"
	"io"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Читаем сообщение от клиента
	body, _ := io.ReadAll(r.Body)
	fmt.Printf("Получено сообщение: %s\n", string(body))

	// Отправляем ответ клиенту
	fmt.Fprintf(w, "Сервер получил твое сообщение!")
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("HTTP Сервер запущен на :8000")
	// Слушаем порт 8000
	http.ListenAndServe(":8000", nil)
}
