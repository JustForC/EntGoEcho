package configs

import (
	"CompanyAPI/ent"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Open() *ent.Client {
	env, errEnv := godotenv.Read()
	if errEnv != nil {
		log.Fatal("Error reading .env!")
	}

	dsn := env["DB_USERNAME"] + ":" + env["DB_PASSWORD"] + "@tcp(" + env["DB_HOST"] + ":" + env["DB_PORT"] + ")/" + env["DB_NAME"] + "?parseTime=True"
	client, err := ent.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	return client

}
