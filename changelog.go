package main

import (
	"fmt"
	"strings"

	"github.com/dronestock/drone"
)

func (p *plugin) changelog() (undo bool, err error) {
	// 防止出现错误：fatal: unsafe repository
	if err = p.git(`config`, `--global`, `--add`, `safe.directory`, p.Folder); nil != err {
		return
	}

	if `` == strings.TrimSpace(p.From) {
		args := []interface{}{
			`describe`,
			`--tags`,
			`--abbrev=0`,
		}
		latestErr := p.Exec(exeGit, drone.Args(args...), drone.String(&p.From), drone.Dir(p.Folder))
		if nil != latestErr {
			p.From = ``
		}
	}

	args := []interface{}{
		`--config`, changelogConfigFilename,
		`--repository-url`, p.Url,
		`--output`, p.Output,
	}

	// JIRA集成
	if `` != strings.TrimSpace(p.Jira.Url) {
		args = append(args, `--jira-url`, p.Jira.Url)
	}
	if `` != strings.TrimSpace(p.Jira.Username) {
		args = append(args, `--jira-username`, p.Jira.Username)
	}
	if `` != strings.TrimSpace(p.Jira.Token) {
		args = append(args, `--jira-token`, p.Jira.Token)
	}

	// 加入标签选择参数
	from := strings.TrimSpace(p.From)
	to := strings.TrimSpace(p.To)
	switch {
	case `` != from && `` != to:
		args = append(args, fmt.Sprintf(`%s..%s`, from, to))
	case `` != from && `` == to:
		args = append(args, fmt.Sprintf(`%s..`, from))
	case `` == from && `` != to:
		args = append(args, fmt.Sprintf(`..%s`, to))
	}

	// 执行命令
	err = p.Exec(exeChangelog, drone.Args(args...), drone.Dir(p.Folder))

	return
}
