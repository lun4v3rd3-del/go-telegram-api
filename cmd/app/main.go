package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"telegram-api-service/internal/infrastructure/telegram"
	"telegram-api-service/internal/usecase"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	t := token()

	fmt.Printf("Loaded token %s\n", t)

	tgClient := telegram.NewHttpClient(t)

	router := telegram.NewRouter()
	re, _ := regexp.Compile(".*")
	router.RegisterHandler(re, telegram.NewGreetHandler(tgClient))

	eventProcessor := usecase.NewEventProcessor(tgClient)

	fetcher := usecase.Fetcher(eventProcessor)
	processor := usecase.Processor(eventProcessor)

	events := fetcher.Fetch()
	fmt.Println(events)
	for _, event := range events {
		processor.Process(event, router.HandlerPool)
	}

	// consumer.Start(fetcher, processor)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func token() string {
	token := os.Getenv("telegram_api_key")
	return token
}
