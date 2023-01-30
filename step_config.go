package main

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
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

func (s *stepConfig) Runnable() bool {
	return true
}

func (s *stepConfig) Run(_ context.Context) (err error) {
	if bytes, me := json.Marshal(s.config); nil != me {
		err = me
	} else {
		err = os.WriteFile(filepath.Join(s.Source, configFilename), bytes, os.ModePerm)
	}

	return
}
