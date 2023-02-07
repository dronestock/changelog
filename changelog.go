package main

import (
	"github.com/goexl/git"
)

const configTpl = `
style: {{ .Style }}
template: CHANGELOG.tpl.md
info:
  title: {{ .Title.Info }}
  repository_url: https://github.com/dronestock/changelog
options:
  commits:
    filters:
      Type:
        - feat
        - fix
        - perf
        - refactor
        - chore
  commit_groups:
    title_maps:
      feat: {{ .Title.Feat }}
      fix: {{ .Title.Fix }}
      perf: {{ .Title.Perf }}
      refactor: {{ .Title.Refactor }}
      chore: {{ .Title.Chore }}
  header:
    pattern: "^(\\w*)(?:\\(([\\w\\$\\.\\-\\*\\s]*)\\))?\\:\\s(.*)$"
    pattern_maps:
      - Type
      - Scope
      - Subject
  notes:
    keywords:
      - BREAKING CHANGE
`

func (p *plugin) changelog() (undo bool, err error) {
	// 防止出现错误：fatal: unsafe repository
	if err = git.SafeDirectory(git.Folder(p.Folder)); nil != err {
		return
	}

}
