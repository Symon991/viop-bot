package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"bot/utils"
	"fmt"
	"time"
)

type CountdownCommand struct {
	interactionCreate messages.InteractionCreate
}

func (d CountdownCommand) Execute() error {

	start := int64(d.interactionCreate.D.Data.Options[0].Value.(float64))

	ticker := time.NewTicker(time.Second)

	interactionCallback := utils.CreateInteractionCallback().
		AddContent(fmt.Sprintf("%d", start)).
		Get()

	discord.PostInteractionCallback(d.interactionCreate.D.ID, d.interactionCreate.D.Token, interactionCallback)

	for {
		<-ticker.C

		interactionCallback = utils.CreateInteractionCallback().
			AddContent(fmt.Sprintf("%d", start)).
			Get()

		discord.EditOriginalInteraction(d.interactionCreate.D.ID, d.interactionCreate.D.Token, "", interactionCallback)

		start = start - 1
		if start <= 0 {
			break
		}
	}

	return nil
}

func (d CountdownCommand) Respond() error {
	return nil
}
