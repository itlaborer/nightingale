package stra

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/didi/nightingale/src/model"
	"github.com/didi/nightingale/src/toolkits/str"

	"github.com/toolkits/pkg/file"
	"github.com/toolkits/pkg/logger"
)

func NewPortCollect(port, step int, tags string) *model.PortCollect {
	return &model.PortCollect{
		CollectType: "port",
		Port:        port,
		Step:        step,
		Tags:        tags,
	}
}

func GetPortCollects() map[int]*model.PortCollect {
	portPath := StraConfig.PortPath
	ports := make(map[int]*model.PortCollect)

	if StraConfig.Enable {
		ports = Collect.GetPorts()
		for _, p := range ports {
			tagsMap := str.DictedTagstring(p.Tags)
			tagsMap["port"] = strconv.Itoa(p.Port)

			p.Tags = str.SortedTags(tagsMap)
		}
	}

	files, err := file.FilesUnder(portPath)
	if err != nil {
		logger.Error(err)
		return ports
	}
	//扫描文件采集配置
	for _, f := range files {
		port, step, err := parsePorFile(f)
		if err != nil {
			logger.Warning(err)
			continue
		}

		service, err := file.ToTrimString(StraConfig.PortPath + "/" + f)
		if err != nil {
			logger.Warning(err)
			continue
		}

		tags := fmt.Sprintf("port=%s,service=%s", strconv.Itoa(port), service)
		p := NewPortCollect(port, step, tags)
		ports[p.Port] = p
	}

	return ports
}

func parsePorFile(fname string) (port, step int, err error) {
	arr := strings.Split(fname, "_")
	if len(arr) < 2 {
		err = fmt.Errorf("illegal filename:%s, split _ < 2", fname)
		return
	}

	step, err = strconv.Atoi(arr[0])
	if err != nil {
		err = fmt.Errorf("illegal filename:%s, error:%v", fname, err)
		return
	}

	port, err = strconv.Atoi(arr[1])
	if err != nil {
		err = fmt.Errorf("illegal filename:%s, error:%v", fname, err)
		return
	}
	return
}
