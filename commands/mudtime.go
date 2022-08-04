package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"bot/utils"
	"fmt"
	"time"
)

type MudTimeCommand struct {
	interactionCreate messages.InteractionCreate
}

func (d MudTimeCommand) Execute() error {

	inputTimeString := d.interactionCreate.D.Data.Options[0].Value.(string)

	location, err := time.LoadLocation("Europe/Rome")
	if err != nil {
		return fmt.Errorf("load location: %w", err)
	}

	mudLocation, err := time.LoadLocation("Europe/Tallinn")
	if err != nil {
		return fmt.Errorf("load location: %w", err)
	}

	inputTime, err := time.ParseInLocation(time.Kitchen, inputTimeString, location)
	if err != nil {
		return fmt.Errorf("parse time: %w", err)
	}

	inputTime.In(mudLocation)

	interactionCallback := utils.CreateInteractionCallback().
		AddContent(fmt.Sprintf("In MudTime: %s", inputTime.String())).
		Get()

	discord.PostInteractionCallback(d.interactionCreate.D.ID, d.interactionCreate.D.Token, interactionCallback)

	return nil
}

func (d MudTimeCommand) Respond() error {
	return nil
}
