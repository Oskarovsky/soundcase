package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Users []User
type Artists []Artist
type Tracks []Track

type TrackService interface {
	GetTrack(id string) (*Track, error)
	GetTracks() ([]*Track, error)
	CreateTrack(track *Track) (*Track, error)
	UpdateTrack(track *Track) (*Track, error)
	DeleteTrack(id string) error
}
