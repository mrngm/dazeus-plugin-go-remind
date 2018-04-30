package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/dazeus/dazeus-go"
)

func handleCommand(dz *dazeus.DaZeus, ev dazeus.Event, what string, regex *regexp.Regexp, params string) {
	switch what {
	case "help":
		handleHelp(dz, ev, params)
	case "unset":
		ev.Reply("matched REGEX_UNSET", true)
	case "set":
		ev.Reply("matched REGEX_SET", true)
	case "remind":
		handleRemind(dz, ev, regex, params)
	case "open":
		ev.Reply("matched REGEX_OPEN", true)
	case "debug":
		ev.Reply("matched REGEX_DEBUG", true)
	default:
		handleDefault(dz, ev)
	}
}

func handleDefault(dz *dazeus.DaZeus, ev dazeus.Event) {
	var hl string
	var err error

	if hl, err = dz.HighlightCharacter(); err != nil {
		panic(err)
	}

	ev.Reply("Sorry, I couldn't quite understand that. Use `"+hl+"remind help` for usage instructions.", true)
}

func handleHelp(dz *dazeus.DaZeus, ev dazeus.Event, params string) {
	if strings.Index(strings.ToLower(params), "regex") != -1 {
		// Send help in regex
		ev.Reply("Alright, sending you the help for haxxors.", true)
		for k, v := range regexes {
			dz.Message(ev.Network, ev.Sender, fmt.Sprintf("%s: %s", strings.ToTitle(k), v))
		}
	} else {
		ev.Reply("Alright, sending you the help.", true)
		for _, v := range helpstrs {
			dz.Message(ev.Network, ev.Sender, v)
		}
	}
}

func handleRemind(dz *dazeus.DaZeus, ev dazeus.Event, regex *regexp.Regexp, params string) {
	who := "$who"
	what := "$what"
	where := "$where"
	whomatch := []byte{}
	whatmatch := []byte{}
	wherematch := []byte{}

	// For each match of the regex in the content.
	for _, submatches := range regex.FindAllStringSubmatchIndex(params, -1) {
		// Apply the captured submatches to the template and append the output
		// to the result.
		whomatch = regex.ExpandString(whomatch, who, params, submatches)
		whatmatch = regex.ExpandString(whatmatch, what, params, submatches)
		wherematch = regex.ExpandString(wherematch, where, params, submatches)
	}

	// XXX: check the boolean conditions
	whom := ev.Sender
	if strings.TrimSpace(string(whomatch)) != "me" && strings.TrimSpace(string(whomatch)) != ev.Sender {
		whom = strings.TrimSpace(string(whomatch))
	}

	fmt.Println("  handleRemind: whom:", whom)

	destination := ev.Channel
	if string(wherematch) == "personally" || whom != ev.Sender {
		destination = whom
	} else if string(wherematch) != "here" {
		destination = string(wherematch)
	}

	fmt.Println("  handleRemind: destination: ", destination)

	dz.Message(ev.Network, destination, fmt.Sprintf("%s: %s", whom, string(whatmatch)))
}
