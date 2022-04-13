package main

import (
	"github.com/goexl/gox"
)

func (p *plugin) push() (undo bool, err error) {
	if undo = !p.Push; undo {
		return
	}

	// 远端名，为了不和原来仓库的远端名冲突，使用随机字符串
	name := gox.RandString(defaultNameLength)

	// 提交
	if err = p.git(`commit`, p.Output, `--message`, p.Message); nil != err {
		return
	}

	// 添加远程仓库地址
	if err = p.git(`name`, `add`, name, p.Remote); nil != err {
		return
	}

	// 推送
	if err = p.git(`push`, `--set-upstream`, `origin`, p.Branch); nil != err {
		return
	}

	return
}
