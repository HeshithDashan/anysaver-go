package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings" 

	"github.com/schollz/progressbar/v3"
)

func main() {
	fmt.Println("AnySaver Downloader Starting... 🚀")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter URL to download: ")

	urlInput, _ := reader.ReadString('\n')
	
	fileUrl := strings.TrimSpace(urlInput)

	fmt.Print("Enter name to save as (e.g. video.mp4): ")
	nameInput, _ := reader.ReadString('\n')
	fileName := strings.TrimSpace(nameInput)

	if fileName == "" {
		fileName = "downloaded_file.mp4"
	}

	fmt.Println("Downloading:", fileUrl)
	fmt.Println("Saving as:", fileName)

	req, _ := http.NewRequest("GET", fileUrl, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Connection Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Error: File not found! Status Code:", resp.StatusCode)
		return
	}

	file, _ := os.Create(fileName)
	defer file.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)

	io.Copy(io.MultiWriter(file, bar), resp.Body)

	fmt.Println("\nවැඩේ ඉවරයි! File Saved. ✅")
}