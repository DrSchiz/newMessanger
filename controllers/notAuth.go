package controllers

import (
	"encoding/json"
	"messanger/http/user"
	"messanger/initializers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var requestLoginUser user.RequestLoginUser
	json.NewDecoder(c.Request.Body).Decode(&requestLoginUser)

	getToken, err := initializers.ClientKeyCloak.Login(c, os.Getenv("KEYCLOAK_CLIENT_ID"), os.Getenv("KEYCLOAK_CLIENT_SECRET"), os.Getenv("KEYCLOAK_USER_REALM_NAME"), requestLoginUser.Username, requestLoginUser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": getToken,
	})
}

func RefreshToken(c *gin.Context) {
	refreshToken := c.Request.Header.Get("Refresh_Token")

	accessToken, err := initializers.ClientKeyCloak.RefreshToken(c, refreshToken, os.Getenv("KEYCLOAK_CLIENT_ID"), os.Getenv("KEYCLOAK_CLIENT_SECRET"), os.Getenv("KEYCLOAK_USER_REALM_NAME"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": accessToken,
	})
}
