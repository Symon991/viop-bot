package commands

import (
	"bot/cache"
	"bot/discord"
	"bot/discord/messages"
	"bot/utils"
	"bot/utils/builders"
	"encoding/json"
	"fmt"
	"strconv"

	pirate "github.com/symon991/pirate/sites"
)

type PirateCommand struct {
	interactionCreate messages.InteractionCreate
}

type PirateEntry struct {
	Metadata []pirate.Metadata
	Site     string
}

func (d PirateCommand) Execute() error {

	search := pirate.GetSearch(d.interactionCreate.D.Data.Options[1].Value.(string))
	site := search.GetName()

	result, err := search.Search(d.interactionCreate.D.Data.Options[0].Value.(string))
	if err != nil {
		return fmt.Errorf("search %s: %w", site, err)
	}

	marshalledResult, err := json.Marshal(PirateEntry{
		Metadata: result,
		Site:     d.interactionCreate.D.Data.Options[1].Value.(string),
	})
	if err != nil {
		return fmt.Errorf("marshal result: %w", err)
	}

	cache.Set(string(d.interactionCreate.D.ID), string(marshalledResult), 0)

	interactionCallback := builders.CreateInteractionCallback()
	selectComponent := builders.CreateSelectComponent(site, "test")

	for i, resultItem := range result[0:utils.Min(len(result), 20)] {

		selectComponent.AddOption(
			builders.CreateOption(
				resultItem.Name[:utils.Min(len(resultItem.Name), 100)],
				fmt.Sprintf("Seeders: %s, Size: %s",
					resultItem.Seeders, resultItem.Size), fmt.Sprintf("%d", i)))
	}

	interactionCallback.AddEmbed(
		builders.CreateEmbed("",
			fmt.Sprintf(
				"Results for \"%s\" on %s",
				d.interactionCreate.D.Data.Options[0].Value.(string),
				site)))

	interactionCallback.AddActionRowComponent(
		builders.CreateActionRowComponent().
			AddComponent(selectComponent))

	_, err = discord.PostInteractionResponse(d.interactionCreate.D.ID, d.interactionCreate.D.Token, interactionCallback.Get())
	if err != nil {
		return fmt.Errorf("post interaction callback: %w", err)
	}

	return nil
}

func (d PirateCommand) Respond() error {

	value, err := cache.Get(string(d.interactionCreate.D.Message.Interaction.ID))
	if err != nil {
		return fmt.Errorf("get value from cache: %w", err)
	}

	var cachePirateEntry PirateEntry
	err = json.Unmarshal([]byte(value), &cachePirateEntry)
	if err != nil {
		return fmt.Errorf("unmarshal value from cache: %w", err)
	}

	search := pirate.GetSearch(cachePirateEntry.Site)

	fmt.Printf("debug selected value %s\n\n", d.interactionCreate.D.Data.Values[0])
	index, _ := strconv.ParseInt(d.interactionCreate.D.Data.Values[0], 10, 64)

	magnet, err := search.GetMagnet(cachePirateEntry.Metadata[index])
	if err != nil {
		return fmt.Errorf("get magnet: %w", err)
	}

	interactionCallback := builders.CreateInteractionCallback().
		AddEmbed(
			builders.CreateEmbed(
				cachePirateEntry.Metadata[index].Name[:utils.Min(len(cachePirateEntry.Metadata[index].Name), 100)],
				fmt.Sprintf("`%s`", magnet)))

	_, err = discord.PostInteractionResponse(d.interactionCreate.D.ID, d.interactionCreate.D.Token, interactionCallback.Get())
	if err != nil {
		return fmt.Errorf("post interaction callback: %w", err)
	}

	return nil
}
