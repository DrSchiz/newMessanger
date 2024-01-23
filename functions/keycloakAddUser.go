package functions

import (
	"log"
	"messanger/database"
	"messanger/initializers"
	"messanger/models"
	"net/http"
	"os"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
)

func KeycloakAddUser(c *gin.Context, register_id string, password string) (string, error) {

	var registerUserStepOne models.RegisterUserStepOne
	database.GormDB.Table("register_user_step_one").Where("registration_id = ?", register_id).Find(&registerUserStepOne)

	var registerUserStepTwo models.RegisterUserStepTwo
	database.GormDB.Table("register_user_step_two").Where("registration_id = ?", register_id).Find(&registerUserStepTwo)

	var registerUserStepThree models.RegisterUserStepThree
	database.GormDB.Table("register_user_step_three").Where("registration_id = ?", register_id).Find(&registerUserStepThree)

	var registerUserStepFour models.RegisterUserStepFour
	database.GormDB.Table("register_user_step_four").Where("registration_id = ?", register_id).Find(&registerUserStepFour)

	gocloakUser := gocloak.User{
		Email:         gocloak.StringP(registerUserStepOne.Email),
		FirstName:     gocloak.StringP(registerUserStepThree.Firstname),
		LastName:      gocloak.StringP(registerUserStepThree.Lastname),
		Username:      gocloak.StringP(registerUserStepThree.Username),
		Enabled:       gocloak.BoolP(true),
		EmailVerified: gocloak.BoolP(true),
	}

	var err error

	keycloakId, err := initializers.ClientKeyCloak.CreateUser(c, initializers.Token.AccessToken, os.Getenv("KEYCLOAK_USER_REALM_NAME"), gocloakUser)
	if err != nil {
		log.Println(err)
		return "", err
	}

	err = initializers.ClientKeyCloak.SetPassword(c, initializers.Token.AccessToken, keycloakId, os.Getenv("KEYCLOAK_USER_REALM_NAME"), password, false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return "", err
	}

	return keycloakId, nil
}
