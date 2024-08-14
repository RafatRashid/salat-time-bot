package cmd

import (
	"telegram-bot/services/telegram"

	"github.com/spf13/cobra"
)

var notifyPrayerCmd = &cobra.Command{
	Use:   "notify-prayer",
	Short: "Notify telegram on each prayer times",
	Long:  "============",
	Run:   runPrayerNotifier,
}

func runPrayerNotifier(cmd *cobra.Command, args []string) {
	telegram.NewTelegramBot().Ping()

}
