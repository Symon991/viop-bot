package commands

import (
	"bot/cache"
	"bot/commands/pirate"
	"bot/discord"
	"bot/discord/messages"
	"bot/utils"
	"encoding/json"
	"fmt"
	"strconv"
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

	interactionCallback := utils.CreateInteractionCallback()
	selectComponent := utils.CreateSelectComponent(site, "test")

	for i, resultItem := range result[0:utils.Min(len(result), 20)] {

		selectComponent.AddOption(
			utils.CreateOption(
				resultItem.Name[:utils.Min(len(resultItem.Name), 100)],
				fmt.Sprintf("Seeders: %s, Size: %s",
					resultItem.Seeders, resultItem.Size), fmt.Sprintf("%d", i)))
	}

	interactionCallback.AddEmbed(
		utils.CreateEmbed("",
			fmt.Sprintf(
				"Results for \"%s\" on %s",
				d.interactionCreate.D.Data.Options[0].Value.(string),
				site)))

	interactionCallback.AddActionRowComponent(
		utils.CreateActionRowComponent().
			AddComponent(selectComponent))

	_, err = discord.PostInteractionCallback(d.interactionCreate.D.ID, d.interactionCreate.D.Token, interactionCallback.Get())
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

	interactionCallback := utils.CreateInteractionCallback().
		AddEmbed(
			utils.CreateEmbed(
				cachePirateEntry.Metadata[index].Name[:utils.Min(len(cachePirateEntry.Metadata[index].Name), 100)],
				fmt.Sprintf("`%s`", search.GetMagnet(cachePirateEntry.Metadata[index]))))

	_, err = discord.PostInteractionCallback(d.interactionCreate.D.ID, d.interactionCreate.D.Token, interactionCallback.Get())
	if err != nil {
		return fmt.Errorf("post interaction callback: %w", err)
	}

	return nil
}
