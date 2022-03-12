package events

import (
	"github.com/bwmarrin/discordgo"
)

type JoinLeaveHandler struct {
}

func NewJoinLeaveHandler() *JoinLeaveHandler {
	return &JoinLeaveHandler{}
}

func (h *JoinLeaveHandler) HandlerJoin(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
}

func (h *JoinLeaveHandler) HandlerLeave(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
}
