package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"fmt"
	"math/rand"
)

func diceCommand(interactionCreate messages.InteractionCreate) {

	faces := int64(interactionCreate.D.Data.Options[0].Value.(float64))
	dices := int64(interactionCreate.D.Data.Options[1].Value.(float64))
	var result []int64

	var i int64
	for i = 0; i < dices; i++ {
		result = append(result, rand.Int63n(faces)+1)
	}

	var interactionCallback messages.InteractionCallback
	interactionCallback.Type = 4
	interactionCallback.Data.Content = fmt.Sprintf("%dd%d, result: %d", dices, faces, result)

	discord.PostInteractionCallback(interactionCreate.D.ID, interactionCreate.D.Token, &interactionCallback)
}
