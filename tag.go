package main

type tag struct {
	// 当前版本
	Name string `json:"name" validate:"required"`
	// 下个版本
	Next string `json:"next"`
}
