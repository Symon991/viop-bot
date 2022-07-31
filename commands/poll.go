package commands

import (
	"bot/cache"
	"bot/discord"
	"bot/discord/messages"
	"bot/utils"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Poll struct {
	Question string
	Options  []Option
	Users    map[string]int
}

type Option struct {
	Description string
	Votes       int
}

type PollCommand struct {
	interactionCreate messages.InteractionCreate
}

func (p *PollCommand) Execute() error {

	question := p.interactionCreate.D.Data.Options[0].Value.(string)
	optionsString := p.interactionCreate.D.Data.Options[1].Value.(string)
	optionsArray := strings.Split(optionsString, ";")

	var options []Option

	for _, option := range optionsArray {

		options = append(options, Option{
			Description: option,
			Votes:       0,
		})
	}

	poll := Poll{
		Question: question,
		Options:  options,
		Users:    make(map[string]int),
	}

	marshalledResult, err := json.Marshal(poll)
	if err != nil {
		return fmt.Errorf("poll command: %s", err)
	}
	cache.Set(string(p.interactionCreate.D.ID), string(marshalledResult), 0)

	interactionCallback := messageFromPoll(poll, false, true)

	discord.PostInteractionCallback(p.interactionCreate.D.ID, p.interactionCreate.D.Token, interactionCallback.Get())

	return nil
}

func pivotOptionUsers(poll Poll) *map[int][]string {

	pivot := make(map[int][]string)

	for user, choice := range poll.Users {
		pivot[choice] = append(pivot[choice], user)
	}

	return &pivot
}

func messageFromPoll(poll Poll, edit bool, showUsers bool) *utils.InteractionCallbackBuilder {

	embed := utils.CreateEmbed("Poll", poll.Question)
	selectComponent := utils.CreateSelectComponent("Your choice:", "test")
	pivot := pivotOptionUsers(poll)

	for i, option := range poll.Options {
		if showUsers {
			embed.AddField(utils.CreateField(option.Description,
				fmt.Sprintf("%d (%s)",
					option.Votes,
					strings.Join((*pivot)[i], ","))))
		} else {
			embed.AddField(utils.CreateField(option.Description, fmt.Sprint(option.Votes)))
		}
		selectComponent.AddOption(utils.CreateOption(option.Description, "", fmt.Sprint(i)))
	}

	interactionCallbackBuilder := utils.CreateInteractionCallbackEdit(edit).
		AddEmbed(embed).
		AddActionRowComponent(
			utils.CreateActionRowComponent().
				AddComponent(selectComponent))

	return interactionCallbackBuilder
}

func followUpFromPoll(poll Poll) *utils.InteractionCallbackBuilder {

	selectComponent := utils.CreateSelectComponent("Your choice:", "test")

	for i, option := range poll.Options {
		selectComponent.AddOption(utils.CreateOption(option.Description, "", fmt.Sprint(i)))
	}

	interactionCallbackBuilder := utils.CreateInteractionCallbackEdit(true).
		AddActionRowComponent(
			utils.CreateActionRowComponent().
				AddComponent(selectComponent))

	return interactionCallbackBuilder
}

func (p *PollCommand) Respond() error {

	value, err := cache.Get(string(p.interactionCreate.D.Message.Interaction.ID))
	if err != nil {
		return fmt.Errorf("get value from cache: %s", err)
	}

	var poll Poll
	err = json.Unmarshal([]byte(value), &poll)
	if err != nil {
		return fmt.Errorf("unmarshal value from cache: %w", err)
	}

	fmt.Print(poll)
	fmt.Print(poll.Users)

	if _, exist := poll.Users[p.interactionCreate.D.Member.User.Username]; !exist {

		selectedValue, err := strconv.ParseInt(p.interactionCreate.D.Data.Values[0], 10, 64)
		if err != nil {
			return fmt.Errorf("parse selected value: %w", err)
		}

		poll.Options[selectedValue].Votes += 1
		poll.Users[p.interactionCreate.D.Member.User.Username] = int(selectedValue)

		fmt.Print(poll.Users)

		marshalledResult, err := json.Marshal(poll)
		if err != nil {
			return fmt.Errorf("poll command: %s", err)
		}
		cache.Set(string(p.interactionCreate.D.Message.Interaction.ID), string(marshalledResult), 0)
	}

	interactionCallback := messageFromPoll(poll, true, true)

	discord.PostInteractionCallback(p.interactionCreate.D.ID, p.interactionCreate.D.Token, interactionCallback.Get())

	return nil
}
