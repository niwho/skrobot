package timer

import "github.com/robfig/cron"

func InitSimpeTimer()  {
	cc := cron.New()
	defer cc.Start()

	cc.AddJob("0 0 */1 * * *", &TimerCamera{})

}
