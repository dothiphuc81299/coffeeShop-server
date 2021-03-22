package model

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	prefixSm = "sm_"
	prefixMd = "md_"
)

// FilePhoto ...
type FilePhoto struct {
	ID         AppID           `bson:"_id" json:"_id"`
	Name       string          `bson:"name" json:"name,omitempty"`
	Dimensions *FileDimensions `bson:"dimensions" json:"dimensions"`
}

// FileSize ...
type FileSize struct {
	Width  int    `json:"width" bson:"width"`
	Height int    `json:"height" bson:"height"`
	URL    string `json:"url" bson:"-"`
}

// ConvertToFilePhoto ...
func (f *FilePhotoRequest) ConvertToFilePhoto() *FilePhoto {
	if f == nil {
		return nil
	}
	return &FilePhoto{
		ID:         util.GetAppIDFromHex(f.ID),
		Name:       f.Name,
		Dimensions: f.Dimensions,
	}
}

// FileDimensions ...
type FileDimensions struct {
	Small  *FileSize `json:"sm" bson:"sm"`
	Medium *FileSize `json:"md" bson:"md"`
}

// FileDefaultPhoto ...
func FileDefaultPhoto() *FilePhoto {
	return generateFileData("default_photo.jpg", 375, 130, 750, 250)
}

// generateFileData ...
func generateFileData(name string, smW int, smH int, mdW int, mdH int) *FilePhoto {
	var response = &FilePhoto{
		ID:   primitive.NewObjectID(),
		Name: name,
	}

	var smallWidthHeight = FileSize{
		Width:  smW,
		Height: smH,
	}

	var mediumWidthHeight = FileSize{
		Width:  mdW,
		Height: mdH,
	}

	var dimensions = &FileDimensions{
		Small:  &smallWidthHeight,
		Medium: &mediumWidthHeight,
	}

	response.Dimensions = dimensions
	return response.GetResponseData()
}

// Validate ...
func (f *FilePhotoRequest) Validate() error {
	if f == nil {
		return nil
	}
	return validation.ValidateStruct(f,
		validation.Field(&f.ID, is.MongoID.Error(locale.CommonKeyInvalidPhoto)),
		validation.Field(&f.Name, validation.Required.Error(locale.CommonKeyInvalidPhoto)),
	)
}

// GetResponseData ...
func (fp *FilePhoto) GetResponseData() *FilePhoto {
	if fp == nil || fp.ID.IsZero() {
		return FileDefaultPhoto()
	}

	fp.Dimensions.GenerateURL(config.GetEnv().FileHost, fp.Name)
	return fp
}

// FilePhotoRequest ...
type FilePhotoRequest struct {
	ID         string          `json:"_id"`
	Name       string          `json:"name"`
	Dimensions *FileDimensions `json:"dimensions"`
}

// GenerateURL ...
func (fd *FileDimensions) GenerateURL(fileHost, filename string) {
	if fd == nil || fd.Small == nil || fd.Medium == nil {
		return
	}
	fd.Small.URL = fileHost + prefixSm + filename
	fd.Medium.URL = fileHost + prefixMd + filename
}
