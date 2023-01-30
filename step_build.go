package main

import (
	"context"
)

type stepBuild struct {
	*plugin
}

func newBuildStep(plugin *plugin) *stepBuild {
	return &stepBuild{
		plugin: plugin,
	}
}

func (s *stepBuild) Runnable() bool {
	return true
}

func (s *stepBuild) Run(_ context.Context) (err error) {
	args := []any{
		"--preset",
		"angular",
		"---infile",
		s.Output,
		"--same-file",
	}
	err = s.Command(changelogExe).Args(args...).Dir(s.Source).Exec()

	return
}
