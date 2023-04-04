package main

import (
	con "auto-clean/controller"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("start monitoring")
	pwdPath, err := os.Getwd()
	if err != nil {
		log.Println("Get pwdPATH ERROR: ", err)
		return
	}
	readRes, err := con.ReadConfig(pwdPath)
	if err != nil {
		log.Println("read config error: ", err)
		return
	}
	configJson, _ := json.MarshalIndent(readRes, "", " ")
	log.Println(string(configJson))

	for k, v := range readRes.Service {
		fmt.Println(k)
		go con.StartMonitoring(v.Cleanpath, v.Cleanname, v.Cyclelength, readRes.Interval)
	}

	for {
		time.Sleep(60 * time.Second)
	}

}
