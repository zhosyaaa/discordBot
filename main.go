package main

import (
	"discordBot/bot"
	"discordBot/clients"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// loadEnv loads environment variables from the .env file.
func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func main() {
	// Load environment variables from .env file
	err := loadEnv()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Retrieve Discord bot token, Weather API key, and Translate API key from environment variables
	token := os.Getenv("TOKEN")
	weatherApiKey := os.Getenv("WEATHER_API_KEY")
	translateApiKey := os.Getenv("TRANSLATE_API_KEY")

	// Create instances of WeatherClient and TranslateClient
	weatherApiClient := clients.NewWeatherClient(weatherApiKey)
	translateApiClient, _ := clients.NewTranslateClient(translateApiKey)

	// Create a new Discord bot instance with the provided token, WeatherClient, and TranslateClient
	discordBot := bot.NewBot(token, weatherApiClient, translateApiClient)
	// Run the Discord bot
	discordBot.Run()
}
