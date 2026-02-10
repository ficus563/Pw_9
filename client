package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// СЮДА ВСТАВЬ ССЫЛКУ ИЗ CODESPACES (вкладка Ports -> Forwarded Address)
	url := "https://твой-адрес-codespaces-8000.app.github.dev/"

	message := []byte("Привет от клиента!")
	
	// Делаем запрос
	resp, err := http.Post(url, "text/plain", bytes.NewBuffer(message))
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ от сервера
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Ответ сервера:", string(body))
}
