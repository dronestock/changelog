package main

type skip struct {
	Bump      *bool `json:"bump"`
	Changelog *bool `json:"changelog"`
	Commit    *bool `default:"true" json:"commit"`
	Tag       *bool `default:"true" json:"tag"`
}
