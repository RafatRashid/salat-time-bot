package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"io"
	"net/http"
	"telegram-bot/dto"
	"telegram-bot/infra/cache"
	"telegram-bot/utils"
	"time"
)

const baseUrl = "http://api.aladhan.com/v1/timings"

func GetDailyPrayerTimes(lat float64, lng float64) (response dto.SalatTimeResponse) {
	today := time.Now().Format("01-02-2006")

	if times, err := getCachedPrayerTimes(today); err == nil {
		return times
	}

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

	go cachePrayerTimes(today, response)

	return response
}

func cachePrayerTimes(today string, response dto.SalatTimeResponse) {
	defer utils.RecoverPanic()

	prayerTimesJson := utils.ToJson(response)
	if prayerTimesJson == "" {
		return
	}

	key := createCacheKey(today)
	if err := cache.SetString(key, prayerTimesJson, 48*time.Hour); err != nil {
		utils.LogError("salat.cachePrayerTimes - error : ,", err.Error())
	}
}

func getCachedPrayerTimes(today string) (response dto.SalatTimeResponse, err error) {
	key := createCacheKey(today)
	timesString, err := cache.GetString(key)
	if err != nil && err.Error() == redis.Nil.Error() {
		return response, err
	}

	if err = json.Unmarshal([]byte(timesString), &response); err != nil {
		utils.LogError("salat.getCachedPrayerTimes - error : ", err.Error())
		return response, err
	}

	utils.LogInfo("serving prayer times from cache [", today, "]")
	return response, err
}

func createCacheKey(today string) string {
	return "salat-times:" + today
}
