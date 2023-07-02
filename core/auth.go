package core

import (
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
	"os"
)

var DiscordOAuth = &oauth2.Config{
	Endpoint:     discord.Endpoint,
	Scopes:       []string{discord.ScopeIdentify, "role_connections.write"},
	RedirectURL:  os.Getenv("DISCORD_REDIRECT_URL"),
	ClientID:     os.Getenv("DISCORD_CLIENT_ID"),
	ClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
}
