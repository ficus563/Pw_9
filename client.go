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
	url := "https://your-codespace-url-8000.app.github.dev/"

	fmt.Print("Ник: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nickname := scanner.Text()

	go func() {
		lastCount := 0
		for {
			resp, err := http.Get(url)
			if err == nil {
				body, _ := io.ReadAll(resp.Body)
				lines := strings.Split(strings.TrimSpace(string(body)), "\n")
				if len(lines) > lastCount && lines[0] != "" {
					for i := lastCount; i < len(lines); i++ {
						fmt.Println(lines[i])
					}
					lastCount = len(lines)
				}
				resp.Body.Close()
			}
			time.Sleep(time.Second)
		}
	}()

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		msg := fmt.Sprintf("[%s]: %s", nickname, text)
		http.Post(url, "text/plain", bytes.NewBufferString(msg))
	}
}
