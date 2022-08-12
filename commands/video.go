package commands

import (
	"bot/discord/messages"
	"fmt"
	"os"
	"os/exec"
)

type VideoCommand struct {
	interactionCreate messages.InteractionCreate
}

func (d VideoCommand) Execute() error {

	cmd := exec.Command("ffmpeg")
	fmt.Println(cmd.String())

	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout

	fmt.Println(cmd.Output())

	return nil
}

func (d VideoCommand) Respond() error {
	return nil
}
