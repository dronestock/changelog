package main

import (
	"path/filepath"

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

	// 头
	Header string `default:"${HEADER=# 更新历史 \n\n}"`
	// 新功能
	Feat string `default:"${FEAT=✨ Features | 新功能}"`
	// 修复
	Fix string `default:"${FIX=🐛 Bug Fixes | Bug 修复}"`
	// 性能优化
	Perf string `default:"${PERF=⚡ Performance Improvements | 性能优化}"`
	// 回退
	Revert string `default:"${REVERT=⏪ Reverts | 回退}"`
	// 回退
	Chore string `default:"${CHORE=📦 Chores | 其他更新}"`
	// 文档
	Docs string `default:"${DOCS=📝 Documentation | 文档}"`
	// 风格
	Style string `default:"${STYLE=💄 Styles | 风格}"`
	// 代码重构
	Refactor string `default:"${REFACTOR=♻ Code Refactoring | 代码重构}"`
	// 测试
	Test string `default:"${TEST=✅ Tests | 测试}"`
	// 构建
	Build string `default:"${BUILD=👷‍ Build System | 构建}"`
	// 持续集成
	Ci string `default:"${CI=🔧 Continuous Integration | CI 配置}"`
	// 地址格式
	Url url `default:"${URL}"`
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(newConfigStep(p)).Name("配置").Build(),
		drone.NewStep(newBuildStep(p)).Name("生成").Build(),
		drone.NewStep(newCleanupStep(p)).Name("清理").Build(),
	}
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("output", p.Output),
	}
}

func (p *plugin) configFilepath() string {
	return filepath.Join(p.Source, configFilename)
}
