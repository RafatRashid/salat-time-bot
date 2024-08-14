package telegram

import (
	"fmt"
	"telegram-bot/dto"
	"telegram-bot/services/api"
	"telegram-bot/utils"
	"time"

	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

type TelegramBot struct {
	apiToken string
}

func NewTelegramBot() TelegramBot {
	return TelegramBot{
		apiToken: viper.GetString("TELEGRAM_BOT_TOKEN"),
	}
}

func (b TelegramBot) Ping() {
	defer utils.RecoverPanic()

	utils.LogInfo("Initiating telegram bot with token: ", b.apiToken)

	bot, err := tgBotApi.NewBotAPI(b.apiToken)
	if err != nil {
		utils.LogError("TelegramBot.Ping - error: ", err.Error())
		return
	}

	bot.Debug = true

	updateConfig := tgBotApi.NewUpdate(0)
	updateConfig.Timeout = 30

	updateChannel := bot.GetUpdatesChan(updateConfig)

	salatTimesForToday := api.GetDailyPrayerTimes(utils.DhakaLat, utils.DhakaLng)

	for update := range updateChannel {
		if update.Message == nil {
			continue
		}

		messageContent := createMessage(update.Message.Time(), salatTimesForToday)

		msg := tgBotApi.NewMessage(update.Message.Chat.ID, messageContent)
		if _, err := bot.Send(msg); err != nil {
			utils.LogError("TelegramBot.Ping - error: ", err.Error())
		}

		utils.LogInfo("sent message")
	}
}

func createMessage(messageTime time.Time, salatTimesForToday dto.SalatTimeResponse) string {
	messageContent := fmt.Sprintf("Salat times for %s\n\n", messageTime.Format("2006-01-02"))

	messageContent += fmt.Sprintf("Fajr - %v\n", salatTimesForToday.Data.Timings.Fajr)
	messageContent += fmt.Sprintf("Sunrise - %v\n", salatTimesForToday.Data.Timings.Sunrise)
	messageContent += fmt.Sprintf("Dhuhr - %v\n", salatTimesForToday.Data.Timings.Dhuhr)
	messageContent += fmt.Sprintf("Asr - %v\n", salatTimesForToday.Data.Timings.Asr)
	messageContent += fmt.Sprintf("Sunset - %v\n", salatTimesForToday.Data.Timings.Sunset)
	messageContent += fmt.Sprintf("Maghrib - %v\n", salatTimesForToday.Data.Timings.Maghrib)
	messageContent += fmt.Sprintf("Isha - %v\n", salatTimesForToday.Data.Timings.Isha)

	return messageContent
}
