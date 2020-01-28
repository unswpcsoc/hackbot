package responder

import (
	"strings"
)

// Ping implements Responder
type Ping struct{}

// OnMessage
func (p *Ping) OnMessage(message string) string {
	if strings.ToLower(message) == "|ping" {
		return "|pong"
	}
	return ""
}
