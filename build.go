package main

import (
	"github.com/dronestock/drone"
)

func (p *plugin) build() (undo bool, err error) {
	args := []any{
		"--preset",
		"angular",
		"---infile",
		p.Output,
		"--same-file",
	}
	err = p.Exec(changelogExe, drone.Args(args...), drone.Dir(p.Source))

	return
}
