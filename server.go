package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	messages []string
	mutex    sync.Mutex
)

func chatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if r.Method == http.MethodPost {
		body, _ := io.ReadAll(r.Body)
		if len(body) > 0 {
			mutex.Lock()
			messages = append(messages, string(body))
			mutex.Unlock()
			fmt.Println("Новое сообщение:", string(body))
		}
		return
	}


	mutex.Lock()
	for _, msg := range messages {
		fmt.Fprintln(w, msg)
	}
	mutex.Unlock()
}

func main() {
	http.HandleFunc("/", chatHandler)
	fmt.Println("Чат-сервер запущен на :8000")
	http.ListenAndServe(":8000", nil)
}
