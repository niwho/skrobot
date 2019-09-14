package subtask

import (
	"github.com/niwho/logs"
	"os"
	"os/exec"
	"time"
)

type PiCamera struct {
}

func (ca *PiCamera) do(opt string) (result []byte, err error) {
	base := "/mnt/c/camera/"
	l, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(l)
	dir := base + now.Format("/2006-01-02/")
	_, err = os.Stat(dir)
	if err != nil {
		os.Mkdir(dir, os.ModePerm)
	}
	imgFile := now.Format("2006-01-02-15-04-05") + "pi.jpg"
	// _ = imgFile
	cc := exec.Command("raspistill", "-t", "2000", "-o", dir+imgFile)
	out, err := cc.Output()
	if err != nil {
		logs.Log(logs.F{"err": err.Error(), "out": string(out)}).Error()
	}
	result = []byte(dir+imgFile)
	return
}

func (ca *PiCamera) Run(opt string) ([]byte, error) {

	//return []byte("camera"), nil
	return ca.do(opt)
}

