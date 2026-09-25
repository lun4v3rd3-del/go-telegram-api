package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"telegram-api-service/internal/entitiy/arguments"
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
	re, _ := regexp.Compile("hello")

	router.RegisterHandler(arguments.HandlerArgs{
		Re: re,
		H:  telegram.NewTestHandler(tgClient),
	})
	router.RegisterHandler(arguments.HandlerArgs{
		H:  telegram.NewCallbackHandler(tgClient),
		Cd: "callback_data_1",
	})

	statesGroup := telegram.TestGet()
	re_1, _ := regexp.Compile("test")

	router.RegisterHandler(arguments.HandlerArgs{
		Re: re_1,
		H:  telegram.NewStateFirstHandler(tgClient),
	})
	router.RegisterHandler(arguments.HandlerArgs{
		H:     telegram.NewStateSecondHandler(tgClient),
		State: statesGroup.State1,
	})

	fsmContext := telegram.NewFSMContext()

	eventProcessor := usecase.NewEventProcessor(tgClient, router, fsmContext)

	fetcher := usecase.Fetcher(eventProcessor)
	processor := usecase.Processor(eventProcessor)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	consumer := usecase.NewConsumer(processor, fetcher)

	consumer.Start(ctx)
	fmt.Println("Consumer успешно запущен...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	fmt.Printf("Получен сигнал %s, начинаем остановку...\n", sig)

	cancel()

	consumer.Stop()
	fmt.Println("Приложение успешно остановлено.")
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
