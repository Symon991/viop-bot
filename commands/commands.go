package commands

import (
	"bot/utils"
	"fmt"
)

func HandleInteraction(message []byte) error {

	command, response, interactionCreate := utils.ReadInteractionCreatePayload(message)

	fmt.Printf("debug interaction: command %s, response %t\n\n", command, response)

	switch command {

	case "hello":
		helloCommand(interactionCreate)

	case "pirate":
		if response {
			return pirateResponse(interactionCreate)
		} else {
			pirateCommand(interactionCreate)
		}

	case "dice":
		diceCommand(interactionCreate)

	case "poll":
		if response {
			pollResponse(interactionCreate)
		} else {
			pollCommmand(interactionCreate)
		}

	case "wolfram":
		wolframCommand(interactionCreate)
	}

	return nil
}
