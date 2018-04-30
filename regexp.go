package main

import (
	"regexp"
)

// Note: Perl syntax: `/regexp/flags`. Go: `(?flags:regexp)`, e.g. `/foo/im` => `(?im:foo)`
var REGEX_HELP = regexp.MustCompile(`(?im:^help((?:\s+in)?\s+regex(es)?)?$)`)
var REGEX_UNSET = regexp.MustCompile(`(?im:^(.+?)\s+no\s+more(\s+(in\s+(.+?)|personally|here))?(\s+(?:to|about)\s+(.+))?$)`)
var REGEX_SET = regexp.MustCompile(`(?im:^(.+?)(\s+(in\s+(.+?)|personally|here))?\s+(?:to|about)\s+("[^"]+"|.+?)\s+` +
	`((tomorrow\s+at|at|this|next|today\s+at)\s+(.+?)|in\s+(\d+.+?)|every\s+(.+?)\s+(from\s+(.+?)\s+)?(to|until)\s+(.+?))((\.|!)+)?$)`)
var REGEX_REMIND = regexp.MustCompile(`(?im:^(.+?)(\s+(in\s+(.+?)|personally|here))?\s+(?:to|about)\s+(.+)$)`)
var REGEX_OPEN = regexp.MustCompile(`^open(\s+here)?$`)
var REGEX_DEBUG = regexp.MustCompile(`^debug(\s+channel(\s+(.+?))?)?(\s+for\s+(.+?))?(\s+remove\s+(\d+)(\s+confirm)?)?$`)
