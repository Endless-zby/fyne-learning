package util

import (
	"errors"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	"os"
)

func CheckAutoRun() {
	// 1. 获取当前程序的可执行文件路径
	exePath, err := os.Executable()
	if err != nil {
		log.Printf("无法获取可执行文件路径: %v", err)
	}

	// 2. 检查是否已注册开机启动
	key, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Printf("无法打开注册表键: %v\n", err)
	}
	defer func(key registry.Key) {
		err := key.Close()
		if err != nil {
			log.Printf("%v", err)
		}
	}(key)

	_, _, err = key.GetStringValue("此处填写程序名(注册表key)")
	if errors.Is(err, registry.ErrNotExist) {
		// 3. 如果未注册，则进行注册
		err = key.SetStringValue("此处填写程序名(注册表key)", exePath)
		if err != nil {
			log.Printf("无法设置注册表值: %v", err)
		}
		fmt.Println("已注册开机启动")
	} else if err != nil {
		log.Printf("无法获取注册表值: %v", err)
	} else {
		// 已注册
		fmt.Println("已经注册开机启动")
	}
}
