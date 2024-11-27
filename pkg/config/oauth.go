package config

import (
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConfig *oauth2.Config

func InitOAuthConfig() {
	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"email","profile"},
		Endpoint:     google.Endpoint,
	}

	if GoogleOAuthConfig.ClientID == "" || GoogleOAuthConfig.ClientSecret == "" {
		log.Fatal("Missing Google OAuth2 credentials. Check .env file.")
	}
}
