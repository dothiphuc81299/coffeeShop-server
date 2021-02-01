package locale

import (
	"log"

	"github.com/magiconair/properties"
)

// Keys
const (
	CommonKeySuccess      = "success"
	CommonKeyBadRequest   = "badRequest"
	CommonKeyUnauthorized = "unauthorized"
	CommonKeyNotFound     = "notFound"
)

type (
	commonLang struct {
		Success      string `properties:"success"`
		BadRequest   string `properties:"badRequest"`
		Unauthorized string `properties:"unauthorized"`
		NotFound     string `properties:"notFound"`
	}
)

var (
	commonVi commonLang
)

func init() {
	// Load properties
	p2 := properties.MustLoadFile(getLocalePath()+"/properties/common.properties", properties.UTF8)
	if err := p2.Decode(&commonVi); err != nil {
		log.Fatal(err)
	}
}

func commonLoadLocales() (response []Locale) {
	// 1-99
	response = []Locale{
		{
			Key: CommonKeySuccess,
			Message: &Message{
				Vi: commonVi.Success,
			},
			Code: 1,
		},
		{
			Key: CommonKeyBadRequest,
			Message: &Message{
				Vi: commonVi.BadRequest,
			},
			Code: 2,
		},
		{
			Key: CommonKeyUnauthorized,
			Message: &Message{
				Vi: commonVi.Unauthorized,
			},
			Code: 3,
		},
		{
			Key: CommonKeyNotFound,
			Message: &Message{
				Vi: commonVi.NotFound,
			},
			Code: 4,
		},
	}
	return response
}
