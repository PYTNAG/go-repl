package commands

import "strings"

type Description struct {
	name    string
	aliases []string
}

func NewDescription(name string, aliases ...string) *Description {
	description := new(Description)

	description.name = name
	description.aliases = aliases

	return description
}

func (d *Description) Name() string {
	return d.name
}

func (d *Description) Aliases() []string {
	return d.aliases
}

func (d *Description) ToStringLine() string {
	description := d.name

	if d.aliases != nil && len(d.aliases) != 0 {
		description += " ("
		description += strings.Join(d.aliases, ", ")
		description += ")"
	}

	return description
}
