package main

import (
	"context"
	"encoding/json"
)

type stepBuild struct {
	*plugin

	types []typ
}

func newBuildStep(plugin *plugin) *stepBuild {
	return &stepBuild{
		plugin: plugin,

		types: []typ{
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
	}
}

func (s *stepBuild) Runnable() bool {
	return true
}

func (s *stepBuild) Run(_ context.Context) (err error) {
	args := []any{
		"--header",
		s.Header,
	}
	if nil != s.Skip.Bump && *s.Skip.Bump {
		args = append(args, "--skip.bump")
	}
	if nil != s.Skip.Changelog && *s.Skip.Changelog {
		args = append(args, "--skip.changelog")
	}
	if nil != s.Skip.Commit && *s.Skip.Commit {
		args = append(args, "--skip.commit")
	}
	if nil != s.Skip.Tag && *s.Skip.Tag {
		args = append(args, "--skip.tag")
	}
	if "" != s.Version {
		args = append(args, "--release-as", s.Version)
	}

	if types, me := json.Marshal(s.types); nil != me {
		err = me
	} else {
		args = append(args, "--types", string(types))
		args = append(args, "--commitUrlFormat", s.Url.Commit)
		args = append(args, "--compareUrlFormat", s.Url.Compare)
		args = append(args, "--issueUrlFormat", s.Url.Issue)
		args = append(args, "--userUrlFormat", s.Url.User)
		err = s.Command(changelogExe).Args(args...).Dir(s.Source).Exec()
	}

	return
}
