package main

import (
	"log"
	"messanger/controllers"
	api "messanger/controllers/api"
	"messanger/database"
	"messanger/initializers"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/tbaehler/gin-keycloak/pkg/ginkeycloak"
)

func init() {
	initializers.EnvInitializer()
	database.ConnectDataBase()
	initializers.KeycloakInitializer()
	initializers.MigrateInitializer()
}

func main() {

	var sbbEndpoint = ginkeycloak.KeycloakConfig{
		Url:           os.Getenv("KEYCLOAK_REALM_URL"),
		Realm:         os.Getenv("KEYCLOAK_USER_REALM_NAME"),
		FullCertsPath: nil,
	}

	r := gin.Default()

	r.Use(ginkeycloak.RequestLogger([]string{"uid"}, "data"))
	r.Use(gin.Recovery())

	r.POST("/authorization", controllers.LoginUser)
	r.GET("/reauth", controllers.RefreshToken)
	r.POST("/registration/step_1", controllers.FirstStepRegistration)
	r.POST("/registration/step_2", controllers.SecondStepRegistration)
	r.POST("/registration/step_3", controllers.ThirdStepRegistration)
	r.POST("/registration/step_4", controllers.FourthStepRegistration)
	r.POST("/send_email_code", controllers.SendVerificationEmail)

	privateGroup := r.Group("/api/v1")
	privateGroup.Use(ginkeycloak.Auth(ginkeycloak.AuthCheck(), sbbEndpoint))

	privateGroup.GET("/user/get_data", api.AuthUser)
	// privateGroup.POST("/user/set_password", api.SetPassword)
	privateGroup.GET("/user/logout", api.Logout)
	privateGroup.POST("/user/send_message", api.SendMessage)

	err := r.Run(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Error of start the program: ", err.Error())
	}
}
