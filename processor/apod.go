package processor

import (
	"github.com/nikitanovikovdev/astrolog/processor/entity"
	dbEntity "github.com/nikitanovikovdev/astrolog/repository/entity"
)

type APODRepoI interface {
	ContentByDate(date string) (dbEntity.APODContent, error)
	AllContent() ([]dbEntity.APODContent, error)
	InsertContent(data dbEntity.APODContent) error
}

type APODProcessor struct {
	apodRepo APODRepoI
}

func NewAPODProcessor(repo APODRepoI) *APODProcessor {
	return &APODProcessor{
		apodRepo: repo,
	}
}

func (p APODProcessor) ContentByDate(date string) (dbEntity.APODContent, error) {
	return p.apodRepo.ContentByDate(date)
}

func (p APODProcessor) AllContent() ([]dbEntity.APODContent, error) {
	return p.apodRepo.AllContent()
}

func (p APODProcessor) InsertData(content entity.APODContent) error {
	data := dbEntity.APODContent{
		Copyright:      content.Copyright,
		Explanation:    content.Explanation,
		HdURL:          content.HdURL,
		MediaType:      content.MediaType,
		ServiceVersion: content.ServiceVersion,
		Title:          content.Title,
		URL:            content.URL,
		Date:           content.Date,
	}

	return p.apodRepo.InsertContent(data)
}
