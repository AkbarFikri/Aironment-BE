package service

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/pkg/model"
)

type AqiService struct{}

func NewAqi() AqiService {
	return AqiService{}
}

func (s *AqiService) FetchAqiData(req model.AqiParam) (model.ServiceResponse, error) {
	latitude, err := strconv.ParseFloat(req.Latitude, 64)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "invalid latitude format",
			Payload: latitude,
		}, err
	}

	longitude, err := strconv.ParseFloat(req.Longitude, 64)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "invalid longitude format",
			Payload: longitude,
		}, err
	}

	postBody, _ := json.Marshal(map[string]map[string]float64{
		"location": {
			"latitude":  latitude,
			"longitude": longitude,
		},
	})

	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("https://airquality.googleapis.com/v1/currentConditions:lookup?key=", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var res model.GoogleResAqi

	if err := json.Unmarshal(body, &res); err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "gagal",
			Payload: err.Error(),
		}, err
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "successfully get data",
		Payload: res,
	}, nil
}
