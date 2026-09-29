package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/lun4v3rd3-del/go-telegram-api/pkg/entitiy/arguments"
	"github.com/lun4v3rd3-del/go-telegram-api/pkg/infrastructure/telegram"
	"github.com/lun4v3rd3-del/go-telegram-api/pkg/usecase"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	t := token()

	fmt.Printf("Loaded token %s\n", t)

	bot := usecase.NewBot(t)

	re, _ := regexp.Compile("hello")

	bot.Router.RegisterHandler(arguments.HandlerArgs{
		Re: re,
		H:  telegram.TestHandler,
	})
	bot.Router.RegisterHandler(arguments.HandlerArgs{
		H:  telegram.CallbackHandler,
		Cd: "callback_data_1",
	})

	statesGroup := telegram.TestGet()
	re_1, _ := regexp.Compile("test")

	bot.Router.RegisterHandler(arguments.HandlerArgs{
		Re: re_1,
		H:  telegram.StateFirstHandler,
	})
	bot.Router.RegisterHandler(arguments.HandlerArgs{
		H:     telegram.StateSecondHandler,
		State: statesGroup.State1,
	})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := bot.Start(ctx); err != nil {
			fmt.Printf("Ошибка при работе бота: %v\n", err)
		}
	}()
	fmt.Println("Bot успешно запущен...")

	<-ctx.Done()
	fmt.Println("Получен сигнал от системы, начинаем graceful shutdown...")

	bot.Stop()
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
