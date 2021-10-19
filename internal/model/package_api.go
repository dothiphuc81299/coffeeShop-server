package model

type Level string
type PackageBody struct {
	Title      string  `json:"title"`
	Level      string  `json:"level"` // easy , medium, hard
	Reward     float64 `json:"reward"`
	MinusPoint float64 `json:"minusPoint"`
}

type PackageAdminResponse struct {
}
