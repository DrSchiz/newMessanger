package models

import "time"

type VerificationEmail struct {
	Id               uint   `gorm:"primaryKey;not null"`
	Email            string `gorm:"unique;not null"`
	VerificationCode string `gorm:"not null"`
	CreatedAt        time.Time
}

func (VerificationEmail) TableName() string {
	return "verification_email"
}
