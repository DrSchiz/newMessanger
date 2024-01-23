package functions

import (
	"log"
	"messanger/database"
	"messanger/models"
	"time"
)

func DatabaseAddUser(register_id string, keycloak_id string, stepThree models.RegisterUserStepThree, stepFour models.RegisterUserStepFour) error {
	var registerUserStepOne models.RegisterUserStepOne
	database.GormDB.Table("register_user_step_one").Where("registration_id = ?", register_id).Find(&registerUserStepOne)

	var registerUserStepTwo models.RegisterUserStepTwo
	database.GormDB.Table("register_user_step_one").Where("registration_id = ?", register_id).Find(&registerUserStepTwo)

	var registerUserStepThree models.RegisterUserStepThree
	database.GormDB.Table("register_user_step_one").Where("registration_id = ?", register_id).Find(&registerUserStepThree)

	var registerUserStepFour models.RegisterUserStepFour
	database.GormDB.Table("register_user_step_one").Where("registration_id = ?", register_id).Find(&registerUserStepFour)

	var messangerUser models.MessangerUser
	messangerUser = models.MessangerUser{
		KeycloakId: keycloak_id,
		Email:      registerUserStepOne.Email,
		Firstname:  stepThree.Firstname,
		Lastname:   stepThree.Lastname,
		Username:   stepThree.Username,
		IsBlocked:  false,
		CreatedAt:  time.Now(),
	}

	err := database.GormDB.Create(&messangerUser)
	if err.Error != nil {
		log.Println(err.Error)
		return err.Error
	}

	return nil
}
