package models

type CurrentBalance struct {
	Current   float64 `json:"current"`
	Withdrawn float64 `json:"withdrawn"`
}
