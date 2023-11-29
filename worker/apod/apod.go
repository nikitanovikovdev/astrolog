package apod

import (
	"github.com/nikitanovikovdev/astrolog/client"
	"github.com/nikitanovikovdev/astrolog/processor"
	"github.com/nikitanovikovdev/astrolog/processor/entity"
	"log"
)

type APODJob struct {
	apodClient    *client.APODClient
	apodProcessor *processor.APODProcessor
}

func NewAPODJob(apodClient *client.APODClient, apodProcessor *processor.APODProcessor) *APODJob {
	return &APODJob{
		apodClient:    apodClient,
		apodProcessor: apodProcessor,
	}
}

func (j APODJob) Do() error {
	log.Println("Starting fetching new content")

	content, err := j.apodClient.FetchContent()
	if err != nil {
		log.Printf("Fetching content error: %v", err)
		return err
	}

	apod := entity.APODContent{
		ID:             content.ID,
		Copyright:      content.Copyright,
		Explanation:    content.Explanation,
		HdURL:          content.HdURL,
		MediaType:      content.MediaType,
		ServiceVersion: content.ServiceVersion,
		Title:          content.Title,
		URL:            content.URL,
		Date:           content.Date,
	}

	err = j.apodProcessor.InsertData(apod)
	if err != nil {
		log.Printf("Insert data error: %v", err)
		return err
	}

	log.Println("Successfully added new content")
	return nil
}

func (j APODJob) Stop() {
	log.Println("Stopping a job")
}
