# Discord Bot

## Overview

This Discord bot provides various functionalities, including weather information retrieval, language translation, and more. It is designed to be extensible, allowing for easy addition of new commands.

## Features

- **Help Command**: A command that lists all available commands and their descriptions.

- **Weather Command**: Retrieve current weather information for a specified city.

- **Translate Command**: Translate text to different languages using the Google Translate API.

- **Poll Command**: Create polls with multiple options for user voting.

- **Language Command**: Get a list of supported languages for translation.

## Setup

### Prerequisites

- Go programming language installed on your machine.
- Discord bot token.
- OpenWeatherMap API key for the weather functionality.
- Google Cloud API key for the translation functionality.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/discord-bot.git
   cd discord-bot
   
2. Create a .env file and add the following configuration:

    ```dotenv
    TOKEN=your_discord_bot_token
    WEATHER_API_KEY=your_openweathermap_api_key
    TRANSLATE_API_KEY=your_google_cloud_api_key

3. Build and run the bot:

    ```bash
    go build
    ./discord-bot

## Usage
To interact with the bot, use the specified command prefix followed by one of the available commands. For example:

- !help: Display a list of available commands.
- !language: Get a list of supported languages for translation.
- !poll [question]? -[option1] -[option2] ...
- !weather [city]
- !translate [language code] [text]
