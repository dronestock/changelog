package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dronestock/drone"
)

func (p *plugin) changelog() (undo bool, err error) {
	if `` == strings.TrimSpace(p.From) {
		err = p.Exec(exeGit, drone.Args(`describe`, `--tags`, `--abbrev=0`), drone.String(&p.From))
	}
	if nil != err {
		return
	}

	args := []interface{}{
		`--repository-url`, p.Remote,
		`--output`, filepath.Join(p.Folder, p.Output),
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
	args = append(args, fmt.Sprintf(`%s..%s`, p.From, p.To))

	// 执行命令
	err = p.Exec(exeChangelog, drone.Args(args...), drone.Dir(p.Folder))

	return
}
