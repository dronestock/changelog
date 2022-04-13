package main

type jira struct {
	// 地址
	Url string `json:"url"`
	// 用户名
	Username string `json:"username"`
	// 令牌
	Token string `json:"token"`
}
