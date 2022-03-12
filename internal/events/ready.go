package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type ReadyHandler struct {
}

func NewReadyHandler() *ReadyHandler {
	return &ReadyHandler{}
}

func (h *ReadyHandler) Handler(s *discordgo.Session, e *discordgo.Ready) {
	fmt.Println("Bot is ready!")
	user, err := s.User("@me")
	if err != nil {
		return
	}
	fmt.Printf("*********************************************************************************************************************************************************************\n\n")
	fmt.Printf("Logged in as %s\n", e.User.String())
	fmt.Println("https://discord.com/api/oauth2/authorize?client_id=" + user.ID + "&permissions=8&scope=bot%20applications.commands")
	fmt.Printf("\n*********************************************************************************************************************************************************************\n\n")
	fmt.Println("Bot is now running. Press CTRL-C to exit...")
}
