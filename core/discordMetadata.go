package core

import (
	"bytes"
	"encoding/json"
	"github.com/tigrouland/api/mongo/entities"
	"net/http"
	"os"
)

type UpdateRequest struct {
	PlatformName     string                 `json:"platform_name"`
	PlatformUsername string                 `json:"platform_username"`
	Metadata         map[string]interface{} `json:"metadata"`
}

func UpdateMetadata(user entities.User) error {
	statistics, err := GetStatistics(user.MinecraftProfile.Username)
	if err != nil {
		return err
	}

	updateRequest := UpdateRequest{
		PlatformName:     "[TL] Events",
		PlatformUsername: user.MinecraftProfile.Username,
		Metadata: map[string]interface{}{
			"kills":  statistics.Kills,
			"deaths": statistics.Deaths,
			"wins":   statistics.Wins,
		},
	}

	marshal, err := json.Marshal(updateRequest)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(marshal)

	request, err := http.NewRequest(
		"PUT",
		"https://discord.com/api/v10/users/@me/applications/"+os.Getenv("DISCORD_CLIENT_ID")+"/role-connection",
		reader,
	)
	if err != nil {
		return err
	}
	request.Header.Set("Authorization", user.OAuth.AccessToken)
	request.Header.Set("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	return nil
}
