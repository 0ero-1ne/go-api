package envresolver

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv(filename string) {
	err := godotenv.Load(filename)

	if err != nil {
		panic(".env not found: " + err.Error())
	}
}

func GetKey(key string) string {
	return os.Getenv(key)
}
