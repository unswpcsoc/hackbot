package responder

import (
	"strings"
)

type Help struct{}

func (h *Help) OnMessage(message string) string {
	if !strings.HasPrefix(strings.Trim(strings.ToLower(message), " "), "|help") {
		return ""
	}
	argv := strings.Split(strings.Trim(message, " "), " ")

	out := "Help is here!\n"
	if len(argv) > 1 && (strings.ToLower(argv[1]) == "-v" || strings.ToLower(argv[1]) == "verbose") {
		for _, res := range responders {
			out += res.Help() + "\n\t" + res.Desc() + "\n\n"
		}
	} else {
		for _, res := range responders {
			out += res.Help() + "\n"
		}
	}

	return out
}

func (h *Help) Desc() string { return "helps describe other commands" }

func (h *Help) Help() string { return "`|help` lists help text of all commands" }
