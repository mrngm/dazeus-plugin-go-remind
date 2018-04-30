package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dazeus/dazeus-go"
)

func sendHelp(ev dazeus.Event, short bool) {
	if short {
		ev.Reply("Sorry, I couldn't quite understand that. Use `!remind help` for usage instructions.", true)
	} else {
		ev.Reply("Well, well well...", true)
	}
}

func handleRemind(params string, ev dazeus.Event) {
	fmt.Printf("handleRemind: got params: '%v'\n", params)
	switch {
	case params == "":
		sendHelp(ev, true)
	case REGEX_HELP.MatchString(params):
		ev.Reply("matched REGEX_HELP", true)
	case REGEX_UNSET.MatchString(params):
		ev.Reply("matched REGEX_UNSET", true)
	case REGEX_SET.MatchString(params):
		ev.Reply("matched REGEX_SET", true)
	case REGEX_REMIND.MatchString(params):
		ev.Reply("matched REGEX_REMIND", true)
	case REGEX_OPEN.MatchString(params):
		ev.Reply("matched REGEX_OPEN", true)
	case REGEX_DEBUG.MatchString(params):
		ev.Reply("matched REGEX_DEBUG", true)
	default:
		ev.Reply("none matched", true)
	}
}

func main() {
	connStr := "unix:/tmp/dazeus.sock"
	if len(os.Args) > 1 {
		connStr = os.Args[1]
	}

	dz, err := dazeus.Connect(connStr)
	if err != nil {
		panic(err)
	}

	_, err = dz.Subscribe(dazeus.EventPrivMsg, func(evt dazeus.Event) {
		fmt.Printf("Got PRIVMSG: %+v\n", evt)
		if len(evt.Params) > 1 {
			// The message sits in evt.Params[0]
			splitted := strings.Split(evt.Params[0], " ")
			if len(splitted) > 0 {
				if splitted[0] == "!remind" {
					handleRemind(strings.Join(splitted[1:], " "), evt)
				}
			}
		}
	})
	if err != nil {
		panic(err)
	}

	dz.Listen()
}
