package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
		return
	}

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"test.txt"}

	for i := 0; i < len(fileArr); i++ {
		file, err := os.Open(fileArr[i])
		if err != nil {
			fmt.Printf("Error opening file: %s\n", err)
			continue
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Printf("Error getting file info: %s\n", err)
			continue
		}

		params := slack.UploadFileV2Parameters{
			Channel:  channelArr[0],
			File:     fileArr[i],
			Filename: fileArr[i],
			Reader:   file,
			FileSize: int(fileInfo.Size()),
		}
		f, err := api.UploadFileV2(params)
		if err != nil {
			fmt.Printf("Error uploading file: %s\n", err)
			continue
		}
		fmt.Printf("File uploaded successfully: %s\n", f.ID)
	}

}