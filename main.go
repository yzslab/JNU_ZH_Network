package main

import (
	"flag"
	"fmt"
	"jnu_network/utils"
	"time"
)

func main() {
	configPath := flag.String("config", "./config.yml", "配置文件地址, 默认为\"./config.yml\"")
	logFlag := flag.Bool("log", false, "日志输出到文件,默认为false即输出到终端, 打开则输入到文件")
	flag.Parse()
	utils.LogInit(*logFlag)
	fmt.Println("使用配置:", *configPath)
	utils.Config.GetConf(*configPath)
	for {
		loginFlag, loginStruct := utils.Login()
		// keep trying to login until success
		if loginFlag == false {
			time.Sleep(time.Second * 1)
			continue // retry login
		}

		// begin heartbeat
		hbErrorCount := 0
		hbcount := 0
		for hbErrorCount < 5 {
			hbFlag := utils.HeartBeat(loginStruct, &hbcount)
			if !hbFlag {
				// if error occurred, retry
				hbErrorCount++
				time.Sleep(time.Second * 1)
			} else {
				hbErrorCount = 0 // reset error count if heartbeat success
				time.Sleep(time.Second * time.Duration(utils.Config.HBTime))
			}
		}
	}
}
