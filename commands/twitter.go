package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"bot/twitter"
	"bot/utils/builders"
	"fmt"
)

type TwitterCommand struct {
	interactionCreate messages.InteractionCreate
}

func (d TwitterCommand) Execute() error {

	subcommand := d.interactionCreate.D.Data.Options[0].Name
	var err error
	var message string

	switch subcommand {
	case "add":
		value := d.interactionCreate.D.Data.Options[0].Options[0].Value.(string)
		tag := d.interactionCreate.D.Data.Options[0].Options[1].Value.(string)
		err = twitter.AddRule(value, tag)
		message = "Added rule"
	case "delete":
		id := d.interactionCreate.D.Data.Options[0].Options[0].Value.(string)
		err = twitter.RemoveRule(id)
		message = "Removed rule"
	case "rules":
		message = "Rules"
	default:
		return fmt.Errorf("subcommand not found")
	}

	if err != nil {
		return fmt.Errorf("subcommand %s: %w", subcommand, err)
	}

	rules, err := twitter.GetRules()
	if err != nil {
		return fmt.Errorf("get rules: %w", err)
	}

	embed := builders.CreateEmbed("Rules", "")
	for _, rule := range rules.Data {
		embed.AddField(builders.CreateField(fmt.Sprintf("%s (%s)", rule.Tag, rule.ID), rule.Value))
	}

	interactionCallback := builders.CreateInteractionCallback().
		AddContent(message).
		AddEmbed(embed)

	discord.PostInteractionResponse(d.interactionCreate.D.ID, d.interactionCreate.D.Token, interactionCallback.Get())

	return nil
}

func (d TwitterCommand) Respond() error {
	return nil
}
