package commands

import (
	"bot/cache"
	"bot/discord"
	"bot/discord/messages"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Poll struct {
	Options []Option
}

type Option struct {
	Description string
	Votes       int
}

func pollCommmand(interactionCreate messages.InteractionCreate) error {

	question := interactionCreate.D.Data.Options[0].Value.(string)
	options := interactionCreate.D.Data.Options[1].Value.(string)
	optionsArray := strings.Split(options, ";")

	var fields []messages.Field
	var optionsSelect []messages.Option
	var optionCache []Option

	for i, option := range optionsArray {

		fields = append(fields, messages.Field{
			Name:  option,
			Value: "0",
		})

		optionsSelect = append(optionsSelect, messages.Option{
			Label: option,
			Value: string(i),
			Emoji: messages.Emoji{
				ID:   "625891304148303894",
				Name: "rogue",
			},
		})

		optionCache = append(optionCache, Option{
			Description: options,
			Votes:       0,
		})
	}

	var interactionCallback messages.InteractionCallback
	interactionCallback.Type = 4
	interactionCallback.Embeds = append(interactionCallback.Embeds, messages.Embed{
		Title:       "Poll",
		Description: question,
		Fields:      fields,
	})
	interactionCallback.Data.Components = append(interactionCallback.Data.Components, messages.Components{
		Type: 1,
		Components: []messages.Component{
			{
				Type:     3,
				Style:    1,
				Label:    "mm",
				CustomID: "1",
				Options:  optionsSelect,
			},
		},
	})

	marshalledResult, err := json.Marshal(Poll{
		Options: optionCache,
	})
	if err != nil {
		return fmt.Errorf("poll command: %s", err)
	}
	cache.Set(string(interactionCreate.D.ID), string(marshalledResult), 0)

	discord.PostInteractionCallback(interactionCreate.D.ID, interactionCreate.D.Token, &interactionCallback)

	return nil
}

func pollResponse(interactionCreate messages.InteractionCreate) error {

	value, err := cache.Get(string(interactionCreate.D.Message.Interaction.ID))
	if err != nil {
		return fmt.Errorf("error getting value from cache: %s", err)
	}

	var poll Poll
	err = json.Unmarshal([]byte(value), &poll)
	if err != nil {
		return fmt.Errorf("error unmarshaling value from cache: %s", err)
	}

	selectedValue, err := strconv.ParseInt(interactionCreate.D.Data.Values[0], 10, 64)
	poll.Options[selectedValue].Votes += 1

	return nil
}
