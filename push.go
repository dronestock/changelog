package main

import (
	"fmt"

	"github.com/goexl/gox"
)

func (p *plugin) push() (undo bool, err error) {
	if undo = !p.Push; undo {
		return
	}

	// 设置用户名
	if err = p.git(`config`, `--global`, `user.name`, p.Author); nil != err {
		return
	}

	// 设置邮箱
	if err = p.git(`config`, `--global`, `user.email`, p.Email); nil != err {
		return
	}

	// 加入更新日志
	if err = p.git(`add`, p.Output); nil != err {
		return
	}

	// 提交
	message := fmt.Sprintf(commitMessageFormatter, p.Message)
	if err = p.git(`commit`, `--message`, message, `--allow-empty`, p.Output); nil != err {
		return
	}

	// 远端名，为了不和原来仓库的远端名冲突，使用随机字符串
	name := gox.RandString(defaultNameLength)
	// 添加远程仓库地址
	if err = p.git(`remote`, `add`, name, p.Remote); nil != err {
		return
	}

	// 推送
	err = p.git(`push`, `--set-upstream`, name, p.Branch)

	return
}
