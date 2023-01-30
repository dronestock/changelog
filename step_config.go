package main

import (
	"context"
	"encoding/json"
	"os"
)

type stepConfig struct {
	*plugin

	config config
}

func newConfigStep(plugin *plugin) *stepConfig {
	return &stepConfig{
		plugin: plugin,

		config: config{
			Header: plugin.Header,
			Types: []typ{
				{Type: "feat", Section: plugin.Feat},
				{Type: "fix", Section: plugin.Fix},
				{Type: "perf", Section: plugin.Perf},
				{Type: "revert", Section: plugin.Revert},
				{Type: "chore", Section: plugin.Chore},
				{Type: "docs", Section: plugin.Docs},
				{Type: "style", Section: plugin.Style},
				{Type: "refactor", Section: plugin.Refactor},
				{Type: "test", Section: plugin.Test},
				{Type: "build", Section: plugin.Build},
				{Type: "ci", Section: plugin.Ci},
			},
		},
	}
}

func (c *stepConfig) Runnable() bool {
	return true
}

func (c *stepConfig) Run(_ context.Context) (err error) {
	if bytes, me := json.Marshal(c.config); nil != me {
		err = me
	} else {
		err = os.WriteFile(c.configFilepath(), bytes, os.ModePerm)
	}

	return
}
