package commands

import (
	"bot/utils"
	"fmt"
)

func HandleInteraction(message []byte) error {

	command, response, interactionCreatePayload := utils.ReadInteractionCreatePayload(message)

	fmt.Printf("debug interaction: command %s, response %t\n\n", command, response)

	switch command {

	case "hello":
		helloCommand(interactionCreatePayload)

	case "pirate":
		if response {
			return pirateResponse(interactionCreatePayload)
		} else {
			pirateCommand(interactionCreatePayload)
		}

	case "dice":
		diceCommand(interactionCreatePayload)

	}

	return nil
}
