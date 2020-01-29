// responder defines an interface for you to build quick and simple bot commands that just respond to messages
package responder

// Responder is a simple interface for simple commands which just respond to a message
type Responder interface {
	OnMessage(message string) string // OnMessage takes a message as a string and responds with some string

	Help() string // Help should provide brief usage instructions
	Desc() string // Desc should provide a longer description of the command, maybe throw some memes in there
}

var responders = []Responder{
	&Ping{},
	&Help{},
}

// Notify notifies all Responders and returns a slice of all the responses
// this is kind of cooked tbh but whatever lmao
func Notify(message string) []string {
	responses := []string{}
	for _, res := range responders {
		resp := res.OnMessage(message)
		if resp != "" {
			responses = append(responses, resp)
		}
	}
	return responses
}
