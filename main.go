package main // import "github.com/niwho/skrobot"

import (
	"fmt"
	"github.com/niwho/logs"
	"github.com/niwho/skrobot/command"
	"github.com/niwho/skrobot/config"
	"github.com/niwho/skrobot/pump"
	"github.com/niwho/skrobot/timer"
	"github.com/nlopes/slack"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.LoadConf("./conf.yml")

	logs.InitLog("skrobot.log", 3, 5)

	timer.InitSimpeTimer()
	go func() {
		pump.InitPump().SetCommandHandler(command.InitCommand()).StartLoop()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	onStop(c, 5*time.Second, func() {
		//kafkaConsumer.Close()
		//_ = local_cache.CacheMng.Close()
	})
}

func onStop(c chan os.Signal, timeout time.Duration, work func()) {
	select {
	case <-c:
		logs.Log(logs.F{"stop": time.Now()}).Info("begin stop")
		go func() {
			if _, ok := <-time.After(timeout); ok {
				logs.Log(logs.F{"stop": time.Now()}).Warn("time out")
				os.Exit(-1)
			}
		}()
		work()
	}
}




func test(){
	api := slack.New("xoxb-760186935424-759729619652-9aCpwBAHi0EC3i5Z6YyHb9Oc")
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// api.SetDebug(true)
	groups, err := api.GetGroups(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, group := range groups {
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	}

	channels, err := api.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, channel := range channels {
		fmt.Println(channel.Name)
		// channel is of type conversation & groupConversation
		// see all available methods in `conversation.go`
	}

	user, err := api.GetUserInfo("UNBM948UU")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)


	users, _:= api.GetUsers()
	for _, user:= range users {
		fmt.Println(user)
	}

	ims, _:= api.GetIMChannels()
	for _, im:= range ims {
		fmt.Println(im)
	}

	//return
	userID := "UNBM948UU"

	_, _, channelID, err := api.OpenIMChannel(userID)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	_ = channelID
	//api.PostMessage(channelID, slack.MsgOptionText("Hello World!", false))
	api.SendMessage("random", slack.MsgOptionText("Hello World!", false),
		slack.MsgOptionAttachments(slack.Attachment{ImageURL:"http://pic16.nipic.com/20111006/6239936_092702973000_2.jpg"}))
}
