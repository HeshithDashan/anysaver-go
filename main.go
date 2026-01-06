package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/kkdai/youtube/v2" 
	"github.com/schollz/progressbar/v3"
)

func main() {
	fmt.Println("AnySaver Ultimate Downloader Starting... 🚀")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Video URL: ")
	urlInput, _ := reader.ReadString('\n')
	url := strings.TrimSpace(urlInput)

	if strings.Contains(url, "youtu") {
		fmt.Println("YouTube Link Detected! 🟥")
		downloadYoutubeVideo(url)
	} else {
		fmt.Println("Direct Link Detected! 🔗")
		downloadDirectFile(url)
	}
}

func downloadYoutubeVideo(videoURL string) {
	client := youtube.Client{}

	video, err := client.GetVideo(videoURL)
	if err != nil {
		fmt.Println("Error getting video info:", err)
		return
	}

	fmt.Printf("Downloading: %s \n", video.Title)

	formats := video.Formats.WithAudioChannels() 
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		fmt.Println("Error getting stream:", err)
		return
	}
	defer stream.Close()

	fileName := "youtube_video.mp4"
	file, _ := os.Create(fileName)
	defer file.Close()

	fmt.Println("Starting Download...")
	bar := progressbar.DefaultBytes(
		int64(formats[0].ContentLength),
		"downloading",
	)
	io.Copy(io.MultiWriter(file, bar), stream)
	fmt.Println("\nYouTube Video Saved! ✅")
}

func downloadDirectFile(fileUrl string) {
	fileName := "direct_download.mp4"

	req, _ := http.NewRequest("GET", fileUrl, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Connection Error:", err)
		return
	}
	defer resp.Body.Close()

	file, _ := os.Create(fileName)
	defer file.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	io.Copy(io.MultiWriter(file, bar), resp.Body)
	fmt.Println("\nFile Saved! ✅")
}