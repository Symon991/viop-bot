package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"fmt"
	"os/exec"
)

type VideoCommand struct {
	interactionCreate messages.InteractionCreate
}

func (d VideoCommand) Execute() error {

	parameters := []string{"https://www.youtube.com/watch?v=gW-N7AHl7dI", "--download-sections", "*00:10-00:15", "-v", "-o", "-"}
	cmd := exec.Command("yt-dlp.exe", parameters...)
	fmt.Println(cmd.String())

	bytes, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	discord.PostInteractionFile(d.interactionCreate.D.ID, d.interactionCreate.D.Token, bytes)

	return nil
}

func (d VideoCommand) Respond() error {
	return nil
}
