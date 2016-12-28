package main

import (
	"flag"
	"fmt"
	"log"

	youtube "google.golang.org/api/youtube/v3"
)

var (
	dirDownloading = flag.String("dirDownloading", "C:\\Users\\gnimgirov\\Downloads", "Name of video file to upload")
)

func main() {
	client, err := buildOAuthHTTPClient([]string{
		youtube.YoutubeUploadScope,
		youtube.YoutubeScope,
		youtube.YoutubepartnerScope,
		youtube.YoutubeForceSslScope,
	})
	if err != nil {
		log.Fatalf("Error building OAuth client: %v", err)
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	videoCategoriesList := service.VideoCategories.List("snippet")

	list, err := videoCategoriesList.Id("Sports").Do()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(list)
	}

}
