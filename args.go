package utils

import (
	"fmt"
	"strings"
)

var validFlags = map[rune]bool{
	'l': true,
	'r': true,
	'R': true,
	'a': true,
	't': true,
}

// options handles single-character flags.
type options struct {
	flags map[rune]bool
}

// Parse parses and validates single-character flags.
func (o *options) Parse(args []string) error {
	fmt.Println(args)
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") && arg == "--l" {
			arg = arg[2:]
		} else if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") && len(arg) >= 2 {
			arg = arg[1:]
		} else {
			return fmt.Errorf("invalid argument: %s", arg)
		}
		for _, r := range arg {
			if !validFlags[r] {
				return fmt.Errorf("invalid flag: %c", r)
			}
			o.flags[r] = true
		}
	}
	return nil
}

func (c *options) IsFlagSet(flag rune) bool {
	return c.flags[flag]
}

// NewCustomFlagSet creates a new instance of CustomFlagSet.
func NewOptions() *options {
	return &options{flags: make(map[rune]bool)}
}
