package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"bot/utils/builders"
	"fmt"
	"time"
)

type CountdownCommand struct {
	interactionCreate messages.InteractionCreate
}

func (d CountdownCommand) Execute() error {

	start := int64(d.interactionCreate.D.Data.Options[0].Value.(float64))

	ticker := time.NewTicker(time.Second)
	<-ticker.C

	interactionCallback := builders.CreateInteractionCallback().
		AddContent(fmt.Sprintf("%d", start)).
		Get()

	discord.PostInteractionResponse(d.interactionCreate.D.ID, d.interactionCreate.D.Token, interactionCallback)

	for {
		start = start - 1

		discord.EditOriginalInteraction(d.interactionCreate.D.ApplicationID, d.interactionCreate.D.Token, &messages.Data{
			Content: fmt.Sprintf("%d", start),
		})

		<-ticker.C

		if start <= 1 {
			break
		}
	}

	discord.EditOriginalInteraction(d.interactionCreate.D.ApplicationID, d.interactionCreate.D.Token, &messages.Data{
		Content: "Let's go!",
	})

	return nil
}

func (d CountdownCommand) Respond() error {
	return nil
}
