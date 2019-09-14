package timer

import "github.com/niwho/skrobot/command/subtask"

type TimerCamera struct {

}

func (TimerCamera) Run(){
	ca := &subtask.PiCamera{}
	ca.Run("")
}
