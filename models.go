package main

import "time"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Track struct {
	Id             int       `json:"id"`
	Title          string    `json:"title"`
	Artist         Artist    `json:"artist"`
	Genre          string    `json:"genre"`
	CreationDate   time.Time `json:"creationDate"`   // date when track has been added to library
	ProductionDate time.Time `json:"productionDate"` // year when track has been created by artist

}

type Artist struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Users []User
type Artists []Artist
type Tracks []Track
