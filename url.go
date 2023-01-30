package main

type url struct {
	Commit  string `default:"{{host}}/{{owner}}/{{repository}}/commit/{{hash}}" json:"commit"`
	Compare string `default:"{{host}}/{{owner}}/{{repository}}/compare/{{previousTag}}...{{currentTag}}" json:"compare"`
	Issue   string `default:"{{host}}/{{owner}}/{{repository}}/issues/{{id}}" json:"issue"`
	User    string `default:"{{host}}/{{user}}" json:"user"`
}
