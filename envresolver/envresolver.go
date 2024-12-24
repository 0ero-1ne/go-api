package envresolver

import (
	"os"

	"github.com/joho/godotenv"
)

func GetKey(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		return "storage.db"
	}

	value := os.Getenv(key)
	os.Clearenv()

	return value
}
