package models

import "time"

type Conversion struct {
	ID              uint      `gorm:"primaryKey"`
	BaseCurrency    string    `gorm:"size:3;not null"`
	TargetCurrency  string    `gorm:"size:3;not null"`
	Amount          float64   `gorm:"type:numeric(18,2);not null"`
	ConvertedAmount float64   `gorm:"type:numeric(18,2);not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
