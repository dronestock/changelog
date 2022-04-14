package main

type title struct {
	Info     string `default:"更新日志" json:"info"`
	Feat     string `default:"新功能" json:"feature"`
	Fix      string `default:"修复" json:"bugfix"`
	Perf     string `default:"性能优化" json:"perf"`
	Refactor string `default:"重构" json:"refactor"`
	Chore    string `default:"持续集成" json:"chore"`
}
