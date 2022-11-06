package applications

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVars() {
	if err := godotenv.Load("gotwitter.env"); err != nil {
		log.Fatalln("Loading gotwitter.env file failed!")
	}
}
