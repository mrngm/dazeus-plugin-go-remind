package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dazeus/dazeus-go"
)

var myCommand string

func handlePrivmsg(dz *dazeus.DaZeus, ev dazeus.Event) {
	if len(ev.Params) > 1 {
		// The message sits in evt.Params[0]
		if idx := strings.Index(ev.Params[0], myCommand); idx > -1 {
			params := strings.Replace(ev.Params[0], myCommand+" ", "", 1)
			handled := false

		regexLoop:
			for k, v := range regexes {
				if v.MatchString(params) {
					handleCommand(dz, ev, k, v, params)
					handled = true
					break regexLoop
				}
			}
			if !handled {
				handleCommand(dz, ev, "", nil, params)
			}
		}
	}
}

func main() {
	connStr := "unix:/tmp/dazeus.sock"
	if len(os.Args) > 1 {
		connStr = os.Args[1]
	}

	dz, err := dazeus.ConnectWithLoggingToStdErr(connStr)
	if err != nil {
		panic(err)
	}

	if hl, hlerr := dz.HighlightCharacter(); hlerr != nil {
		panic(hlerr)
	} else {
		myCommand = hl + "remind"
	}

	_, err = dz.Subscribe(dazeus.EventPrivMsg, func(ev dazeus.Event) {
		fmt.Printf("Got PRIVMSG: %+v\n", ev)
		handlePrivmsg(dz, ev)
	})
	if err != nil {
		panic(err)
	}

	dz.Listen()
}
