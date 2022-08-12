package commands

import (
	"bot/utils"
	"fmt"
	"log"
)

type Command interface {
	Execute() error
	Respond() error
}

func HandleInteraction(message []byte) error {

	var c Command
	var err error
	command, response, interactionCreate := utils.ReadInteractionCreatePayload(message)
	log.Printf("debug interaction: command %s, response %t\n\n", command, response)

	switch command {
	case "hello":
		c = &HelloCommand{interactionCreate}
	case "pirate":
		c = &PirateCommand{interactionCreate}
	case "dice":
		c = &DiceCommand{interactionCreate}
	case "wolfram":
		c = &WolframCommand{interactionCreate}
	case "poll":
		c = &PollCommand{interactionCreate}
	case "twitter":
		c = &TwitterCommand{interactionCreate}
	case "countdown":
		c = &CountdownCommand{interactionCreate}
	case "mudtime":
		c = &MudTimeCommand{interactionCreate}
	case "video":
		c = &VideoCommand{interactionCreate}
	default:
		return fmt.Errorf("command not found")
	}

	if response {
		err = c.Respond()
	} else {
		err = c.Execute()
	}

	if err != nil {
		return fmt.Errorf("command %s: %w", command, err)
	}
	return nil
}
