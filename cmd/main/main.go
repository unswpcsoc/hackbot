package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/unswpcsoc/hackbot/responder"

	"github.com/bwmarrin/discordgo"
)

var (
	dgo *discordgo.Session
)

// init for discord API key
func init() {
	key, ok := os.LookupEnv("DISCORD")
	if !ok {
		errs.Fatalln("Missing Discord API Key: Set env var $DISCORD")
	}

	var err error
	dgo, err = discordgo.New("Bot " + key)
	if err != nil {
		errs.Fatalln(err)
	}

	err = dgo.Open()
	if err != nil {
		errs.Fatalln(err)
	}

	dgo.SyncEvents = false
	log.Println("Logged in as: ", dgo.State.User.ID)
}

func main() {
	defer dgo.Close()
	defer commands.DBClose()

	dgo.UpdateStatus(0, commands.Prefix+handlers.HelpAlias)

	// init loggers
	handlers.InitLogs(dgo)

	// handle responder commands
	dgo.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		responses := responder.Notify(m.Message.Content)
		for res := range responses {
			s.ChannelMessageSend(m.ChannelID, res)
		}
	})

	// keep alive
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	sig := <-sc

	log.Println("Received Signal: " + sig.String())
	log.Println("Bye!")
}
