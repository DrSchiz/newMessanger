package database

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	GormDB *gorm.DB
	Ctx    context.Context
	Rdb    *redis.Client
)

func ConnectDataBase() {
	var err error
	dsn := os.Getenv("DSN")

	GormDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Println("postgresql: ", err)
		return
	}

	log.Println("success postgresql database connect")
}
