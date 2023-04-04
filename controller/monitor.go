package controller

import (
	bf "auto-clean/basefunc"
	"log"
	"time"
)

func StartMonitoring(dir string, filtername string, filtertime string, intervalTime int64) {
	filtertime = "-" + filtertime + "s"
	for {
		time.Sleep(time.Duration(intervalTime) * time.Second)
		resTmp, err := bf.FilterFileList(dir, filtername, filtertime)
		if err != nil {
			log.Println(err)
		}
		bf.RemoveFile(resTmp)
	}
}
