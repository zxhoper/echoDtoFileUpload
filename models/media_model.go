package models

import "mime/multipart"

// File struct define the required property for local file upload.
type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

// Url struct define the required property for remote URL upload.
type Url struct {
	Url string `json:"url,omitempty" validate:"required"`
}
