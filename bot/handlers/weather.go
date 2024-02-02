package handlers

import (
	"discordBot/clients"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

// WeatherCommand represents a command for fetching weather information for a given city.
type WeatherCommand struct {
	Weather *clients.WeatherClient
}

// Execute implements the Command interface for the WeatherCommand.
// It fetches weather information for the specified city and sends the details to the Discord channel.
func (w *WeatherCommand) Execute(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	// Split the words in the message content
	words := strings.Fields(message.Content)

	// Check if there are enough arguments to get weather information
	if len(words) < 2 {
		session.ChannelMessageSend(message.ChannelID, "Usage: !weather [city]")
		return
	}

	// Extract the city from the command arguments
	city := words[1]

	// Get weather data for the specified city
	weather, err := w.Weather.GetWeatherData(city)
	if err != nil {
		log.Printf("Error getting weather information: %v", err)
		session.ChannelMessageSend(message.ChannelID, "Error getting weather information. Please try again later.")
		return
	}
	// Build and send the weather response to the Discord channel
	response := buildWeatherResponse(city, weather)
	session.ChannelMessageSend(message.ChannelID, response)
}

// buildWeatherResponse creates a formatted weather response string based on the provided weather data.
func buildWeatherResponse(city string, weather *clients.WeatherData) string {
	response := fmt.Sprintf("Weather information for %s:\n", city)
	response += fmt.Sprintf("Temperature: %.2f°C\n", weather.Main.Temp)
	response += fmt.Sprintf("Humidity: %.2f%%\n", weather.Main.Humidity)
	if len(weather.Weather) > 0 {
		response += fmt.Sprintf("Description: %s\n", weather.Weather[0].Description)
	}
	return response
}
