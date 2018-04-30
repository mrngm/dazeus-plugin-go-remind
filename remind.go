package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dazeus/dazeus-go"
)

func handlePrivmsg(dz *dazeus.DaZeus, ev dazeus.Event) {
	if len(ev.Params) > 1 {
		// The message sits in evt.Params[0]
		if idx := strings.Index(ev.Params[0], "!remind"); idx > -1 {
			params := strings.Replace(ev.Params[0], "!remind ", "", 1)
			handled := false

		regexLoop:
			for k, v := range regexes {
				if v.MatchString(params) {
					handleRemind(dz, ev, k, params)
					handled = true
					break regexLoop
				}
			}
			if !handled {
				handleRemind(dz, ev, "", params)
			}
		}
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

	_, err = dz.Subscribe(dazeus.EventPrivMsg, func(ev dazeus.Event) {
		fmt.Printf("Got PRIVMSG: %+v\n", ev)
		for i, v := range ev.Params {
			fmt.Printf("-> %d: %s\n", i, v)
		}
		handlePrivmsg(dz, ev)
	})
	if err != nil {
		panic(err)
	}

	dz.Listen()
}
