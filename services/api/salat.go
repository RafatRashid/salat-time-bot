package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"telegram-bot/dto"
	"telegram-bot/utils"
	"time"
)

const baseUrl = "http://api.aladhan.com/v1/timings"

func GetDailyPrayerTimes(lat float64, lng float64) (response dto.SalatTimeResponse) {
	today := time.Now().Format("01-02-2006")

	url := fmt.Sprintf("%s/%s?latitude=%f&longitude=%f&method=4", baseUrl, today, lat, lng)
	resp, err := http.Get(url)

	if err != nil {
		utils.LogError(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.LogError(err)
	}

	if err = json.Unmarshal(body, &response); err != nil {
		utils.LogError(err)
	}

	return response
}
