package initializers

import (
	"context"
	"log"
	"os"

	"github.com/Nerzal/gocloak/v13"
)

var (
	ClientKeyCloak *gocloak.GoCloak
	Ctx            context.Context
	Token          *gocloak.JWT
)

func KeycloakInitializer() {
	ClientKeyCloak = gocloak.NewClient(os.Getenv("KEYCLOAK_REALM_URL"))
	Ctx = context.Background()

	var err error

	Token, err = ClientKeyCloak.LoginAdmin(Ctx, os.Getenv("KEYCLOAK_ADMIN_USERNAME"), os.Getenv("KEYCLOAK_ADMIN_PASSWORD"), os.Getenv("KEYCLOAK_ADMIN_REALM_NAME"))
	if err != nil {
		log.Println("keycloak admin: ", err)
		return
	}
	log.Println("success login admin")
}
