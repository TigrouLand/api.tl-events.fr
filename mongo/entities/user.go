package entities

import "time"

type User struct {
	ID               int64            `json:"id" bson:"id"`
	DiscordProfile   DiscordProfile   `json:"discordProfile" bson:"discordProfile"`
	MinecraftProfile MinecraftProfile `json:"minecraftProfile" bson:"minecraftProfile"`
	OAuth            OAuthUser        `json:"-" bson:"oauth"`
}

type DiscordProfile struct {
	ID        int64     `json:"id" bson:"id"`
	Username  string    `json:"username" bson:"username"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type MinecraftProfile struct {
	UUID      string    `json:"uuid" bson:"uuid"`
	Username  string    `json:"username" bson:"username"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type OAuthUser struct {
	AccessToken  string    `json:"-" bson:"accessToken"`
	RefreshToken string    `json:"-" bson:"refreshToken"`
	ExpiresAt    time.Time `json:"-" bson:"expiresAt"`
}
