package config

import (
	//"fmt"
	"log"
	"os"

	_ "github.com/cloudinary/cloudinary-go"
	_ "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func EnvCloudName() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func EnvCloudAPIKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_API_KEY")
}

func EnvCloudAPISecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_API_SECERT")
}

func EnvCloudUploadFolder() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}