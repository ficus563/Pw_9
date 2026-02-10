package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	url := "ссылкусюда"

	fmt.Print("Введите ваш ник: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()

	go func() {
		lastCount := 0
		for {
			resp, err := http.Get(url)
			if err == nil {
				body, _ := io.ReadAll(resp.Body)
				lines := strings.Split(string(body), "\n")
				if len(lines) > lastCount {
					for i := lastCount; i < len(lines)-1; i++ {
						fmt.Println(lines[i])
					}
					lastCount = len(lines)
				}
				resp.Body.Close()
			}
			time.Sleep(2 * time.Second)
		}
	}()


	for scanner.Scan() {
		text := scanner.Text()
		if text == "" { continue }
		
		fullMsg := fmt.Sprintf("[%s]: %s", username, text)
		http.Post(url, "text/plain", bytes.NewBufferString(fullMsg))
	}
}
