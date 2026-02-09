package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("http://localhost:8080/hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		panic(string(body))
	}

	contentType := resp.Header.Get("Content-Type")

	// 🔹 กรณีเป็น JSON

	switch contentType {
	case "application/json":
		var result Response
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			panic(err)
		}

		fmt.Println("JSON response:", result.Message)
		return
	case "text/plain":
		// 🔹 กรณีเป็น string ธรรมดา
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println("string response:", string(body))
	default:
		fmt.Println("content type is not support")
	}

}
