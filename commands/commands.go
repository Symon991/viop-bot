package commands

import (
	"bot/utils"
	"fmt"
)

type Command interface {
	Execute() error
	Respond() error
}

func HandleInteraction(message []byte) error {

	var c Command
	var err error
	command, response, interactionCreate := utils.ReadInteractionCreatePayload(message)
	fmt.Printf("debug interaction: command %s, response %t\n\n", command, response)

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
