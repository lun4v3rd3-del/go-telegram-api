package usecase

import (
	"context"
	"github.com/lun4v3rd3-del/go-telegram-api/internal/entitiy"
	"github.com/lun4v3rd3-del/go-telegram-api/internal/infrastructure/telegram"
	"log"
	_ "log"
	"sync"
	"time"
)

type Bot struct {
	processor Processor
	fetcher   Fetcher
	producers *sync.WaitGroup
	consumers *sync.WaitGroup
	Router    telegram.Router
	Client    telegram.HttpClient
	channel   chan entitiy.Event
	errChan   chan error
}

func NewBot(token string) *Bot {
	client := telegram.NewHttpClient(token)
	router := telegram.NewRouter()

	fsmContext := telegram.NewFSMContext()

	eventProcessor := NewEventProcessor(client, router, fsmContext)

	fetcher := Fetcher(eventProcessor)
	processor := Processor(eventProcessor)

	return &Bot{
		processor: processor,
		fetcher:   fetcher,
		Router:    *router,
		Client:    *client,
		producers: &sync.WaitGroup{},
		consumers: &sync.WaitGroup{},
		channel:   make(chan entitiy.Event, 100),
		errChan:   make(chan error, 100),
	}
}

func (c *Bot) Start(ctx context.Context) error {
	c.producers.Add(1)

	go func() {
		select {
		case <-ctx.Done():
			return
		case err := <-c.errChan:
			log.Println(err)
		}
	}()

	go func(id int64) {
		defer c.producers.Done()
		for {
			time.Sleep(time.Millisecond * 50)
			select {
			case <-ctx.Done():
				return
			default:
				events, err := c.fetcher.Fetch()

				if err != nil {
					c.errChan <- err
				}

				if len(events) == 0 {
					continue
				}
				for _, e := range events {
					select {
					case <-ctx.Done():
						return
					case c.channel <- e:
					}
				}
			}
		}
	}(1)

	c.consumers.Add(4)

	for i := 0; i < 1; i++ {
		go func(id int64) {
			defer c.consumers.Done()
			for event := range c.channel {
				err := c.processor.Process(event)
				if err != nil {
					c.errChan <- err
				}
			}
		}(int64(i))
	}

	return nil
}

func (c *Bot) Stop() {
	c.producers.Wait()
	c.consumers.Wait()
}
