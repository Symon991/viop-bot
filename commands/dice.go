package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"bot/utils"
	"fmt"
	"math/rand"
	"time"
)

type DiceCommand struct {
	interactionCreate messages.InteractionCreate
}

func (d DiceCommand) Execute() error {

	faces := int64(d.interactionCreate.D.Data.Options[0].Value.(float64))
	dices := int64(d.interactionCreate.D.Data.Options[1].Value.(float64))
	var result []int64

	rand.Seed(time.Now().UnixNano())

	var i int64
	for i = 0; i < dices; i++ {
		result = append(result, rand.Int63n(faces)+1)
	}

	interactionCallback := utils.CreateInteractionCallback().
		AddContent(fmt.Sprintf("%dd%d, result: %d", dices, faces, result)).
		Get()

	discord.PostInteractionCallback(d.interactionCreate.D.ID, d.interactionCreate.D.Token, interactionCallback)

	return nil
}

func (d DiceCommand) Respond() error {
	return nil
}
