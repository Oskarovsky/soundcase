package worker

import (
	"fmt"
	"github.com/dhowden/tag"
	"io/ioutil"
	"log"
	"os"
	"soundcase/models"
	"strings"
)

func ReadFilesFromDirectory(dir string) []string {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var trackList []string
	var tracks []models.Track
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() && strings.Contains(file.Name(), ".mp3") {
			trackList = append(trackList, file.Name())
			track := ReadTrackMetadataFromDirectory(dir + "/" + file.Name())
			tracks = append(tracks, track)
		}
	}

	return trackList
}

func ReadTrackMetadataFromDirectory(fileName string) models.Track {
	fmt.Printf("Trying to fetch metadata for track from %v \n", fileName)
	f, err := os.Open(fileName)

	m, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}

	track := models.Track{
		Title: m.Title(),
		Genre: m.Genre(),
		Artists: []*models.Artist{
			{
				Name: m.Artist(),
			},
		},
	}

	log.Printf("Track: %v \n", track)
	log.Println("=================")
	return track
}
