package client

import (
	"encoding/json"
	"github.com/nikitanovikovdev/astrolog/client/entity"
	"io/ioutil"
	"log"
	"net/http"
)

type APODClient struct {
	url    string
	apiKey string
}

func NewAPODClient(url, apiKey string) *APODClient {
	return &APODClient{
		url:    url,
		apiKey: apiKey,
	}
}

func (c APODClient) FetchContent() (entity.APODClientResponse, error) {
	req, err := http.NewRequest(http.MethodGet, c.url, nil)
	if err != nil {
		log.Printf("Failed to create new request: %v", err)
		return entity.APODClientResponse{}, err
	}

	req.Header.Add("x-api-key", c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed send request: %v", err)
		return entity.APODClientResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed send request: %v", err)
		return entity.APODClientResponse{}, err
	}

	var apodResponse entity.APODClientResponse
	if err = json.Unmarshal(body, &apodResponse); err != nil {
		log.Printf("Failed unmarshal data: %v", err)
		return entity.APODClientResponse{}, err
	}

	return apodResponse, nil
}
