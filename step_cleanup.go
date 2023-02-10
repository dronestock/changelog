package main

import (
	"context"
	"os"
)

type stepCleanup struct {
	*plugin
}

func newCleanupStep(plugin *plugin) *stepCleanup {
	return &stepCleanup{
		plugin: plugin,
	}
}

func (c *stepCleanup) Runnable() bool {
	return true
}

func (c *stepCleanup) Run(_ context.Context) (err error) {
	if ce := os.Remove(c.Filepath.Config); nil != ce {
		err = ce
	} else if te := os.Remove(c.Filepath.Template); nil != te {
		err = te
	}

	return
}
