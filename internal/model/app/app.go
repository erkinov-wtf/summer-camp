package app

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type App struct {
	ID             uuid.UUID `gorm:"types:uuid;primaryKey;default:gen_random_uuid()"`
	Name           string    `gorm:"column:name;type:text"`
	FirstNumber    float64   `gorm:"column:first_number;types:double precision"`
	SecondNumber   float64   `gorm:"column:second_number;types:double precision"`
	ExpectedNumber *float64  `gorm:"column:expected_number;types:uuid"`
	Text           string    `gorm:"column:text;types:text"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// BeforeCreate callbacks
func (finance *App) BeforeCreate(tx *gorm.DB) (err error) {
	finance.ID = uuid.New()
	return err
}
