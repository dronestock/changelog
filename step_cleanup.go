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

func (c *stepCleanup) Run(_ context.Context) error {
	return os.RemoveAll(c.configFilepath())
}
