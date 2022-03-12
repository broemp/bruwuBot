package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/broemp/bruwuBot/internal/commandHandler"
	"github.com/broemp/bruwuBot/internal/commands"
	"github.com/broemp/bruwuBot/internal/config"
	"github.com/broemp/bruwuBot/internal/events"
	"github.com/bwmarrin/discordgo"
)

const defaultConfigFile = "./config/config.json"

var flagConfig = flag.String("c", defaultConfigFile, "Location of config file.")

func main() {

	flag.Parse()
	cfg, err := config.ParseConfig(*flagConfig)

	if err != nil {
		panic(err)
	}

	s, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		panic(err)
	}

	s.Identify.Intents = discordgo.MakeIntent(
		discordgo.IntentsGuildMembers |
			discordgo.IntentGuildMessages |
			discordgo.IntentGuildMessageReactions)

	registerEvents(s)
	RegisterCommands(s, cfg)

	if err = s.Open(); err != nil {
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	s.Close()
}

func registerEvents(s *discordgo.Session) {
	joinLeaveHandler := events.NewJoinLeaveHandler()
	s.AddHandler(joinLeaveHandler.HandlerJoin)
	s.AddHandler(joinLeaveHandler.HandlerLeave)

	s.AddHandler(events.NewReadyHandler().Handler)
	s.AddHandler(events.NewMessageHandler().Handler)
	s.AddHandler(events.NewReactionHandler().Handler)
}

func RegisterCommands(s *discordgo.Session, cfg *config.Config) {
	cmdHandler := commandHandler.NewCommandHandler(cfg.Prefix)
	cmdHandler.OnError = func(err error, ctx *commandHandler.Context) {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			fmt.Sprintf("Command Execution Failed: %s", err.Error()))
	}

	// Command List
	cmdHandler.RegisterCommand(&commands.CmdPing{})
	cmdHandler.RegisterCommand(&commands.CmdMapChooser{})
	cmdHandler.RegisterMiddleware(&commandHandler.MwPermissions{})

	s.AddHandler(cmdHandler.HandleMessage)
}
