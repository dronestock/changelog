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

	// 额外配置
	Jira jira `default:"${JIRA}"`
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
	}
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("output", p.Output),
	}
}
