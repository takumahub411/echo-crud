package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvfile() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading file")
	}
}
