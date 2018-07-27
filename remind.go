package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/dazeus/dazeus-go"
)

var myCommand string

func handleRemindMessage(dz *dazeus.DaZeus, ev dazeus.Event) {
	if len(ev.Params) > 1 {
		// The message sits in evt.Params[0]
		handled := false
	regexLoop:
		for k, v := range regexes {
			if v.MatchString(ev.Params[0]) {
				handleCommand(dz, ev, k, v, ev.Params[0])
				handled = true
				break regexLoop
			}
		}
		if !handled {
			handleCommand(dz, ev, "", nil, ev.Params[0])
		}
	}
}

func main() {
	connStr := "unix:/tmp/dazeus.sock"
	if len(os.Args) > 1 {
		connStr = os.Args[1]
	}
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Paniek! %v\n", p)
			debug.PrintStack()
		}
	}()

	dz, err := dazeus.ConnectWithLoggingToStdErr(connStr)
	if err != nil {
		panic(err)
	}

	if _, hlerr := dz.HighlightCharacter(); hlerr != nil {
		panic(hlerr)
	}

	_, err = dz.SubscribeCommand("remind", dazeus.NewUniversalScope(), func(ev dazeus.Event) {
		handleRemindMessage(dz, ev)
	})
	if err != nil {
		panic(err)
	}

	listenerr := dz.Listen()
	panic(listenerr)
}
