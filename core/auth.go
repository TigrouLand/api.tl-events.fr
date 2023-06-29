package core

import (
	"github.com/ravener/discord-oauth2"
	"github.com/tigrouland/api/mongo/entities"
	"golang.org/x/oauth2"
	"os"
)

type User struct {
	DiscordProfile   DiscordProfile   `json:"discordProfile"`
	MinecraftProfile MinecraftProfile `json:"minecraftProfile"`
}

type DiscordProfile struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type MinecraftProfile struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
}

func PlayerToUser(player entities.Player) User {
	return User{
		DiscordProfile: DiscordProfile{
			ID:       player.ID,
			Username: player.Name,
		},
		MinecraftProfile: MinecraftProfile{
			UUID:     player.DecodedUUID.String(),
			Username: player.Name,
		},
	}
}

var DiscordOAuth = &oauth2.Config{
	Endpoint:     discord.Endpoint,
	Scopes:       []string{discord.ScopeIdentify},
	RedirectURL:  os.Getenv("DISCORD_REDIRECT_URL"),
	ClientID:     os.Getenv("DISCORD_CLIENT_ID"),
	ClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
}
