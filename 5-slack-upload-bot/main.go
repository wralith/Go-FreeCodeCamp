package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while reading env file")
	}
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArray := []string{os.Getenv("CHANNEL_ID")}
	fileArrray := []string{"test.txt"}

	for i := 0; i < len(fileArrray); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArray,
			File:     fileArrray[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
