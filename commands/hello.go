package commands

import (
	"bot/discord"
	"bot/utils"
	"fmt"
)

func helloCommand(interactionCreatePayload discord.InteractionCreatePayload) {

	name := utils.GetUsername(&interactionCreatePayload)

	var interactionCallbackPayload discord.InteractionCallbackPayload
	interactionCallbackPayload.Type = 4
	interactionCallbackPayload.Data.Content = fmt.Sprintf("Hello, %s.", name)

	discord.PostInteractionCallback(interactionCreatePayload.D.ID, interactionCreatePayload.D.Token, &interactionCallbackPayload)
}
