package main

import (
	"regexp"
)

var regexes map[string]*regexp.Regexp

func init() {
	regexes = make(map[string]*regexp.Regexp)

	// Note: Perl syntax: `/regexp/flags`. Go: `(?flags:regexp)`, e.g. `/foo/im` => `(?im:foo)`
	regexes["help"] = regexp.MustCompile(`(?im:^help((?:\s+in)?\s+regex(es)?)?$)`)
	regexes["unset"] = regexp.MustCompile(`(?im:^(.+?)\s+no\s+more(\s+(in\s+(.+?)|personally|here))?(\s+(?:to|about)\s+(.+))?$)`)
	regexes["set"] = regexp.MustCompile(`(?im:^(.+?)(\s+(in\s+(.+?)|personally|here))?\s+(?:to|about)\s+("[^"]+"|.+?)\s+` +
		`((tomorrow\s+at|at|this|next|today\s+at)\s+(.+?)|in\s+(\d+.+?)|every\s+(.+?)\s+(from\s+(.+?)\s+)?(to|until)\s+(.+?))((\.|!)+)?$)`)
	regexes["remind"] = regexp.MustCompile(`(?im:^(.+?)(\s+(in\s+(.+?)|personally|here))?\s+(?:to|about)\s+(.+)$)`)
	regexes["open"] = regexp.MustCompile(`^open(\s+here)?$`)
	regexes["debug"] = regexp.MustCompile(`^debug(\s+channel(\s+(.+?))?)?(\s+for\s+(.+?))?(\s+remove\s+(\d+)(\s+confirm)?)?$`)
}
