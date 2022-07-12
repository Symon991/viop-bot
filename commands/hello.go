package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"bot/utils"
	"fmt"
)

func helloCommand(interactionCreate messages.InteractionCreate) {

	name := utils.GetUsername(&interactionCreate)

	var interactionCallback messages.InteractionCallback
	interactionCallback.Type = 4
	interactionCallback.Data.Content = fmt.Sprintf("Hello, %s.", name)

	discord.PostInteractionCallback(interactionCreate.D.ID, interactionCreate.D.Token, &interactionCallback)
}
