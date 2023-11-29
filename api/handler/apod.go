package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nikitanovikovdev/astrolog/api/entity"
	"github.com/nikitanovikovdev/astrolog/errorutil"
	"github.com/nikitanovikovdev/astrolog/processor"
	"log"
	"net/http"
	"time"
)

type APODHandler struct {
	processor *processor.APODProcessor
}

func NewAPODHandler(processor *processor.APODProcessor) *APODHandler {
	return &APODHandler{
		processor: processor,
	}
}

func (h APODHandler) ByDate(ctx *gin.Context) {
	log.Println("Start processing inbound request")

	date := ctx.Query("date")
	if date == "" {
		log.Println("Invalid request params")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request params"})
		return
	}

	ok, err := isValidDate(date)
	if err != nil {
		log.Printf("Validate date error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	if !ok {
		log.Println("Provided date in the future")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Provided date in the future"})
		return
	}

	content, err := h.processor.ContentByDate(date)
	if err != nil {
		if err == errorutil.ErrNoRecords {
			log.Println("Content doesn't exist")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Content by provided date doesn't exist"})
			return
		}

		log.Printf("Failed to get content by date: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	resp := entity.APODContentResponse{
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

	ctx.JSON(http.StatusOK, gin.H{"message": resp})

	log.Println("Finish processing inbound request")
}

func (h APODHandler) AllContent(ctx *gin.Context) {
	contents, err := h.processor.AllContent()
	if err != nil {
		log.Printf("Failed to get contents: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	contentsResp := make(entity.APODContentsResponse, len(contents))
	for i, v := range contents {
		contentResp := entity.APODContentResponse{
			ID:             v.ID,
			Copyright:      v.Copyright,
			Explanation:    v.Explanation,
			HdURL:          v.HdURL,
			MediaType:      v.MediaType,
			ServiceVersion: v.ServiceVersion,
			Title:          v.Date,
			URL:            v.URL,
			Date:           v.Date,
		}

		contentsResp[i] = contentResp
	}

	ctx.JSON(http.StatusOK, gin.H{"message": contentsResp})
}

func isValidDate(date string) (bool, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false, err
	}
	currentDate := time.Now()

	return parsedDate.Before(currentDate), nil
}
