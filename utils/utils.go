package utils

import (
	"bot/discord/messages"
	"encoding/json"
)

func GetUsername(interactionCreatePayload *messages.InteractionCreate) string {

	name := interactionCreatePayload.D.User.Username
	if name == "" {
		name = interactionCreatePayload.D.Member.User.Username
	}
	return name
}

func ReadInteractionCreatePayload(message []byte) (string, bool, messages.InteractionCreate) {

	var interactionCreate messages.InteractionCreate
	json.Unmarshal(message, &interactionCreate)

	if interactionCreate.D.Message.Interaction.ID != "" {
		return interactionCreate.D.Message.Interaction.Name, true, interactionCreate
	}
	return interactionCreate.D.Data.Name, false, interactionCreate
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
