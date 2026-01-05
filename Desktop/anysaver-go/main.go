package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("AnySaver Engine Starting... 🐹")

	fileUrl := "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"
	
	fileName := "google_logo.png"

	fmt.Println("Downloading:", fileUrl)

	response, err := http.Get(fileUrl)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("Error: File not found online! Status:", response.Status)
		return
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	size, err := io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Printf("වැඩේ ගොඩ! Downloaded a file of size %d bytes ✅\n", size)
}