package telegram

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	offsetFile = "offset.cache"
)

func (h *HttpClient) saveOffset(offset int64) {
	str := strconv.FormatInt(offset, 10)
	tmpFile := offsetFile + ".tmp"

	if err := os.WriteFile(tmpFile, []byte(str), 0644); err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(tmpFile)
	if err == nil {
		err := f.Sync()
		if err != nil {
			return
		}
		err = f.Close()
		if err != nil {
			return
		}
	}

	if err = os.Rename(tmpFile, offsetFile); err != nil {
		fmt.Printf("Ошибка применения атомарного сохранения: %v\n", err)
	}
}

func (h *HttpClient) loadOffset() int64 {
	data, err := os.ReadFile(offsetFile)
	if err != nil {
		return 0
	}

	str := strings.TrimSpace(string(data))
	offset, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return 0
	}

	return offset
}
