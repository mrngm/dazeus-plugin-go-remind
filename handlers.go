package main

import (
	"fmt"
	"strings"

	"github.com/dazeus/dazeus-go"
)

func handleRemind(dz *dazeus.DaZeus, ev dazeus.Event, what, params string) {
	switch what {
	case "help":
		handleHelp(dz, ev, params)
	case "unset":
		ev.Reply("matched REGEX_UNSET", true)
	case "set":
		ev.Reply("matched REGEX_SET", true)
	case "remind":
		ev.Reply("matched REGEX_REMIND", true)
	case "open":
		ev.Reply("matched REGEX_OPEN", true)
	case "debug":
		ev.Reply("matched REGEX_DEBUG", true)
	default:
		handleDefault(dz, ev)
	}
}

func handleDefault(dz *dazeus.DaZeus, ev dazeus.Event) {
	ev.Reply("Sorry, I couldn't quite understand that. Use `!remind help` for usage instructions.", true)
}

func handleHelp(dz *dazeus.DaZeus, ev dazeus.Event, params string) {
	fmt.Printf("  idx: %d, params: %s\n", strings.Index(strings.ToLower(params), "regex"), params)
	if strings.Index(strings.ToLower(params), "regex") != -1 {
		// Send help in regex
		ev.Reply("Alright, sending you the help for haxxors.", true)
		for k, v := range regexes {
			dz.Message(ev.Network, ev.Sender, fmt.Sprintf("%s: %s", strings.ToTitle(k), v))
		}
	} else {
		ev.Reply("Alright, sending you the help.", true)
	}
}
