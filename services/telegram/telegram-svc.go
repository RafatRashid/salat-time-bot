package telegram

import (
	"fmt"
	"telegram-bot/utils"

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

	for update := range updateChannel {
		if update.Message == nil {
			continue
		}

		location := update.Message.Location
		messageTime := update.Message.Time()
		messageContent := fmt.Sprintf("You have sent this message from [lat: %f, lng: %f] @%s", location.Latitude, location.Longitude, messageTime.String())

		msg := tgBotApi.NewMessage(update.Message.Chat.ID, messageContent)
		if _, err := bot.Send(msg); err != nil {
			utils.LogError("TelegramBot.Ping - error: ", err.Error())
		}
		utils.LogInfo("sent message")
	}
}
