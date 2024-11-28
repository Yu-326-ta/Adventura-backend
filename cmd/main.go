package main

import (
	"fmt"
	"log"

	"Adventura/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting server...")
	loadEnv()
	server.StartServer()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("環境変数が読み込み出来ませんでした: %v", err)
	}
}
