package models

type Artist struct {
	Id          uint     `json:"id" gorm:"primary_key"`
	Name        string   `json:"name" gorm:"type:varchar(255);not null"`
	Tracks      []*Track `json:"tracks" gorm:"many2many:artist_track"`
	Description string   `json:"description"`
}

func (artist *Artist) TableName() string {
	return "artist"
}
