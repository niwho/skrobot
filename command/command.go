package command

import (
	"fmt"
	"github.com/niwho/skrobot/command/proto"
	"github.com/niwho/skrobot/command/subtask"
	"regexp"
)

var CommandIns *Command

type Command struct {
	task map[string]proto.ICommandItem
}

func InitCommand() *Command {
	CommandIns = &Command{
		task: map[string]proto.ICommandItem{},
	}

	CommandIns.Register("ph", &subtask.PiCamera{})
	CommandIns.Register("ip", &subtask.LocalIpInfo{})
	return CommandIns
}

// cmd:param1 param2 param3
// 执行结果：xxx
const (
	t = `(<@[\w\d]+>)?(?:\s+)?(\w+)(.*)`
)

var (
	CMDRE = regexp.MustCompile(t)
)

func (cd *Command) Register(cmdName string, command proto.ICommandItem) {
	cd.task[cmdName] = command
}

func (cd *Command) GetCommand(cs string) (cmd, other string) {
	ps := CMDRE.FindStringSubmatch(cs)
	if ps == nil {
		return
	}
	cmd = ps[2]
	other = ps[3]
	for k, pp := range ps {
		fmt.Println(k, pp)
	}
	return
}

func (cd *Command) DoCommand(cs string) (result []byte, err error) {
	cmdName, opt := cd.GetCommand(cs)
	_ = opt
	cmdExute, ok := cd.task[cmdName]
	if !ok {
		return
	}
	result, err = cmdExute.Run(opt)
	return
}
