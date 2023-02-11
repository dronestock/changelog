package main

type filepath struct {
	Config   string `default:"${CONFIG_PATH=config.yml}" json:"config"`
	Template string `default:"${TEMPLATE_PATH=changelog.tpl.md}" json:"template"`
}
