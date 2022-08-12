package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"bot/utils"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type VideoCommand struct {
	interactionCreate messages.InteractionCreate
}

func (d VideoCommand) Execute() error {

	interactionCallback := utils.CreateInteractionCallback().
		AddContent("Will do, bro.")

	discord.PostInteractionCallback(d.interactionCreate.D.ID, d.interactionCreate.D.Token, interactionCallback.Get())

	parameters := []string{d.interactionCreate.D.Data.Options[0].Value.(string), "--download-sections", d.interactionCreate.D.Data.Options[1].Value.(string), "-v", "-o", "-"}
	cmd := exec.Command("yt-dlp", parameters...)
	cmd2 := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "webm", "pipe:1")

	fmt.Println(cmd.String())

	r, w := io.Pipe()

	cmd2.Stdin, _ = cmd.StdoutPipe()

	cmd2.Stdout = w
	cmd.Stderr = os.Stdout
	cmd2.Stderr = os.Stdout

	var buff bytes.Buffer

	go func() {
		io.Copy(&buff, r)
	}()

	cmd2.Start()

	cmd.Run()

	cmd2.Wait()

	log.Print("finished encoding")

	discord.PostInteractionFile(d.interactionCreate.D.ApplicationID, d.interactionCreate.D.Token, buff.Bytes())

	return nil
}

func (d VideoCommand) Respond() error {
	return nil
}
