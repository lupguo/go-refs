package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error load dot env")
	}
	log.Println(os.Getwd())

	s3Bucket := os.Getenv("S3_BUCKET")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")

	log.Printf("s3Bucket=>%s, s3SecerctKey=>%s\n", s3Bucket, s3SecretKey)

	// all env
	for k, v := range os.Environ() {
		log.Printf("%d=>%s\n", k, v)
	}
}
