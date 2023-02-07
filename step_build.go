package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/goexl/git"
	"github.com/goexl/gox"
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
	// 更新日志是基于标签来实现的，应该根据标签的个数来确定相应的逻辑
	// 没有标签，不填充from和to，使用next-tag参数
	// 只有一个标签，只填充to
	// 大于等于两个标签，填充from和to
	var count int64
	if count, err = git.Count(); nil != err {
		return
	}

	switch {
	case 1 == count && `` == strings.TrimSpace(b.To) && `` != strings.TrimSpace(b.Tag):
		b.To, err = git.Tag(git.Dir(b.Source))
	case 2 <= count && `` == strings.TrimSpace(b.From) && `` == strings.TrimSpace(b.Tag):
		b.From, err = git.Tag(git.Dir(b.Source))
	case 2 <= count && `` == strings.TrimSpace(b.From) && `` != strings.TrimSpace(b.Tag):
		b.From, err = git.Tag(git.Second(), git.Dir(b.Source))
	}
	if nil != err {
		return
	}

	// 写入配置文件
	if err = gox.RenderToFile(changelogConfigFilename, configTpl, b.Changelog, false); nil != err {
		return
	}

	args := []interface{}{
		"--config", changelogConfigFilename,
		"--template", changelogTplFilename,
		"--repository-url", b.Changelog.Url,
		"--output", b.Output,
	}

	// JIRA集成
	if `` != strings.TrimSpace(b.Jira.Url) {
		args = append(args, `--jira-url`, b.Jira.Url)
	}
	if `` != strings.TrimSpace(b.Jira.Username) {
		args = append(args, `--jira-username`, b.Jira.Username)
	}
	if `` != strings.TrimSpace(b.Jira.Token) {
		args = append(args, `--jira-token`, b.Jira.Token)
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
