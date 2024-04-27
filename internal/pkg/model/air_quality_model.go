package model

type AqiParam struct {
	Latitude  string `form:"latitude" binding:"required"`
	Longitude string `form:"longitude" binding:"required"`
}

type AqiResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GoogleResAqi struct {
	DateTime   string `json:"dateTime"`
	RegionCode string `json:"regionCode"`
	Indexes    any    `json:"indexes"`
}
