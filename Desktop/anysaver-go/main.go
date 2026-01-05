package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("AnySaver Engine Starting... 🐹")

	// 1. Download කරන්න ඕන ලින්ක් එක (Google Logo - මේක ෂුවර්)
	fileUrl := "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"
	
	// අපි Save කරන නම මෙතන දෙනවා (මේක තමයි කලින් මිස් වෙලා තිබ්බේ)
	fileName := "google_logo.png"

	fmt.Println("Downloading:", fileUrl)

	// 2. Internet එකෙන් File එක ඉල්ලනවා
	response, err := http.Get(fileUrl)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}
	defer response.Body.Close()

	// Check කරමු File එක ඇත්තටම තියෙනවද කියලා
	if response.StatusCode != 200 {
		fmt.Println("Error: File not found online! Status:", response.Status)
		return
	}

	// 3. අපේ මැෂින් එකේ හිස් ෆයිල් එකක් හදනවා
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 4. මැජික් එක: Internet එකෙන් එන Data ටික අපේ ෆයිල් එකට පුරවනවා
	size, err := io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Printf("වැඩේ ගොඩ! Downloaded a file of size %d bytes ✅\n", size)
}