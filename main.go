package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/nikitanovikovdev/astrolog/api/server"
	"github.com/nikitanovikovdev/astrolog/api/shutdown"
	"github.com/nikitanovikovdev/astrolog/client"
	"github.com/nikitanovikovdev/astrolog/database"
	"github.com/nikitanovikovdev/astrolog/processor"
	"github.com/nikitanovikovdev/astrolog/repository"
	"github.com/nikitanovikovdev/astrolog/worker"
	"github.com/nikitanovikovdev/astrolog/worker/apod"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load env variables")
	}

	db, err := database.New()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	apodRepo := repository.NewAPODRepo(db)
	apodProcessor := processor.NewAPODProcessor(apodRepo)

	srv := server.New(apodProcessor)
	srv.InitRoutes()

	go func() {
		if err := srv.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to run a server")
		}
	}()

	apodClient := client.NewAPODClient(os.Getenv("APOD_URL"), os.Getenv("APOD_API_KEY"))
	apodJob := apod.NewAPODJob(apodClient, apodProcessor)

	tickerTime, err := time.ParseDuration(os.Getenv("APOD_THRESHOLD_TIME"))
	if err != nil {
		log.Fatal("Failed to parse apod threshold time")
	}
	apodWorker := worker.NewWorker(apodJob, tickerTime)
	apodWorker.Do()

	err = shutdown.Wait(func() error {
		log.Println("Graceful shutdown start")
		apodWorker.Stop()

		ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return err
		}

		log.Println("Graceful shutdown finish")
		return nil
	})

	if err != nil {
		log.Printf("Graceful shutdown finished with error: %v", err)
	}
}
