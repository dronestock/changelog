package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/goexl/gox/tpl"
)

type stepBuild struct {
	*plugin
}

func newBuildStep(plugin *plugin) *stepBuild {
	return &stepBuild{
		plugin: plugin,
	}
}

func (b *stepBuild) Runnable() bool {
	return true
}

func (b *stepBuild) Run(_ context.Context) (err error) {
	// 写入配置文件
	config := new(config)
	config.Style = b.Style
	config.Title = b.Title
	config.Types = b.Types
	if err = tpl.New(b.Conf).Data(config).Build().File(b.Filepath.Config); nil != err {
		return
	}

	// 写入模板文件
	if err = os.WriteFile(b.Filepath.Template, []byte(b.Template), os.ModePerm); nil != err {
		return
	}

	args := []any{
		"--config", b.Filepath.Config,
		"--template", b.Filepath.Template,
		"--repository-url", b.Url,
		"--output", b.Output,
	}

	// JIRA集成
	if nil != b.Jira {
		args = append(args, "--jira-url", b.Jira.Url)
		args = append(args, "--jira-username", b.Jira.Username)
		args = append(args, "--jira-token", b.Jira.Token)
	}

	// 加入标签选择参数
	from := strings.TrimSpace(b.From)
	to := strings.TrimSpace(b.To)
	switch {
	case "" != from && "" != to:
		args = append(args, fmt.Sprintf("%s..%s", from, to))
	case "" != from && "" == to:
		args = append(args, fmt.Sprintf("%s..", from))
	case "" == from && "" != to:
		args = append(args, fmt.Sprintf("..%s", to))
	}

	// 执行命令
	err = b.Command(changelogExe).Args(args...).Dir(b.Source).Exec()

	return
}
