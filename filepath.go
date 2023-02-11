package main

type filepath struct {
	Config   string `default:"config.yml" json:"config"`
	Template string `default:"changelog.tpl.md" json:"template"`
}
