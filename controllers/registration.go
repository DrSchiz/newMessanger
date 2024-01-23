package controllers

import (
	"encoding/json"
	"log"
	"messanger/database"
	"messanger/functions"
	"messanger/http/mail"
	"messanger/http/user"
	"messanger/math"
	"messanger/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func FirstStepRegistration(c *gin.Context) {
	var requestRegisterUserStepOne user.RequestRegisterUserStepOne

	json.NewDecoder(c.Request.Body).Decode(&requestRegisterUserStepOne)

	validate := validator.New()
	err := validate.Struct(requestRegisterUserStepOne)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.Error(),
		})
		return
	}

	newUserId := math.GenerateId()

	registerUserStepOne := models.RegisterUserStepOne{
		RegistrationId: newUserId,
		Email:          requestRegisterUserStepOne.Email,
	}

	err2 := database.GormDB.Create(&registerUserStepOne)
	if err2.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err2.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": registerUserStepOne.RegistrationId,
	})
}

func SecondStepRegistration(c *gin.Context) {
	var requestRegisterUserStepTwo user.RequestRegisterUserStepTwo
	json.NewDecoder(c.Request.Body).Decode(&requestRegisterUserStepTwo)

	var registerUserStepOne models.RegisterUserStepOne
	database.GormDB.Table("register_user_step_one").Where("registration_id = ?", requestRegisterUserStepTwo.RegisterId).Find(&registerUserStepOne)

	if registerUserStepOne.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "step 1 is not completed",
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(requestRegisterUserStepTwo)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.Error(),
		})
		return
	}

	var verificationEmail models.VerificationEmail

	database.GormDB.Table("verification_email").Where("email = ?", requestRegisterUserStepTwo.Email).Find(&verificationEmail)

	timeNow := time.Now()
	timeCreateVerificationCode := verificationEmail.CreatedAt
	timeDifference := timeNow.Sub(timeCreateVerificationCode)

	log.Println(timeDifference.Minutes())
	if timeDifference.Minutes() > 5 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "code lifespan is over",
		})
		return
	}

	if !functions.CheckValueHash(requestRegisterUserStepTwo.VerificationCode, verificationEmail.VerificationCode) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "codes don't match",
		})
		return
	}

	var registerUserStepTwo models.RegisterUserStepTwo
	registerUserStepTwo.RegistrationId = requestRegisterUserStepTwo.RegisterId
	registerUserStepTwo.VerificationCode = requestRegisterUserStepTwo.VerificationCode

	err2 := database.GormDB.Create(&registerUserStepTwo)
	if err2.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err2.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": registerUserStepTwo.RegistrationId,
	})
}

func ThirdStepRegistration(c *gin.Context) {
	var requestRegisterUserStepThree user.RequestRegisterUserStepThree
	json.NewDecoder(c.Request.Body).Decode(&requestRegisterUserStepThree)

	var registerUserStepTwo models.RegisterUserStepTwo
	database.GormDB.Table("register_user_step_two").Where("registration_id = ?", requestRegisterUserStepThree.RegisterId).Find(&registerUserStepTwo)

	if registerUserStepTwo.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "step 2 is not completed",
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(requestRegisterUserStepThree)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.Error(),
		})
		return
	}

	var registerUserStepThree models.RegisterUserStepThree
	registerUserStepThree.RegistrationId = requestRegisterUserStepThree.RegisterId
	registerUserStepThree.Username = requestRegisterUserStepThree.Username
	registerUserStepThree.Firstname = requestRegisterUserStepThree.Firstname
	registerUserStepThree.Lastname = requestRegisterUserStepThree.Lastname

	err2 := database.GormDB.Create(&registerUserStepThree)
	if err2.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err2.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": registerUserStepThree.RegistrationId,
	})
}

func FourthStepRegistration(c *gin.Context) {
	var requestRegisterUserStepFour user.RequestRegisterUserStepFour
	json.NewDecoder(c.Request.Body).Decode(&requestRegisterUserStepFour)

	var registerUserStepThree models.RegisterUserStepThree
	database.GormDB.Table("register_user_step_three").Where("registration_id = ?", requestRegisterUserStepFour.RegisterId).Find(&registerUserStepThree)

	if registerUserStepThree.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "step 3 is not completed",
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(requestRegisterUserStepFour)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.Error(),
		})
		return
	}

	var registerUserStepFour models.RegisterUserStepFour
	registerUserStepFour.RegistrationId = requestRegisterUserStepFour.RegisterId
	registerUserStepFour.PasswordStatus = true

	err2 := database.GormDB.Create(&registerUserStepFour)
	if err2.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err2.Error,
		})
		return
	}

	keycloakId, err := functions.KeycloakAddUser(c, registerUserStepFour.RegistrationId, requestRegisterUserStepFour.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}

	err4 := functions.DatabaseAddUser(requestRegisterUserStepFour.RegisterId, keycloakId, registerUserStepThree, registerUserStepFour)
	if err4 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err4,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": registerUserStepFour.RegistrationId,
	})
}

func SendVerificationEmail(c *gin.Context) {
	var email mail.RequestSendMail
	json.NewDecoder(c.Request.Body).Decode(&email)

	validate := validator.New()
	err := validate.Struct(email)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.Error(),
		})
		return
	}

	var verificationEmail models.VerificationEmail
	database.GormDB.Table("verification_email").Where("email = ?", email.Email).Find(&verificationEmail)

	notHashedCode := math.GenerateCode()
	hashedCode := notHashedCode

	var err2 error
	verificationEmail.Email = email.Email
	verificationEmail.VerificationCode, err2 = functions.HashValue(hashedCode)
	if err2 != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err2.Error,
		})
		return
	}

	if verificationEmail.Id == 0 {
		database.GormDB.Create(&verificationEmail)
	} else {
		database.GormDB.Table("verification_email").Where("id = ?", verificationEmail.Id).Update("verification_code", verificationEmail.VerificationCode)
		database.GormDB.Table("verification_email").Where("id = ?", verificationEmail.Id).Update("created_at", time.Now())
	}

	functions.SendEmail(notHashedCode, email.Email)

	c.JSON(http.StatusOK, gin.H{
		"success": "email is sent",
	})
}
