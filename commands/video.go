package commands

import (
	"bot/discord"
	"bot/discord/messages"
	"bot/utils"
	"fmt"
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

	parameters := []string{ytLinkString, "-f", "best[height=720]", "--downloader-args", "ffmpeg_o:-f webm", "--download-sections", rangeString, "--force-keyframes-at-cuts", "-v", "-o", "-"}
	ytCommand := exec.Command("yt-dlp", parameters...)

	ytCommand.Stderr = os.Stdout
	outBytes, err := ytCommand.Output()
	if err != nil {
		return fmt.Errorf("run yt-dlp: %w", err)
	}

	interaction := utils.CreateInteractionCallback().AddContent("[Youtube]")

	discord.PostFollowUpWithFile(
		d.interactionCreate.D.Token,
		outBytes,
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
	index := strings.LastIndex(link, "/")
	if index > -1 {
		return link[index:]
	}
	return ""
}

func (d VideoCommand) Respond() error {
	return nil
}
