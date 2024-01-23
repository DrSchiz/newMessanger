package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvInitializer() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error of load the .env file!")
		return
	}
	log.Println("success .env file load")
}
