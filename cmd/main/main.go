package main

import (
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
		log.Fatalln("Missing Discord API Key: Set env var $DISCORD")
	}

	var err error
	dgo, err = discordgo.New("Bot " + key)
	if err != nil {
		log.Fatalln(err)
	}

	err = dgo.Open()
	if err != nil {
		log.Fatalln(err)
	}

	dgo.SyncEvents = false
	log.Println("Logged in as: ", dgo.State.User.ID)
}

func main() {
	defer dgo.Close()

	dgo.UpdateStatus(0, "hacking...")

	// handle responder commands
	dgo.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot {
			return
		}

		responses := responder.Notify(m.Message.Content)
		for _, res := range responses {
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
