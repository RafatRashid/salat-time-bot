package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	RootCmd = &cobra.Command{
		Use: "telegram-bot [command]",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
)

func Execute(args []string) error {
	loadEnvironment()

	RootCmd.AddCommand(consoleCmd)
	RootCmd.AddCommand(notifyPrayerCmd)

	if len(args) == 0 {
		RootCmd.SetArgs([]string{notifyPrayerCmd.Name()})
	}

	return RootCmd.Execute()
}

func loadEnvironment() {
	if _, err := os.Stat(".env"); errors.Is(err, os.ErrNotExist) {
		viper.AutomaticEnv()
	} else {
		viper.SetConfigFile(".env")

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
	}
}
