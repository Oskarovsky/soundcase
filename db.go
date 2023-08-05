package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

var currentTrackId int
var currentUserId int

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	HandleError(err)
	return c
}

// Give us some data
// function named init() will be run automatically upon running
func init() {
	CreateTrack(Track{})
	CreateTrack(Track{})
}

func FindAllTracks() Tracks {
	var tracks Tracks

	c := RedisConnect()
	defer c.Close()

	keys, err := c.Do("KEYS", "track:*")
	HandleError(err)

	for _, k := range keys.([]interface{}) {
		var track Track

		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)

		if err := json.Unmarshal(reply.([]byte), &track); err != nil {
			panic(err)
		}

		tracks = append(tracks, track)
	}
	return tracks
}

func FindTrack(id int) Track {
	var track Track

	c := RedisConnect()
	defer c.Close()

	reply, err := c.Do("GET", "track:"+strconv.Itoa(id))
	HandleError(err)

	fmt.Println("GET OK")

	if err := json.Unmarshal(reply.([]byte), &track); err != nil {
		panic(err)
	}
	return track
}

func CreateTrack(track Track) {
	currentTrackId += 1
	currentUserId += 1

	track.Id = currentTrackId
	track.CreationDate = time.Now()

	c := RedisConnect()
	defer c.Close()

	b, err := json.Marshal(track)
	HandleError(err)

	// Save json blob to redis
	reply, err := c.Do("SET", "track:"+strconv.Itoa(track.Id), b)
	HandleError(err)

	fmt.Println("GET", reply)
}
