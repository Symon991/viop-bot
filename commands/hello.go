package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"bot/utils"
	"bot/utils/builders"
	"fmt"
)

type HelloCommand struct {
	interactionCreate messages.InteractionCreate
}

func (d HelloCommand) Execute() error {

	name := utils.GetUsername(&d.interactionCreate)

	interactionCallback := builders.CreateInteractionCallback().
		AddContent(
			fmt.Sprintf("Hello, %s.", name))

	discord.PostInteractionResponse(d.interactionCreate.D.ID, d.interactionCreate.D.Token, interactionCallback.Get())

	return nil
}

func (d HelloCommand) Respond() error {
	return nil
}
