package responder

import (
	"strings"
)

// Ping implements Responder
type Ping struct{}

func (p *Ping) OnMessage(message string) string {
	if strings.ToLower(message) == "|ping" {
		return "|pong"
	}
	return ""
}

func (p *Ping) Desc() string { return "Pong!" }

func (p *Ping) Help() string { return "`|ping` returns a `|pong`" }
