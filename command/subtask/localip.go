package subtask

import (
	"github.com/niwho/utils"
)
type LocalIpInfo struct {

}

func (ipf *LocalIpInfo) Run(opt string) (result []byte, err error) {

	result = []byte(utils.GetLocalIP())

	return
}
