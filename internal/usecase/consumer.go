package usecase

import (
	"context"
	"sync"
	"telegram-api-service/internal/entitiy"
	"time"
)

type Consumer struct {
	processor Processor
	fetcher   Fetcher
	producers *sync.WaitGroup
	consumers *sync.WaitGroup
	channel   chan entitiy.Event
}

func NewConsumer(processor Processor, fetcher Fetcher) *Consumer {
	return &Consumer{
		processor: processor,
		fetcher:   fetcher,
		producers: &sync.WaitGroup{},
		consumers: &sync.WaitGroup{},
		channel:   make(chan entitiy.Event, 100),
	}
}

func (c *Consumer) Start(ctx context.Context) {
	c.producers.Add(1)

	go func(id int64) {
		defer c.producers.Done()
		for {
			time.Sleep(time.Millisecond * 50)
			select {
			case <-ctx.Done():
				return
			default:
				events := c.fetcher.Fetch()
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
				c.processor.Process(event)
			}
		}(int64(i))
	}
}

func (c *Consumer) Stop() {
	c.producers.Wait()
	c.consumers.Wait()
}
