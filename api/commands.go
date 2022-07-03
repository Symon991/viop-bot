package api

import (
	"fmt"
	"math/rand"

	"github.com/symon991/pirate/sites"
)

func diceCommand(interactionCreatePayload InteractionCreatePayload) {

	faces := int64(interactionCreatePayload.D.Data.Options[0].Value.(float64))
	dices := int64(interactionCreatePayload.D.Data.Options[1].Value.(float64))
	var result []int64

	var i int64
	for i = 0; i < dices; i++ {
		result = append(result, rand.Int63n(faces)+1)
	}

	var interactionCallbackPayload InteractionCallbackPayload
	interactionCallbackPayload.Type = 4
	interactionCallbackPayload.Data.Content = fmt.Sprintf("%dd%d, result: %d", dices, faces, result)

	postInteractionCallback(interactionCreatePayload.D.ID, interactionCreatePayload.D.Token, &interactionCallbackPayload)
}

func pirateCommand(interactionCreatePayload InteractionCreatePayload) {

	var result []sites.Metadata
	var site string

	switch interactionCreatePayload.D.Data.Options[1].Value {
	case "piratebay":
		site = "Pirate Bay"
		result = sites.SearchTorrent(interactionCreatePayload.D.Data.Options[0].Value.(string))
	case "nyaa":
		site = "Nyaa"
		result = sites.SearchNyaa(interactionCreatePayload.D.Data.Options[0].Value.(string))
	}

	var interactionCallbackPayload InteractionCallbackPayload
	interactionCallbackPayload.Type = 4

	var fields []Field
	for _, resultItem := range result[0:min(len(result), 5)] {
		fields = append(fields, Field{Name: fmt.Sprintf("%s - %s", resultItem.Name, resultItem.Seeders), Value: sites.GetMagnet(resultItem, sites.PirateBayTrackers())})
	}

	interactionCallbackPayload.Data.Embeds = append(interactionCallbackPayload.Data.Embeds,
		Embed{
			Description: fmt.Sprintf("Results for \"%s\" on %s", interactionCreatePayload.D.Data.Options[0].Value, site),
			Fields:      fields,
		})

	postInteractionCallback(interactionCreatePayload.D.ID, interactionCreatePayload.D.Token, &interactionCallbackPayload)
}

func helloCommand(interactionCreatePayload InteractionCreatePayload) {

	name := getUsername(&interactionCreatePayload)

	var interactionCallbackPayload InteractionCallbackPayload
	interactionCallbackPayload.Type = 4
	interactionCallbackPayload.Data.Content = fmt.Sprintf("Hello, %s.", name)

	postInteractionCallback(interactionCreatePayload.D.ID, interactionCreatePayload.D.Token, &interactionCallbackPayload)
}
