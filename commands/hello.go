package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"bot/utils"
	"fmt"
)

type HelloCommand struct {
	interactionCreate messages.InteractionCreate
}

func (d HelloCommand) Execute() error {

	name := utils.GetUsername(&d.interactionCreate)

	interactionCallback := utils.CreateInteractionCallback().
		AddContent(
			fmt.Sprintf("Hello, %s.", name))

	discord.PostInteractionCallback(d.interactionCreate.D.ID, d.interactionCreate.D.Token, interactionCallback.Get())

	return nil
}

func (d HelloCommand) Respond() error {
	return nil
}
