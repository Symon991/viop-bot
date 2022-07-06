package utils

import (
	"bot/discord"
	"encoding/json"
)

func GetUsername(interactionCreatePayload *discord.InteractionCreatePayload) string {

	name := interactionCreatePayload.D.User.Username
	if name == "" {
		name = interactionCreatePayload.D.Member.User.Username
	}
	return name
}

func ReadInteractionCreatePayload(message []byte) (string, bool, discord.InteractionCreatePayload) {

	var interactionCreatePayload discord.InteractionCreatePayload
	json.Unmarshal(message, &interactionCreatePayload)

	if interactionCreatePayload.D.Message.Interaction.ID != "" {
		return interactionCreatePayload.D.Message.Interaction.Name, true, interactionCreatePayload
	}
	return interactionCreatePayload.D.Data.Name, false, interactionCreatePayload
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
