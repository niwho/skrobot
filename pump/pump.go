package pump

import (
	"fmt"
	"github.com/niwho/logs"
	"github.com/niwho/skrobot/command/proto"
	"github.com/niwho/skrobot/config"
	"github.com/nlopes/slack"
	"strings"
)

var PumpIns *Pump

type Pump struct {
	cmd    proto.ICommand
	client *slack.Client
}

func InitPump() *Pump {
	p := &Pump{
		client: slack.New(
			config.Conf.AppToken, //"xoxb-760186935424-759729619652-9aCpwBAHi0EC3i5Z6YyHb9Oc",
			//slack.OptionDebug(true),
			//slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
		),
	}
	PumpIns = p
	return p
}

func (pp *Pump) SetCommandHandler(command proto.ICommand) *Pump {
	pp.cmd = command
	return pp
}

func (pp *Pump) StartLoop() {
	api := pp.client

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			continue
			// Ignore hello

		case *slack.ConnectedEvent:
			logs.Log(logs.F{"Infos:": ev.Info, "Connection counter:": ev.ConnectionCount}).Info()
			continue
			// Replace C2147483705 with your Channel ID
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "C2147483705"))

		case *slack.MessageEvent:
			// 过滤不处理bot发送的消息
			if ev.SubType == "bot_message" {
				continue
			}
			logs.Log(logs.F{"Message": ev, "channel": ev.Channel}).Info()
			_, err := api.GetChannelInfo(ev.Channel)
			if err == nil {
				if !strings.HasPrefix(ev.Text, fmt.Sprintf("<@%s>", config.Conf.BotId)) {
					continue
				}
			}

			go func() {
				result, err := pp.cmd.DoCommand(ev.Text)
				if err != nil {
					logs.Log(logs.F{"cmd": ev.Text, "err": err.Error()}).Error()
					api.SendMessage(ev.Channel,
						slack.MsgOptionAsUser(true),
						slack.MsgOptionUser(config.Conf.BotId),
						slack.MsgOptionText(string(err.Error()), false))
				} else {
					api.SendMessage(ev.Channel,
						slack.MsgOptionAsUser(true),
						slack.MsgOptionUser(config.Conf.BotId),
						slack.MsgOptionText(string(result), false))
				}
			}()

		case *slack.PresenceChangeEvent:
			//fmt.Printf("Presence Change: %v\n", ev)
			logs.Log(logs.F{"ev": ev}).Info("Presence Change")

		case *slack.LatencyReport:
			//fmt.Printf("Current latency: %v\n", ev.Value)
			logs.Log(logs.F{"ev": ev}).Debug("Current latency")

		case *slack.RTMError:
			//fmt.Printf("Error: %s\n", ev.Error())
			logs.Log(logs.F{"ev": ev.Error()}).Error("RTMError")

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			logs.Log(logs.F{"ev": "Invalid credentials"}).Error("RTMError")
			return

		default:

			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
