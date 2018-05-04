package main

import (
	"regexp"
)

var regexes map[string]*regexp.Regexp
var helpstrs []string

func init() {
	regexes = make(map[string]*regexp.Regexp)

	// Note: Perl syntax: `/regexp/flags`. Go: `(?flags:regexp)`, e.g. `/foo/im` => `(?im:foo)`
	regexes["help"] = regexp.MustCompile(`(?im:^help((?:\s+in)?\s+regex(es)?)?$)`)
	regexes["unset"] = regexp.MustCompile(`(?im:^(.+?)\s+no\s+more(\s+(in\s+(.+?)|personally|here))?(\s+(?:to|about)\s+(.+))?$)`)
	regexes["set"] = regexp.MustCompile(`(?im:^(?P<who>.+?)(\s+(in\s+(?P<where>(.+?)|personally|here)))?\s+(?:to|about)\s+` +
		`(?P<what>("[^"]+"|.+?))\s+` +
		`((?P<datespec>(tomorrow\s+at|at|this|next|today\s+at)\s+(.+?))|` +
		`in\s+(?P<durationspec>\d+.+?)|` +
		`every\s+(?P<repeatspec>.+?)\s+(from\s+(?P<repeatfrom>.+?)\s+)?(to|until)\s+(?P<repeatuntil>.+?))((\.|!)+)?$)`)
	regexes["remind"] = regexp.MustCompile(`(?im:^(?P<who>.+?)(\s+(in\s+(?P<where>(.+?)|personally|here)))?\s+(?:to|about)\s+(?P<what>.+)$)`)
	regexes["open"] = regexp.MustCompile(`^open(\s+here)?$`)
	regexes["debug"] = regexp.MustCompile(`^debug(\s+channel(\s+(.+?))?)?(\s+for\s+(.+?))?(\s+remove\s+(\d+)(\s+confirm)?)?$`)

	helpstrs = []string{
		"The " + myCommand + " command is used to create automatic reminders. They should not be regarded as reliable in any way.",
		"To create a new reminder use `" + myCommand + " <who> to <message>`. This will however send the reminder right away.",
		"To create a reminder at a specific time you can use a number of formats:",
		"- You can create a reminder with an interval from the current time: `" + myCommand + " John to wash the dishes in 20 minutes` (the keyword here is 'in').",
		"- You can create a reminder at a specific time using: `" + myCommand + " me to wash the dishes at 20:00` (the keyword here is 'at').",
		"- Some alternative formats include:",
		"  - `" + myCommand + " me to wash the dishes tomorrow at 13:00`",
		"  - `" + myCommand + " Jane to wash the dishes next thursday`",
		"- You can also set a repeated (some limits are applied) reminder using: `" + myCommand + " me to wash the dishes every 10 minutes until 20:00`.",
		"- Repeated reminders can also be run in between two times: `" + myCommand + " John to wash the dishes every day from monday to friday`.",
		"- You can also send a reminder to a specific channel: `" + myCommand + " me in #example to wash the dishes next friday`.",
		"- You can also send a reminder in a query chat: `" + myCommand + " me personally to wash the dishes in 1 week`.",
		"You can cancel all your reminders in the current channel by using `" + myCommand + " me no more` or `" + myCommand + " John no more` for another person.",
		"You can also remove only reminders with a specific message using `" + myCommand + " me no more to wash the dishes`",
	}
}
