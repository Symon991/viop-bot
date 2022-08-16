package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"bot/utils"
	"bytes"
	"fmt"
	"io"
	"net/url"
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

	parameters := []string{ytLinkString, "--download-sections", rangeString, "--force-keyframes-at-cuts", "-v", "-o", "-"}
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

	interaction := utils.CreateInteractionCallback().AddContent("[Youtube]")

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

	parsed, err := url.Parse(link)
	if err == nil {
		return parsed.Query().Get("v")
	}
	index := strings.Index(link, "/")
	if index > -1 {
		return link[index:]
	}
	return ""
}

func (d VideoCommand) Respond() error {
	return nil
}
