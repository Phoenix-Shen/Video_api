package model

import (
	"mime/multipart"

	"github.com/jinzhu/gorm"
)

type VideoEffect struct {
	gorm.Model
	File             *multipart.FileHeader
	Effect           string
	LocalPath        string
	PathAfterProcess string
	CoverPath        string
}
