package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"telegram-api-service/internal/entitiy"
)

type HttpClient struct {
	Host    string
	SubPath string
	client  *http.Client
	offset  int64
}

func (h *HttpClient) OffsetUpdate(newOffset int64) {
	if newOffset > h.offset {
		h.offset = newOffset
	}
	h.saveOffset(h.offset)
}

func NewHttpClient(token string) *HttpClient {
	client := http.Client{
		Transport: &http.Transport{
			Proxy: func(r *http.Request) (*url.URL, error) {
				proxy := os.Getenv("PROXY")

				return url.Parse(proxy)
			},
		},
	}

	h := HttpClient{
		Host:    "https://api.telegram.org",
		SubPath: getSubpath(token),
		client:  &client,
	}

	offset := h.loadOffset()
	h.offset = offset

	return &h
}

func getSubpath(token string) string {
	return "bot" + token
}

func (c *HttpClient) path(method string) string {
	path := c.Host + "/" + c.SubPath + "/" + method + "?offset=" + strconv.FormatInt(c.offset+1, 10)

	fmt.Println(path)

	return path
}

func (c *HttpClient) Updates() []byte {
	r, err := c.client.Get(c.path("getUpdates"))

	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		return nil
	}

	fmt.Println(string(resp))

	return resp
}

func (c *HttpClient) SendMessage(query entitiy.SendMessageQuery) {
	jsonData, err := json.Marshal(query)
	if err != nil {
		log.Printf("Ошибка маршалинга JSON: %v", err)
		return
	}

	fmt.Println(string(jsonData))

	resp, err := c.client.Post(c.path("sendMessage"), "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Printf("Ошибка отправки запроса: %v", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Printf("Telegram API вернул статус: %d", resp.StatusCode)
	}
}
