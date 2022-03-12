package commandHandler

import (
	"github.com/bwmarrin/discordgo"
)

type MwPermissions struct{}

func (mw *MwPermissions) Exec(ctx *Context, cmd Command) (next bool, err error) {
	if !cmd.AdminRequired() {
		next = true
		return
	}

	defer func() {
		if !next && err == nil {
			_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
				"You don't have permission to use this command!")
		}
	}()

	guild, err := ctx.Session.Guild(ctx.Message.GuildID)
	if err != nil {
		return
	}

	if guild.OwnerID == ctx.Message.Author.ID {
		next = true
		return
	}

	roleMap := make(map[string]*discordgo.Role)
	for _, role := range guild.Roles {
		roleMap[role.ID] = role
	}

	for _, roleID := range ctx.Message.Member.Roles {
		if role, roleExistsInMap := roleMap[roleID]; roleExistsInMap && role.Permissions&discordgo.PermissionAdministrator > 0 {
			next = true
			break
		}
	}

	return
}
