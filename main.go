package main

import (
	"bookstore/app/api/server"
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func main() {
	serve := server.NewServer()
	serve.Start()
}
