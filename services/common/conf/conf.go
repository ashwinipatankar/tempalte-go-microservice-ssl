package conf

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

//CONF For reading .env file
func CONF(key string) string {

	p := "/Users/ashwinipatankar/sample.conf"
	err := godotenv.Load(p)
	if err != nil {
		log.Fatalf("Error in loading config file:\n%v", err)
	}
	return os.Getenv(key)
}
