package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"telegram-bot/utils"
	"time"
)

const baseUrl = "http://api.aladhan.com/v1/timings"

func GetTodaysPrayerTimes(lat float64, lng float64) any {
	url := fmt.Sprintf("%s/%s?latitude=%f&longitude=%f&method=4", baseUrl, time.Now().Format("01-02-2006"), lat, lng)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	utils.LogError(sb)

	return nil
}
