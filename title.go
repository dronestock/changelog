package main

type title struct {
	// 新功能
	Feat string `default:"✨ Features | 新功能"`
	// 修复
	Fix string `default:"🐛 Bug Fixes | Bug修复"`
	// 性能优化
	Perf string `default:"📈 Performance Improvements | 性能优化"`
	// 回退
	Revert string `default:"⏪ Reverts | 回退"`
	// 回退
	Chore string `default:"📦 Chores | 其他更新"`
	// 文档
	Docs string `default:"📝 Documentation | 文档"`
	// 风格
	Style string `default:"🌈 Styles | 风格"`
	// 代码重构
	Refactor string `default:"🔄 Code Refactoring | 代码重构"`
	// 测试
	Test string `default:"✅ Tests | 测试"`
	// 构建
	Build string `default:"👷‍ Build System | 构建"`
}
