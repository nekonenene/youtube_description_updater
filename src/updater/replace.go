package updater

import (
	"fmt"
	"log"
	"strings"
)

func ReplaceDescription() {
	service := getService()

	nextPageToken := ""
	counter := 0

	for {
		searchCall := service.Search.List("id").ForMine(true).Type("video").Order("date").MaxResults(maxResult).PageToken(nextPageToken)

		searchResponse, err := searchCall.Do()
		if err != nil {
			log.Fatalf("Error making API call to search: %v", err.Error())
		}

		var videoIds []string
		for _, video := range searchResponse.Items {
			videoId := video.Id.VideoId
			videoIds = append(videoIds, videoId)
		}

		videoListCall := service.Videos.List("id, snippet").Id(strings.Join(videoIds, ",")).MaxResults(maxResult)

		videoListResponse, err := videoListCall.Do()
		if err != nil {
			log.Fatalf("Error making API call to list videos: %v", err.Error())
		}

		for _, video := range videoListResponse.Items {
			videoId := video.Id
			title := video.Snippet.Title
			description := video.Snippet.Description

			if strings.Contains(description, params.TargetString) {
				counter += 1
				fmt.Printf("%v (%v)\n", title, videoId)

				description = strings.Replace(description, params.TargetString, params.ReplacementString, -1)
				video.Snippet.Description = description
				_, err := service.Videos.Update("snippet", video).Do()
				if err != nil {
					log.Fatalf("Error making API call to update video: %v", err.Error())
				}
			}
		}

		nextPageToken = searchResponse.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	fmt.Printf("Updated %d videos\n", counter)
}
