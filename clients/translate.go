package clients

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/translate/v2"
)

// TranslateClient represents a client for interacting with the Google Cloud Translate API.
type TranslateClient struct {
	service *translate.Service
}

// NewTranslateClient creates a new instance of TranslateClient with the provided API key.
func NewTranslateClient(apiKey string) (*TranslateClient, error) {
	ctx := context.Background()

	// Initialize the Translate service client with the provided API key
	service, err := translate.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("error creating Translate service client: %v", err)
	}
	// Return a new TranslateClient instance
	return &TranslateClient{
		service: service,
	}, nil
}

// TranslateText translates the given text to the target language using the Google Cloud Translate API.
func (c *TranslateClient) TranslateText(text, targetLanguage string) (string, error) {
	// Perform translation request to the Translate API
	resp, err := c.service.Translations.List([]string{text}, targetLanguage).Do()
	if err != nil {
		return "", fmt.Errorf("error translating text: %v", err)
	}

	// Check if translations are available in the response
	if len(resp.Translations) > 0 {
		// Return the translated text
		return resp.Translations[0].TranslatedText, nil
	}
	// Return an error if no translations are found
	return "", fmt.Errorf("no translations found")
}
