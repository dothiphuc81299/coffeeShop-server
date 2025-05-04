 package locale

// import (
// 	"encoding/json"
// 	"path"
// 	"runtime"
// 	"strings"

// 	"github.com/thoas/go-funk"
// )

// // Lang ...
// const (
// 	LangVi = "vi"
// )

// type (
// 	// Locale ...
// 	Locale struct {
// 		Key     string
// 		Code    int      `json:"code"`
// 		Message *Message `json:"message"`
// 	}

// 	// Message ...
// 	Message struct {
// 		Vi      string
// 		Display string
// 	}
// )

// func getLocalePath() string {
// 	_, filename, _, ok := runtime.Caller(0)
// 	if !ok {
// 		panic("No caller information")
// 	}
// 	return path.Dir(filename)
// }

// // GetDisplay return text with language
// func (msg *Message) GetDisplay(lang string) {
// 	text := funk.Get(msg, strings.Title(lang))
// 	if text != nil {
// 		msg.Display = text.(string)
// 	} else {
// 		msg.Display = "N/A"
// 	}
// }

// // MarshalJSON ...
// func (msg *Message) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(msg.Display)
// }

// var notFoundKey = Locale{
// 	Key: "NotFound",
// 	Message: &Message{
// 		Vi: "Không tìm thấy",
// 	},
// 	Code: -1,
// }

// // Default key for each error
// const (
// 	Default200 = CommonKeySuccess
// 	Default400 = CommonKeyBadRequest
// 	Default401 = CommonKeyUnauthorized
// 	Default404 = CommonKeyNotFound
// )

// // GetByKey give key and receive message + code
// func GetByKey(lang string, key string) Locale {
// 	item := funk.Find(list, func(item Locale) bool {
// 		return item.Key == key
// 	})

// 	if item == nil {
// 		return notFoundKey
// 	}

// 	return item.(Locale)
// }

// var list = make([]Locale, 0)

// // LoadProperties ...
// func LoadProperties() {
// 	// Assign locales
// 	// 1-99
// 	list = append(list, commonLoadLocales()...)

// 	//100-199
// 	list = append(list, drinkLoadLocales()...)

// 	// 200-299
// 	list = append(list, categoryLoadLocales()...)
// }
