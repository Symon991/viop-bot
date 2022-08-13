package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"bot/utils"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type VideoCommand struct {
	interactionCreate messages.InteractionCreate
}

func (d VideoCommand) Execute() error {

	ytLinkString := d.interactionCreate.D.Data.Options[0].Value.(string)
	rangeString := d.interactionCreate.D.Data.Options[1].Value.(string)

	interactionCallback := utils.CreateInteractionCallback().
		AddContent("Will do, bro.")

	discord.PostInteractionCallback(d.interactionCreate.D.ID, d.interactionCreate.D.Token, interactionCallback.Get())

	parameters := []string{ytLinkString, "--download-sections", rangeString, "-v", "-o", "-"}
	ytCommand := exec.Command("yt-dlp", parameters...)
	ffmpegCommand := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "webm", "pipe:1")

	r, w := io.Pipe()

	ffmpegCommand.Stdin, _ = ytCommand.StdoutPipe()
	ffmpegCommand.Stdout = w
	ffmpegCommand.Stderr = os.Stdout

	ytCommand.Stderr = os.Stdout

	var buff bytes.Buffer
	go func() {
		io.Copy(&buff, r)
	}()

	ffmpegCommand.Start()
	ytCommand.Run()
	ffmpegCommand.Wait()

	interaction := utils.CreateInteractionCallback().AddContent(fmt.Sprintf("[%s][%s]", ytLinkString, ytCommand))

	discord.PostFollowUpWithFile(
		d.interactionCreate.D.Token,
		buff.Bytes(),
		fmt.Sprintf(
			"%s-%s.webm",
			codeFromLink(ytLinkString),
			strings.ReplaceAll(strings.ReplaceAll(rangeString, "*", ""), ":", "")),
		interaction)

	return nil
}

func codeFromLink(link string) string {

	i := strings.LastIndex(link, "/")
	if i == -1 {
		return link
	}
	return link[i:]
}

func (d VideoCommand) Respond() error {
	return nil
}
