package main

import (
	"strings"

	"github.com/goexl/git"
)

func (p *plugin) netrc() (undo bool, err error) {
	if undo = `` == strings.TrimSpace(p.Username) || `` == strings.TrimSpace(p.Password); undo {
		return
	}
	err = git.Auth(git.Netrc(p.Remote, p.Username, p.Password))

	return
}
