package events

import "github.com/bwmarrin/discordgo"

type ReactionHandler struct {
}

func NewReactionHandler() *ReactionHandler {
	return &ReactionHandler{}
}

func (h *ReactionHandler) Handler(s *discordgo.Session, e *discordgo.MessageReactionAdd) {
}
