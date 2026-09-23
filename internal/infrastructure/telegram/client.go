package telegram

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type HttpClient struct {
	Host    string
	SubPath string
	client  *http.Client
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

	return &HttpClient{
		Host:    "https://api.telegram.org",
		SubPath: getSubpath(token),
		client:  &client,
	}
}

func getSubpath(token string) string {
	return "bot" + token
}

func (c *HttpClient) path(method string) string {
	path := c.Host + "/" + c.SubPath + "/" + method

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

	resp := make([]byte, 12800)

	_, err = r.Body.Read(resp)
	if err != nil {
		return nil
	}

	return resp
}

func (c *HttpClient) SendMessage(chat_id int64, text string) {
	_, err := c.client.PostForm(c.path("sendMessage"), url.Values{
		"chat_id": []string{strconv.FormatInt(chat_id, 10)},
		"text":    []string{text},
	})

	if err != nil {
		log.Fatal(err)
	}
}
