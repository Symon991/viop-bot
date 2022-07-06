package commands

import (
	"bot/cache"
	"bot/discord"
	"bot/utils"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/symon991/pirate/sites"
)

func pirateCommand(interactionCreatePayload discord.InteractionCreatePayload) {

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

	marshalledResult, _ := json.Marshal(discord.CachePirateEntry{
		Metadata: result,
		Site:     interactionCreatePayload.D.Data.Options[1].Value.(string),
	})

	cache.Set(string(interactionCreatePayload.D.ID), string(marshalledResult), 0)

	var interactionCallbackPayload discord.InteractionCallbackPayload
	interactionCallbackPayload.Type = 4

	var options []discord.Option
	for i, resultItem := range result[0:utils.Min(len(result), 5)] {

		options = append(options, discord.Option{
			Label:       resultItem.Name[:utils.Min(len(resultItem.Name), 100)],
			Value:       fmt.Sprintf("%d", i),
			Description: fmt.Sprintf("Seeders: %s, Size: %s", resultItem.Seeders, resultItem.Size),
			Emoji: discord.Emoji{
				ID:   "625891304148303894",
				Name: "rogue",
			},
		})
	}

	interactionCallbackPayload.Data.Embeds = append(interactionCallbackPayload.Data.Embeds,
		discord.Embed{
			Description: fmt.Sprintf("Results for \"%s\" on %s", interactionCreatePayload.D.Data.Options[0].Value.(string), site),
		})
	interactionCallbackPayload.Data.Components = append(interactionCallbackPayload.Data.Components, discord.Components{
		Type: 1,
		Components: []discord.Component{
			{
				Type:     3,
				Style:    1,
				Label:    site,
				CustomID: "1",
				Options:  options,
			},
		},
	})

	discord.PostInteractionCallback(interactionCreatePayload.D.ID, interactionCreatePayload.D.Token, &interactionCallbackPayload)
}

func pirateResponse(interactionCreatePayload discord.InteractionCreatePayload) error {

	value, err := cache.Get(string(interactionCreatePayload.D.Message.Interaction.ID))
	if err != nil {
		return fmt.Errorf("error getting value from cache: %s", err)
	}

	var cachePirateEntry discord.CachePirateEntry
	err = json.Unmarshal([]byte(value), &cachePirateEntry)
	if err != nil {
		return fmt.Errorf("error unmarshaling value from cache: %s", err)
	}

	fmt.Printf("debug selected value %s\n\n", interactionCreatePayload.D.Data.Values[0])
	index, _ := strconv.ParseInt(interactionCreatePayload.D.Data.Values[0], 10, 64)

	var interactionCallbackPayload discord.InteractionCallbackPayload
	interactionCallbackPayload.Type = 4

	var trackers []string
	switch cachePirateEntry.Site {
	case "piratebay":
		trackers = sites.PirateBayTrackers()
	case "nyaa":
		trackers = sites.NyaaTrackers()
	}

	/*tinyUrlRequest, err := json.Marshal(discord.TinyUrlRequest{
		Url:    sites.GetMagnet(cachePirateEntry.Metadata[index], trackers),
		Domain: "tiny.one",
	})
	if err != nil {
		return fmt.Errorf("error marshaling tinyurl rquest: %s", err)
	}
	response, err := http.Post("https://api.tinyurl.com/create?api_token=Ex3deMVSNtPhqlRtHtyk6OxFV4Ubyv5pW7rV9zTS3eRoPq9kW7IqiFAqOmvH", "application/json", bytes.NewBuffer(tinyUrlRequest))
	if err != nil {
		return fmt.Errorf("error creating tinyurl: %s", err)
	}

	tinyUrlResponseBytes, _ := io.ReadAll(response.Body)
	var tinyUrlResponse discord.TinyUrlResponse
	fmt.Printf("%s\n\n", tinyUrlResponseBytes)
	json.Unmarshal(tinyUrlResponseBytes, &tinyUrlResponse)*/

	interactionCallbackPayload.Data.Embeds = append(interactionCallbackPayload.Data.Embeds,
		discord.Embed{
			Title:       cachePirateEntry.Metadata[index].Name[:utils.Min(len(cachePirateEntry.Metadata[index].Name), 100)],
			Description: fmt.Sprintf("`%s`", sites.GetMagnet(cachePirateEntry.Metadata[index], trackers)),
		})
	/*interactionCallbackPayload.Data.Components = append(interactionCallbackPayload.Data.Components, discord.Components{
		Type: 1,
		Components: []discord.Component{
			{
				Type:  2,
				Style: 5,
				Label: "Magnet",
				Url:   "",
			},
		},
	})*/

	err = discord.PostInteractionCallback(interactionCreatePayload.D.ID, interactionCreatePayload.D.Token, &interactionCallbackPayload)
	if err != nil {
		return fmt.Errorf("error unmarshaling value from cache: %s", err)
	}

	return nil
}
