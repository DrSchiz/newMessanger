package controllers

import (
	"messanger/database"
	"messanger/initializers"
	"messanger/models"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthUser(c *gin.Context) {
	ginToken := c.Request.Header.Get("Authorization")
	nameElement := strings.Split(ginToken, " ")

	decodedToken, _, err := initializers.ClientKeyCloak.DecodeAccessToken(c, nameElement[1], os.Getenv("KEYCLOAK_USER_REALM_NAME"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims := decodedToken.Claims.(jwt.MapClaims)
	id := claims["sub"]

	var user models.MessangerUser

	database.GormDB.Where("keycloak_id = ?", id).Find(&user)
	if user.Id == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user is not found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"User": user,
	})
}

func Logout(c *gin.Context) {
	refreshToken := c.Request.Header.Get("Refresh_Token")

	err := initializers.ClientKeyCloak.Logout(c, os.Getenv("KEYCLOAK_CLIENT_ID"), os.Getenv("KEYCLOAK_CLIENT_SECRET"), os.Getenv("KEYCLOAK_USER_REALM_NAME"), refreshToken)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "logout is complete",
	})
}

// func SetPassword(c *gin.Context) {
// 	var decodedData user.RequestSetPassword
// 	json.NewDecoder(c.Request.Body).Decode(&decodedData)

// 	validate := validator.New()
// 	err := validate.Struct(decodedData)
// 	if err != nil {
// 		errors := err.(validator.ValidationErrors)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": errors.Error(),
// 		})
// 		return
// 	}

// 	if decodedData.Password != decodedData.RepeatPassword {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "passwords don't match",
// 		})
// 		return
// 	}

// 	client := gocloak.NewClient(os.Getenv("KEYCLOAK_REALM_URL"))

// 	err = client.SetPassword(c, initializers.Token.AccessToken, decodedData.UserId, os.Getenv("KEYCLOAK_USER_REALM_NAME"), decodedData.Password, false)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"success": "the password has been successfully set",
// 	})
// }
