package main

import (
	"github.com/dronestock/drone"
)

func (p *plugin) git(args ...interface{}) error {
	return p.Exec(exeGit, drone.Args(args...), drone.Dir(p.Folder))
}
