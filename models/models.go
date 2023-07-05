package models

import (
	"time"
)

type Armament struct {
	ID uint `gorm:"primary_key"`
	SpaceshipID uint `gorm:"spaceship_id"`
	Title     string `gorm:"title" json:"title"`
	Qty       string `gorm:"qty" json:"qty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Spaceship struct {
	ID        uint `gorm:"primary_key"`
	Name      string     `gorm:"name" json:"name"`
	Class     string     `gorm:"class" json:"class"`
	Armaments []Armament `gorm:"foreignKey:spaceship_id" json:"armaments"`
	Crew      int        `gorm:"crew" json:"crew"`
	Image     string     `gorm:"image" json:"image"`
	Value     float32    `gorm:"value" json:"value"`
	Status    string     `gorm:"status" json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Validation struct {
	Value string
	Valid string
}

func (s *Spaceship) IsValidSpaceship() bool {
	if s.Name == "" || s.Class == ""  {
		return false
	}
	return true
}
