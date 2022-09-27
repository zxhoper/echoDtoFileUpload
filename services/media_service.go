package services

import (
	"echoDtoFileUpload/helper"
	"echoDtoFileUpload/models"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUpload(file models.File) (string, error)
	RemoteUpload(url models.Url) (string, error)
}

type media struct{}

func NewMediaUpload() mediaUpload {
	return &media{}
}

func (*media) FileUpload(file models.File) (string, error) {
	// validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	// upload
	uploadUrl, err := helper.ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (*media) RemoteUpload(url models.Url) (string, error) {
	// validate
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	// upload
	uploadUrl, err := helper.ImageUploadHelper(url.Url)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}
