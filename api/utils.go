package api

import "encoding/json"

func getUsername(interactionCreatePayload *InteractionCreatePayload) string {

	name := interactionCreatePayload.D.User.Username
	if name == "" {
		name = interactionCreatePayload.D.Member.User.Username
	}
	return name
}

func readInteractionCreatePayload(message []byte) (string, InteractionCreatePayload) {

	var interactionCreatePayload InteractionCreatePayload
	json.Unmarshal(message, &interactionCreatePayload)
	return interactionCreatePayload.D.Data.Name, interactionCreatePayload
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
