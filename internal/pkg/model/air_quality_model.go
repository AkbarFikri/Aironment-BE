package model

type AqiParam struct {
	Latitude  string `form:"latitude" binding:"required"`
	Longitude string `form:"longitude" binding:"required"`
}

type AqiResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
