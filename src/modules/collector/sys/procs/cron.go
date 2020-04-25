package procs

import (
	"time"

	"github.com/didi/nightingale/src/modules/collector/stra"
)

func Detect() {
	detect()
	go loopDetect()
}

func detect() {
	ps := stra.GetProcCollects()
	DelNoProcCollect(ps)
	AddNewProcCollect(ps)
}

func loopDetect() {
	for {
		time.Sleep(time.Second * 10)
		detect()
	}
}
