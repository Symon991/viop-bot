package commands

import (
	"bot/discord/messages"
	"fmt"
	"os/exec"
)

type VideoCommand struct {
	interactionCreate messages.InteractionCreate
}

func (d VideoCommand) Execute() error {

	cmd := exec.Command("ffmpeg")
	fmt.Println(cmd.String())
	fmt.Println(cmd.CombinedOutput())

	return nil
}

func (d VideoCommand) Respond() error {
	return nil
}
