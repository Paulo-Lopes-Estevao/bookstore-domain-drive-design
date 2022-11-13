package repository_test

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func Env() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}
