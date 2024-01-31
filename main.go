package main

import (
	"discordBot/bot"
	"fmt"
)

func main() {
	token := "MTIwMjI5MDI2NjE1Mzk0NzIwOA.GdXquL.5_gTjchHc-0ueTgY4rRPtnFw8YmgPvJbdsn-2s"
	fmt.Println(token)
	newBot := bot.NewBot(token)
	newBot.Run()
}
