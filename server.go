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

func handleChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodPost {
		body, _ := io.ReadAll(r.Body)
		if len(body) > 0 {
			mutex.Lock()
			messages = append(messages, string(body))
			mutex.Unlock()
			fmt.Println(string(body))
		}
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	for _, msg := range messages {
		fmt.Fprintln(w, msg)
	}
}

func main() {
	http.HandleFunc("/", handleChat)
	http.ListenAndServe(":8000", nil)
}
