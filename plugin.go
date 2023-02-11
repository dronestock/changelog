package main

import (
	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 源文件
	Source string `default:"${SOURCE=.}"`
	// 输出
	Output string `default:"${OUTPUT=CHANGELOG.md}" validate:"required"`
	// 开始版本
	From string `default:"${FROM}"`
	// 结束版本
	To string `default:"${TO}"`
	// 当前版本
	Tag string `default:"${TAG}"`
	// 下个版本
	Next string `default:"${NEXT}"`

	// 仓库地址
	Url string `default:"${URL=${DRONE_REPO_LINK}}"`
	// 配置
	Conf string `default:"${CONFIG}"`
	// 样式
	Style string `default:"${STYLE=github}"`
	// 主题
	Subject string `default:"# 更新历史 \n\n"`
	// 标题
	Title title `default:"${TITLE}"`
	// 可以导出的类型
	Types []string `default:"['feat', 'fix', 'pref', 'refactor', 'chore']"`
	// 模板
	Template string `default:"${TEMPLATE}"`
	// 路径
	Filepath filepath `default:"${FILEPATH}"`
	// 额外配置
	Jira *jira `default:"${JIRA}"`
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(newBuildStep(p)).Name("生成").Build(),
		drone.NewStep(newCleanupStep(p)).Name("清理").Build(),
	}
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("output", p.Output),
	}
}
