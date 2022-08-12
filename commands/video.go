package commands

import (
	"bot/discord"
	"bot/discord/messages"
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

	parameters := []string{"https://www.youtube.com/watch?v=gW-N7AHl7dI", "--download-sections", "*00:10-00:15", "-v", "-o", "-"}
	cmd := exec.Command("yt-dlp", parameters...)
	cmd2 := exec.Command("ffmpeg.exe", "-i", "pipe:0", "-f", "webm", "pipe:1")

	fmt.Println(cmd.String())

	r, w := io.Pipe()

	cmd2.Stdin, _ = cmd.StdoutPipe()

	cmd2.Stdout = w
	cmd2.Stderr = os.Stdout

	var buff bytes.Buffer

	go func() {
		io.Copy(&buff, r)
	}()

	cmd2.Start()

	cmd.Run()

	cmd2.Wait()

	log.Print("finished encoding")

	discord.PostInteractionFile(d.interactionCreate.D.ID, d.interactionCreate.D.Token, buff.Bytes())

	return nil
}

func (d VideoCommand) Respond() error {
	return nil
}
