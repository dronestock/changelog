package main

import (
	"fmt"

	"github.com/goexl/git"
	"github.com/goexl/gox"
)

func (p *plugin) push() (undo bool, err error) {
	if undo = !p.pushable(); undo {
		return
	}

	// 设置全局用户
	if err = git.User(p.Author, p.Email, git.Dir(p.Folder)); nil != err {
		return
	}

	// 只提交输出文件，不提交其它多余的文件
	message := fmt.Sprintf(commitMessageFormatter, p.Message)
	if err = git.Commit(message, git.Filenames(p.Output)); nil != err {
		return
	}

	// 远端名，为了不和原来仓库的远端名冲突，使用随机字符串
	name := gox.RandString(defaultNameLength)
	// 添加远程仓库地址
	if err = git.Remote(p.Remote, git.Name(name)); nil != err {
		return
	}

	// 推送
	err = git.Push(git.Name(name), git.Branch(p.Branch))

	return
}
