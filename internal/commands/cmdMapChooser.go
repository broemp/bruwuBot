package commands

import (
	"strings"
	"time"

	"github.com/broemp/bruwuBot/internal/commandHandler"
	"github.com/bwmarrin/discordgo"
)

type CmdMapChooser struct{}

func (c *CmdMapChooser) Invokes() []string {
	return []string{"mapchooser", "mc"}
}

func (c *CmdMapChooser) Description() string {
	return "Helps Choosing a Map for CSGO!"
}

func (c *CmdMapChooser) AdminRequired() bool {
	return false
}

func (c *CmdMapChooser) Exec(ctx *commandHandler.Context) (err error) {

	var csgoMapsAll = []string{"Ancient", "Dust II", "Inferno", "Mirage", "Nuke", "Overpass", "Vertigo", "Agency", "Cache", "Climb", "Iris", "Train", "Office"}
	var csgoMapsBomb = []string{"Ancient", "Dust II", "Inferno", "Mirage", "Nuke", "Overpass", "Vertigo", "Cache", "Iris", "Train"}
	var csgoMapsActive = []string{"Ancient", "Dust II", "Inferno", "Mirage", "Nuke", "Overpass", "Vertigo"}
	var csgoMapsReserve = []string{"Agency", "Cache", "Climb", "Iris", "Train", "Office"}
	var csgoMapsHostage = []string{"Agency", "Climb", "Office"}

	var message discordgo.MessageEmbed

	// Build Message
	message.Type = discordgo.EmbedTypeRich
	message.Color = 1
	message.Title = "CSGO Map Chooser"
	message.Description = "Helps you decide what Map to Play!"

	var csgoMapsFieldAll discordgo.MessageEmbedField
	csgoMapsFieldAll.Name = "1️⃣ All Maps"
	csgoMapsFieldAll.Value = strings.Join(csgoMapsAll, ", ")
	csgoMapsFieldAll.Inline = false

	var csgoMapsFieldBomb discordgo.MessageEmbedField
	csgoMapsFieldBomb.Name = "2️⃣ Bomb Maps"
	csgoMapsFieldBomb.Value = strings.Join(csgoMapsBomb, ", ")
	csgoMapsFieldBomb.Inline = false

	var csgoMapsFieldActive discordgo.MessageEmbedField
	csgoMapsFieldActive.Name = "3️⃣ Active Maps"
	csgoMapsFieldActive.Value = strings.Join(csgoMapsActive, ", ")
	csgoMapsFieldActive.Inline = false

	var csgoMapsFieldReserve discordgo.MessageEmbedField
	csgoMapsFieldReserve.Name = "4️⃣ Reserve Maps"
	csgoMapsFieldReserve.Value = strings.Join(csgoMapsReserve, ", ")
	csgoMapsFieldReserve.Inline = false

	var csgoMapsFieldHostage discordgo.MessageEmbedField
	csgoMapsFieldHostage.Name = "5️⃣ Hostage Maps"
	csgoMapsFieldHostage.Value = strings.Join(csgoMapsHostage, ", ")
	csgoMapsFieldHostage.Inline = false

	var csgoMapsFieldCustom discordgo.MessageEmbedField
	csgoMapsFieldCustom.Name = "6️⃣ Custom Maps"
	csgoMapsFieldCustom.Value = "Create a Custom Map Pool!"
	csgoMapsFieldCustom.Inline = false

	message.Fields = append(message.Fields, &csgoMapsFieldAll, &csgoMapsFieldBomb, &csgoMapsFieldActive, &csgoMapsFieldReserve, &csgoMapsFieldHostage, &csgoMapsFieldCustom)

	sendMessage, err := ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, &message)
	if err != nil {
		return
	}

	var emojiList = []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣"}

	for _, emoji := range emojiList {
		err = ctx.Session.MessageReactionAdd(sendMessage.ChannelID, sendMessage.ID, emoji)
		if err != nil {
			return
		}
	}

	time.Sleep(5 * time.Second)

	userList := make(map[string][]*discordgo.User)

	for _, emoji := range emojiList {
		users, err := ctx.Session.MessageReactions(ctx.Message.ChannelID, sendMessage.ID, emoji, 100, "", "")
		if err != nil {
			return err
		}

		userList[emoji] = users

	}

	winningUserList := userList[emojiList[0]]
	var winningEmoji string

	// Need to address Tie
	for emoji, user := range userList {
		if len(user) > len(winningUserList) {
			winningUserList = user
			winningEmoji = emoji
		}
	}

	if winningEmoji == "" {
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Please Vote Next Time!")
		if err != nil {
			return
		}
	}

	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "The Winner is "+winningEmoji)
	if err != nil {
		return
	}

	return
}
