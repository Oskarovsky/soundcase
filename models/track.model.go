package models

import (
	"gorm.io/gorm"
	"time"
)

type Track struct {
	gorm.Model
	Id             uint      `json:"id" gorm:"primary_key"`
	Title          string    `json:"title" gorm:"type:varchar(500);not null"`
	Artists        []*Artist `json:"artists" gorm:"many2many:artist_track"`
	Version        string    `json:"version"`
	Length         int       `json:"length"`
	Genre          string    `json:"genre"`
	CreationDate   time.Time `json:"creationDate" gorm:"column:creation_date"`     // date when track has been added to library
	ProductionDate time.Time `json:"productionDate" gorm:"column:production_date"` // year when track has been created by artist
}

func (track *Track) TableName() string {
	return "track"
}
