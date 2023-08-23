package worker

import (
	"fmt"
	"github.com/dhowden/tag"
	"github.com/tcolgate/mp3"
	"io"
	"io/ioutil"
	"log"
	"os"
	"soundcase/models"
	"strings"
	"time"
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
	file01, err := os.Open(fileName)
	file02, err := os.Open(fileName)

	length := ReadTrackLength(file02)

	m, _ := tag.ReadFrom(file01)
	if err != nil {
		log.Fatal(err)
	}

	if m != nil {
		track := models.Track{
			Title:          m.Title(),
			Genre:          m.Genre(),
			ProductionDate: YearTime(m.Year()),
			Length:         length,
			Artists: []*models.Artist{
				{
					Name: m.Artist(),
				},
			},
		}

		log.Printf("Track: %v \n", track)
		log.Println("================= SUCCESS")
		return track
	}

	log.Println("================= ERROR")
	return models.Track{}
}

func YearTime(y int) time.Time {
	// convert int to Time - use the last day of the year, which is 31st December
	t := time.Date(y, time.December, 31, 0, 0, 0, 0, time.Local)
	return t
}

func ReadTrackLength(file io.Reader) int {
	t := 0.0
	d := mp3.NewDecoder(file)
	var frame mp3.Frame
	skipped := 0

	for {
		if err := d.Decode(&frame, &skipped); err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}

		t = t + frame.Duration().Seconds()
	}
	return int(t)

}
