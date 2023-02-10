package main

type jira struct {
	// 地址
	Url string `json:"url" validate:"required"`
	// 用户名
	Username string `json:"username" validate:"required"`
	// 令牌
	Token string `json:"token" validate:"required"`
}
