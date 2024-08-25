package telegram

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"telegram-bot/dto"
	"telegram-bot/infra/cache"
	"telegram-bot/services/api"
	"telegram-bot/utils"
	"time"

	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

type Bot struct {
	bot *tgBotApi.BotAPI
}

func NewTelegramBot() Bot {
	apiToken := viper.GetString("TELEGRAM_BOT_TOKEN")

	utils.LogInfo("Initiating telegram bot with token: ", apiToken)

	bot, err := tgBotApi.NewBotAPI(apiToken)
	if err != nil {
		utils.LogError("TelegramBot.Ping - error: ", err.Error())
		panic(fmt.Sprintf("could not initialize bot - %s", err.Error()))
	}

	bot.Debug = true

	return Bot{
		bot,
	}
}

func (b Bot) Run() {
	go b.SubscribeForNotification()
	b.Ping()

	gracefullyShutdown()
}

func gracefullyShutdown() {
	terminator := make(chan os.Signal, 1)
	signal.Notify(terminator, os.Kill)
	<-terminator
	utils.LogInfo("shutting down bot")
}

func (b Bot) SubscribeForNotification() {
	defer utils.RecoverPanic()

	updateConfig := tgBotApi.NewUpdate(0)
	updateConfig.Timeout = 30
	updateChannel := b.bot.GetUpdatesChan(updateConfig)

	utils.LogInfo("listening for subscriptions...")

	for update := range updateChannel {
		utils.LogInfo("GOT UPDATE ----> %v", utils.ToJson(update))

		if update.Message == nil {
			continue
		}

		command := update.Message.Command()
		switch command {
		case "start":
			subscribeChat(update.Message.Chat.ID)

		case "stop":
			unsubscribeChat(update.Message.Chat.ID)
		}
	}

	utils.LogInfo("shutting down subscription listener")
}

func subscribeChat(chatId int64) {
	if err := cache.SetString(getSubscriptionCacheKey(strconv.FormatInt(chatId, 10)), "1", -1); err != nil {
		utils.LogError("Bot.subscribeChat - error [", err.Error(), "] on subscribing chat id: ", chatId)
	}
}

func unsubscribeChat(chatId int64) {
	chatIdString := fmt.Sprintf("%d", chatId)

	if err := cache.RemoveString(getSubscriptionCacheKey(strconv.FormatInt(chatId, 10))); err != nil {
		utils.LogInfo("Bot.unsubscribeChat - chat id [", chatIdString, "] is not in database")
		return
	}
}

func getSubscriptionCacheKey(chatId string) string {
	return fmt.Sprintf("%s:%s", "chat-ids", chatId)
}

func (b Bot) Ping() {
	defer utils.RecoverPanic()

	salatTimesForToday := api.GetDailyPrayerTimes(utils.DhakaLat, utils.DhakaLng)

	tick := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-tick.C:
				b.sendSalatTimes(salatTimesForToday)
			}
		}
	}()
}

func (b Bot) sendSalatTimes(salatTimesForToday dto.SalatTimeResponse) {
	key := getSubscriptionCacheKey("*")
	chatIds, err := cache.GetFolderElements(key)
	if err != nil {
		utils.LogInfo("no subscribers")
	}

	msg := createMessage(salatTimesForToday)
	for _, idString := range chatIds {
		var id int64
		if splitted := strings.Split(idString, ":"); len(splitted) == 2 {
			id, _ = strconv.ParseInt(splitted[1], 10, 64)
		}

		if _, err := b.bot.Send(tgBotApi.NewMessage(id, msg)); err != nil {
			utils.LogError("Bot.sendSalatTimes -  error: ", err.Error())
		}
	}

	utils.LogInfo("Bot.sendSalatTimes - sent times to ", len(chatIds), " users")
}

func createMessage(salatTimesForToday dto.SalatTimeResponse) string {
	messageContent := fmt.Sprintf("Fajr - %v\n", salatTimesForToday.Data.Timings.Fajr)
	messageContent += fmt.Sprintf("Sunrise - %v\n", salatTimesForToday.Data.Timings.Sunrise)
	messageContent += fmt.Sprintf("Dhuhr - %v\n", salatTimesForToday.Data.Timings.Dhuhr)
	messageContent += fmt.Sprintf("Asr - %v\n", salatTimesForToday.Data.Timings.Asr)
	messageContent += fmt.Sprintf("Sunset - %v\n", salatTimesForToday.Data.Timings.Sunset)
	messageContent += fmt.Sprintf("Maghrib - %v\n", salatTimesForToday.Data.Timings.Maghrib)
	messageContent += fmt.Sprintf("Isha - %v\n", salatTimesForToday.Data.Timings.Isha)

	return messageContent
}
