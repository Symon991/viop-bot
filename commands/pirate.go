package commands

import (
	"bot/cache"
	"bot/discord"
	"bot/discord/messages"
	"bot/utils"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/symon991/pirate/sites"
)

func pirateCommand(interactionCreatePayload messages.InteractionCreate) error {

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

	marshalledResult, err := json.Marshal(cache.PirateEntry{
		Metadata: result,
		Site:     interactionCreatePayload.D.Data.Options[1].Value.(string),
	})
	if err != nil {
		return fmt.Errorf("pirateCommand: %s", err)
	}

	cache.Set(string(interactionCreatePayload.D.ID), string(marshalledResult), 0)

	var interactionCallbackPayload messages.InteractionCallback
	interactionCallbackPayload.Type = 4

	var options []messages.Option
	for i, resultItem := range result[0:utils.Min(len(result), 5)] {

		options = append(options, messages.Option{
			Label:       resultItem.Name[:utils.Min(len(resultItem.Name), 100)],
			Value:       fmt.Sprintf("%d", i),
			Description: fmt.Sprintf("Seeders: %s, Size: %s", resultItem.Seeders, resultItem.Size),
			Emoji: messages.Emoji{
				ID:   "625891304148303894",
				Name: "rogue",
			},
		})
	}

	interactionCallbackPayload.Data.Embeds = append(interactionCallbackPayload.Data.Embeds,
		messages.Embed{
			Description: fmt.Sprintf("Results for \"%s\" on %s", interactionCreatePayload.D.Data.Options[0].Value.(string), site),
		})
	interactionCallbackPayload.Data.Components = append(interactionCallbackPayload.Data.Components, messages.Components{
		Type: 1,
		Components: []messages.Component{
			{
				Type:     3,
				Style:    1,
				Label:    site,
				CustomID: "1",
				Options:  options,
			},
		},
	})

	err = discord.PostInteractionCallback(interactionCreatePayload.D.ID, interactionCreatePayload.D.Token, &interactionCallbackPayload)
	if err != nil {
		return fmt.Errorf("pirateCommand: %s", err)
	}

	return nil
}

func pirateResponse(interactionCreatePayload messages.InteractionCreate) error {

	value, err := cache.Get(string(interactionCreatePayload.D.Message.Interaction.ID))
	if err != nil {
		return fmt.Errorf("error getting value from cache: %s", err)
	}

	var cachePirateEntry cache.PirateEntry
	err = json.Unmarshal([]byte(value), &cachePirateEntry)
	if err != nil {
		return fmt.Errorf("error unmarshaling value from cache: %s", err)
	}

	fmt.Printf("debug selected value %s\n\n", interactionCreatePayload.D.Data.Values[0])
	index, _ := strconv.ParseInt(interactionCreatePayload.D.Data.Values[0], 10, 64)

	var interactionCallbackPayload messages.InteractionCallback
	interactionCallbackPayload.Type = 4

	var trackers []string
	switch cachePirateEntry.Site {
	case "piratebay":
		trackers = sites.PirateBayTrackers()
	case "nyaa":
		trackers = sites.NyaaTrackers()
	}

	interactionCallbackPayload.Data.Embeds = append(interactionCallbackPayload.Data.Embeds,
		messages.Embed{
			Title:       cachePirateEntry.Metadata[index].Name[:utils.Min(len(cachePirateEntry.Metadata[index].Name), 100)],
			Description: fmt.Sprintf("`%s`", sites.GetMagnet(cachePirateEntry.Metadata[index], trackers)),
		})

	err = discord.PostInteractionCallback(interactionCreatePayload.D.ID, interactionCreatePayload.D.Token, &interactionCallbackPayload)
	if err != nil {
		return fmt.Errorf("error unmarshaling value from cache: %s", err)
	}

	return nil
}
