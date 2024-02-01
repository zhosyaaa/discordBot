package main

import (
	"discordBot/bot"
	"discordBot/clients"
	"fmt"
	"time"
)

func main() {
	token := ""
	weatherApiKey := ""
	weatherApiClient := clients.NewWeatherClient(weatherApiKey)
	fmt.Println(token)
	newBot := bot.NewBot(token, weatherApiClient)
	newBot.Run()
	go func() {
		for {
			newBot.ReminderManager.CheckReminders(newBot.Session)
			time.Sleep(1 * time.Minute)
		}
	}()

}
