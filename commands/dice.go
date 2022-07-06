package commands

import (
	"bot/discord"
	"fmt"
	"math/rand"
)

func diceCommand(interactionCreatePayload discord.InteractionCreatePayload) {

	faces := int64(interactionCreatePayload.D.Data.Options[0].Value.(float64))
	dices := int64(interactionCreatePayload.D.Data.Options[1].Value.(float64))
	var result []int64

	var i int64
	for i = 0; i < dices; i++ {
		result = append(result, rand.Int63n(faces)+1)
	}

	var interactionCallbackPayload discord.InteractionCallbackPayload
	interactionCallbackPayload.Type = 4
	interactionCallbackPayload.Data.Content = fmt.Sprintf("%dd%d, result: %d", dices, faces, result)

	discord.PostInteractionCallback(interactionCreatePayload.D.ID, interactionCreatePayload.D.Token, &interactionCallbackPayload)
}
