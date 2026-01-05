package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/schollz/progressbar/v3" 
)

func main() {
	fmt.Println("AnySaver Video Downloader Starting... 🎥")

	fileUrl := "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerBlazes.mp4"
	fileName := "my_video.mp4"

	fmt.Println("Downloading Video:", fileName)

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

	fmt.Println("\nවැඩේ ඉවරයි! Video එක Download වුණා. Play කරලා බලන්න. ✅")
}