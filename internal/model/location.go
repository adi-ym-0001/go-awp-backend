package model

type Location struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Floor    string
	Area     string
	Drawings []Drawing `gorm:"many2many:location_drawings;"`
}
