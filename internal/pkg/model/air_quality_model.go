package model

type AqiParam struct {
	Latitude  string `form:"latitude" binding:"required"`
	Longitude string `form:"longitude" binding:"required"`
}
