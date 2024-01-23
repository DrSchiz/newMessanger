package models

import (
	"time"

	"gorm.io/gorm"
)

type MessangerUser struct {
	Id         uint   `gorm:"primaryKey;not null"`
	KeycloakId string `gorm:"unique;not null"`
	Email      string `gorm:"not null"`
	Firstname  string
	Lastname   string
	Username   string `gorm:"not null"`
	IsBlocked  bool   `gorm:"not null"`
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (MessangerUser) TableName() string {
	return "messanger_user"
}

type RegisterUserStepOne struct {
	Id             uint   `gorm:"primaryKey;not null"`
	RegistrationId string `gorm:"unique;not null"`
	Email          string `gorm:"not null"`
}

func (RegisterUserStepOne) TableName() string {
	return "register_user_step_one"
}

type RegisterUserStepTwo struct {
	Id               uint   `gorm:"primaryKey;not null"`
	RegistrationId   string `gorm:"unique;not null"`
	VerificationCode string `gorm:"not null"`
}

func (RegisterUserStepTwo) TableName() string {
	return "register_user_step_two"
}

type RegisterUserStepThree struct {
	Id             uint   `gorm:"primaryKey"`
	RegistrationId string `gorm:"unique"`
	Username       string `gorm:"not null"`
	Firstname      string `gorm:"default:null"`
	Lastname       string `gorm:"default:null"`
}

func (RegisterUserStepThree) TableName() string {
	return "register_user_step_three"
}

type RegisterUserStepFour struct {
	Id             uint   `gorm:"primaryKey"`
	RegistrationId string `gorm:"unique;not null"`
	PasswordStatus bool   `gorm:"not null"`
}

func (RegisterUserStepFour) TableName() string {
	return "register_user_step_four"
}
