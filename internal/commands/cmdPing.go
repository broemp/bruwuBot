package commands

import "github.com/broemp/bruwuBot/internal/commandHandler"

type CmdPing struct{}

func (c *CmdPing) Invokes() []string {
	return []string{"ping", "p"}
}

func (c *CmdPing) Description() string {
	return "Pong!"
}

func (c *CmdPing) AdminRequired() bool {
	return true
}

func (c *CmdPing) Exec(ctx *commandHandler.Context) (err error) {
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Pong!")
	return
}
